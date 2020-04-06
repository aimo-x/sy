package app

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/aimo-x/sy/conf"
	"github.com/aimo-x/sy/lib/jwt"
	"github.com/aimo-x/sy/lib/wechat"
	"github.com/aimo-x/sy/lib/wechat/cache"
	"github.com/aimo-x/sy/model"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const wxmpredirectURI = "https://iuu.pub/v2/api/wxmp/callback/"

// Wxmp ...
type Wxmp struct {
	OpenID     string // OpenID
	UserID     uint   // 用户ID
	CampaignID uint   // 活动ID
}

// Route ...
func (uw *Wxmp) Route(r *gin.RouterGroup) {
	r.Any("callback/:CampaignID", uw.CallBack)
	r.GET("oauthurl", uw.OauthURL)
	r.GET("authorization_token", uw.UseCodeToToken)
}

// MiddleWare 中间件
func (uw *Wxmp) MiddleWare(c *gin.Context) {
	jwtToken := c.GetHeader("Authorization")
	if len(jwtToken) < 20 {
		rwErr("没有授权信息", errors.New("Not Find Authorization"), c)
		c.Abort()
		return
	}
	var j jwt.Wxmp
	OpenID, UserID, CampaignID, err := j.Verify(jwtToken)
	if err != nil {
		rwErr("验证授权失败", err, c)
		c.Abort()
		return
	}
	uid, err := strconv.Atoi(UserID)
	if err != nil {
		rwErr("UserID ERROR", err, c)
		c.Abort()
		return
	}
	cid, err := strconv.Atoi(CampaignID)
	if err != nil {
		rwErr("CampaignID ERROR", err, c)
		c.Abort()
		return
	}
	uw.OpenID = OpenID
	uw.UserID = uint(uid)
	uw.CampaignID = uint(cid)
	c.Next()
}

// CallBack ...
func (uw *Wxmp) CallBack(c *gin.Context) {
	CampaignID := c.Param("CampaignID")
	wx := uw.GetWeChat()
	oauth := wx.GetOauth()
	rat, err := oauth.GetUserAccessToken(c.Request.FormValue("code"))
	if err != nil {
		rwErr("授权错误", err, c)
		return
	}
	userInfo, err := oauth.GetUserInfo(rat.AccessToken, rat.OpenID)
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + userInfo.ErrMsg + "</h1>"))
		return
	}

	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}

	db, err := Mysql()
	if err != nil {
		// rwErr("数据库连接错误", err, c)
		c.Writer.Write([]byte("<title>数据库连接错误</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	defer db.Close()
	// 开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	defer tx.Close()
	CampaignIDInt, err := strconv.Atoi(CampaignID)
	if err != nil {
		tx.Rollback()
		c.Writer.Write([]byte("<title>CampaignID 错误</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	var ca model.Campaign
	row := tx.Where("id = ?", uint(CampaignIDInt)).First(&ca)
	if row.Error != nil || row.RecordNotFound() {
		tx.Rollback()
		c.Writer.Write([]byte("<title>找不到此活动</title><h1>" + fmt.Sprint(row.Error) + "</h1>"))
		return
	}
	var u model.User
	row = tx.Where("wx_open_id = ?", userInfo.OpenID).First(&u)
	if row.RecordNotFound() {
		u.CampaignID = uint(CampaignIDInt)
		u.UserUUID = uuid.New().String()
		u.WxOpenID = userInfo.OpenID
		u.HeadImg = userInfo.HeadImgURL
		u.NickName = base64.StdEncoding.EncodeToString([]byte(userInfo.Nickname))
		u.Sex = userInfo.Sex
		u.City = userInfo.City
		u.Country = userInfo.Country
		u.Province = userInfo.Province
		u.Unionid = userInfo.Unionid
		err = tx.Create(&u).Error
		if err != nil {
			tx.Rollback()
			c.Writer.Write([]byte("<title>用户写入错误</title><h1>" + fmt.Sprint(err) + "</h1>"))
			return
		}
		row.Error = nil
	}
	CampaignID = strconv.Itoa(int(u.CampaignID))
	if row.Error != nil {
		tx.Rollback()
		c.Writer.Write([]byte("<title>查找用户时，数据库错误</title><h1>" + fmt.Sprint(row.Error) + "</h1>"))
		return
	}
	var j jwt.Wxmp
	token, err := j.Token(u.WxOpenID, strconv.Itoa(int(u.ID)), CampaignID)
	if err != nil {
		tx.Rollback()
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	code, err := uw.useTokenToCode(token)
	if err != nil {
		tx.Rollback()
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	tx.Commit()
	state := c.Request.FormValue("state")
	if strings.Index(state, "?") == -1 {
		c.Redirect(302, state+"?oauth=wechat&&code="+code)
	} else {
		c.Redirect(302, state+"&oauth=wechat&&code="+code)
	}

}

// OauthURL ...
func (uw *Wxmp) OauthURL(c *gin.Context) {
	wx := uw.GetWeChat()
	oauth := wx.GetOauth()
	CampaignID := c.Request.FormValue("CampaignID")
	var scope, state = "snsapi_userinfo", c.Request.FormValue("state")
	uri, err := oauth.GetRedirectURL(wxmpredirectURI+CampaignID, scope, state)
	if err != nil {
		rwErr("获取授权地址错误", err, c)
		return
	}
	rwSus(uri, c)
}

// useTokenToCode 使用token 存入兑换码
func (uw *Wxmp) useTokenToCode(token string) (code string, err error) {
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err = client.Ping().Result()
	if err != nil {
		return code, err
	}
	code = uw.RandomCode(16) + strconv.FormatInt(time.Now().Unix(), 10)
	_, err = client.Set("useCodeToToken"+code, token, time.Minute*5).Result()
	if err != nil {
		return code, err
	}
	return code, err
}

// UseCodeToToken 使用code 换取token
func (uw *Wxmp) UseCodeToToken(c *gin.Context) {
	code := c.Request.FormValue("code")
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	token, err := client.Get("useCodeToToken" + code).Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	_, err = client.Del("useCodeToToken" + code).Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	rwSus(token, c)
}

// RandomCode 随机码
func (uw *Wxmp) RandomCode(n int) string {
	str := "0123456789asdfghjklqwertyuiopzxcvbnm"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetWeChat 示例
func (uw *Wxmp) GetWeChat() (wx *wechat.Wechat) {
	var opts cache.RedisOpts
	opts.Host = conf.Redis().Addr
	opts.Password = conf.Redis().Password
	opts.Database = conf.Redis().DB
	Redis := cache.NewRedis(&opts)
	var cfg wechat.Config
	cfg.AppID = conf.GetConf().WeChat.AppID
	cfg.AppSecret = conf.GetConf().WeChat.AppSecret
	cfg.Cache = Redis
	wx = wechat.NewWechat(&cfg)
	return wx

	/*
		var opts cache.RedisOpts
		opts.Host = conf.Redis().Addr
		opts.Password = conf.Redis().Password
		opts.Database = 1
		Redis := cache.NewRedis(&opts)
		var cfg wechat.Config
		cfg.AppID = "wxbdb9cd64895da3d3"                   // "wxa67a64f664dfba26"                   // conf.GetConf().WeChat.AppID         // "wxa67a64f664dfba26"
		cfg.AppSecret = "25295943ffeaa1e9e8b7de4c8588eaf0" // "2ddcf1edc8a54ca2de8b2ebd8f15fcca" // conf.GetConf().WeChat.AppSecret //  "2ddcf1edc8a54ca2de8b2ebd8f15fcca"
		cfg.Cache = Redis
		wx = wechat.NewWechat(&cfg)
		return wx
	*/
}

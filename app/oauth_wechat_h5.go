package app

import (
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

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

const redirectURI = "https://iuu.pub/v2/api/oauth_wechat_h5/callback"

// OauthWechatH5 ...
type OauthWechatH5 struct {
	OpenID   string
	AppID    string
	NickName string
}

// Route ...
func (uw *OauthWechatH5) Route(r *gin.RouterGroup) {
	r.Any("callback", uw.CallBack)
	r.GET("oauthurl", uw.OauthURL)
	r.GET("authorization_token", uw.UseCodeToToken)
}

// MiddleWare 中间件
func (uw *OauthWechatH5) MiddleWare(c *gin.Context) {
	jwtToken := c.GetHeader("Authorization")
	if len(jwtToken) < 20 {
		rwErr("没有授权信息", errors.New("Not Find Authorization"), c)
		c.Abort()
		return
	}
	var j jwt.WeChat
	openid, appid, nickname, err := j.Verify(jwtToken)
	if err != nil {
		rwErr("验证授权失败", err, c)
		c.Abort()
		return
	}
	fmt.Println(openid, appid, nickname, err)

	uw.OpenID = openid
	uw.AppID = appid
	uw.NickName = nickname
	c.Next()
}

// CallBack ...
func (uw *OauthWechatH5) CallBack(c *gin.Context) {
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
	var j jwt.WeChat
	token, err := j.Token(userInfo.OpenID, wx.Context.AppID, userInfo.Nickname)
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	code, err := uw.useTokenToCode(token)
	if err != nil {
		c.Writer.Write([]byte("<title>授权登陆失败</title><h1>" + fmt.Sprint(err) + "</h1>"))
		return
	}
	state := c.Request.FormValue("state")
	if strings.Index(state, "?") == -1 {
		c.Redirect(302, state+"?oauth=wechat&&code="+code)
	} else {
		c.Redirect(302, state+"&oauth=wechat&&code="+code)
	}

}

// OauthURL ...
func (uw *OauthWechatH5) OauthURL(c *gin.Context) {
	wx := uw.GetWeChat()
	oauth := wx.GetOauth()
	var scope, state = "snsapi_userinfo", c.Request.FormValue("state")
	uri, err := oauth.GetRedirectURL(redirectURI, scope, state)
	if err != nil {
		rwErr("获取授权地址错误", err, c)
		return
	}
	rwSus(uri, c)
}

// useTokenToCode 使用token 存入兑换码
func (uw *OauthWechatH5) useTokenToCode(token string) (code string, err error) {
	client := redis.NewClient(conf.Redis())
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
func (uw *OauthWechatH5) UseCodeToToken(c *gin.Context) {
	code := c.Request.FormValue("code")
	client := redis.NewClient(conf.Redis())
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
func (uw *OauthWechatH5) RandomCode(n int) string {
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
func (uw *OauthWechatH5) GetWeChat() (wx *wechat.Wechat) {
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

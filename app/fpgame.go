package app

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/aimo-x/sy/conf"
	"github.com/aimo-x/sy/lib/jwt"
	"github.com/aimo-x/sy/log"
	"github.com/aimo-x/sy/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
	"github.com/skip2/go-qrcode"
)

// FpGame ...
type FpGame struct {
}

// Route 游戏路由
func (f *FpGame) Route(r *gin.RouterGroup) {
	m := r.Group("")
	// m.Use(f.W.MiddleWare)                                   // 认证中间件
	m.Use(f.CampaignIDGWF)                                  // 验证身份与与接口不一致用户
	m.GET("getNowDayUserGameInfo", f.getNowDayUserGameInfo) // 获取今日游戏信息
	p := m.Group("")
	p.Use(f.signatureVerification)
	p.PUT("putAchievement", f.putAchievement)       // 提交成绩返回得分 body {Time:1.11}
	m.PUT("getShareGameCount", f.getShareGameCount) // 分享增加游戏次数
	m.GET("getQrCode", f.getQrCode)                 // 获取用户UUID
	m.GET("getSelf", f.getSelf)                     // 获取排名信息
	m.Any("isGoGame", f.isGoGame)                   // 验证游戏
	r.GET("getNowDayRank", f.getNowDayRank)         // 日排行
	r.GET("getRank", f.getRank)                     // 总排行
	r.GET("getQrCodeInfo", f.getQrCodeInfo)         // 查看二维码信息
	m.GET("getTime", f.getTime)
	m.POST("do", f.do)
}
func (f *FpGame) wxTime(c *gin.Context) {
	/*
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	*/
}

const key = "tqexAUG4Tb3EnfzCadhCwwIyvViy0YPh"

func getUserInfo(c *gin.Context) (ROpenID string, RUserID, RCampaignID uint, err error) {
	jwtToken := c.GetHeader("Authorization")
	if len(jwtToken) < 20 {
		err = errors.New("Not Find Authorization")
		return
	}
	var j jwt.Wxmp
	OpenID, UserID, CampaignID, err := j.Verify(jwtToken)
	if err != nil {
		return
	}
	uid, err := strconv.Atoi(UserID)
	if err != nil {
		return
	}
	cid, err := strconv.Atoi(CampaignID)
	if err != nil {
		return
	}
	ROpenID = OpenID
	RUserID = uint(uid)
	RCampaignID = uint(cid)
	return
}

// 游戏动作日志v
func (f *FpGame) do(c *gin.Context) {
	CampaignIDInt, err := strconv.Atoi(c.Request.FormValue("CampaignID"))
	if err != nil {
		rwErr("CampaignID ERROR", err, c)
		c.Abort()
		return
	}
	m := c.Request.FormValue("m")
	v := c.Request.FormValue("v")
	if m == "init" || m == "fpmp3" || m == "cwmp3" || m == "cgmp3" || m == "jsmp3" || m == "time2" || m == "time" || strings.Index(m, "a") != -1 {
		db, err := Mysql()
		if err != nil {
			rwErr("数据库连接错误", err, c)
			return
		}
		defer db.Close()
		_, UserID, _, err := getUserInfo(c)
		if err != nil {
			rwErr("认证错误", err, c)
			return
		}
		var do model.FpGameDo
		do.CampaignID = uint(CampaignIDInt)
		do.Do = m
		do.UserID = UserID
		do.Value = v
		err = db.Create(&do).Error
		if err != nil {
			rwErr("sign expired", err, c)
		}
	}
	rwSus("success", c)
}

// 验证请求有效期 30 秒
func (f *FpGame) signatureVerification(c *gin.Context) {
	// var mapResult map[string]interface{}
	noncestr := c.Request.FormValue("noncestr")
	timestampStr := c.Request.FormValue("timestamp")
	sign := c.Request.FormValue("sign")
	timestamp, err := strconv.Atoi(timestampStr)
	if err != nil {
		rwErr("timestamp error", err, c)
		c.Abort()
		return
	}
	if time.Now().Unix() > int64(timestamp+30) {
		rwErr("sign expired", err, c)
		c.Abort()
		return
	}

	// defer c.Request.Body.Close()
	// io.Copy(res, c.Request.Body)

	// 把request的内容读取出来
	var body []byte

	body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		rwErr("request.Body error ", err, c)
		c.Abort()
		return
	}
	// 把刚刚读出来的再写进去
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	str := "data=" + string(body) + "&noncestr=" + noncestr + key + "&timestamp=" + timestampStr
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	v := hex.EncodeToString(cipherStr)
	if sign == v {
		c.Next()
		return
	}
	rwErr(string(body), v, c)
	c.Abort()
}

// getQrCodeInfo 查看二维码信息
func (f *FpGame) getQrCodeInfo(c *gin.Context) {
	uuid := c.Request.FormValue("uuid")
	var user model.User
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	row := db.Where("user_uuid = ?", uuid).First(&user)
	if row.RecordNotFound() || row.Error != nil {
		rwErr("uuid First()", row.Error, c)
		return
	}
	var fg model.FpGame
	row = db.Where("user_id = ?", user.ID).First(&fg)
	if row.RecordNotFound() || row.Error != nil {
		rwErr("FpGame First()", row.Error, c)
		return
	}
	// decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	nickname, err := base64.StdEncoding.DecodeString(user.NickName)
	if err != nil {
		rwErr("base64 解码错误", err, c)
		return
	}
	type Result struct {
		UserID uint
		Rownum int
	}

	var result Result
	err = db.Raw(`SELECT
						* 
					FROM
						(
							SELECT
							obj.user_id,
							obj.time,
							obj.campaign_id,
						CASE	
								WHEN @rowtotal = obj.time THEN
								@rownum 
								WHEN @rowtotal := obj.time THEN
								@rownum := @rownum + 1 
								WHEN @rowtotal = 0 THEN
								@rownum := @rownum + 1 
							END AS rownum 
						FROM
							( SELECT user_id, time, campaign_id FROM fp_games WHERE time < 999 AND campaign_id = `+strconv.Itoa(int(user.CampaignID))+` ORDER BY time ) AS obj,
							( SELECT @rownum := 0, @rowtotal := NULL ) r 
						) c
					WHERE
						user_id = ?`, user.ID).Scan(&result).Error

	body := "<h2>用户ID：" + strconv.Itoa(int(user.ID)) + "<br></h2>" + "<h2>游戏名次：" + strconv.Itoa(result.Rownum) + "<br></h2>" + "<h2>昵称：" + string(nickname) + "<br></h2>" + "<h2>游戏版本：" + strconv.Itoa(int(fg.CampaignID)) + "<br></h2>" + "<h2>游戏得分：" + fmt.Sprint(fg.Time) + "<br></h2>" + "<h2>openid：" + user.WxOpenID + "<br></h2>" + "<img src='https" + user.HeadImg[4:] + "' width='100%'>"
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, body)
}

// func
func (f *FpGame) isGoGame(c *gin.Context) {
	OpenID, _, _, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err = client.Ping().Result()
	if err != nil {
		rwErr("Redis 链接错误", err, c)
		return
	}
	var TM time.Time
	err = client.Get("gameTime" + OpenID).Scan(&TM) // 找到数据
	if err != nil {
		err = client.Set("gameTime"+OpenID, time.Now(), time.Minute*10).Err()
		if err != nil {
			rwErr("Redis数据异常", err, c)
			return
		}
		rwSus("通过验证", c)
		return
	}
	rwErr("数据异常，请刷新页面后重试再试", err, c)
}

// 查询自己的二维码
func (f *FpGame) getQrCode(c *gin.Context) {
	var user model.User
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	_, UserID, _, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	row := db.Where("id = ?", UserID).First(&user)
	if row.RecordNotFound() || row.Error != nil {
		rwErr("user First()", row.Error, c)
		return
	}
	var png []byte
	png, err = qrcode.Encode("https://iuu.pub/v2/api/fpgame/getQrCodeInfo?uuid="+user.UserUUID, qrcode.Medium, 256)
	if err != nil {
		rwErr("qrcode.Encode(", err, c)
		return
	}
	encodeString := base64.StdEncoding.EncodeToString(png)
	rwSus(encodeString, c)
}

// CampaignIDGWF 游戏中间件 自动注册游戏拦截身份不同的请求
func (f *FpGame) CampaignIDGWF(c *gin.Context) {

	CampaignIDInt, err := strconv.Atoi(c.Request.FormValue("CampaignID"))
	if err != nil {
		rwErr("CampaignID ERROR", err, c)
		c.Abort()
		return
	}
	// uint(CampaignIDInt) = uint(CampaignIDInt)
	// 识别H5游戏发来的ID 与 数据库是否匹配，不匹配就拒绝
	_, _, CampaignID, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	if uint(CampaignIDInt) != CampaignID {
		rwErr("您已参与星得斯的另一款游戏", "无需重复参与", c)
		c.Abort()
		return
	}
	var ca model.Campaign
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		c.Abort()
		return
	}
	defer db.Close()
	row := db.Where("id = ?", uint(CampaignIDInt)).First(&ca)
	if row.Error != nil {
		rwErr("CampaignID First()", row.Error, c)
		c.Abort()
		return
	}
	if ca.StartAt.Unix() > time.Now().Unix() {
		rwErr("活动未开始", err, c)
		c.Abort()
		return
	}
	if ca.ExpAt.Unix() < time.Now().Unix() {
		rwErr("活动已过期", err, c)
		c.Abort()
		return
	}
	c.Next()
}

// 获取今日用户信息游戏，如果没有，就自动注册
func (f *FpGame) getNowDayUserGameInfo(c *gin.Context) {
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		rwErr("系统错误", err, c)
		return
	}
	OpenID, UserID, CampaignID, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	err = client.Del("gameTime" + OpenID).Err()
	if err != nil {
		// 记录异常
		var log = log.New()
		log.Debug("gameTime"+OpenID, err, c.Request.URL.String())
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
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

	CampaignIDInt, err := strconv.Atoi(c.Request.FormValue("CampaignID"))
	if err != nil {
		tx.Rollback()
		rwErr("CampaignID ERROR", err, c)
		c.Abort()
		return
	}
	// uint(CampaignIDInt) = uint(CampaignIDInt)
	// 识别H5游戏发来的ID 与 数据库是否匹配，不匹配就拒绝
	// 查询用户进入
	var fg model.FpGame
	row := tx.Where("user_id = ?", UserID).First(&fg)
	if row.RecordNotFound() {
		/*
			_, err := client.Set("mygame"+OpenID, "ff", time.Hour*24).Result()
			if err != nil {
				tx.Rollback()
				rwErr("系统错误", err, c)
				return
			}*/
		// 自动注册
		fg.CampaignID = CampaignID
		fg.UserID = UserID
		fg.Count = 1
		fg.Share = 1
		fg.Time = 999999
		err = db.Create(&fg).Error
		if err != nil {
			tx.Rollback()
			rwErr("插入游戏数据表时，数据库错误", err, c)
			return
		}
		row.Error = nil
	}
	if row.Error != nil {
		tx.Rollback()
		rwErr("查询游戏数据表时，数据库错误 First()", row.Error, c)
		return
	}
	// 查询今日是否有玩游戏
	var fgl model.FpGameLog
	tn := time.Now()
	row = tx.Where("user_id = ? AND updated_at > ?", UserID, time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location())).First(&fgl)
	if row.Error != nil || row.RecordNotFound() {
		fmt.Println("日志为空")
		// 没有玩游戏，更新游戏次数
		fg.Count = 1
		fg.Share = 1
		row := tx.Model(&fg).Where("user_id = ?", UserID).Updates(map[string]interface{}{"count": fg.Count, "share": fg.Share})
		if row.Error != nil {
			tx.Rollback()
			rwErr("数据库更新游戏次数时发生错误 update()", row.Error, c)
			return
		}
		// 写入新的日志
		fgl.CampaignID = uint(CampaignIDInt)
		fgl.UserID = UserID
		fgl.Time = 999999

		row = tx.Create(&fgl)
		if row.Error != nil {
			tx.Rollback()
			rwErr("日志更新错误 update()", row.Error, c)
			return
		}
		fmt.Println("日志新增成功", fgl)
		row.Error = nil
	}

	/*
		if fgl.UpdatedAt.Unix() < time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location()).Unix() {

			row = tx.Model(&fg).Where("user_id = ? updated_at > ?", UserID, time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location()).Unix()).Updates(map[string]interface{}{"count": fg.Count, "share": fg.Share})
			if row.Error != nil {
				tx.Rollback()
				rwErr("数据库更新游戏次数时发生错误 update()", row.Error, c)
				return
			}
		}
	*/
	var user model.User
	row = tx.Where("id = ?", UserID).First(&user)
	if row.Error != nil {
		tx.Rollback()
		rwErr("查询用户数据出现错误 First", row.Error, c)
		return
	}
	tx.Commit()
	data := gin.H{
		"FpGame": fg,
		"User":   user,
		"fgl":    fgl,
	}
	rwSus(data, c)
}

// getNowDayRank 获取当日排行
func (f *FpGame) getNowDayRank(c *gin.Context) {
	CampaignIDInt, err := strconv.Atoi(c.Request.FormValue("CampaignID"))
	if err != nil {
		rwErr("CampaignID ERROR", err, c)
		c.Abort()
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var fgls []model.FpGameLog
	tn := time.Now()
	row := db.Where("campaign_id = ? AND updated_at > ? AND time < ?", uint(CampaignIDInt), time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location()), 999).Limit(20).Offset(0).Order("time").Find(&fgls)
	if row.Error != nil {
		rwErr("查询游戏数据列表时，数据库错误 Find()", row.Error, c)
		return
	}
	type List struct {
		NickName string
		HeadImg  string
		Time     float64
	}
	var list []List
	for _, fgl := range fgls {
		var user model.User
		row = db.Where("id = ?", fgl.UserID).First(&user)
		if row.Error != nil {
			rwErr("查询游戏数据列表时，数据库错误 Find()", row.Error, c)
			return
		}
		list = append(list, List{NickName: user.NickName, HeadImg: user.HeadImg, Time: fgl.Time})
	}
	rwSus(list, c)
}

// getNowDayRank 获取排行 1/2
func (f *FpGame) getRank(c *gin.Context) {

	CampaignIDInt, err := strconv.Atoi(c.Request.FormValue("CampaignID"))
	if err != nil {
		rwErr("CampaignID ERROR", err, c)
		c.Abort()
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var fgs []model.FpGame
	row := db.Where("campaign_id = ? AND time < ?", uint(CampaignIDInt), 999).Limit(20).Offset(0).Order("time").Find(&fgs)
	if row.Error != nil {
		rwErr("查询游戏数据列表时，数据库错误 Find()", row.Error, c)
		return
	}
	type List struct {
		NickName string
		HeadImg  string
		Time     float64
	}
	var list []List
	for _, fg := range fgs {
		var user model.User
		row = db.Where("id = ?", fg.UserID).First(&user)
		if row.Error != nil {
			rwErr("查询游戏数据列表时，数据库错误 Find()", row.Error, c)
			return
		}
		list = append(list, List{NickName: user.NickName, HeadImg: user.HeadImg, Time: fg.Time})
	}
	rwSus(list, c)
}

// getShareGameCount 分享增加游戏次数
func (f *FpGame) getShareGameCount(c *gin.Context) {
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	_, UserID, _, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	var fg model.FpGame
	row := db.Where("user_id = ?", UserID).First(&fg)
	if row.RecordNotFound() || row.Error != nil {
		rwErr("查找用户数据发生错误", row.Error, c)
		return
	}
	if fg.Share > 0 {
		fg.Count = fg.Count + 50
		fg.Share = fg.Share - 1
		up := map[string]interface{}{"count": fg.Count, "share": fg.Share}
		row = db.Model(&fg).Where("user_id = ?", UserID).Updates(up)
		if row.Error != nil {
			rwErr("更新游戏数据错误 Update()", row.Error, c)
			return
		}
	}
	rwSus(fg, c)
}

// putAchievement 提交成绩
func (f *FpGame) putAchievement(c *gin.Context) {
	OpenID, UserID, CampaignID, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err = client.Ping().Result()
	if err != nil {
		rwErr("Redis 链接错误", err, c)
		return
	}
	var TM time.Time
	err = client.Get("gameTime" + OpenID).Scan(&TM) // 找到数据
	if err != nil {
		rwErr("异常的提交", err, c)
		return
	}
	l := "10000000"

	TMN := time.Now().UnixNano() - TM.UnixNano()

	str := strconv.FormatInt(TMN, 10)
	// 精确至两位小数点
	tmStr := str[:len(str)-len(l)+1]
	fmt.Println("tmStr", tmStr)
	tmDE, err := decimal.NewFromString(tmStr)
	if err != nil {
		rwErr("异常的提交: NewFromString()", err, c)
		return
	}
	tmf64, _ := tmDE.Div(decimal.NewFromInt(100)).Float64()

	/*
		type resJSON struct {
			Time float64 `json:"time"`
		}

		var res resJSON

		err = c.BindJSON(&res)
		if err != nil {
			rwErr("BindJSON error", err, c)
			return
		}
	*/
	// 使用服务器计算时间 删除

	if tmf64 < 9 { // 小于9秒
		rwErr("异常的提交", err, c)
		return
	}

	err = client.Del("gameTime" + OpenID).Err()
	if err != nil {
		rwErr("异常的提交", err, c)
		return
	}

	// res.Time = tmf64
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	defer tx.Close()
	var fg model.FpGame
	row := tx.Where("user_id = ?", UserID).First(&fg)
	if row.RecordNotFound() || row.Error != nil {
		tx.Rollback()
		rwErr("查找用户数据发生错误", row.Error, c)
		return
	}
	if fg.Count < 1 {
		tx.Rollback()
		rwErr("游戏次数已用完", row.Error, c)
		return
	}
	if tmf64 < fg.Time { // 本次成绩比记录的成绩优异时
		fg.Time = tmf64
	}
	fg.Count = fg.Count - 1
	up := map[string]interface{}{"time": fg.Time, "count": fg.Count}
	row = tx.Model(&fg).Where("user_id = ?", UserID).Updates(up)
	if row.Error != nil {
		tx.Rollback()
		rwErr("更新游戏数据错误 Update()", row.Error, c)
		return
	}

	var fgs model.FpGameLog
	// tn := time.Now()
	// "campaign_id = ? AND updated_at > ? AND time < ?", uint(CampaignIDInt), time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location())
	tn := time.Now()
	row = tx.Where("user_id = ? AND updated_at > ?", UserID, time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location())).First(&fgs)
	if row.Error != nil || row.RecordNotFound() {
		tx.Rollback()
		rwErr("查询游戏数据错误 First()", row.Error, c)
		return
	}
	if fgs.Time > tmf64 {
		fgs.Time = tmf64
		row = tx.Model(&fgs).Where("user_id = ? AND updated_at > ?", UserID, time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location())).Update("time", fgs.Time)
		if row.Error != nil || row.RecordNotFound() {
			tx.Rollback()
			rwErr("查询游戏数据错误 Update()", row.Error, c)
			return
		}
	}

	type Result struct {
		UserID uint
		Rownum int
		Time   float64
	}

	var result []Result
	sql := `
	SELECT
	obj.user_id,
	obj.time,
	obj.campaign_id,
CASE	
		WHEN @rowtotal = obj.time THEN
		@rownum 
		WHEN @rowtotal := obj.time THEN
		@rownum := @rownum + 1 
		WHEN @rowtotal = 0 THEN
		@rownum := @rownum + 1 
	END AS rownum 
FROM
	( SELECT user_id, time, campaign_id FROM fp_games WHERE time < 999 AND campaign_id = ` + strconv.Itoa(int(CampaignID)) + ` ORDER BY time ) AS obj,
	( SELECT @rownum := 0, @rowtotal := NULL ) r 
		`
	err = tx.Raw(sql).Scan(&result).Error
	if err != nil {
		tx.Rollback()
		rwErr("查询用户所在排名出错", err, c)
		return
	}
	// 查询总数
	N := len(result)
	var MC int
	for _, rv := range result {
		if tmf64 > rv.Time {
			MC = rv.Rownum
		}
	}

	tx.Commit()
	data := gin.H{"FpGame": fg, "MC": N - MC, "N": N, "time": tmf64}
	rwSus(data, c)
}
func (f *FpGame) getTime(c *gin.Context) {
	OpenID, _, _, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err = client.Ping().Result()
	if err != nil {
		rwErr("Redis 链接错误", err, c)
		return
	}
	var TM time.Time
	err = client.Get("gameTime" + OpenID).Scan(&TM) // 找到数据
	if err != nil {
		rwErr("异常的提交", err, c)
		return
	}
	l := "10000000"
	TMN := time.Now().UnixNano() - TM.UnixNano()
	str := strconv.FormatInt(TMN, 10)
	// 精确至两位小数点
	tmStr := str[:len(str)-len(l)+1]
	fmt.Println("tmStr", tmStr)
	tmDE, err := decimal.NewFromString(tmStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmf64, _ := tmDE.Div(decimal.NewFromInt(100)).Float64()

	rwSus(tmf64, c)
}
func (f *FpGame) getSelf(c *gin.Context) {

	_, _, CampaignID, err := getUserInfo(c)
	if err != nil {
		rwErr("认证错误", err, c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	// 查询本次打败多少对手
	// Scan
	type Result struct {
		UserID uint
		Rownum int
		Time   int
	}

	var result []Result
	sql := `
	SELECT
	obj.user_id,
	obj.time,
	obj.campaign_id,
CASE	
		WHEN @rowtotal = obj.time THEN
		@rownum 
		WHEN @rowtotal := obj.time THEN
		@rownum := @rownum + 1 
		WHEN @rowtotal = 0 THEN
		@rownum := @rownum + 1 
	END AS rownum 
FROM
	( SELECT user_id, time, campaign_id FROM fp_games WHERE time < 999 AND campaign_id = ` + strconv.Itoa(int(CampaignID)) + ` ORDER BY time ) AS obj,
	( SELECT @rownum := 0, @rowtotal := NULL ) r 
	`
	/*
		err = db.Raw(`SELECT
							*
						FROM
							(
							SELECT
								obj.user_id,
								obj.time,
							CASE
									WHEN @rowtotal = obj.time THEN
									@rownum
									WHEN @rowtotal := obj.time THEN
									@rownum := @rownum + 1
									WHEN @rowtotal = 0 THEN
									@rownum := @rownum + 1
								END AS rownum
							FROM
								( SELECT user_id, time FROM fp_games ORDER BY time ) AS obj,
								( SELECT @rownum := 0, @rowtotal := NULL ) r
							) c
						WHERE
							user_id = ?`, UserID).Scan(&result).Error
	*/
	err = db.Raw(sql).Scan(&result).Error
	if err != nil {
		rwErr("查询用户所在排名出错", err, c)
		return
	}
	// 查询总数
	N := len(result)
	var MC int
	data := gin.H{"MC": N - MC, "N": N}
	rwSus(data, c)
}

func (f *FpGame) csv(c *gin.Context) {
	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	defer db.Close()
	var mfcs []ConstructionDevelopmentH5VoteData
	err = db.Model(&ConstructionDevelopmentH5VoteData{}).Where("construction_development_h5_vote_uuid = ?", c.Request.FormValue("construction_development_h5_vote_uuid")).Find(&mfcs).Error
	if err != nil {
		rwErr("db.Where.Delete", err, c)
		return
	}
	// fmt.Println(mfcs)
	header := []string{"城市", "小区", "单元号", "编号", "姓名", "手机", "地址", "菜名", "主图地址", "过程图", "是否审核", "提交时间（格林尼治）", "票数", "openid"} //标题
	columns := [][]string{
		header,
	}
	for i := 0; i < len(mfcs); i++ {
		columns = append(columns, []string{mfcs[i].City, mfcs[i].Community, mfcs[i].UnitNo, strconv.Itoa(mfcs[i].Number), mfcs[i].Name, mfcs[i].Phone, mfcs[i].Address, mfcs[i].DishName, mfcs[i].CompleteImage, mfcs[i].ProcessImages, strconv.FormatBool(mfcs[i].Status), mfcs[i].CreatedAt.String(), strconv.Itoa(mfcs[i].VoteCount), mfcs[i].OpenID})
	}
	path := "usr/csv/" + strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().Minute()) + strconv.Itoa(time.Now().Second()) + ".csv"
	var ex Export
	err = ex.CSV(path, columns)
	if err != nil {
		rwErr(" db.Where.Delete", err, c)
		return
	}
	c.Redirect(302, conf.GetConf().Host+"/"+path)
}

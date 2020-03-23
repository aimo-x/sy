// Package app This app
package app

// 建发投票
import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/aimo-x/sy/conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ConstructionDevelopment struct
type ConstructionDevelopment struct {
	OauthWechatH5 *OauthWechatH5
}

// AutoMigrate sql...
func (cd *ConstructionDevelopment) AutoMigrate() error {
	db, err := Mysql()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.AutoMigrate(
	// &ConstructionDevelopmentH5Vote{},
	// &ConstructionDevelopmentH5VoteViewIPLog{},
	// &ConstructionDevelopmentH5VoteData{},
	).Error
	return err
}

// Route ...
func (cd *ConstructionDevelopment) Route(r *gin.RouterGroup) {
	r.POST("h5/vote/view/ip/log", cd.CreateViewIPLog)
	r.GET("h5/vote/data/all", cd.GetVoteDataAll)
	r.DELETE("h5/vote/data", cd.DeleteVoteData)
	r.PUT("h5/vote/data/examine", cd.ExamineVoteData)
	r.GET("h5/vote/data/all/status", cd.GetVoteDataAllStatus)
	r.GET("h5/vote/data/export_csv", cd.ExportCSV)
	rm := r
	rm.Use(cd.OauthWechatH5.MiddleWare)
	rm.POST("h5/vote/data", cd.CreateVoteData)
	rm.GET("h5/vote/data", cd.GetVoteData)
	rm.PUT("h5/vote/data", cd.UpdatesVoteData)
	rm.PUT("h5/vote/cont", cd.VoteCont)
}

// ConstructionDevelopmentH5Vote 建发投票
type ConstructionDevelopmentH5Vote struct {
	gorm.Model
	UUID             string // 活动UUID
	Name             string // 活动名称
	PageView         int    // 点击量
	UniqueVisitor    int    // 独立访客
	InternetProtocol int    // IP数 可读取 ConstructionDevelopmentH5VoteViewIPLog 记录条数
	VoteCont         int    // 总投票数
}

// ConstructionDevelopmentH5VoteViewUserLog 建发投票
type ConstructionDevelopmentH5VoteViewUserLog struct {
	gorm.Model
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
	UserUUID                          string `json:"user_uuid"`                             // openid 日志
}

// ConstructionDevelopmentH5VoteViewIPLog 建发投票
type ConstructionDevelopmentH5VoteViewIPLog struct {
	gorm.Model
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
	IP                                string // ip地址记录
}

// CreateViewIPLog 日志
func (cd *ConstructionDevelopment) CreateViewIPLog(c *gin.Context) {
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvil ConstructionDevelopmentH5VoteViewIPLog
	cdhvil.ConstructionDevelopmentH5VoteUUID = c.Request.FormValue("construction_development_h5_vote_uuid")
	cdhvil.IP = c.GetHeader("X-Real-IP")
	err = db.Create(&cdhvil).Error
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	rwSus(cdhvil, c)
}

// ConstructionDevelopmentH5VoteData 建发投票
type ConstructionDevelopmentH5VoteData struct {
	gorm.Model
	VoteCount int `json:"vote_count"` // 票数
	// UUID                              string `json:"uuid"`
	OpenID                            string `json:"open_id"`
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
	Status                            bool   `json:"status"`                                // 是否可以被投票
	Number                            int    `json:"number"`                                // 编号
	Show                              string `json:"show"`                                  // show YES NO
	ConstructionDevelopmentH5VoteDataPost
}

// ConstructionDevelopmentH5VoteDataPost data
type ConstructionDevelopmentH5VoteDataPost struct {
	Name          string `json:"name"`           // 姓名
	Phone         string `json:"phone"`          // 手机
	City          string `json:"city"`           // 城市
	Community     string `json:"community"`      // 小区
	UnitNo        string `json:"unit_no"`        // 单元号
	Address       string `json:"address"`        // 地址
	DishName      string `json:"dish_name"`      // 菜名
	CompleteImage string `json:"complete_image"` // 主图
	ProcessImages string `json:"process_images"` // 逗号分隔符
}

// CreateVoteData 参与投票 post ConstructionDevelopmentH5VoteDataPost
func (cd *ConstructionDevelopment) CreateVoteData(c *gin.Context) {

	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
	}

	rwErr("活动已截止提交信息", "活动已截止提交信息", c)
	return
	/*
		db, err := Mysql()
		if err != nil {
			rwErr("数据库连接错误", err, c)
			return
		}
		defer db.Close()
		var cdhvdp ConstructionDevelopmentH5VoteDataPost
		c.BindJSON(&cdhvdp)
		if err != nil {
			rwErr("JSON 数据错误", err, c)
			return
		}
		var Getcdhvds []ConstructionDevelopmentH5VoteData
		var count int
		row0 := db.Where("city = ? AND construction_development_h5_vote_uuid = ?", cdhvdp.City, c.Request.FormValue("construction_development_h5_vote_uuid")).Find(&Getcdhvds).Count(&count)
		if row0.Error != nil && !row0.RecordNotFound() {
			rwErr("数据库查询错误", row0.Error, c)
			return
		}
		var Getcdhvd ConstructionDevelopmentH5VoteData
		var cdhvd ConstructionDevelopmentH5VoteData
		row := db.Where("open_id = ? AND construction_development_h5_vote_uuid = ?", cd.OauthWechatH5.OpenID, c.Request.FormValue("construction_development_h5_vote_uuid")).First(&Getcdhvd)
		if row.RecordNotFound() {
			cdhvd.VoteCount = 0
			cdhvd.Number = count + 1
			cdhvd.OpenID = cd.OauthWechatH5.OpenID
			// cdhvd.UUID = uuid.New().String()
			cdhvd.ConstructionDevelopmentH5VoteUUID = c.Request.FormValue("construction_development_h5_vote_uuid")
			cdhvd.ConstructionDevelopmentH5VoteDataPost = cdhvdp
			cdhvd.Status = false
			row2 := db.Create(&cdhvd)
			if row2.Error != nil {
				rwErr("数据库写入错误", err, c)
				return
			}
			rwSusMsg("提交成功", cdhvd, c)
			return
		}
		if row.Error != nil {
			rwErr("数据库查询错误", row.Error, c)
			return
		}
		// 增加更新
		msi := map[string]interface{}{
			"status": false,
			"name":   cdhvdp.Name,
			"phone":  cdhvdp.Phone,
			// "city":           cdhvdp.City,
			"community":      cdhvdp.Community,
			"unit_no":        cdhvdp.UnitNo,
			"address":        cdhvdp.Address,
			"dish_name":      cdhvdp.DishName,
			"complete_image": cdhvdp.CompleteImage,
			"process_images": cdhvdp.ProcessImages,
		}
		row3 := db.Model(&cdhvd).Where("open_id = ? AND construction_development_h5_vote_uuid = ?", cd.OauthWechatH5.OpenID, c.Request.FormValue("construction_development_h5_vote_uuid")).Updates(msi, false)
		if row3.Error != nil || row3.RecordNotFound() {
			rwErr("Not Find OpenID", row3.Error, c)
			return
		}
		rwSusMsg("更新成功", Getcdhvd, c)
		// rwSus(cdhvd, c)
		// rwErr("请不要重复参与", Getcdhvd, c)
		// rwSus(Getcdhvd, c)
	*/
}

// ExamineVoteData 变更审核状态
func (cd *ConstructionDevelopment) ExamineVoteData(c *gin.Context) {
	if len(c.Request.FormValue("construction_development_h5_vote_uuid")) < 1 || len(c.Request.FormValue("id")) < 1 {
		rwErr("缺少活动uui/目标id", "construction_development_h5_vote_uuid", c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvd ConstructionDevelopmentH5VoteData
	status := false
	if c.Request.FormValue("status") == "ok" {
		status = true
	}
	row := db.Model(&cdhvd).Where("construction_development_h5_vote_uuid = ? AND id = ?", c.Request.FormValue("construction_development_h5_vote_uuid"), c.Request.FormValue("id")).Update("status", status)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find OpenID", row.Error, c)
		return
	}
	rwSus(cdhvd, c)
}

// UpdatesVoteData 更新个信息 post ConstructionDevelopmentH5VoteDataPost
func (cd *ConstructionDevelopment) UpdatesVoteData(c *gin.Context) {
	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvdp ConstructionDevelopmentH5VoteDataPost
	c.BindJSON(&cdhvdp)
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	var cdhvd ConstructionDevelopmentH5VoteData
	cdhvd.ConstructionDevelopmentH5VoteDataPost = cdhvdp
	// cdhvd.Name, cdhvd.Phone, cdhvd.City, cdhvd.Community, cdhvd.Address, cdhvd.DishName, cdhvd.CompleteImage, cdhvd.ProcessImages = cdhvdp.Name, cdhvdp.Phone, cdhvdp.City, cdhvdp.Community, cdhvdp.Address, cdhvdp.DishName, cdhvdp.CompleteImage, cdhvdp.ProcessImages
	row := db.Model(&cdhvd).Where("open_id = ? AND construction_development_h5_vote_uuid = ?", cd.OauthWechatH5.OpenID, c.Request.FormValue("construction_development_h5_vote_uuid")).Updates(cdhvd)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find OpenID", row.Error, c)
		return
	}
	rwSus(cdhvd, c)
}

// VoteCont 进行投票 更新投票信息
func (cd *ConstructionDevelopment) VoteCont(c *gin.Context) {
	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
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
	// 查询数据库 投票记录条数
	var cl ConstructionDevelopmentH5VoteLog
	var timeNow = time.Now()
	var n int
	row1 := tx.Model(&cl).Where("created_at > ? AND from_open_id = ?", time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location()), cd.OauthWechatH5.OpenID).Count(&n)
	if row1.Error != nil { // 非没有记录
		tx.Rollback()
		rwErr("Not Find Log", row1.Error, c)
		return
	}
	if n > 4 {
		tx.Rollback()
		rwErr("今日投票已达上限", row1.Error, c)
		return
	}

	id, err := strconv.Atoi(c.Request.FormValue("id"))
	if err != nil {
		tx.Rollback()
		rwErr("strconv.Atoi error", err, c)
		return
	}
	// 写入日志
	cl.ConstructionDevelopmentH5VoteUUID = c.Request.FormValue("construction_development_h5_vote_uuid")
	cl.FromNickName = base64.StdEncoding.EncodeToString([]byte(cd.OauthWechatH5.NickName))
	cl.FromOpenID = cd.OauthWechatH5.OpenID
	cl.ConstructionDevelopmentH5VoteDataID = uint(id)
	cl.IP = c.ClientIP()
	err = tx.Create(&cl).Error
	if err != nil {
		tx.Rollback()
		rwErr("tx.Create(&cl)", err, c)
		return
	}
	// 更新投票数
	var cdhvd ConstructionDevelopmentH5VoteData
	row := tx.Where("id = ?", uint(id)).First(&cdhvd)
	if row.Error != nil || row.RecordNotFound() {
		tx.Rollback()
		rwErr("Not Find ID", row.Error, c)
		return
	}
	cdhvd.VoteCount = cdhvd.VoteCount + 1
	row2 := tx.Save(&cdhvd)
	if row2.Error != nil || row2.RecordNotFound() {
		tx.Rollback()
		rwErr("Not Find ID", row2.Error, c)
		return
	}
	tx.Commit()
	rwSus(cdhvd, c)
}

// GetVoteData 获取自己的投票信息/获取到可以进行更新操作
func (cd *ConstructionDevelopment) GetVoteData(c *gin.Context) {
	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvd ConstructionDevelopmentH5VoteData
	row := db.Where("open_id = ? AND construction_development_h5_vote_uuid = ?", cd.OauthWechatH5.OpenID, c.Request.FormValue("construction_development_h5_vote_uuid")).First(&cdhvd)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find OpenID", row.Error, c)
		// rwSus("没有数据", c)
		return
	}
	rwSus(cdhvd, c)
}

// GetVoteDataAll 获取自己的投票信息/获取到可以进行更新操作
func (cd *ConstructionDevelopment) GetVoteDataAll(c *gin.Context) {
	if c.Request.FormValue("pass") != "11112222" {
		rwErr("没有权限", "not auth", c)
		return
	}
	city := c.Request.FormValue("city")
	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
	}
	limit, err := strconv.Atoi(c.Request.FormValue("limit"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
		return
	}
	offset, err := strconv.Atoi(c.Request.FormValue("offset"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvds []ConstructionDevelopmentH5VoteData
	TI := time.Date(2020, time.March, 15, 0, 0, 0, 0, time.Now().Location())
	row := db.Where("construction_development_h5_vote_uuid = ? AND city = ? AND created_at > ?", c.Request.FormValue("construction_development_h5_vote_uuid"), city, TI).Limit(limit).Offset(offset).Find(&cdhvds)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find Data", row.Error, c)
		return
	}
	rwSus(cdhvds, c)
}

// ExportCSV 到处
func (cd *ConstructionDevelopment) ExportCSV(c *gin.Context) {
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

// GetVoteDataAllStatus 获取可以投票的列表
func (cd *ConstructionDevelopment) GetVoteDataAllStatus(c *gin.Context) {

	city := c.Request.FormValue("city")
	if c.Request.FormValue("construction_development_h5_vote_uuid") == "" {
		rwErr("缺少活动ID", "construction_development_h5_vote_uuid", c)
		return
	}
	limit, err := strconv.Atoi(c.Request.FormValue("limit"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
		return
	}
	offset, err := strconv.Atoi(c.Request.FormValue("offset"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvds []ConstructionDevelopmentH5VoteData
	TI := time.Date(2020, time.March, 15, 0, 0, 0, 0, time.Now().Location())
	var row *gorm.DB
	row = db.Where("construction_development_h5_vote_uuid = ? AND city = ? AND created_at > ? AND status = ?", c.Request.FormValue("construction_development_h5_vote_uuid"), city, TI, true).Limit(limit).Offset(offset).Find(&cdhvds)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find Data", row.Error, c)
		return
	}
	rwSus(cdhvds, c)
}

// DeleteVoteData 删除
func (cd *ConstructionDevelopment) DeleteVoteData(c *gin.Context) {
	if c.Request.FormValue("pass") != "5566" {
		rwErr("当前不允许删除", "not auth", c)
		return
	}
	if c.Request.FormValue("pass") != "11112222" {
		rwErr("没有权限", "not auth", c)
		return
	}
	if c.Request.FormValue("id") == "" {
		rwErr("缺少活动ID", "id", c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("数据库连接错误", err, c)
		return
	}
	defer db.Close()
	var cdhvd ConstructionDevelopmentH5VoteData
	row := db.Where("id = ?", c.Request.FormValue("id")).Delete(&cdhvd)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find Data", row.Error, c)
		return
	}
	rwSus(cdhvd, c)
}

// ConstructionDevelopmentH5VoteLog 投票日志
type ConstructionDevelopmentH5VoteLog struct {
	gorm.Model
	ConstructionDevelopmentH5VoteDataID uint   `json:"construction_development_h5_vote_data_id"` // 被投票的人id
	ConstructionDevelopmentH5VoteUUID   string `json:"construction_development_h5_vote_uuid"`    // uuid
	FromOpenID                          string `json:"from_open_id"`                             // id
	FromNickName                        string `json:"from_nick_name"`                           // 昵称
	IP                                  string `json:"IP"`                                       // IP
}

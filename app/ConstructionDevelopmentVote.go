// Package app This app
package app

// 建发投票
import (
	"strconv"

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
		&ConstructionDevelopmentH5Vote{},
		&ConstructionDevelopmentH5VoteViewIPLog{},
		&ConstructionDevelopmentH5VoteData{},
	).Error
	return err
}

// Route ...
func (cd *ConstructionDevelopment) Route(r *gin.RouterGroup) {
	r.POST("h5/vote/view/ip/log", cd.CreateViewIPLog)
	r.GET("h5/vote/data/all", cd.GetVoteDataAll)
	r.DELETE("h5/vote/data", cd.DeleteVoteData)
	r.PUT("h5/vote/data/examine", cd.ExamineVoteData)
	rm := r
	rm.Use(cd.OauthWechatH5.MiddleWare)
	rm.POST("h5/vote/data", cd.CreateVoteData)
	rm.GET("h5/vote/data", cd.GetVoteData)
	rm.PUT("h5/vote/data", cd.UpdatesVoteData)
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
	row := db.Where("open_id = ? AND construction_development_h5_vote_uuid = ?", cd.OauthWechatH5.OpenID, c.Request.FormValue("construction_development_h5_vote_uuid")).First(&Getcdhvd)
	if row.RecordNotFound() {
		var cdhvd ConstructionDevelopmentH5VoteData
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
		rwSus(cdhvd, c)
		return
	}
	if row.Error != nil {
		rwErr("数据库查询错误", row.Error, c)
		return
	}
	rwErr("请不要重复参与", Getcdhvd, c)
	// rwSus(Getcdhvd, c)

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
	cdhvd.Name, cdhvd.Phone, cdhvd.City, cdhvd.Community, cdhvd.Address, cdhvd.DishName, cdhvd.CompleteImage, cdhvd.ProcessImages = cdhvdp.Name, cdhvdp.Phone, cdhvdp.City, cdhvdp.Community, cdhvdp.Address, cdhvdp.DishName, cdhvdp.CompleteImage, cdhvdp.ProcessImages
	row := db.Model(&cdhvd).Where("open_id = ? AND construction_development_h5_vote_uuid = ?", cd.OauthWechatH5.OpenID, c.Request.FormValue("construction_development_h5_vote_uuid")).Updates(cdhvd)
	if row.Error != nil || row.RecordNotFound() {
		rwErr("Not Find OpenID", row.Error, c)
		return
	}
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
	row := db.Where("construction_development_h5_vote_uuid = ?", c.Request.FormValue("construction_development_h5_vote_uuid")).Limit(limit).Offset(offset).Find(&cdhvds)
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
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"`           // uuid
	ConstructionDevelopmentH5VoteData string `json:"construction_development_h5_vote_uuid_data_uuid"` // uuid
	FromOpenID                        string `json:"from_open_id"`                                    // id
	FromNickName                      string `json:"from_nick_name"`                                  // 昵称
	IP                                string `json:"IP"`                                              // IP
}

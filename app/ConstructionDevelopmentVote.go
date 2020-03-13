package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// ConstructionDevelopment struct
type ConstructionDevelopment struct {
}

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
	r.POST("h5/vote/view/ip/log", cd.ViewIPLogCreate)

}

// ConstructionDevelopmentH5Vote 建发投票
type ConstructionDevelopmentH5Vote struct {
	gorm.Model
	UUID             string // 活动
	PageView         int    // 点击量
	UniqueVisitor    int    // 独立访客
	InternetProtocol int    // IP数 读取ConstructionDevelopmentH5VoteViewIPLog 记录条数
}

// ConstructionDevelopmentH5VoteViewIPLog 建发投票
type ConstructionDevelopmentH5VoteViewIPLog struct {
	gorm.Model
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
	IP                                string // ip地址记录
}

// ViewIPLogCreate 日志
func (cd *ConstructionDevelopment) ViewIPLogCreate(c *gin.Context) {
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
	VoteCount                         int // 票数
	UUID                              string
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
	ConstructionDevelopmentH5VoteDataPost
}

// ConstructionDevelopmentH5VoteDataPost data
type ConstructionDevelopmentH5VoteDataPost struct {
	OpenID        string `json:"open_id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	Address       string `json:"address"`
	CompleteImage string `json:"complete_image"`
	ProcessImages string `json:"process_images"`
}

// VoteAdd 投票数据
func (cd *ConstructionDevelopment) CreateVoteData(c *gin.Context) {
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
	cdhvd.VoteCount = 0
	cdhvd.UUID = uuid.New().String()
	cdhvd.ConstructionDevelopmentH5VoteUUID = c.Request.FormValue("construction_development_h5_vote_uuid")
	cdhvd.ConstructionDevelopmentH5VoteDataPost = cdhvdp
	err = db.Create(&cdhvd).Error
	if err != nil {
		rwErr("数据库写入错误", err, c)
		return
	}
	rwSus(cdhvd, c)
}

// ConstructionDevelopmentH5VoteLog 投票日志
type ConstructionDevelopmentH5VoteLog struct {
	gorm.Model
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"`           //  uuid
	ConstructionDevelopmentH5VoteData string `json:"construction_development_h5_vote_uuid_data_uuid"` // uuid
	FromOpenID                        string `json:"from_open_id"`                                    // id
	FromNickName                      string `json:"from_nick_name"`                                  // 昵称
	IP                                string `json:"IP"`                                              // IP
}

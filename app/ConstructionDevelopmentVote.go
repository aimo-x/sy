package app

import "github.com/jinzhu/gorm"
// ConstructionDevelopment struct
type ConstructionDevelopment struct{
}

// Route ...
func (cd *ConstructionDevelopment) Route(r *gin.RouterGroup) {
	r.POST("h5/vote/view/ip/log", cd.ViewIPLogCreate)
}

// ViewLog ... 
func (cd *ConstructionDevelopment) ViewIPLogCreate(c *gin.Context){
	type Data struct {
		ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
		IP                                string // ip地址记录
	}
	var data Data
	err := c.BindJSON(&data)
	if err != nil {
		rwErr("BindJSON error", err, c)
		return
	}
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

// ConstructionDevelopmentH5VoteData 建发投票
type ConstructionDevelopmentH5VoteData struct {
	gorm.Model
	VoteCount int // 票数
	UUID      string
	ConstructionDevelopmentH5VotePost
}

// ConstructionDevelopmentH5VotePost data
type ConstructionDevelopmentH5VotePost struct {
	ConstructionDevelopmentH5VoteUUID string `json:"construction_development_h5_vote_uuid"` //  uuid
	OpenID                            string `json:"open_id"`
	Name                              string `json:"name"`
	Phone                             string `json:"phone"`
	City                              string `json:"city"`
	Address                           string `json:"address"`
	CompleteImage                     string `json:"complete_image"`
	ProcessImages                     string `json:"process_images"`
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

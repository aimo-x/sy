package model

import "github.com/jinzhu/gorm"

// FpGame 游戏关联的ID
type FpGame struct {
	gorm.Model
	CampaignID uint
	UserID     uint
	Count      int     // 今日可玩游戏次数
	Share      int     // 当日可用分享次数
	Time       float64 // 游戏完成时间
}

// FpGameLog 游戏日志
type FpGameLog struct {
	gorm.Model
	CampaignID uint
	UserID     uint
	Time       float64 // 游戏完成时间
}

// FpGameDo 游戏关联的ID
type FpGameDo struct {
	gorm.Model
	CampaignID uint
	UserID     uint
	Do         string // 本次游戏点击了啥
	Value      string // 值
}

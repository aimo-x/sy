package model

import "github.com/jinzhu/gorm"

// User 用户表
type User struct {
	gorm.Model
	CampaignID uint
	WxOpenID   string // 微信OpenID
	UserUUID   string // UUID
	HeadImg    string
	NickName   string
	Sex        int32
	City       string
	Name       string
	Phone      string
	Address    string
	Country    string
	Province   string
	Unionid    string
}

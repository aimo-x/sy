package model

import (
	"time"

	"github.com/jinzhu/gorm"
)
// Campaign 活动
type Campaign struct{
	gorm.Model
	Name string
	StartAt time.Time // 活动开始时间
	ExpAt time.Time // 活动过期时间 
}
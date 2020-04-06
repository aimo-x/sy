package app

import (
	"github.com/aimo-x/sy/conf"
	"github.com/aimo-x/sy/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	//	_ "github.com/jinzhu/gorm/dialects/sqlite" // 正确的包
)

// Sqlite3 ...
func Sqlite3() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "sqlite3.sy.db")
	return
}

// Mysql ...
func Mysql() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", conf.GetConf().Mysql.User+":"+conf.GetConf().Mysql.Password+"@tcp("+conf.GetConf().Mysql.Host+":"+conf.GetConf().Mysql.Port+")/"+conf.GetConf().Mysql.Name+"?charset=utf8&parseTime=True&loc=Local")
	return db, err
}

// AutoMigrate ...
func AutoMigrate() error {
	db, err := Mysql()
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.AutoMigrate(
		// &GacmotorVoiceHeartContent{},
		// &MugedaFormContentDB{},
		&ConstructionDevelopmentH5VoteData{},
		&ConstructionDevelopmentH5VoteLog{},
		&XlServiceOrder{},
		&model.Campaign{},
		&model.FpGame{},
		&model.FpGameLog{},
		&model.User{},
		&model.FpGameDo{},
	).Error
	return err
}

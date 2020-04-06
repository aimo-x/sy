package main

import (
	"fmt"

	"github.com/aimo-x/sy/app"
	"github.com/aimo-x/sy/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func xmain() {
	db, err := app.Mysql()
	if err != nil {
		fmt.Println("数据库连接错误", err)
		return
	}
	defer db.Close()
	// 查询本次打败多少对手
	// Scan
	type Result struct {
		UserID uint
		Rownum float64
	}

	var result Result
	err = db.Raw("SELECT * FROM ( SELECT obj.user_id, obj.time, CASE WHEN @rowtotal = obj.time THEN @rownum WHEN @rowtotal := obj.time THEN @rownum := @rownum + 1 WHEN @rowtotal = 0 THEN @rownum := @rownum + 1 END AS rownum FROM ( SELECT user_id, time FROM fp_games ORDER BY time ) AS obj, ( SELECT @rownum := 0, @rowtotal := NULL ) r ) c WHERE user_id = ?", 3).Scan(&result).Error
	if err != nil {
		fmt.Println("查询用户所在排名出错", err)
		return
	}
	fmt.Println(result)
	/*
		// 查询总数
		var N float64
		var fg model.FpGame
		row := db.Model(&fg).Where("campaign_id = ?", 1).Count(&N)
		if row.Error != nil {
			fmt.Println("查询游戏数据总数，数据库错误 Find()", row.Error)
			return
		}
		bfb := result.rowtotal / N
		data := gin.H{"FpGame": fg, "bfb": bfb}
		fmt.Println(data, result)
	*/
}

func BX() {
	var err error
	type Result struct {
		UserID   uint
		Rowtotal int
	}

	// var res []Result
	var c []map[string]string
	engine, err = xorm.NewEngine("mysql", conf.GetConf().Mysql.User+":"+conf.GetConf().Mysql.Password+"@tcp("+conf.GetConf().Mysql.Host+":"+conf.GetConf().Mysql.Port+")/"+conf.GetConf().Mysql.Name+"?charset=utf8&parseTime=True&loc=Local")
	sql := "SELECT * FROM ( SELECT obj.user_id, obj.time, CASE WHEN @rowtotal = obj.time THEN @rownum WHEN @rowtotal := obj.time THEN @rownum := @rownum + 1 WHEN @rowtotal = 0 THEN @rownum := @rownum + 1 END AS rownum FROM ( SELECT user_id, time FROM fp_games ORDER BY time ) AS obj, ( SELECT @rownum := 0, @rowtotal := NULL ) r ) c WHERE user_id = 3;"

	result, err := engine.DB().Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	c, err = result.ToMapString()
	// c[0]["rowtotal"]
	for _, v := range c {
		fmt.Println("rownum", v["time"])
	}
	// fmt.Println("rowtotal", c[0]["rowtotal"])

}

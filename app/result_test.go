package app

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/aimo-x/sy/model"
)

/*
func TestCSV(T *testing.T) {
	var cid = 2
	db, err := Mysql()
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	// defer db.Close()
	var fgls []model.FpGame
	// tn := time.Now()
	err = db.Where("campaign_id = ?", cid).Order("time").Offset(0).Limit(200).Find(&fgls).Error
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println(fgls)
	// fmt.Println(mfcs)
	header := []string{"排名", "昵称", "游戏时间", "用户ID", "OpenID", "头像地址", "提交时间"} //标题
	columns := [][]string{
		header,
	}
	var user model.User

	for i := 0; i < len(fgls); i++ {

		// 参与时间长度大于 2 小时，得分高于 9.16，游戏次数不低于 30 次
		if fgls[i].Time > 9.6 && fgls[i].UpdatedAt.Unix()-fgls[i].CreatedAt.Unix() > 5 {
			//	var user model.User
			row := db.Model(&model.User{}).Where("id = ?", fgls[i].UserID).Scan(&user)
			if row.Error != nil {
				fmt.Println(row.Error, "ID", fgls[i].ID)
				panic(row.Error)
				return
			}
			name, err := base64.StdEncoding.DecodeString(user.NickName)
			if err != nil {
				fmt.Println(err, "ID", fgls[i].ID)
				panic(err)
				return
			}
			columns = append(columns, []string{strconv.Itoa(i + 1), string(name), strconv.FormatFloat(fgls[i].Time, 'g', -1, 64), strconv.Itoa(int(user.ID)), user.WxOpenID, user.HeadImg, fgls[i].CreatedAt.String()})
			// 参与时间长度大于 10 分钟，得分高于 9.16，游戏次数低于 15 次
		}

		// columns = append(columns, []string{mfcs[i].Shape, mfcs[i].Branch, mfcs[i].Name, mfcs[i].JobCode, mfcs[i].Phone, mfcs[i].Mouthpiece, mfcs[i].MouthpieceDesc, mfcs[i].Models, mfcs[i].Slogan, mfcs[i].Originality, mfcs[i].CreatedAt.String()})
	}
	var who string
	if cid == 1 {
		who = "红色"
	} else {
		who = "蓝色"
	}
	path := "./" + who + strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().Minute()) + strconv.Itoa(time.Now().Second()) + ".csv"
	var ex Export
	err = ex.CSV(path, columns)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println(path)
}
*/

func TestCSV(T *testing.T) {
	var cid = 1
	db, err := Mysql()
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	// defer db.Close()
	var fgls []model.FpGameLog
	tn := time.Now()
	ct := time.Date(tn.Year(), tn.Month(), tn.Day(), 0, 0, 0, 0, tn.Location())
	qt := time.Date(tn.Year(), tn.Month(), tn.Day()-1, 0, 0, 0, 0, tn.Location())
	err = db.Where("created_at < ? AND created_at > ? AND campaign_id = ?", ct, qt, cid).Order("time").Offset(0).Limit(500).Find(&fgls).Error
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println(fgls)
	// fmt.Println(mfcs)
	header := []string{"排名", "昵称", "游戏时间", "用户ID", "OpenID", "头像地址", "提交时间"} //标题
	columns := [][]string{
		header,
	}
	var user model.User
	var fg model.FpGame
	for i := 0; i < len(fgls); i++ {
		row := db.Model(&model.FpGame{}).Where("user_id = ?", fgls[i].UserID).Scan(&fg)
		if row.Error != nil {
			fmt.Println(row.Error, "ID", fgls[i].ID)
			panic(row.Error)
			return
		}
		// 参与时间长度大于 2 小时，得分高于 9.16，游戏次数不低于 30 次
		if fgls[i].UpdatedAt.Unix()-fgls[i].CreatedAt.Unix() > 7200 && fgls[i].Time > 9.9 {
			//	var user model.User
			row := db.Model(&model.User{}).Where("id = ?", fgls[i].UserID).Scan(&user)
			if row.Error != nil {
				fmt.Println(row.Error, "ID", fgls[i].ID)
				panic(row.Error)
				return
			}
			name, err := base64.StdEncoding.DecodeString(user.NickName)
			if err != nil {
				fmt.Println(err, "ID", fgls[i].ID)
				panic(err)
				return
			}
			columns = append(columns, []string{strconv.Itoa(i + 1), string(name), strconv.FormatFloat(fgls[i].Time, 'g', -1, 64), strconv.Itoa(int(user.ID)), user.WxOpenID, user.HeadImg, fgls[i].CreatedAt.String()})
			// 参与时间长度大于 10 分钟，得分高于 9.16，游戏次数低于 15 次
		} else if fgls[i].UpdatedAt.Unix()-fgls[i].CreatedAt.Unix() > 600 && fgls[i].Time > 10 {
			//	var user model.User
			row := db.Model(&model.User{}).Where("id = ?", fgls[i].UserID).Scan(&user)
			if row.Error != nil {
				fmt.Println(row.Error, "ID", fgls[i].ID)
				panic(row.Error)
				return
			}
			name, err := base64.StdEncoding.DecodeString(user.NickName)
			if err != nil {
				fmt.Println(err, "ID", fgls[i].ID)
				panic(err)
				return
			}
			columns = append(columns, []string{strconv.Itoa(i + 1), string(name), strconv.FormatFloat(fgls[i].Time, 'g', -1, 64), strconv.Itoa(int(user.ID)), user.WxOpenID, user.HeadImg, fgls[i].CreatedAt.String()})
			// 参与时间长度大于 5 分钟，得分高于 9.16，游戏次数低于 5 次
		} else if fgls[i].UpdatedAt.Unix()-fgls[i].CreatedAt.Unix() > 180 && fgls[i].Time > 13 {
			//	var user model.User
			row := db.Model(&model.User{}).Where("id = ?", fgls[i].UserID).Scan(&user)
			if row.Error != nil {
				fmt.Println(row.Error, "ID", fgls[i].ID)
				panic(row.Error)
				return
			}
			name, err := base64.StdEncoding.DecodeString(user.NickName)
			if err != nil {
				fmt.Println(err, "ID", fgls[i].ID)
				panic(err)
				return
			}
			columns = append(columns, []string{strconv.Itoa(i + 1), string(name), strconv.FormatFloat(fgls[i].Time, 'g', -1, 64), strconv.Itoa(int(user.ID)), user.WxOpenID, user.HeadImg, fgls[i].CreatedAt.String()})
		}

		// columns = append(columns, []string{mfcs[i].Shape, mfcs[i].Branch, mfcs[i].Name, mfcs[i].JobCode, mfcs[i].Phone, mfcs[i].Mouthpiece, mfcs[i].MouthpieceDesc, mfcs[i].Models, mfcs[i].Slogan, mfcs[i].Originality, mfcs[i].CreatedAt.String()})
	}
	var who string
	if cid == 1 {
		who = "红色"
	} else {
		who = "蓝色"
	}
	path := "./" + who + strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().Minute()) + strconv.Itoa(time.Now().Second()) + ".csv"
	var ex Export
	err = ex.CSV(path, columns)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println(path)
}

/*
func TestRSS(T *testing.T) {
	client := redis.NewClient(conf.Redis())
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	OpenID := "ssss"

	var TM time.Time
	err = client.Get("gameTime" + OpenID).Scan(&TM)
	if err != nil {
		token, err := client.Set("gameTime"+OpenID, time.Now(), time.Minute*5).Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("token:", token)
		return
	}
	// 调用多次出现错误
	fmt.Println("TM:", TM.UnixNano())
	// 纳秒转秒 10000000
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
	f, b := tmDE.Div(decimal.NewFromInt(100)).Float64()
	fmt.Println("f:", f, "b:", b)
	// tmF64, err :=
	// fmt.Println("tmF64", tmF64)

			db, err := Mysql()
			if err != nil {
				panic(err)
				return
			}
			defer db.Close()
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
			( SELECT user_id, time, campaign_id FROM fp_games WHERE time < 999 AND campaign_id = 1 ORDER BY time ) AS obj,
			( SELECT @rownum := 0, @rowtotal := NULL ) r
				`
			err = db.Raw(sql).Scan(&result).Error
			if err != nil {
				panic(err)
				return
			}
			// 查询总数
			N := len(result)
			var MC int
			for _, rv := range result {
				if 10 < rv.Time {
					MC = rv.Rownum
				}
			}
			fmt.Println(MC,N,result)

}
*/

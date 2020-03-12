package app

import (
	"strconv"
	"time"

	"github.com/aimo-x/sy/conf"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GacmotorVoiceHeart 木疙瘩表单内容
type GacmotorVoiceHeart struct {
}

// GacmotorVoiceHeartContent 木疙瘩表单内容
type GacmotorVoiceHeartContent struct {
	gorm.Model
	GacmotorVoiceHeartContentPOST
}

// GacmotorVoiceHeartContentPOST 木疙瘩表单内容
type GacmotorVoiceHeartContentPOST struct {
	GacmotorVoiceHeartID string `json:"gacmotor_voice_heart_id"` // 本次活动ID fa857b22-e92e-52ce-b428-8c2ed217a6ea
	Branch               string `json:"branch"`                  // 部门
	Shape                string `json:"shape"`                   // 形式
	Name                 string `json:"name"`                    // 姓名
	JobCode              string `json:"job_code"`                // 工号
	Phone                string `json:"phone"`                   // 手机
	Mouthpiece           string `json:"mouthpiece"`              // 代言人
	MouthpieceDesc       string `json:"mouthpiece_desc"`         //	代言人说明
	Models               string `json:"models"`                  // 车型
	Slogan               string `json:"slogan"`                  // 标语
	Originality          string `json:"originality"`             // 创意说明
}

// Route ...
func (mfc *GacmotorVoiceHeart) Route(r *gin.RouterGroup) {
	r.POST("", mfc.Create)
	r.DELETE("", mfc.Delete)
	r.GET("data", mfc.Data)
	r.GET("export/csv", mfc.ExportCSV)
}

// Create content
func (mfc *GacmotorVoiceHeart) Create(c *gin.Context) {
	db, err := Mysql()
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	defer db.Close()
	/*
		mfid, err := strconv.Atoi(c.Request.FormValue("mf_id"))
		if err != nil {
			rwErr("strconv.Atoi error", err, c)
			return
		}
	*/
	var mfcp GacmotorVoiceHeartContentPOST
	var mfcd GacmotorVoiceHeartContent
	err = c.BindJSON(&mfcp)
	if err != nil {
		rwErr("BindJSON error", err, c)
		return
	}
	mfcd.GacmotorVoiceHeartContentPOST = mfcp
	err = db.Create(&mfcd).Error
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	rwSus(mfcd, c)
}

// Delete 删除
func (mfc *GacmotorVoiceHeart) Delete(c *gin.Context) {
	mfcid, err := strconv.Atoi(c.Request.FormValue("gacmotor_voice_heart_content_id"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	defer db.Close()
	var mfcd GacmotorVoiceHeartContent
	err = db.Where("id = ?", mfcid).Delete(&mfcd).Error
	if err != nil {
		rwErr("db.Where.Delete", err, c)
		return
	}
	rwSus(mfcd, c)
}

// ExportCSV 删除
func (mfc *GacmotorVoiceHeart) ExportCSV(c *gin.Context) {

	db, err := Mysql()
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	defer db.Close()
	var mfcs []GacmotorVoiceHeartContent
	err = db.Model(&GacmotorVoiceHeartContent{}).Where("gacmotor_voice_heart_id = ?", c.Request.FormValue("gacmotor_voice_heart_id")).Find(&mfcs).Error
	if err != nil {
		rwErr("db.Where.Delete", err, c)
		return
	}
	// fmt.Println(mfcs)
	header := []string{"参赛形式", "参赛部门", "员工姓名", "工号", "联系方式", "代言人", "代言人说明", "代言车型", "广告语", "创意说明", "提交时间"} //标题
	columns := [][]string{
		header,
	}
	for i := 0; i < len(mfcs); i++ {
		columns = append(columns, []string{mfcs[i].Shape, mfcs[i].Branch, mfcs[i].Name, mfcs[i].JobCode, mfcs[i].Phone, mfcs[i].Mouthpiece, mfcs[i].MouthpieceDesc, mfcs[i].Models, mfcs[i].Slogan, mfcs[i].Originality, mfcs[i].CreatedAt.String()})
	}
	path := "usr/" + strconv.Itoa(time.Now().Year()) + strconv.Itoa(int(time.Now().Month())) + strconv.Itoa(time.Now().Day()) + strconv.Itoa(time.Now().Hour()) + strconv.Itoa(time.Now().Minute()) + strconv.Itoa(time.Now().Second()) + ".csv"
	var ex Export
	err = ex.CSV(path, columns)
	if err != nil {
		rwErr(" db.Where.Delete", err, c)
		return
	}
	c.Redirect(302, conf.GetConf().Host+"/"+path)
}

// Data 数据
func (mfc *GacmotorVoiceHeart) Data(c *gin.Context) {
	mfid := c.Request.FormValue("gacmotor_voice_heart_id")
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
		rwErr("mysql error", err, c)
		return
	}
	var mfcs []GacmotorVoiceHeartContent
	defer db.Close()
	err = db.Where("gacmotor_voice_heart_id = ?", mfid).Limit(limit).Offset(offset).Find(&mfcs).Error
	if err != nil {
		rwErr(" db.Where.Delete", err, c)
		return
	}
	rwSus(mfcs, c)
}

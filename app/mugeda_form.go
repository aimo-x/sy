package app

import (
	"strconv"
	"time"

	"github.com/aimo-x/sy/conf"
	"github.com/aimo-x/sy/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MugedaForm 木疙瘩表单
type MugedaForm struct {
	gorm.Model
	URL  string `json:"url"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// MugedaFormContent 木疙瘩表单内容
type MugedaFormContent struct {
}

// MugedaFormContentDB 木疙瘩表单内容
type MugedaFormContentDB struct {
	gorm.Model
	MugedaFormID uint   `json:"mugeda_form_id"` // 表单ID
	Name         string `json:"name"`           // 姓名
	Position     string `json:"position"`       // 职位
	Company      string `json:"company"`        // 公司
	Email        string `json:"email"`          // 电子邮件
}

// MugedaFormContentPOST 木疙瘩表单内容
type MugedaFormContentPOST struct {
	MugedaFormID uint   `json:"mugeda_form_id"` // 表单ID
	Name         string `json:"name"`           // 姓名
	Position     string `json:"position"`       // 职位
	Company      string `json:"company"`        // 公司
	Email        string `json:"email"`          // 电子邮件
}

// Route ...
func (mfc *MugedaFormContent) Route(r *gin.RouterGroup) {
	r.POST("", mfc.Create)
	r.DELETE("", mfc.Delete)
	r.GET("data", mfc.Data)
	r.GET("export/csv", mfc.ExportCSV)
}

// Create content
func (mfc *MugedaFormContent) Create(c *gin.Context) {
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
	var mfcp MugedaFormContentPOST
	var mfcd MugedaFormContentDB
	err = c.BindJSON(&mfcp)
	if err != nil {
		rwErr("BindJSON error", err, c)
		return
	}
	mfcd.MugedaFormID = mfcp.MugedaFormID
	mfcd.Name = mfcp.Name
	mfcd.Position = mfcp.Position
	mfcd.Company = mfcp.Company
	mfcd.Email = mfcp.Email
	err = db.Create(&mfcd).Error
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	// 通知邮件
	var sm tools.SMTP
	sm.MailTo = []string{"243928004@qq.com", c.Request.FormValue("mailto")}
	sm.Subject = c.Request.FormValue("subject")
	sm.Title = c.Request.FormValue("title")
	sm.Desc = c.Request.FormValue("desc")
	sm.Name = mfcd.Name
	sm.Company = mfcd.Company
	sm.Position = mfcd.Position
	sm.Email = mfcd.Email
	sm.BtnLink = "https://iuu.pub/v2/usr/mugeda_form/view/index.html"
	sm.BinText = "查看所有"
	err = sm.Send()
	if err != nil {
		rwErr("sm.Send error", err, c)
		return
	}
	rwSus(mfcd, c)
}

// Delete 删除
func (mfc *MugedaFormContent) Delete(c *gin.Context) {
	mfcid, err := strconv.Atoi(c.Request.FormValue("mfc_id"))
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
	var mfcd MugedaFormContentDB
	err = db.Where("id  = ?", mfcid).Delete(&mfcd).Error
	if err != nil {
		rwErr(" db.Where.Delete", err, c)
		return
	}
	rwSus(mfcd, c)
}

// ExportCSV 删除
func (mfc *MugedaFormContent) ExportCSV(c *gin.Context) {
	mfid, err := strconv.Atoi(c.Request.FormValue("mf_id"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
		return
	}
	db, err := Mysql()
	if err != nil {
		rwErr("mysql error", err, c)
		return
	}
	var mfcs []MugedaFormContentDB
	defer db.Close()
	err = db.Model(&MugedaFormContentDB{}).Where("mugeda_form_id = ?", mfid).Find(&mfcs).Error
	if err != nil {
		rwErr("db.Where.Delete", err, c)
		return
	}
	// fmt.Println(mfcs)
	header := []string{"姓名", "职业", "公司", "邮箱", "提交时间"} //标题
	columns := [][]string{
		header,
	}
	for i := 0; i < len(mfcs); i++ {
		columns = append(columns, []string{mfcs[i].Name, mfcs[i].Position, mfcs[i].Company, mfcs[i].Email, mfcs[i].CreatedAt.String()})
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
func (mfc *MugedaFormContent) Data(c *gin.Context) {
	mfid, err := strconv.Atoi(c.Request.FormValue("mf_id"))
	if err != nil {
		rwErr("strconv.Atoi error", err, c)
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
		rwErr("mysql error", err, c)
		return
	}
	var mfcs []MugedaFormContentDB
	defer db.Close()
	err = db.Where("mugeda_form_id = ?", mfid).Limit(limit).Offset(offset).Find(&mfcs).Error
	if err != nil {
		rwErr(" db.Where.Delete", err, c)
		return
	}
	rwSus(mfcs, c)
}

package sy

import (
	"github.com/aimo-x/sy/app"
	"github.com/gin-gonic/gin"
)

var mfc app.MugedaFormContent
var GacmotorVoiceHeart app.GacmotorVoiceHeart
var ConstructionDevelopment app.ConstructionDevelopment

func api(r *gin.RouterGroup) {
	mfc.Route(r.Group("mugeda_form_content"))
	GacmotorVoiceHeart.Route(r.Group("gacmotor_voice_heart"))
	// 注册cd route
	if err := ConstructionDevelopment.AutoMigrate(); err != nil {
		panic(err)
	}
	ConstructionDevelopment.Route(r.Group("construction_development"))

}

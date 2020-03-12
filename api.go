package sy

import (
	"github.com/aimo-x/sy/app"
	"github.com/gin-gonic/gin"
)

var GacmotorVoiceHeart app.GacmotorVoiceHeart
var mfc app.MugedaFormContent

func api(r *gin.RouterGroup) {

	//mfcr := r.Group("mugeda_form_content")
	//mfc.Route(r.Group("mugeda_form_content"))
	mfc.Route(r.Group("mugeda_form_content"))
	GacmotorVoiceHeart.Route(r.Group("gacmotor_voice_heart"))
}

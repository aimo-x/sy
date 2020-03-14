package sy

import (
	"github.com/aimo-x/sy/app"
	"github.com/gin-gonic/gin"
)

var oauthWechatH5 app.OauthWechatH5
var mfc app.MugedaFormContent
var gacmotorVoiceHeart app.GacmotorVoiceHeart
var constructionDevelopment app.ConstructionDevelopment

func api(r *gin.RouterGroup) {
	// 授权登录组件
	oauthWechatH5.Route(r.Group("oauth_wechat_h5"))
	mfc.Route(r.Group("mugeda_form_content"))
	gacmotorVoiceHeart.Route(r.Group("gacmotor_voice_heart"))
	constructionDevelopment.OauthWechatH5 = &oauthWechatH5
	// 注册cd route
	if err := constructionDevelopment.AutoMigrate(); err != nil {
		panic(err)
	}
	constructionDevelopment.Route(r.Group("construction_development"))

}

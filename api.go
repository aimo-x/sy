package sy

import (
	"github.com/aimo-x/sy/app"
	"github.com/gin-gonic/gin"
)

func api(r *gin.RouterGroup) {

	var oauthWechatH5 app.OauthWechatH5
	var wxmp app.Wxmp
	var fg app.FpGame
	var mfc app.MugedaFormContent
	var gacmotorVoiceHeart app.GacmotorVoiceHeart
	var constructionDevelopment app.ConstructionDevelopment
	// 授权登录组件
	oauthWechatH5.Route(r.Group("oauth_wechat_h5"))
	wxmp.Route(r.Group("wxmp"))
	// fg.W = &wxmp
	fg.Route(r.Group("fpgame"))
	mfc.Route(r.Group("mugeda_form_content"))
	gacmotorVoiceHeart.Route(r.Group("gacmotor_voice_heart"))
	constructionDevelopment.OauthWechatH5 = &oauthWechatH5
	// 注册cd route
	if err := constructionDevelopment.AutoMigrate(); err != nil {
		panic(err)
	}
	constructionDevelopment.Route(r.Group("construction_development"))

}

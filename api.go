package sy

import (
	"github.com/aimo-x/sy/app"
	"github.com/gin-gonic/gin"
)

var mfc app.MugedaFormContent

func api(r *gin.RouterGroup) {

	//mfcr := r.Group("mugeda_form_content")
	//mfc.Route(r.Group("mugeda_form_content"))
	mfc.Route(r.Group("mugeda_form_content"))
}

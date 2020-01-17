package sy

import (
	"github.com/aimo-x/sy/conf"
	"github.com/gin-gonic/gin"
)

var cf = conf.GetConf()
var err error

// New 开始
func New() {
	HTTPServe()

}

// HTTPServe port
func HTTPServe() {
	e := gin.New()
	web(e)
	api(e.Group("api"))
	err = e.Run(":" + cf.Port)
	if err != nil {
		panic(err)
	}
	if cf.IsSsl {
		go HTTPServeTLS(e)
	}
}

// HTTPServeTLS ...
func HTTPServeTLS(e *gin.Engine) {
	err = e.RunTLS(":443", cf.SslPem, cf.SslKey)
	if err != nil {
		panic(err)
	}
}

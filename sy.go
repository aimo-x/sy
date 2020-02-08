package sy

import (
	"github.com/aimo-x/sy/app"
	"github.com/aimo-x/sy/conf"
	"github.com/gin-gonic/gin"
)

var cf = conf.GetConf()
var err error

// New 开始
func New() {

	err = app.AutoMigrate()
	if err != nil {
		panic(err)
	}

	HTTPServe()
}

// HTTPServe port
func HTTPServe() {
	e := gin.Default()
	e.Use(middleware)
	web(e)

	e.StaticFS("v2/usr", gin.Dir("./usr", false))
	api(e.Group("v2/api"))
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

func middleware(c *gin.Context) {
	c.Writer.Header().Set("taobao", "acad.taobao.com")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Token")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

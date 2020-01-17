package sy

import "github.com/gin-gonic/gin"

func web(e *gin.Engine) {
	e.StaticFS("usr", gin.Dir("./usr", false))
}

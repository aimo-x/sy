package app

import (
	"fmt"
	"net/http"

	"github.com/aimo-x/sy/log"

	"github.com/gin-gonic/gin"
)

func rwErr(msg, err interface{}, c *gin.Context) {
	var log = log.New()
	log.Debug(msg, err, c.Request.URL.String())
	c.JSON(http.StatusOK, gin.H{"code": "error", "msg": msg, "err": fmt.Sprint(err)})
}
func rwSus(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": "success", "msg": "请求成功", "data": data})
}
func rwSusMsg(msg, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": "success", "msg": msg, "data": data})
}

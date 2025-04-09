package sven

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sven/layout/internal/base/conf"
	"time"
)

func newEngine(
	web *conf.Web,
// todo
) *gin.Engine {
	var engine *gin.Engine
	gin.SetMode(web.Mode)
	web.Default = false
	if web.Default {
		engine = gin.Default()
	} else {
		engine = gin.New()
	}
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong Now:"+time.Now().Format("2006-01-02 15:04:05"))
	})
	//todo register router

	return engine
}

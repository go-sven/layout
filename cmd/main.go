package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	sven "github.com/go-sven/layout/internal/base/app"
	"github.com/go-sven/layout/internal/base/conf"
	"os"
	"syscall"
)

var (
	// Name is the name of the project
	Name = "sven"
	// Version is the version of the project
	Version = "1.0.1"
)

func Run() {
	//load config
	config, err := conf.Load("")
	if err != nil {
		panic("load config err" + err.Error())
	}
	//init app
	app := newApp(config)
	//run app
	fmt.Println("Name :", Name, " Version:", Version)
	err = app.Run(context.Background())
	if err != nil {
		panic(err)
	}
}

func newApp(c *conf.AppConfig) *sven.Application {
	var engine *gin.Engine
	gin.SetMode(c.Web.Mode)
	if c.Web.Default {
		engine = gin.Default()
	} else {
		engine = gin.New()
	}

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	//addr 从配置文件中读取
	server := sven.NewServer(engine, c.Web.Addr)
	return sven.NewApp(
		sven.WithName(Name),
		sven.WithVersion(Version),
		sven.WithSignals([]os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}),
		sven.WithServer(server),
	)
}

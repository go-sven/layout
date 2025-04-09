package cmd

import (
	"context"
	"fmt"
	sven "github.com/go-sven/layout/internal/base/app"
	"github.com/go-sven/layout/internal/base/conf"
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
	server := sven.NewServer(c.Web)
	return sven.NewApp(
		sven.WithName(Name),
		sven.WithVersion(Version),
		sven.WithServer(server),
	)
}

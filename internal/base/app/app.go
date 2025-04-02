package sven

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Application struct
type Application struct {
	id      string
	name    string
	version string
	server  IServer
	signals []os.Signal
}

// NewApp return app
func NewApp(opts ...Option) *Application {
	app := new(Application)
	for _, opt := range opts {
		opt(app)
	}
	// default accept signals
	if len(app.signals) == 0 {
		app.signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	}
	return app
}

type Option func(*Application)

// WithId add id for app
func WithId(id string) Option {
	return func(app *Application) {
		app.id = id
	}
}

// WithName add name for app
func WithName(name string) Option {
	return func(app *Application) {
		app.name = name
	}
}

// WithVersion add version for app
func WithVersion(name string) Option {
	return func(app *Application) {
		app.name = name
	}
}

// WithServer add server for app to run
func WithServer(server IServer) Option {
	return func(app *Application) {
		app.server = server
	}
}

// WithSignals add signals for app
func WithSignals(signals []os.Signal) Option {
	return func(app *Application) {
		app.signals = signals
	}
}

func (app *Application) Run(ctx context.Context) error {
	if app.server == nil {
		return errors.New("The app has no server to run ")
	}
	quit := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(quit, app.signals...)
	go func(service IServer) {
		if err := service.Start(); err != nil {
			fmt.Println("Failed to start server ,err:" + err.Error())
			errChan <- err
		}
	}(app.server)

	select {
	case err := <-errChan:
		fmt.Println("app stop by err:", err)
		_ = app.Stop()
		return err
	case <-ctx.Done():
		fmt.Println("app stop by ctx done")
		return app.Stop()
	case sign := <-quit:
		fmt.Println("app stop by signal:", sign)
		return app.Stop()
	}
}

func (app *Application) Stop() error {
	go func(service IServer) {
		if err := service.Stop(); err != nil {
			fmt.Println("err:", err)
		}
	}(app.server)
	return nil
}

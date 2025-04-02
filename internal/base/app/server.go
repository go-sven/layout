package sven

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type IServer interface {
	Start() error
	Stop() error
}

type Server struct {
	httpServer *http.Server
}

func NewServer(engine *gin.Engine, addr string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: engine,
		}}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

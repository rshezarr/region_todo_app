package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/internal/config"
)

type Server struct {
	srv               *http.Server
	ServerErrorNotify chan error
}

func NewServer(ctx context.Context, cfg *config.Config, router *gin.Engine) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           cfg.API.Port,
			Handler:        router,
			MaxHeaderBytes: cfg.API.MaxHeaderBytes << 20,
			WriteTimeout:   cfg.API.Timeout,
			ReadTimeout:    cfg.API.Timeout,
		},
		ServerErrorNotify: make(chan error, 1),
	}
}

// run server while sending errors to error channel
func (s *Server) Start() {
	s.ServerErrorNotify <- s.srv.ListenAndServe()
}

// accept errors into channel
func (s *Server) ServerErrNotify() <-chan error {
	return s.ServerErrorNotify
}

// shutdown function
func (s *Server) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}

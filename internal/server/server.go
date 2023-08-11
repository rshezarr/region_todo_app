package server

import (
	"context"
	"net/http"
	"time"
	"todo_list/internal/config"
)

type Server struct {
	srv *http.Server
}

func NewServer(ctx *context.Context, cfg *config.Config) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           ":9090",
			Handler:        nil,
			MaxHeaderBytes: 1 << 20,
			WriteTimeout:   time.Duration(5 * time.Second),
			ReadTimeout:    time.Duration(5 * time.Second),
		},
	}
}

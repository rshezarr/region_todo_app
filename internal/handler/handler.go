package handler

import (
	"github.com/gin-gonic/gin"
	"todo_list/internal/service"
)

type Handler struct {
	svc    *service.Service
	router *gin.Engine
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc:    svc,
		router: gin.New(),
	}
}

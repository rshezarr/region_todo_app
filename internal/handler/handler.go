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

func (h *Handler) InitRoutes() *gin.Engine {
	todoList := h.router.Group("/api/todo-list")
	{
		todoList.POST("/tasks", nil)
		todoList.GET("/tasks:status", nil)
		todoList.PUT("/tasks/:id", nil)
		todoList.DELETE("/tasks/:id", nil)
	}

	return h.router
}

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
		todoList.POST("/tasks", h.CreateListHandler)
		todoList.GET("/tasks", h.GetListsHandler)
		todoList.PUT("/tasks/:id", h.UpdateListHandler)
		todoList.DELETE("/tasks/:id", h.DeleteListHandler)
	}

	return h.router
}

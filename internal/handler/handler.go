package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "todo_list/docs"
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
	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	todoList := h.router.Group("/api/todo-list")
	{
		todoList.POST("/tasks", h.CreateListHandler)
		todoList.GET("/tasks", h.GetListsHandler)
		todoList.PUT("/tasks/:id", h.UpdateListHandler)
		todoList.PUT("/tasks/:id/done", h.UpdateListStatusHandler)
		todoList.DELETE("/tasks/:id", h.DeleteListHandler)
	}

	return h.router
}

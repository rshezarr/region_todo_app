package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/internal/model"
)

func (h *Handler) CreateListHandler(c *gin.Context) {
	list := model.List{}

	if err := c.BindJSON(&list); err != nil {
		c.Error(err)
		return
	}

	id, err := h.svc.TodoList.CreateList(c.Copy(), list)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"list_id": id,
	})
}

func (h *Handler) GetListsHandler(c *gin.Context) {

}

func (h *Handler) UpdateListHandler(c *gin.Context) {

}

func (h *Handler) DeleteListHandler(c *gin.Context) {

}

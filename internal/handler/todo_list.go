package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo_list/internal/model"
)

func (h *Handler) CreateListHandler(c *gin.Context) {
	var list model.List

	if err := c.BindJSON(&list); err != nil {
		c.Error(err)
		log.Printf("error while bind json: %v", err)
		return
	}

	id, err := h.svc.TodoList.CreateList(c.Copy(), list)
	if err != nil {
		c.Error(err)
		log.Printf("error while creating list: %v", err)
		return
	}

	_, err = c.Writer.WriteString(id)
	if err != nil {
		c.Error(err)
	}
}

func (h *Handler) GetListsHandler(c *gin.Context) {
	status := c.Query("status")

	lists, err := h.svc.TodoList.GetList(c.Copy(), status)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lists": lists,
	})
}

func (h *Handler) UpdateListHandler(c *gin.Context) {
	id := c.GetString("id")

	newList := model.List{
		ID: id,
	}

	if err := c.BindJSON(&newList); err != nil {
		c.Error(err)
		return
	}

	id, err := h.svc.TodoList.UpdateList(c.Copy(), newList)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, map[string]interface{}{
		"updated_list_id": id,
	})
}

func (h *Handler) DeleteListHandler(c *gin.Context) {
	id := c.GetString("id")

	if err := h.svc.TodoList.DeleteList(c.Copy(), id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

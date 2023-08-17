package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/internal/dto"
)

func (h *Handler) CreateListHandler(c *gin.Context) {
	list := dto.List{}

	if err := c.BindJSON(&list); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err := h.svc.TodoList.CreateList(c.Copy(), list)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) GetListsHandler(c *gin.Context) {
	status := c.Query("status")

	lists, err := h.svc.TodoList.GetList(c.Copy(), status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lists": lists,
	})
}

func (h *Handler) UpdateListHandler(c *gin.Context) {
	paramId := c.Param("id")

	newList := dto.List{}

	if err := c.BindJSON(&newList); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.svc.TodoList.UpdateList(c.Copy(), paramId, newList)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) UpdateListStatusHandler(c *gin.Context) {
	paramId := c.Param("id")

	err := h.svc.TodoList.UpdateStatus(c.Copy(), paramId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteListHandler(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.TodoList.DeleteList(c.Copy(), id); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

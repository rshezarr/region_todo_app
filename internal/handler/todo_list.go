package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_list/internal/dto"
)

// @Summary Create a new list
// @Description Create a new list with the provided data
// @Tags Lists
// @Accept json
// @Produce json
// @Param list body dto.List true "List data"
// @Success 204 {integer} integer "No Content"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /api/todo-list/tasks [post]
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

// @Summary Get lists
// @Description Retrieve lists based on the provided status
// @Tags Lists
// @Accept json
// @Produce json
// @Param status query string false "Status of lists (optional)"
// @Success 200 {object} dto.List "OK"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /api/todo-list/tasks [get]
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

// @Summary Update a list
// @Description Update an existing list with new data
// @Tags Lists
// @Accept json
// @Produce json
// @Param id path string true "ID of the list"
// @Param list body dto.List true "Updated list data"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} errorResponse "Bad Request"
// @Failure 404 {object} errorResponse "Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /api/todo-list/tasks/{id} [put]
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

// @Summary Update list status
// @Description Update the status of an existing list
// @Tags Lists
// @Produce json
// @Param id path string true "ID of the list"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} errorResponse "Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /api/todo-list/tasks/{id}/done [put]
func (h *Handler) UpdateListStatusHandler(c *gin.Context) {
	paramId := c.Param("id")

	err := h.svc.TodoList.UpdateStatus(c.Copy(), paramId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Delete a list
// @Description Delete an existing list by ID
// @Tags Lists
// @Produce json
// @Param id path string true "ID of the list"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} errorResponse "Not Found"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /api/todo-list/tasks/{id} [delete]
func (h *Handler) DeleteListHandler(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.TodoList.DeleteList(c.Copy(), id); err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

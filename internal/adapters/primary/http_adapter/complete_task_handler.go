package http_adapter

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) CompleteTaskHandler(c echo.Context) error {
	str := c.Param("id")
	id, err := uuid.Parse(str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := r.taskService.CompleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{})
}

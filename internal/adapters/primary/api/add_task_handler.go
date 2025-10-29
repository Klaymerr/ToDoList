package api

import (
	"ToDoList/internal/domain/valueobject"
	"ToDoList/internal/domain/valueobject/dto"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) AddTaskHandler(c echo.Context) error {
	var request dto.TextRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	text, err := valueobject.NewText(request.Text)
	if err != nil {
		return err
	}
	if err := r.taskService.CreateTask(*text); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]any{})
}

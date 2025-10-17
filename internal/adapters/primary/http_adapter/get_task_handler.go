package http_adapter

import (
	"ToDoList/internal/domain/valueobject/dto"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

func (r *Router) GetTaskHandler(c echo.Context) error {
	str := c.Param("id")
	id, err := uuid.Parse(str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	task, err := r.taskService.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	text := task.Text()
	return c.JSON(http.StatusOK, dto.TaskResponse{
		ID:        task.ID(),
		Text:      text.Value(),
		Completed: task.Completed(),
	})
}

func (r *Router) GetTasksHandler(c echo.Context) error {
	tasks, err := r.taskService.GetTasks()
	ret := make([]dto.TaskResponse, len(tasks))
	for i, task := range tasks {
		text := task.Text()
		ret[i] = dto.TaskResponse{
			ID:        task.ID(),
			Text:      text.Value(),
			Completed: task.Completed(),
		}
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, ret)
}

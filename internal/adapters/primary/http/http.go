package http

import (
	"ToDoList/internal/domain/service"
	"github.com/labstack/echo"
)

type http struct {
	taskService service.TaskService
}

func Newhttp(ts service.TaskService) *http {
	return &http{taskService: ts}
}

func (http *http) Start() error {
	e := echo.New()

	e.
}

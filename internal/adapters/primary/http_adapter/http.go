package http_adapter

import (
	"ToDoList/internal/domain/service"
	"github.com/labstack/echo"
	"net/http"
)

type Router struct {
	taskService service.TaskService
}

func Newhttp(ts service.TaskService) *Router {
	return &Router{taskService: ts}
}

func (r *Router) Start() error {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{})
	})

	return nil
}

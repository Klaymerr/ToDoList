package api

import (
	"ToDoList/internal/ports/primary"
	"github.com/labstack/echo"
)

type Router struct {
	taskService primary.TaskService
}

func NewRouter(ts primary.TaskService) *Router {
	return &Router{taskService: ts}
}

func (r *Router) Start() error {
	e := echo.New()

	e.POST("/", r.AddTaskHandler)
	e.GET("/", r.GetTasksHandler)
	e.GET("/:id", r.GetTaskHandler)
	e.PATCH("/complete/:id", r.CompleteTaskHandler)

	e.Start(":8080")

	return nil
}

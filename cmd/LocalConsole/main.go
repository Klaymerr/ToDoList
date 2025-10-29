package main

import (
	"ToDoList/internal/adapters/primary/console"
	"ToDoList/internal/adapters/secondary/localMemory"
	"ToDoList/internal/domain/service"
)

func main() {
	taskRepository := localMemory.NewLocalMap()
	taskService := service.NewTaskService(taskRepository)
	primaryAdapter := console.NewConsole(taskService)
	primaryAdapter.Start()
}

package main

import (
	"ToDoList/internal/adapters/primary/console"
	"ToDoList/internal/adapters/secondary/local_memory"
	"ToDoList/internal/domain/service"
)

func main() {
	taskRepository := local_memory.NewLocalMap()
	taskService := service.NewTaskService(taskRepository)
	primaryAdapter := console.NewConsole(taskService)
	primaryAdapter.Start()
}

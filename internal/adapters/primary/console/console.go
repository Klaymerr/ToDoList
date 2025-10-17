package console

import (
	"ToDoList/internal/domain/valueobject"
	"ToDoList/internal/ports/primary"
	"fmt"
	"github.com/google/uuid"
)

type Console struct {
	taskService primary.TaskService
}

func NewConsole(taskService primary.TaskService) *Console {
	return &Console{
		taskService,
	}
}

func (c *Console) Start() {
	for {
		operation := ""
		_, scan := fmt.Scan(&operation)
		if scan != nil {
			return
		}
		switch operation {
		case "Add":
			str := ""
			_, err := fmt.Scan(&str)
			if err != nil {
				return
			}
			text, err := valueobject.NewText(str)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = c.taskService.CreateTask(*text)
			if err != nil {
				return
			}

		case "Complete":
			str := ""
			_, err := fmt.Scan(&str)
			if err != nil {
				return
			}
			id, err := uuid.Parse(str)
			if err != nil {
				fmt.Println(err)
			}

			err = c.taskService.CompleteTask(id)
			if err != nil {
				fmt.Println(err)
			}

		case "Get":
			str := ""
			_, err := fmt.Scan(&str)
			if err != nil {
				return
			}
			id, err := uuid.Parse(str)
			if err != nil {
				fmt.Println(err)
			}
			task, err := c.taskService.GetTask(id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(task)

		case "All":
			tasks, err := c.taskService.GetTasks()
			if err != nil {
				return
			}
			for _, task := range tasks {
				fmt.Println(task.String())
			}
		}
	}
}

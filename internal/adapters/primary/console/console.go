package console

import (
	"ToDoList/internal/domain/valueobject"
	"ToDoList/internal/ports/primary"
	"fmt"
	"strconv"
)

type Console struct {
	taskService primary.TaskService
	maxid       int
}

func NewConsole(taskService primary.TaskService) *Console {
	return &Console{
		taskService,
		0,
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
			err = c.taskService.CreateTask(c.maxid, *text)
			if err != nil {
				return
			}
			c.maxid++

		case "Complete":
			str := ""
			_, err := fmt.Scan(&str)
			if err != nil {
				return
			}
			id, err := strconv.Atoi(str)
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
			id, err := strconv.Atoi(str)
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

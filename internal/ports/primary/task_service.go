package primary

import (
	"ToDoList/internal/domain/entity"
	"ToDoList/internal/domain/valueobject"
)

type TaskService interface {
	CreateTask(id int, text valueobject.Text) error
	CompleteTask(id int) error
	GetTask(id int) (entity.Task, error)
	GetTasks() ([]entity.Task, error)
}

package primary

import (
	"ToDoList/internal/domain/entity"
	"ToDoList/internal/domain/valueobject"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(text valueobject.Text) error
	CompleteTask(id uuid.UUID) error
	GetTask(id uuid.UUID) (entity.Task, error)
	GetTasks() ([]entity.Task, error)
}

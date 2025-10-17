package secondary

import (
	"ToDoList/internal/domain/entity"
	"github.com/google/uuid"
)

type TaskRepository interface {
	CreateTask(task entity.Task) error
	CompleteTask(id uuid.UUID) error
	GetTask(id uuid.UUID) (entity.Task, error)
	GetTasks() ([]entity.Task, error)
}

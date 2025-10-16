package secondary

import "ToDoList/internal/domain/entity"

type TaskRepository interface {
	CreateTask(task entity.Task) error
	CompleteTask(id int) error
	GetTask(id int) (entity.Task, error)
	GetTasks() ([]entity.Task, error)
}

package postgresql

import (
	"ToDoList/internal/adapters/secondary/postgresql/ent"
	"ToDoList/internal/domain/entity"
	"ToDoList/internal/domain/valueobject"
	"context"
	"github.com/google/uuid"
)

func mapEntTaskToDomain(task *ent.Task) (entity.Task, error) {
	text, err := valueobject.NewText(task.Text)
	if err != nil {
		return entity.Task{}, err
	}

	valtask, err := entity.NewTask(
		task.ID,
		*text,
		task.Completed)
	if err != nil {
		return entity.Task{}, err
	}
	return *valtask, nil
}

func (p PostgresqlRepository) GetTask(id uuid.UUID) (entity.Task, error) {
	enttask, err := p.client.Task.Get(context.Background(), id)
	if err != nil {
		return entity.Task{}, err
	}

	task, err := mapEntTaskToDomain(enttask)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (p PostgresqlRepository) GetTasks() ([]entity.Task, error) {
	enttasks, err := p.client.Task.Query().All(context.Background())
	if err != nil {
		return []entity.Task{}, err
	}
	tasks := make([]entity.Task, len(enttasks))
	for i, task := range enttasks {
		tasks[i], err = mapEntTaskToDomain(task)
		if err != nil {
			return []entity.Task{}, err
		}
	}
	return tasks, nil
}

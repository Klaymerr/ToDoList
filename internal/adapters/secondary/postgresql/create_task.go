package postgresql

import (
	"ToDoList/internal/domain/entity"
	"context"
)

func (p PostgresqlRepository) CreateTask(t entity.Task) error {
	text := t.Text()
	_, err := p.client.Task.Create().
		SetID(t.ID()).
		SetCompleted(t.Completed()).
		SetText(text.Value()).
		Save(context.Background())
	return err
}

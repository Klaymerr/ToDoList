package postgresql

import (
	"context"
	"github.com/google/uuid"
)

func (p PostgresqlRepository) CompleteTask(id uuid.UUID) error {
	_, err := p.client.Task.UpdateOneID(id).
		SetCompleted(true).
		Save(context.Background())
	return err
}

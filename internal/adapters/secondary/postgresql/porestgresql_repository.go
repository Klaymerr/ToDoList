package postgresql

import (
	"ToDoList/internal/adapters/secondary/postgresql/ent"
	"context"
	_ "github.com/lib/pq"
)

type PostgresqlRepository struct {
	client *ent.Client
}

func NewPostgresqlRepository(ds string) (*PostgresqlRepository, error) {
	client, err := ent.Open("postgres", ds)

	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	return &PostgresqlRepository{client: client}, nil
}

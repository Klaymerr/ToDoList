package postgresql

import "ToDoList/internal/adapters/secondary/postgresql/ent"

type PostgresqlRepository struct {
	client *ent.Client
}

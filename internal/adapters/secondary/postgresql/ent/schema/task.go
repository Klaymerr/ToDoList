package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Text("text").
			NotEmpty(),
		field.Bool("completed").
			Default(false),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}

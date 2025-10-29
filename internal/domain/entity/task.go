package entity

import (
	"ToDoList/internal/domain/valueobject"
	"fmt"
	"github.com/google/uuid"
)

type Task struct {
	id        uuid.UUID
	text      valueobject.Text
	completed bool
}

func NewTask(id uuid.UUID, t valueobject.Text, c bool) (*Task, error) {
	return &Task{
		id:        id,
		text:      t,
		completed: c,
	}, nil
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) String() string {
	return fmt.Sprintf("id: %v, text: %s, completed: %t", t.id, t.text, t.completed)
}

func (t *Task) ID() uuid.UUID {
	return t.id
}

func (t *Task) Text() valueobject.Text {
	return t.text
}

func (t *Task) Completed() bool {
	return t.completed
}

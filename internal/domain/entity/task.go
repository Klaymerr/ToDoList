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

func NewTask(text valueobject.Text) *Task {
	return &Task{
		id:        uuid.New(),
		text:      text,
		completed: false,
	}
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

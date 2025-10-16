package entity

import (
	"ToDoList/internal/domain/valueobject"
	"fmt"
)

type Task struct {
	id        int
	text      valueobject.Text
	completed bool
}

func NewTask(id int, text valueobject.Text) *Task {
	return &Task{
		id:        id,
		text:      text,
		completed: false,
	}
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) String() string {
	return fmt.Sprintf("id: %d, text: %s, completed: %t", t.id, t.text, t.completed)
}

func (t *Task) ID() int {
	return t.id
}

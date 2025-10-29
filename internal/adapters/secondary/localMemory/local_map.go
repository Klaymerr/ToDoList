package localMemory

import (
	"ToDoList/internal/domain/entity"
	"errors"
	"github.com/google/uuid"
)

type LocalMap struct {
	mp map[uuid.UUID]*entity.Task
}

func NewLocalMap() *LocalMap {
	return &LocalMap{
		mp: make(map[uuid.UUID]*entity.Task),
	}
}

func (lm *LocalMap) CreateTask(task entity.Task) error {
	_, ok := lm.mp[task.ID()]
	if ok {
		return errors.New("Already exists")
	}
	lm.mp[task.ID()] = &task
	return nil
}

func (lm *LocalMap) CompleteTask(id uuid.UUID) error {
	val, ok := lm.mp[id]
	if !ok {
		return errors.New("Not exists")
	}
	val.Complete()
	return nil
}

func (lm *LocalMap) GetTask(id uuid.UUID) (entity.Task, error) {
	val, ok := lm.mp[id]
	if !ok {
		return entity.Task{}, errors.New("Not exists")
	}
	return *val, nil
}

func (lm *LocalMap) GetTasks() ([]entity.Task, error) {
	ret := make([]entity.Task, 0, len(lm.mp))
	for _, v := range lm.mp {
		ret = append(ret, *v)
	}
	return ret, nil
}

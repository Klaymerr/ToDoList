package local_memory

import (
	"ToDoList/internal/domain/entity"
	"errors"
)

type LocalMap struct {
	mp map[int]*entity.Task
}

func NewLocalMap() *LocalMap {
	return &LocalMap{
		mp: make(map[int]*entity.Task),
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

func (lm *LocalMap) CompleteTask(id int) error {
	val, ok := lm.mp[id]
	if !ok {
		return errors.New("Not exists")
	}
	val.Complete()
	return nil
}

func (lm *LocalMap) GetTask(id int) (entity.Task, error) {
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

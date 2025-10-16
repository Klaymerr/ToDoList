package service

import (
	"ToDoList/internal/domain/entity"
	"ToDoList/internal/domain/valueobject"
	"ToDoList/internal/ports/secondary"
)

type TaskService struct {
	TaskRepo secondary.TaskRepository
}

func NewTaskService(taskRepo secondary.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: taskRepo,
	}
}

func (ts *TaskService) CreateTask(id int, text valueobject.Text) error {
	task := entity.NewTask(id, text)
	err := ts.TaskRepo.CreateTask(*task)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) CompleteTask(id int) error {
	err := ts.TaskRepo.CompleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) GetTask(id int) (entity.Task, error) {
	Task, err := ts.TaskRepo.GetTask(id)
	if err != nil {
		return entity.Task{}, err
	}
	return Task, nil
}

func (ts *TaskService) GetTasks() ([]entity.Task, error) {
	Tasks, err := ts.TaskRepo.GetTasks()
	if err != nil {
		return []entity.Task{}, err
	}
	return Tasks, nil
}

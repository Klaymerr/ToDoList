package service

import (
	"ToDoList/internal/domain/entity"
	"ToDoList/internal/domain/valueobject"
	"ToDoList/internal/ports/secondary"
	"github.com/google/uuid"
)

type TaskService struct {
	TaskRepo secondary.TaskRepository
}

func NewTaskService(taskRepo secondary.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: taskRepo,
	}
}

func (ts *TaskService) CreateTask(text valueobject.Text) error {
	task, err := entity.NewTask(uuid.New(), text, false)
	if err != nil {
		return err
	}
	if err := ts.TaskRepo.CreateTask(*task); err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) CompleteTask(id uuid.UUID) error {
	err := ts.TaskRepo.CompleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TaskService) GetTask(id uuid.UUID) (entity.Task, error) {
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

package app

import (
	"ToDoList/internal/adapters/secondary/postgresql"
	"ToDoList/internal/domain/service"
	"ToDoList/internal/envvar"
	"ToDoList/internal/ports/primary"
	"ToDoList/internal/ports/secondary"
)

type ServiceProvider struct {
	taskRepo secondary.TaskRepository

	taskService primary.TaskService
}

func newServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) TaskRepository() secondary.TaskRepository {
	if s.taskRepo == nil {
		repo, err := postgresql.NewPostgresqlRepository(envvar.GetPostgreEnv())
		if err != nil {
			panic(err)
		}
		s.taskRepo = repo
	}
	return s.taskRepo
}

func (s *ServiceProvider) TaskService() primary.TaskService {
	if s.taskService == nil {
		ts := service.NewTaskService(s.TaskRepository())
		s.taskService = ts
	}
	return s.taskService
}

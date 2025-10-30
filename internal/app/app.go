package app

import (
	"ToDoList/internal/adapters/primary/api"
	"ToDoList/internal/adapters/secondary/postgresql"
	"ToDoList/internal/domain/service"
	"ToDoList/internal/envvar"
	"fmt"
)

type App struct {
	endpoint   *api.Router
	service    *service.TaskService
	repository *postgresql.PostgresqlRepository
}

func NewApp() *App {
	app := App{}

	err := app.initDeps()
	if err != nil {
		panic(err)
	}
	return &app
}

func (a *App) Run() error {
	err := a.endpoint.Start()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initRepository,
		a.initService,
		a.initEndpoint,
	}

	for _, f := range inits {
		err := f()
		if err != nil {
			return fmt.Errorf("init deps: %w", err)
		}
	}
	return nil
}

func (a *App) initRepository() error {
	var err error
	a.repository, err = postgresql.NewPostgresqlRepository(envvar.GetPostgreEnv())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initService() error {
	a.service = service.NewTaskService(a.repository)
	return nil
}

func (a *App) initEndpoint() error {
	a.endpoint = api.NewRouter(a.service)
	return nil
}

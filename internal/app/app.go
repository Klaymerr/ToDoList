package app

import (
	"ToDoList/internal/adapters/primary/api"
	"fmt"
)

type App struct {
	endpoint *api.Router
	sp       ServiceProvider
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
		a.initServiceProvider,
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

func (a *App) initServiceProvider() error {
	a.sp = *newServiceProvider()
	return nil
}

func (a *App) initEndpoint() error {
	a.endpoint = api.NewRouter(a.sp.TaskService())
	return nil
}

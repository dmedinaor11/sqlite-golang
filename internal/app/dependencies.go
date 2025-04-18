package app

import (
	exitCommand "github.com/dmedinaor11/sqlite-golang/internal/domain/exit-command"
	selectCommand "github.com/dmedinaor11/sqlite-golang/internal/domain/select-command"
)

func setUpDependencies(a *app) error {

	// services

	// handlers
	exitHandler := exitCommand.NewHandler()
	selectHandler := selectCommand.NewHandler()

	a.registryCommand(exitCommand.Command, exitHandler.Handle)
	a.registryCommand(selectCommand.Command, selectHandler.Handle)
	return nil
}

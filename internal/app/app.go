package app

import (
	"context"
	"fmt"
	"strings"
)

type (
	commandHandler func(c context.Context, statement string) error

	app struct {
		commandsRegistry map[string]commandHandler
	}
)

func NewApp() *app {
	return &app{
		commandsRegistry: make(map[string]commandHandler),
	}
}

func (a *app) registryCommand(command string, handler commandHandler) {
	a.commandsRegistry[command] = handler
}

func (a *app) processInput(input string) {
	err := a.validateInput(input)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	ctx := context.Background()
	command := a.getCommandFromInput(input)
	handler, ok := a.commandsRegistry[command]

	if !ok {
		fmt.Println("error: command not supported")
		return
	}
	handler(ctx, input)
}

func (a *app) getCommandFromInput(input string) string {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

func (a *app) validateInput(input string) error {
	if input == "" {
		return fmt.Errorf("input cannot be empty")
	}

	return nil
}

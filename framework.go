// framework.go
package main

import (
	"fmt"
	"os"
)

// Command defines the interface for command execution.
type Command interface {
	Run()
}

// CommandHandler is responsible for handling and executing commands.
type CommandHandler struct {
	commands map[string]Command
}

// NewCommandHandler creates a new CommandHandler instance.
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		commands: make(map[string]Command),
	}
}

// RegisterCommand registers a command with the CommandHandler.
func (ch *CommandHandler) RegisterCommand(name string, command Command) {
	ch.commands[name] = command
}

// ExecuteCommand executes the command with the given name.
func (ch *CommandHandler) ExecuteCommand(name string) {
	if command, ok := ch.commands[name]; ok {
		command.Run()
	} else {
		fmt.Printf("Error: Command '%s' not found\n", name)
		os.Exit(1)
	}
}

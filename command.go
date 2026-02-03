package main

import (
	"errors"
	"fmt"

	"github.com/Averagejoestudent/Gator/internal/config"
	"github.com/Averagejoestudent/Gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("the state struct is nil")
	}
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("the unkown command name %s", cmd.Name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	if name == "" {
		return
	}
	c.registeredCommands[name] = f
}

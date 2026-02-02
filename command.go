package main

import (
	"errors"
	"fmt"

	"github.com/Averagejoestudent/Gator/internal/config"
)
type state struct{
	cfg  *config.Config
}

type command struct{
	name string 
	args []string
}

type commands struct {
	variable map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error{
	if s == nil {
		return errors.New("the state struct is nil")
	}
	handler , ok := c.variable[cmd.name] 
	if !ok {
		return fmt.Errorf("the unkown command name %s",cmd.name)
	}
	return handler(s,cmd)
}

func (c *commands) register(name string, f func(*state, command) error){
	if name == "" {
		return 
	}
	c.variable[name] = f
}
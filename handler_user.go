package main

import(
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}	
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	} 
	fmt.Printf("The %s username is set", s.cfg.CurrentUser)
	return nil
}
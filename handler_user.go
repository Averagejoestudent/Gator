package main

import (
    "context"
    "errors"
    "fmt"
    "time"

    "github.com/Averagejoestudent/Gator/internal/database"
    "github.com/google/uuid"
)




func handlerLogin(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: %s <name>", cmd.Name)
    }
    name := cmd.Args[0]

    if _, err := s.db.GetUser(context.Background(), name); err != nil {
        return fmt.Errorf("couldn't find user: %w", err)
    }

    if err := s.cfg.SetUser(name); err != nil {
        return fmt.Errorf("couldn't set current user: %w", err)
    }

    fmt.Println("User switched successfully!")
    return nil
}

func handlerRegister(s *state, cmd command) error {
	
	if len(cmd.Args) == 0 {
		return errors.New("username is required")
	}
	usersname := cmd.Args[0]
	if _ , err := s.db.GetUser(context.Background(), usersname); err == nil {
		return errors.New("user already exists")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
    ID:        uuid.New(),
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
    Name:      usersname,
	})
	if err != nil {
    	return fmt.Errorf("couldn't create user: %w", err)
	}
	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User created successfully!")
	return nil
}

func handlerReset(s *state, cmd command) error {
    err := s.db.DelUsers(context.Background())
    if err != nil {
        return fmt.Errorf("Users were not deleted due to :%w",err)
    }

    fmt.Println("All users were successfully deleted")
    return nil
}

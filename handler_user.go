package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MYAzrak/mya-gator/internal/database"
	"github.com/google/uuid"
)

// handlerLogin switches the current active user in the config if the user exists in the database.
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("%s couldn't be found: %w", name, err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User switched to '%s' successfully!\n", name)
	fmt.Printf("User Data: %+v\n", user)
	return nil
}

// handlerRegister creates a new user in the database and sets them as the current active user.
func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User '%s' registered successfully!\n", name)
	fmt.Printf("User Data: %+v\n", user)
	return nil
}

// handlerReset deletes all user records from users table.
// This is useful for testing without the need for goose postgres db_url down then up
// REMOVE FROM PRODUCTION
func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.db.TruncateUsersTable(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't truncate users table: %w", err)
	}

	fmt.Printf("Users table truncated successfully!\n")
	return nil
}

// handlerUsers lists all registered users in the database and highlights the currently active user.
func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get all users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}
	return nil
}

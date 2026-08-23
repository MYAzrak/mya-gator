package main

import (
	"context"
	"fmt"

	"github.com/MYAzrak/mya-gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("%s couldn't be found: %w", s.cfg.CurrentUserName, err)
		}
		return handler(s, cmd, user)
	}
}

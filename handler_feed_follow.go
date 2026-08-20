package main

import (
	"fmt"
)

// handlerFollow
func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	fmt.Println(url)
	return nil
}

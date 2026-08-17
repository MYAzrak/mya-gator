package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MYAzrak/mya-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		nameHandlerMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args

	if len(args) < 2 {
		log.Fatal("not enough arguments were provided")
	}

	commandName := args[1]

	cmd := command{
		name: commandName,
		args: args[2:],
	}
	cmds.run(s, cmd)
}

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		log.Fatal("The login handler expects a single argument, the username")
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("The user has been set to", username)

	return nil
}

type commands struct {
	nameHandlerMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return fmt.Errorf("State does not exist")
	}

	handler, ok := c.nameHandlerMap[cmd.name]
	if !ok {
		return fmt.Errorf("Unknown command: %s", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.nameHandlerMap[name] = f
}

package main

import (
	"fmt"
	"log"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
	descriptions       map[string]string
}

func (c *commands) register(name, description string, f func(*state, command) error) {
	c.registeredCommands[name] = f
	c.descriptions[name] = description
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return f(s, cmd)
}

func commandsContextWrapper(cmds commands, f func(s *state, cmd command, cmds commands) error) func(s *state, cmd command) error {

	return func(s *state, cmd command) error {
		err := f(s, cmd, cmds)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}

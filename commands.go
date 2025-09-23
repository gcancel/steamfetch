package main

import (
	"fmt"
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

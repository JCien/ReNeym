package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(command) error
}

func (c *commands) register(name string, f func(command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("Command does not exist")
	}

	return f(cmd)
}

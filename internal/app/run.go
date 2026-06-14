package app

import (
	"fmt"
	"os"

	"github.com/shmcgrath/todo.txt-cli/config"
)

type Command func(*config.Config, []string) error

var commands = map[string]Command{
	"list": List,
	"new":  New,
	"help": Help,
}

func Run(cfg *config.Config) error {
	fmt.Println("todo file:", cfg.TodoFile)

	if len(os.Args) < 2 {
		return List(cfg, os.Args[2:])
	}

	name := os.Args[1]

	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}

	return cmd(cfg, os.Args[2:])
}

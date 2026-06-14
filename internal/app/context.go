package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Context(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Context]:", cfg.TodoFile)

	return nil
}

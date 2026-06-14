package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Archive(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Archive]:", cfg.TodoFile)

	return nil
}

package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Priority(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Priority]:", cfg.TodoFile)

	return nil
}

package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Complete(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Complete]:", cfg.TodoFile)

	return nil
}

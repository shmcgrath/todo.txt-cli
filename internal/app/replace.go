package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Replace(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Replace]:", cfg.TodoFile)

	return nil
}

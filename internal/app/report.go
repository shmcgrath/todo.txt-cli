package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Report(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Report]:", cfg.TodoFile)

	return nil
}

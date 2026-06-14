package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Project(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Project]:", cfg.TodoFile)

	return nil
}

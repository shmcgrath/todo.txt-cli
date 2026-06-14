package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Move(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Move]:", cfg.TodoFile)

	return nil
}

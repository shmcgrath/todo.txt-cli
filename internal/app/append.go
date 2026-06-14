package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Append(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Append]:", cfg.TodoFile)

	return nil
}

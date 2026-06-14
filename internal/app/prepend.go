package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Prepend(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Prepend]:", cfg.TodoFile)

	return nil
}

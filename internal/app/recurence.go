package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Recur(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Recur]:", cfg.TodoFile)

	return nil
}

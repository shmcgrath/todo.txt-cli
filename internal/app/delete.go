package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Delete(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Delete]:", cfg.TodoFile)

	return nil
}

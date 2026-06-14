package app

import (
	"fmt"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func Deduplicate(cfg *config.Config, args []string) error {
	fmt.Println("todo file [Deduplicate]:", cfg.TodoFile)

	return nil
}

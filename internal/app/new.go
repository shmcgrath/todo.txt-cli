package app

import (
	"fmt"
	// "os"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func New(cfg *config.Config, args []string) error {
	fmt.Println("todo file [run]:", cfg.TodoFile)
	fmt.Println(args)

	return nil
}

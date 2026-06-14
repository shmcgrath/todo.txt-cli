package app

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/shmcgrath/todo.txt-cli/config"
)

func List(cfg *config.Config, args []string) error {
	file, err := os.Open(cfg.TodoFile)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return nil
}

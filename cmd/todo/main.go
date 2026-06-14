package main

import (
	"log"

	"github.com/shmcgrath/todo.txt-cli/config"
	"github.com/shmcgrath/todo.txt-cli/internal/app"
)

func main() {
	path, err := config.DefaultPath()
	if err != nil {
		log.Println(err)
		return
	}
	cfg, err := config.Load(path)
	if err != nil {
		log.Println(err)
		return
	}
	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}

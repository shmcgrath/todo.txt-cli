package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/shmcgrath/todo.txt-cli/config"
)

type Command func(*config.Config, []string) error

type CommandInfo struct {
	Run         Command
	Description string
}

var commands = map[string]CommandInfo{
	"list": {
		Run:         List,
		Description: "List all tasks",
	},
	"new": {
		Run:         New,
		Description: "Create a new task",
	},
}

func Help(cfg *config.Config, args []string) error {
	fmt.Println("========== Configuration Information ==========")
	fmt.Println("config path:", cfg.CfgPath)
	fmt.Println("todo directory", cfg.TodoDir)
	fmt.Println("todo file:", cfg.TodoFile)
	fmt.Println("done file:", cfg.DoneFile)
	fmt.Println("===============================================")
	fmt.Println("Usage:")
	fmt.Println("    todo <command>")
	fmt.Println()
	fmt.Println("Commands:")

	for name, cmd := range commands {
		fmt.Printf("    %-8s %s\n", name, cmd.Description)
	}

	fmt.Printf("    %-8s %s\n", "help", "Show help")

	return nil
}

func Run(cfg *config.Config) error {

	if len(os.Args) < 2 {
		Help(cfg, nil)
		return nil
	}

	name := strings.ToLower(os.Args[1])

	if name == "help" || name == "-h" || name == "--help" {
		Help(cfg, os.Args[2:])
		return nil
	}

	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}

	return cmd.Run(cfg, os.Args[2:])
}

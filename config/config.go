package config

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	TodoDir    string
	TodoFile   string
	DoneFile   string
	ReportFile string
}

func DefaultPath() (string, error) {
	// if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
	// 	return filepath.Join(xdg, "todo", "todo-cfg.json")
	// }
	//
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	return ""
	// }
	//
	// return filepath.Join(home, ".config", "todo", "todo-cfg.json")
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return filepath.Join(dir, "config", "todo-cfg.json"), nil
}

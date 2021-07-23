package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	return cfg, nil
}

func main() {
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}

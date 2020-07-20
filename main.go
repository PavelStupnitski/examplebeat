package main

import (
	"os"

	"github.com/PavelStupnitsky/examplebeat/cmd"

	_ "github.com/PavelStupnitsky/examplebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

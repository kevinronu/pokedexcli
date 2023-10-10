package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error {
	fmt.Println("Thanks for using PokedexCLI")
	os.Exit(0)
	return nil
}

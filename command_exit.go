package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("Thanks for using PokedexCLI")
	os.Exit(0)
	return nil
}

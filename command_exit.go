package main

import (
	"fmt"
	"os"
	"time"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing... Goodbye!")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
	os.Exit(0)
	return nil
}

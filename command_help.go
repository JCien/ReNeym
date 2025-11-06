package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println(
		"Reneym (Re-Name) is a bulk rename tool that uses an Excel spreadsheet to rename files.",
	)
	fmt.Println("Commands:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

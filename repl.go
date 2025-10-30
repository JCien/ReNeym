package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JCien/ReNeym/internal/reneymapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	reneymapiClient reneymapi.Client
	activeSheet     string
	sheetData       map[string][][]string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the app",
			callback:    commandExit,
		},
		"read": {
			name:        "read",
			description: "Reads an Excel doc to get the workign sheet(s)",
			callback:    commandRead,
		},
		"data": {
			name:        "data",
			description: "Lists the data in the selected/active sheet",
			callback:    commandData,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("RN > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		usrInput := input[0]

		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		command, exists := getCommands()[usrInput]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unkown command: %v\n", usrInput)
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	clean := strings.Fields(lower)
	return clean
}

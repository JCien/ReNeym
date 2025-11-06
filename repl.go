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
	activeDoc       string
	sheetData       map[string][][]string
	sheets          []string
	fileNames       []string
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
		"scan": {
			name:        "scan",
			description: "Scans an Excel doc to get the working Sheet(s)",
			callback:    commandScan,
		},
		"rename": {
			name:        "rename",
			description: "Renames files according to selected Sheet",
			callback:    commandRename,
		},
		"sheets": {
			name:        "sheets",
			description: "List the Sheets in the scanned Excel doc",
			callback:    commandSheets,
		},
	}
}

func startRepl(cfg *config) {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome to the Reneym tool.")
	fmt.Println("Type 'help' for a list of available commands.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		activeDoc := ""
		if cfg.activeDoc != "" {
			activeDoc = fmt.Sprintf("(%s)", cfg.activeDoc)
		}
		fmt.Printf("RN %s> ", activeDoc)
		scanner.Scan()

		ogInput := scanner.Text()
		input := cleanInput(ogInput)
		if len(input) == 0 {
			continue
		}

		// This checks if the command is scan and saves the spreadsheet name
		if input[0] == "scan" && len(input) == 2 {
			doc := getActiveDoc(ogInput)[1]
			_, err := os.Stat(doc)
			if err == nil {
				cfg.activeDoc = doc
			}
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

func getActiveDoc(text string) []string {
	return strings.Fields(text)
}

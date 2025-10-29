package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func commandRead(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: read <Spreadsheet file>")
	}

	spreadsheet := args[0]

	file, err := excelize.OpenFile(spreadsheet)
	if err != nil {
		return err
	}
	defer file.Close()

	sheets := []string{}
	for _, name := range file.GetSheetList() {
		visible, _ := file.GetSheetVisible(name)
		if visible {
			sheets = append(sheets, name)
		}
	}

	if len(sheets) > 1 {
		fmt.Println("Multiple Sheets Detected")
		fmt.Println("Pick the sheet with the Taxonomy:")
		for i, sheet := range sheets {
			fmt.Printf("%v) %s\n", i, sheet)
		}

		fmt.Print("Enter number: ")
		def := 0
		input := getInput(len(sheets), &def)
		cfg.activeSheet = sheets[input]

		fmt.Printf("Selected Sheet: %s\n", cfg.activeSheet)
		// Add spreadsheet data to SheetData
		cfg.sheetData[cfg.activeSheet], err = file.GetCols(cfg.activeSheet)
		if err != nil {
			return err
		}

	} else {
		cfg.activeSheet = sheets[0]
		fmt.Printf("Using sheet: %s\n", cfg.activeSheet)
		cfg.sheetData[cfg.activeSheet], err = file.GetCols(cfg.activeSheet)
		if err != nil {
			return err
		}
	}

	return nil
}

func getInput(numSheets int, defaultValue *int) int {
	scanner := bufio.NewReader(os.Stdin)

	for {
		input, _ := scanner.ReadString('\n')
		selection := strings.TrimSpace(input)

		// Quit option
		if strings.EqualFold(selection, "q") || strings.EqualFold(selection, "quit") {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		// Default value
		if selection == "" && defaultValue != nil {
			return *defaultValue
		}

		// Parse integer
		index, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Printf("%v is not a valid number.\n", selection)
			fmt.Print("Enter a valid entry: ")
			continue
		}

		// Range check
		if index < 0 || index >= numSheets {
			fmt.Printf("%v is out of range (0-%v).\n", selection, numSheets-1)
			fmt.Print("Enter a valid entry: ")
			continue
		}

		return index
	}
}

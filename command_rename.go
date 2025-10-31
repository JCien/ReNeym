package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func commandRename(cfg *config, args ...string) error {
	// Pick Sheet to work on if there are multiple sheets
	if len(cfg.sheets) > 1 {
		fmt.Println("Multiple Sheets Detected")
		fmt.Println("Pick the sheet to work on:")
		for i, sheet := range cfg.sheets {
			fmt.Printf("%v) %s\n", i, sheet)
		}

		fmt.Print("Enter number: ")
		def := 0
		input := getInput(len(cfg.sheets), &def)
		cfg.activeSheet = cfg.sheets[input]

		fmt.Printf("Selected Sheet: %s\n", cfg.activeSheet)

	} else {
		cfg.activeSheet = cfg.sheets[0]
		fmt.Printf("Using sheet: %s\n", cfg.activeSheet)
	}

	// Start
	cols := cfg.sheetData[cfg.activeSheet]
	fnColumn, _, err := getColumnIndex(cfg.activeSheet, cols)
	if err != nil {
		return err
	}

	cfg.fileNames = []string{}
	for c, col := range cols {
		for _, rowCell := range col {
			if rowCell == "" {
				continue
			}
			if c == fnColumn {
				cellValue := strings.FieldsFunc(rowCell, func(r rune) bool {
					return r == '\n'
				})
				cfg.fileNames = append(cfg.fileNames, cellValue...)
			}
		}
	}
	fmt.Printf("There are %v files to rename.\n", len(cfg.fileNames)-1)
	cfg.activeSheet = ""
	return nil
}

func getColumnIndex(activeSheet string, sheet [][]string) (int, int, error) {
	fmt.Printf("Paste the column name in %s for Taxonomy: ", activeSheet)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		columnNameCap := scanner.Text()
		columnName := strings.ToLower(columnNameCap)
		if columnName == "" {
			fmt.Println("Column name can't be blank.")
			fmt.Print("Please enter valid column name: ")
			continue
		}

		if !strings.Contains(columnName, "taxonomy") {
			fmt.Printf("Are you sure %v is the file naming column?: ", columnNameCap)
			scanner.Scan()
			youSure := strings.ToLower(scanner.Text())
			if youSure == "y" || youSure == "yes" {
				c, r := calculateIndex(sheet, columnName)
				return c, r, nil
			}
			fmt.Print("Please enter correct column name: ")
			continue
		}
		c, r := calculateIndex(sheet, columnName)
		return c, r, nil

	}
}

func calculateIndex(sheet [][]string, columnName string) (int, int) {
	for c, col := range sheet {
		for r, rowCell := range col {
			if rowCell == "" {
				continue
			}
			if strings.Contains(strings.ToLower(rowCell), columnName) {
				return c, r
			}
		}
	}
	return 0, 0
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

package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func commandScan(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: scan <Spreadsheet file>")
	}

	spreadsheet := args[0]

	// Opens the spreadsheet and closes it when done
	file, err := excelize.OpenFile(spreadsheet)
	if err != nil {
		return err
	}
	defer file.Close()

	// This gets all the active sheets in the spreadsheet and saves them to a slice
	cfg.sheets = []string{}
	for _, name := range file.GetSheetList() {
		visible, _ := file.GetSheetVisible(name)
		if visible {
			cfg.sheets = append(cfg.sheets, name)
		}
	}

	// Populating sheetData with scanned Spreadsheet
	for _, sheet := range cfg.sheets {
		cfg.sheetData[sheet], err = file.GetCols(sheet)
		if err != nil {
			return err
		}
	}

	// This will clear the screen and put prompt at the top
	fmt.Print("\033[H\033[2J")

	// This saves the active spreadsheet for display
	// cfg.activeDoc = spreadsheet

	// if len(cfg.sheets) > 1 {
	//	fmt.Println("Multiple Sheets Detected")
	//	fmt.Println("Pick the sheet to work on:")
	//	for i, sheet := range cfg.sheets {
	//		fmt.Printf("%v) %s\n", i, sheet)
	//	}

	//	fmt.Print("Enter number: ")
	//	def := 0
	//	input := getInput(len(cfg.sheets), &def)
	//	cfg.activeSheet = cfg.sheets[input]

	//	fmt.Printf("Selected Sheet: %s\n", cfg.activeSheet)
	//	// Add spreadsheet data to SheetData
	//	cfg.sheetData[cfg.activeSheet], err = file.GetCols(cfg.activeSheet)
	//	if err != nil {
	//		return err
	//	}

	//} else {
	//	cfg.activeSheet = cfg.sheets[0]
	//	fmt.Printf("Using sheet: %s\n", cfg.activeSheet)
	//	cfg.sheetData[cfg.activeSheet], err = file.GetCols(cfg.activeSheet)
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

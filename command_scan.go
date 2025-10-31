package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func commandScan(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: read <Spreadsheet file>")
	}

	spreadsheet := args[0]

	file, err := excelize.OpenFile(spreadsheet)
	if err != nil {
		return err
	}
	defer file.Close()

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

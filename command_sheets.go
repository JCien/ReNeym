package main

import "fmt"

func commandSheets(cfg *config, args ...string) error {
	if len(cfg.sheetData) < 1 {
		fmt.Println("No sheets detected, please scan a Spreadsheet")
		return nil
	}

	for _, sheet := range cfg.sheets {
		fmt.Printf("%s\n", sheet)
	}
	return nil
}

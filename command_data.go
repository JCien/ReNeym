package main

import (
	"fmt"
	"strings"
)

func commandData(cfg *config, args ...string) error {
	cols := cfg.sheetData[cfg.activeSheet]
	for _, col := range cols {
		for _, rowCell := range col {
			if rowCell == "" {
				continue
			}
			if strings.Contains(strings.ToLower(rowCell), "taxonomy") {
				fmt.Println(rowCell)
			}
		}
	}
	return nil
}

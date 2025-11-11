package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FileReneym(cfg *config) int {
	path := "./Test"
	var numFilesMoved int
	sheetFolder := filepath.Join(path, cfg.activeSheet)
	// Reading files in the Test directory
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return numFilesMoved
	}

	// If directory doesn't exist create a folder for the active tab to put in the renamed files
	if _, err := os.Stat(sheetFolder); os.IsNotExist(err) {
		err := os.Mkdir(sheetFolder, 0750)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return numFilesMoved
		}
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		} else {
			// Checking only MP4 and MOV files
			if strings.Contains(entry.Name(), ".mp4") || strings.Contains(entry.Name(), ".mov") {
				// Extracting Ad-ID, Destination key words and extension from files
				ext := filepath.Ext(entry.Name())
				fileName := strings.TrimSuffix(entry.Name(), ext)
				fileNameParts := strings.Split(fileName, "__")
				adID := fileNameParts[0]
				destSlice := strings.Split(fileNameParts[1], "_")
				if len(destSlice) == 2 {
					destSlice = append(destSlice, "ss")
				}
				SSorTP := ""
				switch destSlice[len(destSlice)-1] {
				case "3rdParty":
					SSorTP = "3rd"
				case "DCM":
					SSorTP = "dcm"
				case "Apps":
					SSorTP = "3p"
				case "ProRes":
					SSorTP = "prores"
				case "SiteServed":
					SSorTP = "ss"
				case "ss":
					SSorTP = "ss"
				default:
					SSorTP = ""
				}
				for _, newFileName := range cfg.fileNames {
					newCheck := strings.Split(newFileName, " x ")
					for i, key := range destSlice[:len(destSlice)-1] {
						tmpName := strings.ToLower(newCheck[len(newCheck)-1])
						if strings.Contains(tmpName, strings.ToLower(key)) && strings.Contains(tmpName, strings.ToLower(adID)) && strings.Contains(tmpName, SSorTP) {
							if i == len(destSlice)-2 {
								// fmt.Printf("%v -> %v\n", fileName, newFileName)
								err := moveFile(filepath.Join(path, fileName+ext), newFileName+ext, filepath.Join(sheetFolder, destSlice[1]))
								if err != nil {
									fmt.Println(err)
									continue
								}
								numFilesMoved += 1
							} else {
								continue
							}
						} else {
							break
						}
					}
				}
			}
		}
	}

	return numFilesMoved
}

func moveFile(src, newFileName, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0755); err != nil {
			return fmt.Errorf("Error creating directory: %v", err)
		}
	}

	// Create destination path
	dst := filepath.Join(dest, newFileName)

	// Move file
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("Error renaming file: %v", err)
	}
	return nil
}

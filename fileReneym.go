package main

import (
	"fmt"
	"os"
	"strings"
)

func FileReneym(newFilenames []string) {
	entries, err := os.ReadDir("./Test")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		} else {
			if strings.Contains(entry.Name(), ".mp4") || strings.Contains(entry.Name(), ".mov") {
				fmt.Println(entry.Name())
			}
		}
	}
}

package main

import (
	"time"

	"github.com/JesusC-XR/ReNeym/internal/reneymapi"
)

func main() {
	rnClient := reneymapi.NewClient(time.Minute * 5)

	cfg := &config{
		sheetData:       map[string][][]string{},
		reneymapiClient: rnClient,
	}

	startRepl(cfg)
}

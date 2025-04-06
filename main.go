package main

import (
	"control-center-git/tui"
	"log"
)

func main() {
	if err := tui.SetupUI().Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}

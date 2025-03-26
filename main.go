package main

import (
	"control-center-git/tui"
	"log"
)

func main() {
	app := tui.SetupUI()
	if err := app.Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}

package util

import (
	"log"
	"os/exec"
)

func FocusEditor() {
	// time.Sleep(50 * time.Millisecond) // can add this if renderer window doesn't open up quickly
	cmd := exec.Command("../focus-on-editor.sh")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)



func main() {
if os.Getenv("HELP_DEBUG") != "" {
		f, err := tea.LogToFile("debug.log", "help")
		if err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		}
		defer f.Close() // nolint:errcheck
	}
    if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
        fmt.Printf("Could not start program: %v", err)
        os.Exit(1)
    }}

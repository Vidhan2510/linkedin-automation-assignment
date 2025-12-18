// LinkedIn Messenger Automation Tool
// ==================================
// This is the entry point of the application.
// It initializes and runs the CLI interface.
//
// Author: Vidhan (Internship Assignment)
// Purpose: Demonstrate problem-solving and quick learning ability
//
// How to run:
//   go run main.go demo     - Run demo with sample contacts
//   go run main.go start    - Start automation
//   go run main.go stats    - View statistics
//   go run main.go templates - View message templates

package main

import (
	"linkedin-automation/cmd"
)

func main() {
	// Execute the CLI - all logic is in cmd package
	cmd.Execute()
}

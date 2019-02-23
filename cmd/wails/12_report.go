package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/wailsapp/wails/cmd"
)

func init() {

	commandDescription := `Generates a report to help resolve issues.`

	versionCommand := app.Command("report", "Generates Report").
		LongDescription(commandDescription)

	versionCommand.Action(func() error {
		modules := os.Getenv("GO111MODULE")
		if modules == "" {
			modules = "(Not set)"
		}
		fmt.Println("Please copy and paste this report when creating issues")
		fmt.Println("------------------------------------------------------")
		fmt.Printf("Wails Version: %s\n", cmd.Version)
		fmt.Println("Go:")
		fmt.Printf("  GOOS    : %s\n", runtime.GOOS)
		fmt.Printf("  GOARCH  : %s\n", runtime.GOARCH)
		fmt.Printf("  Version : %s\n", runtime.Version())
		fmt.Println("Environment:")
		fmt.Printf("  Shell      : %s\n", getShell())
		fmt.Printf("  GO111MODULE: %s\n", modules)
		fmt.Println("------------------------------------------------------")
		return nil
	})
}

func getShell() string {
	switch runtime.GOOS {
	case "windows":
		return os.Getenv("COMSPEC")
	case "linux":
		return os.Getenv("SHELL")
	case "darwin":
		return os.Getenv("SHELL")
	default:
		return ""
	}
}

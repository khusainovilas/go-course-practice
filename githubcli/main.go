package main

import (
	"fmt"
	"githubcli/api"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no arguments provided")
		fmt.Println("Run 'githubcli help' for usage information")
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "help" {
		printHelp()
		return
	}

	url := strings.TrimSpace(command)
	if !strings.HasPrefix(url, "https://github.com/") {
		fmt.Println("Error: invalid URL format")
		os.Exit(1)
	}

	parts := strings.Split(strings.TrimPrefix(url, "https://github.com/"), "/")
	if len(parts) < 2 {
		fmt.Println("Error: invalid URL format")
		os.Exit(1)
	}
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", parts[0], parts[1])

	flags := []string{}
	for i := 2; i < len(os.Args); i++ {
		flags = append(flags, strings.TrimPrefix(os.Args[i], "-"))
	}

	api.GetInfo(apiURL, flags)
}

func printHelp() {
	fmt.Println("GitHub CLI - Usage")
	fmt.Println("Commands:")
	fmt.Println("  githubcli <repo>   Show information about the repository")
	fmt.Println("  githubcli help     Show this help message")

	fmt.Println("Flags for <repo> command:")
	fmt.Println("  -name      Display repository name")
	fmt.Println("  -desc      Display repository description")
	fmt.Println("  -star      Display number of stars")
	fmt.Println("  -forks     Display number of forks")
	fmt.Println("  -date      Display creation date")

	fmt.Println("Example:")
	fmt.Println("  githubcli https://github.com/mimi-net/miminet -name -star")
}

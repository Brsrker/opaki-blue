package main

import (
	"flag"
	"fmt"
	"main/internal/host"
	"os"
	"path/filepath"
)

func main() {

	installFlag := flag.Bool("i", false, "Install as environment variable")
	flag.Parse()

	// Check if the -i flag was received
	if *installFlag {
		// Install the executable as an environment variable
		err := InstallAsEnvVariable()
		if err != nil {
			fmt.Println("Failed to install as environment variable:", err)
			return
		}

		fmt.Println("Installed as environment variable successfully.")
		return
	} else {
		// Check if the required number of arguments were provided
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run script.go <ipAddress> <domainName>")
			return
		}

		// IP address and domain name to add or update
		ipAddress := os.Args[1]
		domainName := os.Args[2]

		host.UpdateHostFile(ipAddress, domainName)
	}

}

func InstallAsEnvVariable() error {
	// Get the absolute path to the executable
	executablePath, err := filepath.Abs(os.Args[0])
	fmt.Printf("executablePath %s\n", executablePath)
	if err != nil {
		return err
	}

	// Set the environment variable
	err = addToPath(executablePath)
	if err != nil {
		return err
	}
	return nil
}

// AddToPath adds the given value to the PATH environment variable.
func addToPath(value string) error {
	// Get the current value of the PATH environment variable
	currentPath := os.Getenv("PATH")

	// Add the new value to the PATH environment variable
	newPath := fmt.Sprintf("%s%c%s", currentPath, filepath.ListSeparator, value)

	// Set the PATH environment variable to the new value
	return os.Setenv("PATH", newPath)
}

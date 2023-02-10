package host

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func UpdateHostFile(ipAddress string, domainName string) {

	// Path to the hosts file
	hostsPath := getHostsFilePath()

	// Check if the line already exists in the file
	lineToUpdate := ipAddress + " " + domainName
	lineExists := false
	file, err := os.Open(hostsPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == lineToUpdate {
			lineExists = true
			break
		}
	}

	// Add or update the line in the file
	file, err = os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	if !lineExists {
		if _, err = file.WriteString(lineToUpdate + "\n"); err != nil {
			fmt.Println(err)
			return
		}
	} else {
		oldLine := ipAddress + " " + domainName
		newLine := ipAddress + " " + domainName
		if err := updateLine(file, oldLine, newLine); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func getHostsFilePath() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32\drivers\etc\hosts`
	} else {
		return "/etc/hosts"
	}
}

func updateLine(file *os.File, oldLine, newLine string) error {
	tempFile, err := os.Create(file.Name() + ".tmp")
	if err != nil {
		return err
	}
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(tempFile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == oldLine {
			line = newLine
		}
		if _, err = tempFile.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := os.Remove(file.Name()); err != nil {
		return err
	}

	if err := os.Rename(tempFile.Name(), file.Name()); err != nil {
		return err
	}

	return nil
}

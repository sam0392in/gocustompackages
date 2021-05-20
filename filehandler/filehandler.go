package filehandler

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Readlines(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func Writelines(file string, content []string, writeMode string) {

	// Declare file handler and error variable
	var f *os.File
	var err error

	if writeMode == "overwrite" {
		f, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	} else if writeMode == "append" {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		fmt.Println("\nWrite Mode not given correctly. Valide modes are: overwrite, append")
		os.Exit(0)
	}

	if err != nil {
		log.Fatalf("failed writing in file: %s", err)
	}

	datawriter := bufio.NewWriter(f)

	for _, data := range content {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	f.Close()
}

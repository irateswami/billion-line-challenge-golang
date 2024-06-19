package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	exitFail = 1
)

func run() error {
	path := os.Getenv("BILLION_LINE_FILE")
	if path == "" {
		return errors.New("no path found")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	var count int64

	// Read the file line by line
	for scanner.Scan() {
		count++
	}

	// Check for errors during the scanning process
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Printf("lines: %d\n", count)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

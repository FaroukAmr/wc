package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	cFlag := flag.Bool("c", false, "sets the byte count")
	lFlag := flag.Bool("l", false, "sets the line count")

	flag.Parse()
	var args []string = flag.Args()
	fileName, err := handleGetFileName(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	absPath, err := handleGetAbsPath(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
	}

	file, err := handleOpenFile(absPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if *cFlag {
		size, err := handleCountBytes(file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v %v\n", size, fileName)
		return
	}
	if *lFlag {
		numberOfLines := handleGetNumberOfLines(file)
		fmt.Printf("%v %v\n", numberOfLines, fileName)
		return
	}
}

func handleGetFileName(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no arguments provided")
	}
	return args[0], nil
}

func handleGetAbsPath(fileName string) (string, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return "", fmt.Errorf("error getting absolute path: %w", err)
	}

	return absPath, nil
}

func handleOpenFile(absPath string) (*os.File, error) {
	file, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	return file, nil
}

func handleCountBytes(file *os.File) (int, error) {
	contents, err := io.ReadAll(file)

	if err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	return len(contents), nil
}

func handleGetNumberOfLines(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var numberOfLines int = 0
	for scanner.Scan() {
		numberOfLines++
	}
	return numberOfLines
}

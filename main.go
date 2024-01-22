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
	wFlag := flag.Bool("w", false, "sets the word count")

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
			return
		}
		fmt.Printf("%v %v\n", size, fileName)
		return
	}
	if *lFlag {
		numberOfLines, err := handleGetNumberOfLines(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v %v\n", numberOfLines, fileName)
		return
	}
	if *wFlag {
		numberOfWords, err := handleGetNumberOfWords(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v %v\n", numberOfWords, fileName)
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

func handleGetNumberOfLines(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)

	if scanner.Err() != nil {
		return 0, fmt.Errorf("error scanning file: %w", scanner.Err())
	}
	var numberOfLines int = 0
	for scanner.Scan() {
		numberOfLines++
	}
	return numberOfLines, nil
}

func handleGetNumberOfWords(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordsCount := 0
	for scanner.Scan() {
		wordsCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}
	return wordsCount, nil
}

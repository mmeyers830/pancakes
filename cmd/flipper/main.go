package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mmeyers830/pancakes/internal/pancakes"
)

func main() {
	// if we are worried about space, we probably don't want to read the whole file at once, instead
	// processing it line by line, but for the sake of this problem, we are ignoring that fact.
	pancakeStacks, err := parseFile()
	if err != nil {
		// if we can't parse the file, we need to just exit
		panic(err)
	}
	flipper := pancakes.Flipper{
		HappyChar: '+',
		PlainChar: '-',
	}

	for i, stack := range pancakeStacks {
		count, err := flipper.CountFlips(stack)
		if err != nil {
			fmt.Printf("Case #%d: failed with error: %v\n", i+1, err)
		} else {
			fmt.Printf("Case #%d: %d\n", i+1, count)
		}
	}
}

// parseFile will parse our whole input file and return our slice of test cases
func parseFile() ([][]rune, error) {
	fileName := flag.String("file", "test.txt", "the test input file to be parsed")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s for reading: %w", *fileName, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	pancakeStacks, err := readLines(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input: %w", err)
	}

	return pancakeStacks, nil
}

// readLines reads all the test cases from the reader
func readLines(reader *bufio.Reader) ([][]rune, error) {
	numTestsStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read line: %w", err)
	}
	numTestsStr = strings.TrimRight(numTestsStr, "\r\n")
	numTests, err := strconv.Atoi(numTestsStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert string %s to int: %w", numTestsStr, err)
	}

	pancakeStacks := make([][]rune, numTests)
	for i := 0; i < numTests; i++ {
		stack, err := reader.ReadString('\n')
		// if it's EOF, we could still have the last line to return
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("failed to read line: %w", err)
		}
		stack = strings.TrimRight(stack, "\r\n")
		pancakeStacks[i] = []rune(stack)
	}

	return pancakeStacks, nil
}

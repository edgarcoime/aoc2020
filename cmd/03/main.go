package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FILEPATH   = "cmd/03/input/input.txt"
	RISE       = 1
	RUN        = 3
	TREE_CHAR  = '#'
	BLANK_CHAR = '.'
)

func main() {
	treeMap, err := processInput(FILEPATH)
	if err != nil {
		fmt.Println("Error processing input:", err)
		os.Exit(1)
	}

	p1 := NewPart1(treeMap)

	p1.Run()
}

func processInput(filepath string) ([][]rune, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var treeMap [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		var chars []rune
		for _, char := range line {
			chars = append(chars, char)
		}
		treeMap = append(treeMap, chars)
	}

	return treeMap, nil
}

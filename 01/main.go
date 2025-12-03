package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	filepath = "input.txt"
	target   = 2020
)

// https://adventofcode.com/2020/day/1
func main() {
	// Process input.txt into an array of numbers
	nums, err := processInput(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	part1(nums, target)
	part2(nums, target)
}

func part2(orig []int, target int) {
	nums := make([]int, len(orig))
	copy(nums, orig)
	sort.Ints(nums)
}

func part1(orig []int, target int) {
	nums := make([]int, len(orig))
	copy(nums, orig)
	sort.Ints(nums)

	// Two sum solution only for sorted array
	l, r := 0, len(nums)-1
	for l < r {
		res := nums[l] + nums[r]

		// Early break here so that else (==) doesn't evaluate
		if res == target {
			break
		}

		if res < target {
			l++
		} else {
			r--
		}
	}

	answer := nums[l] * nums[r]
	fmt.Println("Part 1: ", answer)
}

func processInput(filepath string) ([]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		err := fmt.Errorf("Error opening file: +%v", err)
		return nil, err
	}

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// "  124 "
		line := scanner.Text()
		line = strings.TrimSpace(line)
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Could not convert string to int:", err)
		}
		nums = append(nums, num)
	}

	return nums, nil
}

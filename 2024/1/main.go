package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Day 1: Historian Hysteria
// Given an input file containing 2 lists of values, pair up the numbers and measure how far apart they are.
// Pair up the smallest number in the left list with the smallest number in the right list,
// then the second-smallest left number with the second-smallest right number, and so on.

const FILE_PATH = "input.txt"

func main() {
	left, right := returnStructuredInput()

	fmt.Println(findTotalDistance(left, right))
}

func findTotalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	l := len(left)

	var totalDiff int

	for i := range l {
		fmt.Printf("Diff set (left: %d - right: %d = total %d)\n", left[i], right[i], int(math.Abs(float64(left[i]-right[i]))))
		totalDiff = totalDiff + int(math.Abs(float64(left[i]-right[i])))
	}

	return totalDiff
}

func returnStructuredInput() ([]int, []int) {
	content, err := os.ReadFile(FILE_PATH)
	if err != nil {
		panic(err)
	}

	numbers := strings.Fields(string(content))

	var left []int
	var right []int

	for i, num := range numbers {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}

	return left, right
}

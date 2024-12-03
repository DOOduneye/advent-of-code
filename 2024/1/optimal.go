//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	distance := calculateDistance(left, right)
	fmt.Println("Total distance:", distance)
}

func parseInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) != 2 {
			continue // Skip invalid lines
		}

		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("invalid number in input")
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculateDistance(left, right []int) int {
	if len(left) != len(right) {
		panic("lists must have equal length")
	}

	sort.Ints(left)
	sort.Ints(right)

	var totalDiff int
	for i := range left {
		totalDiff += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDiff
}

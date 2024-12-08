package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(filepath string) [][]int {
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	var result [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strNums := strings.Fields(line)
		var intRow []int
		for _, strNum := range strNums {
			num, err := strconv.Atoi(strNum)
			check(err)
			intRow = append(intRow, num)
		}

		result = append(result, intRow)
	}

	check(scanner.Err())

	return result
}

func isSorted(arr []int) bool {
	asc, desc := true, true
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			asc = false
		}
		if arr[i-1] <= arr[i] {
			desc = false
		}
	}
	//XOR to return true if its either increasing or decreasing
	return asc != desc
}
func isSortedAndSafe(report []int) bool {
	if !isSorted(report) {
		return false
	}

	for j := 0; j < len(report)-1; j++ {
		diff := math.Abs(float64(report[j] - report[j+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func part1() {
	reports := readInput("day2.input")
	safeCount := 0

	for _, report := range reports {
		if isSorted(report) {
			isSafe := true
			for j := 0; j < len(report)-1; j++ {
				diff := math.Abs(float64(report[j] - report[j+1]))
				if diff < 1 || diff > 3 {
					isSafe = false
					break
				}
			}
			if isSafe {
				safeCount++
			}
		}
	}

	fmt.Println("Safe count:", safeCount)
}

func part2() {
	reports := readInput("day2.input")
	safeCount := 0

	for _, report := range reports {
		if isSortedAndSafe(report) {
			safeCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			//copy elements before and after current element
			modifiedReport := append([]int{}, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)

			if isSortedAndSafe(modifiedReport) {
				safeCount++
				break
			}
		}
	}

	fmt.Println("Safe count:", safeCount)
}

func main() {
	part1()
	part2()
}

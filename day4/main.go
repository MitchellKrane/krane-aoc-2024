package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file string) ([][]rune, error) {
	input, err := os.ReadFile(file)
	check(err)
	lines := regexp.MustCompile("\n").Split(string(input), -1)
	var result [][]rune
	for _, line := range lines {

		result = append(result, []rune(line))
	}

	return result, nil
}

func part1() {
	//Part 1 - Brute Force
	filePath := "day4.input"
	array, err := readInput(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	directions := map[string][2]int{
		"NW": {-1, -1},
		"NE": {-1, 1},
		"SW": {1, -1},
		"SE": {1, 1},
		"N":  {-1, 0},
		"S":  {1, 0},
		"E":  {0, 1},
		"W":  {0, -1},
	}

	count := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if string(array[i][j]) == "X" {
				for _, offset := range directions {
					word := "X"
					for step := 1; step <= 3; step++ {
						newRow := i + offset[0]*step
						newCol := j + offset[1]*step
						if newRow >= 0 && newRow < len(array) && newCol >= 0 && newCol < len(array[i]) {
							neighbor := array[newRow][newCol]
							word += string(neighbor)
							if word == "XMAS" {
								count++
							}
							//debug
							//fmt.Printf("Element at (%d, %d) = '%c', Direction: %s, Step: %d, Neighbor: '%c' Count: %d \n", i, j, array[i][j], direction, step, neighbor, count)
						} else {

							break
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	//Part 2 - Brute Force
	filePath := "day4.input"
	array, err := readInput(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	count := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if string(array[i][j]) == "A" {
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(array) && j+1 < len(array[i]) {
					left := []string{string(array[i-1][j-1]), string(array[i+1][j+1])}
					right := []string{string(array[i+1][j-1]), string(array[i-1][j+1])}
					leftstring := strings.Join(left, "")
					rightstring := strings.Join(right, "")
					if (leftstring == "MS" || leftstring == "SM") && (rightstring == "MS" || rightstring == "SM") {
						count++
					}

				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}

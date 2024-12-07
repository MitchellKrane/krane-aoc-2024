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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(file string) ([]int, []int) {
	input, err := os.Open(file)
	check(err)
	defer input.Close()

	var first []int
	var second []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "   ")

		if val, err := strconv.Atoi(columns[0]); err == nil {
			first = append(first, val)
		} else {
			fmt.Println("Error converting first column:", err)
		}

		if val, err := strconv.Atoi(columns[1]); err == nil {
			second = append(second, val)
		} else {
			fmt.Println("Error converting second column:", err)
		}
	}

	sort.Ints(first)
	sort.Ints(second)

	return first, second
}

func part1() {
	left, right := readInput("day1.input")
	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))

	}
	fmt.Println(distance)
}

func part2() {
	left, right := readInput("day1.input")
	//covert to map for lookup speed and count tracking
	rightMap := make(map[int]int)
	for _, val := range right {
		rightMap[val]++
	}
	total := 0

	for _, val := range left {
		if count, exists := rightMap[val]; exists {
			total += val * count
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()

}

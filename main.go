package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func part1(file string) {
	input, err := os.ReadFile(file)
	check(err)
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(input), -1)
	var total int
	for _, match := range matches {
		if len(match) > 2 {
			a, _ := strconv.ParseUint(match[1], 10, 32)
			b, _ := strconv.ParseUint(match[2], 10, 32)
			product := int(a) * int(b)
			total += product
		}
	}
	fmt.Println("Part 1 Total:", total)
}

func part2(file string) {
	input, err := os.ReadFile(file)
	check(err)

	dontDoPattern := `(?s)don't\(\).*?do\(\)`
	mulPattern := `mul\((\d+),(\d+)\)`

	reDontDo := regexp.MustCompile(dontDoPattern)
	cleanedInput := reDontDo.ReplaceAllString(string(input), "")

	reMul := regexp.MustCompile(mulPattern)

	var total int
	matches := reMul.FindAllStringSubmatch(cleanedInput, -1)

	for _, match := range matches {
		a, _ := strconv.ParseUint(match[1], 10, 32)
		b, _ := strconv.ParseUint(match[2], 10, 32)
		product := int(a) * int(b)
		total += product
	}

	fmt.Println("Part 2 Total:", total)
}

func main() {
	filePath := "day3.input"
	part1(filePath)
	part2(filePath)
}

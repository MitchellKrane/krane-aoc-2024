package main

import (
	"fmt"
	"os"
	"regexp"
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
func isAntenna(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
}

type Coordinate struct {
	X, Y int
}

func part1() {
	input, err := readInput("day8.input")
	check(err)

	antennasByFreq := make(map[rune][]Coordinate)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if isAntenna(input[y][x]) {
				freq := input[y][x]
				antennasByFreq[freq] = append(antennasByFreq[freq], Coordinate{X: x, Y: y})
			}
		}
	}

	rows := len(input)
	cols := len(input[0])
	antinodes := make(map[Coordinate]struct{})

	for _, coords := range antennasByFreq {
		if len(coords) < 2 {
			continue
		}

		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				A := coords[i]
				B := coords[j]
				a1 := Coordinate{X: 2*A.X - B.X, Y: 2*A.Y - B.Y}
				a2 := Coordinate{X: 2*B.X - A.X, Y: 2*B.Y - A.Y}

				if a1.X >= 0 && a1.X < cols && a1.Y >= 0 && a1.Y < rows {
					antinodes[a1] = struct{}{}
				}
				if a2.X >= 0 && a2.X < cols && a2.Y >= 0 && a2.Y < rows {
					antinodes[a2] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func main() {
	part1()
}

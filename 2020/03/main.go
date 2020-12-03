package main

import (
	"bufio"
	"log"
	"os"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}

func part1(lines []string, originalX int, originalY int, slopeX int, slopeY int, treesFound int) int {
	if slopeY >= len(lines) {
		return treesFound
	}
	ground := string(lines[slopeY][slopeX])
	if ground == "#" {
		return part1(lines, originalX, originalY, slopeX+originalX, slopeY+originalY, treesFound+1)
	}
	return part1(lines, originalX, originalY, slopeX+originalX, slopeY+originalY, treesFound)
}

func main() {
	lines := readFile("input.txt")
	var tobogganMap []string
	mapSize := 100
	for _, l := range lines {
		ground := ""
		for i := 0; i < mapSize; i++ {
			ground += l
		}
		tobogganMap = append(tobogganMap, ground)
	}

	result1 := part1(tobogganMap, 3, 1, 0, 0, 0)
	println("part1 result :", result1)

	result2A := part1(tobogganMap, 1, 1, 0, 0, 0)
	println("part2 result for slope x1,y1 :", result2A)
	result2B := part1(tobogganMap, 3, 1, 0, 0, 0)
	println("part2 result for slope x3,y1 :", result2B)
	result2C := part1(tobogganMap, 5, 1, 0, 0, 0)
	println("part2 result for slope x5,y1 :", result2C)
	result2D := part1(tobogganMap, 7, 1, 0, 0, 0)
	println("part2 result for slope x7,y1 :", result2D)
	result2E := part1(tobogganMap, 1, 2, 0, 0, 0)
	println("part2 result for slope x1,y2 :", result2E)
	println("part2 result :", result2A*result2B*result2C*result2D*result2E)

}

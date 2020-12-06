package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func part1(lines []string) int {
	var total int
	s := map[string]bool{}

	for _, line := range lines {
		if line == "" {
			total += len(s)
			s = make(map[string]bool)
		}
		for _, letter := range line {
			s[string(letter)] = true
		}
	}
	return total
}

func part2(lines []string) int {
	var total int
	s := map[string]int{}
	var members int

	for _, line := range lines {
		if line == "" {
			if members == 1 {
				total += len(s)
			} else {
				for _, value := range s {
					if value > 1 {
						total++
					}
				}
			}
			members = 0
			s = make(map[string]int)
		} else {
			for _, letter := range line {
				s[string(letter)]++
			}
			members++
		}
	}
	return total
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	lines := readFile("input.txt")

	result1 := part1(lines)
	println("part1 result :", result1)

	result2 := part2(lines)
	println("part2 result :", result2)

}

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

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func part1(lines []string) int {
	for i, valueA := range lines {
		a := convertToInt(valueA)

		for j, valueB := range lines {
			b := convertToInt(valueB)

			if i != j {
				if a+b == 2020 {
					println("### part 1 found", a, b)
					return a * b
				}
			}
		}
	}
	return -1
}

func part2(lines []string) int {
	for i, valueA := range lines {
		a := convertToInt(valueA)

		for j, valueB := range lines {
			b := convertToInt(valueB)

			for k, valueC := range lines {
				c := convertToInt(valueC)

				if (i - j - k) != i*-1 {
					if a+b+c == 2020 {
						println("### part 2 found", a, b, c)
						return a * b * c
					}
				}
			}
		}
	}
	return -1
}

func main() {
	lines := readFile("input.txt")
	result1 := part1(lines)
	println("part1 result :", result1)
	result2 := part2(lines)
	println("part2 result :", result2)

}

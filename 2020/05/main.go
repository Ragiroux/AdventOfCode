package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

func part1(lines []string) []int {

	var seatId []int

	for _, seat := range lines {
		var lowerRows = 0
		var upperRows = 128
		var lowerColumn = 0
		var upperColumn = 8
		var rows = seat[0 : len(seat)-3]
		var column = seat[len(seat)-3:]

		for _, letter := range rows {
			var l = string(letter)
			var newBound = (lowerRows + upperRows) / 2
			if l == "B" {
				lowerRows = newBound
			} else if l == "F" {
				upperRows = newBound
			}
		}

		for _, letter := range column {
			var l = string(letter)
			var newBound = (lowerColumn + upperColumn) / 2
			if l == "R" {
				lowerColumn = newBound
			} else if l == "L" {
				upperColumn = newBound
			}
		}
		seatId = append(seatId, ((upperRows-1)*8)+(upperColumn-1))
	}
	return seatId
}

func part2(ids []int) int {
	var XOR = 0
	var n = len(ids)

	for i := 0; i < n; i++ {
		XOR ^= (ids[i] + 1)
	}
	return XOR
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
	sort.Ints(result1[:])
	for _, i := range result1 {
		println(i)
	}
	println("part1 result :", result1)

	result2 := part2(result1)
	println("part2 result :", result2)

}

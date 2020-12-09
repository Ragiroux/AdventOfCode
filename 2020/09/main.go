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

func part1(xmas []int64, preamble int64) int64 {
	var k int64
	for k = preamble; k < int64(len(xmas)); k++ {
		if validate(xmas[k], xmas[k-preamble:k]) == false {
			return xmas[k]
		}
	}
	return -1
}

func validate(n int64, series []int64) bool {
	for i := int64(0); i < int64(len(series)); i++ {
		for j := int64(0); j < int64(len(series)); j++ {
			if (i != j) && int64(series[i]+series[j]) == n {
				return true
			}
		}
	}
	return false
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func part2(xmas []int64, preamble int64) int64 {
	var k int64
	for k = preamble; k < int64(len(xmas)); k++ {
		if validate(xmas[k], xmas[k-preamble:k]) == false {
			for i := int64(0); i < k; i++ {
				isValid, segment := findWeakness(xmas[k], xmas[i:k])
				if isValid {
					//finalSlice := xmas[i : i+int64(indexFound)]
					sort.Slice(segment, func(i, j int) bool { return segment[i] < segment[j] })
					return segment[0] + segment[len(segment)-1]
				}
			}
		}
	}
	return -1
}

func findWeakness(n int64, slice []int64) (bool, []int64) {
	var sum int64
	for i, add := range slice {
		sum += add
		if sum > n {
			return false, nil
		}
		if sum == n {
			return true, slice[0 : i+1]
		}
	}
	return false, nil
}

func convertToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	lines := readFile("input.txt")
	var xmas []int64

	for _, line := range lines {
		xmas = append(xmas, convertToInt(line))
	}

	result1 := part1(xmas, 25)
	println("part1 result :", result1)

	result2 := part2(xmas, 25)
	println("part2 result :", result2)

}

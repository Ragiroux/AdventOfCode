package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Jolt struct {
	rating int
	delta  int
}

func (j Jolt) new(rating int) Jolt {
	j.rating = rating
	j.delta = 0
	return j
}

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

func part1(jolts []Jolt, effectiveRating int, delta1 int, delta3 int) (int, int) {

	if len(jolts) == 0 {
		return delta1, delta3+1
	}

	resetDelta(jolts)

	for i := 0; i < len(jolts); i++ {
		jolts[i].delta = jolts[i].rating - effectiveRating
	}

	sort.SliceStable(jolts, func(i, j int) bool {
		return jolts[i].delta < jolts[j].delta
	})

	lowest := jolts[0]
	if lowest.delta == 1 {
		return part1(jolts[1:], effectiveRating+lowest.delta, delta1+1, delta3)
	} else if lowest.delta == 3 {
		return part1(jolts[1:], effectiveRating+lowest.delta, delta1, delta3+1)
	} else {
		return part1(jolts[1:], effectiveRating+lowest.delta, delta1, delta3)
	}
}

func resetDelta(jolts []Jolt) {
	for i := 0; i < len(jolts); i++ {
		jolts[i].delta = 0
	}
}

func part2(jolts []Jolt) int {
	return -1
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
	var jolts []Jolt

	for _, line := range lines {
		var j Jolt
		j = j.new(convertToInt(line))
		jolts = append(jolts, j)
	}

	delta1, delta3 := part1(jolts, 0, 0, 0)
	println(delta1, delta3)
	println("part1 result :", delta1*delta3)

	result2 := part2(jolts)
	println("part2 result :", result2)

}

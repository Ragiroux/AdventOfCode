package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	valid := 0
	for _, input := range lines {
		policy := strings.Split(input, ":")
		policyPassword := strings.Split(policy[0], " ")
		passwordRules := strings.Split(policyPassword[0], "-")
		passwordChar := policyPassword[1]
		lowerBound := convertToInt(passwordRules[0])
		uppperBound := convertToInt(passwordRules[1])
		password := strings.Trim(policy[1], " ")
		occurences := strings.Count(password, passwordChar)

		if occurences >= lowerBound && occurences <= uppperBound {
			valid++
		}
	}
	return valid
}

func part2(lines []string) int {
	valid := 0
	for _, input := range lines {
		policy := strings.Split(input, ":")
		policyPassword := strings.Split(policy[0], " ")
		passwordRules := strings.Split(policyPassword[0], "-")
		passwordChar := policyPassword[1]
		indexA := convertToInt(passwordRules[0])
		indexB := convertToInt(passwordRules[1])
		password := strings.Trim(policy[1], " ")
		slice := password[(indexA - 1):(indexB)]
		isValid := xor(string(slice[0]) == passwordChar, string(slice[len(slice)-1]) == passwordChar)

		if isValid {
			valid++
		}
	}
	return valid
}

func xor(x bool, y bool) bool {
	return x != y
}

func main() {
	lines := readFile("input.txt")
	result1 := part1(lines)
	println("part1 result :", result1)
	result2 := part2(lines)
	println("part2 result :", result2)

}

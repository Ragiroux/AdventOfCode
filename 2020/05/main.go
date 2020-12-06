package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
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

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p Passport) reset() Passport {
	p.byr = ""
	p.iyr = ""
	p.eyr = ""
	p.hgt = ""
	p.hcl = ""
	p.ecl = ""
	p.pid = ""
	p.cid = ""
	return p
}

func (p Passport) isValid() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p Passport) addInfo(info []string) Passport {
	switch info[0] {
	case "byr":
		value := convertToInt(info[1])
		if value >= 1920 && value <= 2002 {
			p.byr = info[1]
		}
		break
	case "iyr":
		value := convertToInt(info[1])
		if value >= 2010 && value <= 2020 {
			p.iyr = info[1]
		}
		break
	case "eyr":
		value := convertToInt(info[1])
		if value >= 2020 && value <= 2030 {
			p.eyr = info[1]
		}
		break
	case "hgt":
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		if re.MatchString(info[1]) {
			submatchall := re.FindAllString(info[1], -1)
			var height = convertToInt(submatchall[0])
			if strings.Contains(info[1], "cm") {
				if height >= 150 && height <= 193 {
					p.hgt = info[1]
				}
			} else if strings.Contains(info[1], "in") {
				if height >= 59 && height <= 76 {
					p.hgt = info[1]
				}
			}
		}
		break
	case "hcl":
		re := regexp.MustCompile(`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`)
		if re.MatchString(info[1]) {
			p.hcl = info[1]
		}
		break
	case "ecl":
		re := regexp.MustCompile(`^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$`)
		if re.MatchString(info[1]) {
			p.ecl = info[1]
		}
		break
	case "pid":
		re := regexp.MustCompile(`^[0-9]{9}$`)
		if re.MatchString(info[1]) {
			p.pid = info[1]
		}
		break
	case "cid":
		p.cid = info[1]
		break
	}
	return p
}

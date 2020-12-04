package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
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

func importPassport(lines []string) []Passport {
	var passports []Passport
	var p Passport
	for _, line := range lines {
		l := strings.Split(line, " ")

		for _, info := range l {
			passportInfo := strings.Split(info, ":")
			p = p.addInfo(passportInfo)
		}

		if line == "" {
			passports = append(passports, p)
			p = p.reset()
		}
	}
	return passports
}

func part1(passports []Passport) int {
	var count = 0
	for _, p := range passports {
		if p.isValid() {
			count++
		}
	}
	return count
}

func part2(lines []string) int {
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
	passports := importPassport(lines)

	result1 := part1(passports)
	println("part1 result :", result1)

	result2 := part2(lines)
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

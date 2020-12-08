package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	name    string
	size    int
	isShiny bool
	next    *bag
}

type linkedList struct {
	length int
	head   *bag
	tail   *bag
}

func (l linkedList) Len() int {
	return l.length
}

func (l linkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.name)
		l.head = l.head.next
	}
	fmt.Println()
}

func (l linkedList) findShiny() int {
	var total int
	for l.head != nil {
		if l.head.isShiny {
			total++
		}
		l.head = l.head.next
	}
	return total
}

func (l *linkedList) PushBack(n *bag) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
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

func part1(bags []linkedList) int {
	var total int
	for _, bag := range bags {
		total += bag.findShiny()
	}

	return total
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

func isBagShiny(s string) bool {
	return strings.Contains(s, "shiny")
}

func formatBag(n string) string {
	n = strings.Trim(n, " ")
	if strings.Contains(n, ".") {
		n = strings.TrimRight(n, ".")
	}
	return n
}

func main() {
	lines := readFile("input.txt")

	var bags []linkedList

	for _, line := range lines {
		list := linkedList{}
		bagsTxt := strings.Split(line, "contain")
		isShiny := isBagShiny(bagsTxt[0])
		mainBag := &bag{name: bagsTxt[0], size: -1, isShiny: isShiny}
		list.PushBack(mainBag)
		smallBagsTxt := strings.Split(bagsTxt[1], ",")
		extractNumber := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		for _, b := range smallBagsTxt {
			size := extractNumber.FindAllString(b, -1)
			if size != nil {
				t := strings.Replace(b, size[0], "", -1)
				smallBag := &bag{name: formatBag(t), size: convertToInt(size[0]), isShiny: isBagShiny(b)}
				list.PushBack(smallBag)
			}
		}
		bags = append(bags, list)
	}

	result1 := part1(bags)
	println("part1 result :", result1)

	result2 := part2(lines)
	println("part2 result :", result2)

}

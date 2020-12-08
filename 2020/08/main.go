package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	name        string
	accumulator int
	operation   int
	exec        int
}

func (i instruction) new(name string, accumulator int, operation int, exec int) instruction {
	i.name = name
	i.accumulator = accumulator
	i.operation = operation
	i.exec = exec
	return i
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

func execution(instructions []instruction, executionStep int, acc int) int {

	i := instructions[executionStep]
	instructions[executionStep].exec++

	if instructions[executionStep].exec > 1 {
		return acc
	}

	switch i.name {
	case "nop":
		return execution(instructions, executionStep+1, acc)
	case "acc":
		return execution(instructions, executionStep+1, acc+i.operation)
	case "jmp":
		return execution(instructions, executionStep+i.operation, acc)
	}
	return acc
}

func execution2(instructions []instruction, executionStep int, acc int, indexChanged int) int {

	if executionStep >= len(instructions) {
		return acc
	}

	i := instructions[executionStep]
	instructions[executionStep].exec++

	if instructions[executionStep].exec > 1 {
		return -1
	}

	if indexChanged != -1 && indexChanged == executionStep && (instructions[indexChanged].name == "nop" || instructions[indexChanged].name == "jmp") {
		if i.name == "nop" {
			//do jmp instead
			return execution2(instructions, executionStep+i.operation, acc, -1)
		} else {
			//do nop instead
			return execution2(instructions, executionStep+1, acc, -1)
		}
	}

	switch i.name {
	case "nop":
		return execution2(instructions, executionStep+1, acc, indexChanged)
	case "acc":
		return execution2(instructions, executionStep+1, acc+i.operation, indexChanged)
	case "jmp":
		return execution2(instructions, executionStep+i.operation, acc, indexChanged)
	}
	return acc
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

	var instructions []instruction

	for _, line := range lines {
		var intMove int
		var instr instruction

		operations := strings.Split(line, " ")

		extractNumber := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		move := extractNumber.FindAllString(operations[1], -1)[0]
		intMove = convertToInt(move)
		instr = instr.new(operations[0], 0, intMove, 0)

		instructions = append(instructions, instr)
	}

	result1 := execution(instructions, 0, 0)
	println("part1 result :", result1)

	for i := 0; i < len(instructions); i++ {
		result2 := execution2(instructions, 0, 0, i)
		if result2 > -1 {
			println("part2 result :", result2)
			break
		}
		//reset program flag
		for j := 0; j < len(instructions); j++ {
			instructions[j].exec = 0
		}
	}
}

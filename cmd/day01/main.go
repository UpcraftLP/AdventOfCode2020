package main

import (
	"errors"
	"github.com/upcraftlp/AdventOfCode2020/internal/aoc"
	"github.com/upcraftlp/go-utils/pkg/std"
	"log"
	"strconv"
)

func main() {
	lines := aoc.GetInput(1)
	aoc.Part1()
	a, b, _ := findSumTo(lines, 2020)
	log.Printf("found: %v + %v; multiplied: %v\n", a, b, a*b)

	aoc.Part2()
	for _, currentLine := range lines {
		num, _ := strconv.Atoi(currentLine)
		toFind := 2020 - num
		a, b, err := findSumTo(lines, toFind)
		if err == nil {
			log.Printf("found: %v + %v + %v; multiplied: %v\n", num, a, b, num*a*b)
			break
		}
	}

	aoc.End()
}

func findSumTo(lines []string, targetSum int) (int, int, error) {
	for _, currentLine := range lines {
		num, _ := strconv.Atoi(currentLine)
		toFind := targetSum - num
		if std.Contains(lines, strconv.Itoa(toFind)) {
			return num, toFind, nil
		}
	}
	return 0, 0, errors.New("no number found")
}

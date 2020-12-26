package main

import (
	"github.com/upcraftlp/AdventOfCode2020/internal/aoc"
	"log"
	"strings"
)

func main() {
	lines := aoc.GetInputKeepEmpty(6)
	aoc.Part1()
	counter := 0
	current := make(map[string]struct{})
	for _, line := range lines {
		line = strings.TrimSpace(strings.ReplaceAll(line, "\r", ""))
		if line == "" {
			counter += len(current)
			current = make(map[string]struct{})
		} else {
			for _, s := range strings.Split(line, "") {
				if s != "" {
					current[s] = struct{}{}
				}
			}
		}
	}
	log.Printf("Amount of 'yes' answers: %d\n", counter)

	aoc.Part2()
	counter = 0
	sum := 0
	test := make(map[string]bool)
	for _, line := range lines {
		line = strings.TrimSpace(strings.ReplaceAll(line, "\r", ""))
		if line == "" {
			counter = 0
			for _, b := range test {
				if b {
					sum++
				}
			}
			test = make(map[string]bool)
			counter = 0
		} else {
			if counter == 0 {
				for _, s := range strings.Split(line, "") {
					if s != "" {
						test[s] = true
					}
				}
			} else {
				for k, v := range test {
					test[k] = v && strings.Contains(line, k)
				}
			}
			counter++
		}
	}

	counter = 0
	log.Printf("Amount of 'ALL yes' answers: %d\n", sum)

	aoc.End()
}

package main

import (
	"github.com/upcraftlp/AdventOfCode2020/internal/aoc"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := aoc.GetInput(8)
	aoc.Part1()
	play(lines, -1)

	aoc.Part2()
	for j := 0; j < len(lines); j++ {
		if !play(lines, j) {
			break
		}
	}
	aoc.End()
}

func play(lines []string, flip int) (looped bool) {
	visited := make(map[int]struct{})
	global := 0
	i := 0
	for ; i < len(lines); i++ {
		if _, ok := visited[i]; ok {
			log.Printf("looped after %v instructions, accumulator value: %v, current instruction:%v\n", len(visited), global, i)
			return true
		}
		split := strings.Split(lines[i], " ")
		amount, _ := strconv.Atoi(split[1])
		if i == flip {
			switch split[0] {
			case "jmp":
				split[0] = "nop"
				break
			case "nop":
				split[0] = "jmp"
				break
			}
		}
		switch split[0] {
		case "acc":
			global += amount
			break
		case "jmp":
			i += amount - 1
			break
		case "nop":
			break
		}
		visited[i] = struct{}{}
	}
	log.Printf("terminated after %v instructions, accumulator value: %v, current instruction:%v\n", len(visited), global, i)
	return false
}

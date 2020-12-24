package main

import (
	util2 "github.com/upcraftlp/AdventOfCode2020/internal/util"
	"log"
)

func main() {
	lines := util2.GetInput(03)
	util2.Part1()
	trees := findTrees(lines, 3, 1)
	log.Printf("Trees: %v\n", trees)

	util2.Part2()

	a := findTrees(lines, 1, 1)
	b := trees
	c := findTrees(lines, 5, 1)
	d := findTrees(lines, 7, 1)
	e := findTrees(lines, 1, 2)

	log.Printf("Trees:\n\tA: %v\n\tB: %v\n\tC: %v\n\tD: %v\n\tE: %v\n", a, b, c, d, e)
	log.Println()
	log.Printf("Total: %v, Mul: %v", a+b+c+d+e, a*b*c*d*e)

	util2.End()
}

func findTrees(lines []string, xInc, yInc int) int {
	x := 0
	y := 0
	trees := 0
	for {
		x += xInc
		y += yInc
		if y >= len(lines) {
			break
		}
		if getCharAt(lines, x, y) == "#" {
			trees++
		}
	}
	return trees
}

func getCharAt(lines []string, x, y int) string {
	xMod := x % len(lines[y])
	return lines[y][xMod : xMod+1]
}

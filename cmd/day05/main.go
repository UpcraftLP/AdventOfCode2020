package main

import (
	util2 "github.com/upcraftlp/AdventOfCode2020/internal/util"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := util2.GetInput(05)

	var seats = make(map[uint]Seat)
	for _, line := range lines {
		rData := strings.ReplaceAll(strings.ReplaceAll(line[:7], "F", "0"), "B", "1")
		cData := strings.ReplaceAll(strings.ReplaceAll(line[7:10], "L", "0"), "R", "1")
		row, _ := strconv.ParseInt(rData, 2, 8)
		col, _ := strconv.ParseInt(cData, 2, 4)
		id := uint(row)*8 + uint(col)
		seats[id] = Seat{
			ID:  id,
			Row: uint(row),
			Col: uint(col),
		}
	}
	util2.Part1()
	seat := Seat{}
	for _, s := range seats {
		if s.ID > seat.ID {
			seat = s
		}
	}
	log.Printf("Highest seat ID: %v (r: %v, c: %v)\n", seat.ID, seat.Row, seat.Col)
	util2.Part2()
	var mySeat uint
	for id := uint(8); id < 127*8; id++ {
		if _, ok := seats[id]; !ok { // seat does not exist
			if _, ok := seats[id-1]; ok {
				if _, ok := seats[id+1]; ok {
					mySeat = id
					break
				}
			}
		}
	}
	log.Printf("Missing Seat: %v\n", mySeat)
	util2.End()
}

type Seat struct {
	ID  uint
	Row uint
	Col uint
}

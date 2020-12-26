package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetInput(day int) []string {
	return getInputInternal(day, false)
}

func getInputInternal(day int, keepEmpty bool) []string {
	if day < 1 || day > 24 {
		log.Fatalln("day out of range")
	}
	data, err := ioutil.ReadFile(fmt.Sprintf("input/day%02d.txt", day))
	if err != nil {
		log.Fatalln("unable to read input", err)
	}
	lines := strings.Split(string(data), "\n")
	if keepEmpty {
		return lines
	}
	var ret []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			ret = append(ret, trimmed)
		}
	}
	return ret
}

func GetInputKeepEmpty(day int) []string {
	return getInputInternal(day, true)
}

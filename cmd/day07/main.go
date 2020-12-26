package main

import (
	"github.com/upcraftlp/AdventOfCode2020/internal/aoc"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	r := regexp.MustCompile("(?P<container>.*) bags contain (?P<contents>.*).")
	r2 := regexp.MustCompile("(?P<amount>\\d+) (?P<color>.*) bags?")
	lines := aoc.GetInput(7)

	mapping := make(map[string]map[string]int)
	for _, line := range lines {
		parsed := parse(r, line)
		container := parsed["container"]
		contents := parsed["contents"]
		if _, ok := mapping[container]; !ok {
			mapping[container] = make(map[string]int)
		}
		if contents != "no other bags" {
			current := mapping[container]
			for _, s := range strings.Split(contents, ", ") {
				parsedContents := parse(r2, s)
				amount, _ := strconv.Atoi(parsedContents["amount"])
				current[parsedContents["color"]] = amount
			}
		}
	}
	aoc.Part1()
	goldenBois := make(map[string]struct{})
	for k, v := range mapping { // step 1: find all colors that can directly contain shiny gold bags
		if _, ok := v["shiny gold"]; ok {
			goldenBois[k] = struct{}{}
		}
	}

	count := len(goldenBois)
	for { // step 2: recursively find all parent elements
		for color := range goldenBois {
			for k, v := range mapping {
				if _, ok := v[color]; ok {
					goldenBois[k] = struct{}{}
				}
			}
		}
		if count == len(goldenBois) {
			break
		}
		count = len(goldenBois)
	}
	log.Printf("Colors that can contain shiny gold: %v\n", count)
	aoc.Part2()
	log.Printf("Shiny gold bags contain a total of %v bags.\n", lookup(mapping, "shiny gold") - 1)
	aoc.End()
}

func lookup(mapping map[string]map[string]int, color string) int {
	found := mapping[color]
	count := 1
	for k, v := range found {
		count += v * lookup(mapping, k)
	}
	return count
}

func parse(regExp *regexp.Regexp, input string) map[string]string {
	match := regExp.FindStringSubmatch(input)
	ret := make(map[string]string)

	for i, name := range regExp.SubexpNames() {
		if i > 0 && i < len(match) {
			ret[name] = match[i]
		}
	}

	return ret
}

package main

import (
	"github.com/upcraftlp/AdventOfCode2020/internal/aoc"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := aoc.GetInput(2)
	aoc.Part1()

	validPasswords1 := 0
	validPasswords2 := 0
	for _, current := range lines {
		split := strings.SplitN(current, "-", 2)
		min, _ := strconv.Atoi(split[0])
		split = strings.SplitN(split[1], " ", 2)
		max, _ := strconv.Atoi(split[0])
		letter := split[1][:1]
		password := split[1][2:]
		policy := PasswordPolicy{
			Min:     min,
			Max:     max,
			ToCheck: letter,
		}
		if policy.applies1(password) {
			validPasswords1++
		}
		if policy.applies2(password) {
			validPasswords2++
		}
	}
	log.Printf("Valid Passwords: %v\n", validPasswords1)

	aoc.Part2()
	log.Printf("Valid Passwords: %v\n", validPasswords2)

	aoc.End()
}

type PasswordPolicy struct {
	Min     int
	Max     int
	ToCheck string
}

func (p *PasswordPolicy) applies1(password string) bool {
	amount := strings.Count(password, p.ToCheck)
	return amount >= p.Min && amount <= p.Max
}

func (p *PasswordPolicy) applies2(password string) bool {
	substr1 := password[p.Min : p.Min+1]
	substr2 := password[p.Max : p.Max+1]
	isA := substr1 == p.ToCheck
	isB := substr2 == p.ToCheck
	return isA != isB
}

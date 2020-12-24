package main

import (
	"errors"
	util2 "github.com/upcraftlp/AdventOfCode2020/internal/util"
	"github.com/upcraftlp/go-utils/pkg/std"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := util2.GetInputKeepEmpty(04)
	var passports []Passport = nil

	var currentLines []string = nil
	for _, line := range lines {
		line = strings.TrimSpace(strings.ReplaceAll(line, "\r", ""))
		if line == "" {
			// parse new passport
			passports = append(passports, parse(currentLines))
			currentLines = nil
		} else {
			currentLines = append(currentLines, line)
		}
	}
	util2.Part1()
	complete := 0
	for _, currentPass := range passports {
		err := currentPass.checkFieldsPresent()
		if err != nil {
			if isCidError(err) {
				complete++
			}
			continue
		}
		complete++
	}
	log.Printf("Complete Passports: %v\n", complete)

	util2.Part2()
	valid := 0
	for _, currentPass := range passports {
		err := currentPass.validate()
		if err != nil {
			if isCidError(err) {
				valid++
			}
			continue
		}
		valid++
	}
	log.Printf("Valid Passports: %v\n", valid)

	util2.End()
}

func (p *Passport) validate() error {
	if p.BirthYear == "" {
		return errors.New("no birth year specified")
	} else {
		byr, err := strconv.Atoi(p.BirthYear)
		if err != nil {
			return err
		}
		if byr < 1920 || byr > 2002 {
			return errors.New("invalid birth year")
		}
	}

	if p.IssueYear == "" {
		return errors.New("no issue year specified")
	} else {
		iyr, err := strconv.Atoi(p.IssueYear)
		if err != nil {
			return err
		}
		if iyr < 2010 || iyr > 2020 {
			return errors.New("invalid issue year")
		}
	}

	if p.ExpirationYear == "" {
		return errors.New("no expiration year specified")
	} else {
		eyr, err := strconv.Atoi(p.ExpirationYear)
		if err != nil {
			return err
		}
		if eyr < 2020 || eyr > 2030 {
			return errors.New("invalid expiration year")
		}
	}

	if p.Height == "" {
		return errors.New("no height specified")
	} else {
		if strings.HasSuffix(p.Height, "cm") {
			hgt, err := strconv.Atoi(p.Height[:len(p.Height)-2])
			if err != nil {
				return err
			}
			if hgt < 150 || hgt > 193 {
				return errors.New("invalid height")
			}
		} else if strings.HasSuffix(p.Height, "in") {
			hgt, err := strconv.Atoi(p.Height[:len(p.Height)-2])
			if err != nil {
				return err
			}
			if hgt < 59 || hgt > 76 {
				return errors.New("invalid height")
			}
		} else {
			return errors.New("invalid height")
		}
	}

	if len(p.HairColor) != 7 || !strings.HasPrefix(p.HairColor, "#") {
		return errors.New("no hair color specified")
	} else {
		if _, err := strconv.ParseInt(p.HairColor[1:], 16, 32); err != nil {
			return err
		}
	}

	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	if p.EyeColor == "" {
		return errors.New("no eye color specified")
	} else if !std.Contains(validEyeColors, p.EyeColor) {
		return errors.New("invalid eye color")
	}

	if p.PassportID == "" {
		return errors.New("no passport ID specified")
	} else {
		if len(p.PassportID) != 9 {
			return errors.New("invalid passport ID")
		}
		digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		for _, s := range strings.Split(p.PassportID, "") {
			if s != "" && !std.Contains(digits, s) {
				return errors.New("invalid passport ID")
			}
		}
	}

	if p.CountryID == "" {
		return errors.New("no country ID specified")
	}
	return nil
}

func (p *Passport) checkFieldsPresent() error {
	if p.BirthYear == "" {
		return errors.New("no birth year specified")
	}
	if p.IssueYear == "" {
		return errors.New("no issue year specified")
	}
	if p.ExpirationYear == "" {
		return errors.New("no expiration year specified")
	}
	if p.Height == "" {
		return errors.New("no height specified")
	}
	if p.HairColor == "" {
		return errors.New("no hair color specified")
	}
	if p.EyeColor == "" {
		return errors.New("no eye color specified")
	}
	if p.PassportID == "" {
		return errors.New("no passport ID specified")
	}
	if p.CountryID == "" {
		return errors.New("no country ID specified")
	}
	return nil
}

func isCidError(err error) bool {
	return err.Error() == "no country ID specified"
}

func parse(data []string) Passport {
	mp := make(map[string]string)
	for _, d1 := range data {
		for _, pair := range strings.Split(d1, " ") {
			split := strings.Split(pair, ":")
			mp[split[0]] = split[1]
		}
	}
	return Passport{
		BirthYear:      mp["byr"],
		IssueYear:      mp["iyr"],
		ExpirationYear: mp["eyr"],
		Height:         mp["hgt"],
		HairColor:      mp["hcl"],
		EyeColor:       mp["ecl"],
		PassportID:     mp["pid"],
		CountryID:      mp["cid"],
	}
}

type Passport struct {
	BirthYear      string // byr
	IssueYear      string // iyr
	ExpirationYear string // eyr
	Height         string // hgt
	HairColor      string // hcl
	EyeColor       string // ecl
	PassportID     string // pid
	CountryID      string // cid; OPTIONAL
}

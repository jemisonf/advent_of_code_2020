package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type PassportEntry struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func isHexCharacter(character string) bool {
	return strings.Contains("0123456789abcdef", character)
}

func (p *PassportEntry) PushToken(token string) error {
	results := strings.Split(token, ":")
	key := results[0]
	value := results[1]

	switch key {
	case "byr":
		p.BirthYear = value
	case "iyr":
		p.IssueYear = value
	case "eyr":
		p.ExpirationYear = value
	case "hgt":
		p.Height = value
	case "hcl":
		p.HairColor = value
	case "ecl":
		p.EyeColor = value
	case "pid":
		p.PassportID = value
	case "cid":
		p.CountryID = value
	default:
		return fmt.Errorf("unknown key %s", key)
	}

	return nil
}

func (p *PassportEntry) Validate() bool {
	return p.BirthYear != "" && p.IssueYear != "" && p.ExpirationYear != "" && p.Height != "" && p.HairColor != "" && p.EyeColor != "" && p.PassportID != ""
}

func (p *PassportEntry) ValidatePart2() (bool, error) {
	if len(p.BirthYear) == 4 {
		birthYear, err := strconv.Atoi(p.BirthYear)
		if err != nil {
			return false, fmt.Errorf("error parsing birth year: %v", err)
		}
		if birthYear <= 1920 || birthYear > 2002 {
			return false, nil
		}
	} else {
		return false, nil
	}

	if len(p.IssueYear) == 4 {
		issueYear, err := strconv.Atoi(p.IssueYear)
		if err != nil {
			return false, fmt.Errorf("error parsing expiration year: %v", err)
		}
		if issueYear < 2010 || issueYear > 2020 {
			return false, nil
		}
	} else {
		return false, nil
	}

	if len(p.ExpirationYear) == 4 {
		expYear, err := strconv.Atoi(p.ExpirationYear)
		if err != nil {
			return false, fmt.Errorf("error parsing expiration year: %v", err)
		}
		if expYear < 2020 || expYear > 2030 {
			return false, nil
		}
	} else {
		return false, nil
	}

	if p.Height != "" {
		var height int
		var units string
		fmt.Sscanf(p.Height, "%d%s", &height, &units)

		switch units {
		case "cm":
			if height < 150 || height > 193 {
				return false, nil
			}
		case "in":
			if height < 59 || height > 76 {
				return false, nil
			}
		default:
			return false, nil
		}
	}

	if len(p.HairColor) > 0 {
		if p.HairColor[0] != '#' {
			return false, nil
		}
		for _, char := range p.HairColor[1:] {
			if !isHexCharacter(string(char)) {
				return false, nil
			}
		}
	} else {
		return false, nil
	}

	foundValidColor := false
	for _, validColor := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if p.EyeColor == validColor {
			foundValidColor = true
		}
	}

	if !foundValidColor {
		return false, nil
	}

	if len(p.PassportID) != 9 {
		return false, nil
	}

	if _, err := strconv.Atoi(p.PassportID); err != nil {
		return false, nil
	}

	return true, nil
}

func ParseFileIntoPassports(fileContents string) ([]PassportEntry, error) {
	entries := strings.Split(fileContents, "\n\n")
	passports := []PassportEntry{}

	for _, entry := range entries {
		passport := PassportEntry{}
		entry = strings.Replace(entry, "\n", " ", -1)
		tokens := strings.Split(entry, " ")

		for _, token := range tokens {
			if err := passport.PushToken(token); err != nil {
				return nil, err
			}
		}

		passports = append(passports, passport)
	}

	return passports, nil
}

func main() {
	arguments := args.ParseArgs()

	contents, err := io.ReadFileAsString(arguments.File)

	if err != nil {
		log.Fatalf("error opening file to read: %v", err)
	}

	passports, err := ParseFileIntoPassports(contents)

	if err != nil {
		log.Fatalf("error parsing input file: %v", err)
	}

	validPassports := 0
	for _, passport := range passports {
		valid := false

		if arguments.Part == 1 {
			valid = passport.Validate()
		} else if arguments.Part == 2 {
			valid, err = passport.ValidatePart2()

			if err != nil {
				log.Fatalf("error validating passport: %v", err)
			}
		}

		if valid {
			validPassports++
		}
	}

	fmt.Printf("%d passports are valid\n", validPassports)
}

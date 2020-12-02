package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type Policy struct {
	Min      int
	Max      int
	Char     rune
	Password string
}

func ParsePolicies(lines []string) ([]Policy, error) {
	policies := []Policy{}
	for _, line := range lines {
		newPolicy := Policy{}
		fmt.Sscanf(line, "%d-%d %c: %s\n", &newPolicy.Min, &newPolicy.Max, &newPolicy.Char, &newPolicy.Password)
		policies = append(policies, newPolicy)
	}
	return policies, nil
}

func checkPasswordPart1(policy Policy) bool {
	count := strings.Count(policy.Password, string(policy.Char))
	return count >= policy.Min && count <= policy.Max
}

func checkPasswordPart2(policy Policy) bool {
	var firstChar rune = 0
	var secondChar rune = 1

	if len(policy.Password) > policy.Min-1 {
		firstChar = rune(policy.Password[policy.Min-1])
	}
	if len(policy.Password) > policy.Max-1 {
		secondChar = rune(policy.Password[policy.Max-1])
	}

	if firstChar == policy.Char || secondChar == policy.Char {
		return firstChar != secondChar
	}

	return false
}

func main() {
	args := args.ParseArgs()

	lines, err := io.ReadFileAsLines(args.File)

	if err != nil {
		log.Fatalf("error reading file %s: %v", args.File, err)
	}

	policies, err := ParsePolicies(lines)

	if err != nil {
		log.Fatalf("error parsing policies: %v", err)
	}

	validCount := 0

	for _, policy := range policies {
		var valid bool

		if args.Part == 1 {
			valid = checkPasswordPart1(policy)
		} else if args.Part == 2 {
			valid = checkPasswordPart2(policy)
		}

		if valid {
			validCount++
		}
	}

	fmt.Printf("%d password(s) are valid\n", validCount)
}

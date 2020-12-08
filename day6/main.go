package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

const questions = "abcdefghijklmnopqrstuvwxyz"

func main() {
	arguments := args.ParseArgs()

	fileContents, err := io.ReadFileAsString(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	groups := strings.Split(fileContents, "\n\n")

	totalQuestions := 0
	for _, group := range groups {
		group = strings.Trim(group, "\n")
		foundQuestions := 0
		if arguments.Part == 1 {
			group = strings.ReplaceAll(group, " ", "")
			group = strings.ReplaceAll(group, "\n", "")
			for _, question := range questions {
				if strings.Contains(group, string(question)) {
					foundQuestions++
				}
			}
		} else {
			numLines := len(strings.Split(group, "\n"))
			for _, question := range questions {
				if strings.Count(group, string(question)) == numLines {
					foundQuestions++
				}
			}
		}
		totalQuestions += foundQuestions
	}

	fmt.Printf("found questions: %d\n", totalQuestions)
}

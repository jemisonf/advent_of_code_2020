package io

import (
	"bufio"
	"os"
)

func ReadFileAsLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	entries := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		entries = append(entries, line)
	}

	return entries, nil
}

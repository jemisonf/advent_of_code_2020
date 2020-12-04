package io

import (
	"bufio"
	"io/ioutil"
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

func ReadFileAsString(fileName string) (string, error) {
	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(contents), nil
}

package io

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
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

func ReadFileAsInts(fileName string) ([]int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	entries := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		entries = append(entries, num)
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

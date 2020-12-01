package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// read input from file
	flag.Parse()
	fileName := flag.Arg(0)

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %v", err)
	}
	// parse into a sorted list of ints

	scanner := bufio.NewScanner(file)
	entries := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to convert value %s to int", line)
		}
		entries = append(entries, val)
	}

	sort.Ints(entries)

	// for each entry
	for i, entry := range entries {
		var searchSpace []int

		if entry+entry >= 2020 {
			searchSpace = entries[:i]
		} else {
			searchSpace = entries[i+1:]
		}

		for _, compare := range searchSpace {
			target := 2020 - entry - compare

			min := 0
			max := len(entries)

			for min <= max {
				pivot := (min + max) / 2

				if entries[pivot] == target {
					product := entry * compare * target
					fmt.Printf("| %d\t| %d \t| %d \t| %d\t|\n", product, entry, compare, target)
					os.Exit(0)
				}

				if entries[pivot] < target {
					min = pivot + 1
				} else {
					max = pivot - 1
				}
			}
		}
	}

	log.Fatal("Couldn't find matching number")
}

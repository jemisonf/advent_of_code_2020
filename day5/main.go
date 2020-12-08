package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/jemisonf/advent_of_code_2020/day5/partition"
	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	largestID := 0

	if arguments.Part == 1 {
		for _, line := range lines {
			seat, err := partition.Search(line[:7], line[7:], 128, 8)

			if err != nil {
				log.Fatalf("error parsing line %s: %v", line, err)
			}

			if seat.SeatID > largestID {
				largestID = seat.SeatID
			}
		}

		fmt.Printf("largest seat ID: %d\n", largestID)
		os.Exit(0)
	}

	ids := []int{}

	for _, line := range lines {
		seat, err := partition.Search(line[:7], line[7:], 128, 8)
		if err != nil {
			log.Fatalf("error parsing line %s: %v", line, err)
		}

		ids = append(ids, seat.SeatID)
	}

	sort.Ints(ids)

	for index, id := range ids {
		if index != len(ids)-1 && ids[index+1] == id+2 {
			fmt.Printf("your ID: %d\n", id+1)
		}
	}
}

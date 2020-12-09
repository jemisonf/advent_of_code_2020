package main

import (
	"fmt"
	"log"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

func FindInvalidNums(nums []int, preambleSize int) []int {
	badNums := []int{}
	for index, num := range nums[preambleSize:] {
		found := false
		realIndex := index + preambleSize
		for i, n := range nums[realIndex-preambleSize : realIndex] {
			for j, m := range nums[realIndex-preambleSize : realIndex] {
				if i != j && n+m == num {
					found = true
				}
			}
		}
		if !found {
			badNums = append(badNums, num)
		}
	}

	return badNums
}

func FindContiguousSum(nums []int, target int) []int {
	for index, num := range nums {
		if num == target {
			continue
		}

		sum := nums[index]
		originalIndex := index
		for sum < target {
			index++
			sum += nums[index]
		}

		if sum == target {
			return nums[originalIndex : index+1]
		}
	}

	return nil
}

func main() {
	arguments := args.ParseArgs()

	if arguments.Number == 0 {
		arguments.Number = 25
	}

	nums, err := io.ReadFileAsInts(arguments.File)

	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	badNums := FindInvalidNums(nums, arguments.Number)

	fmt.Println(badNums)

	p2Nums := FindContiguousSum(nums, badNums[0])
	smallest, largest := 10000000000000, 0

	for _, num := range p2Nums {
		if num <= smallest {
			smallest = num
		}

		if num >= largest {
			largest = num
		}
	}

	fmt.Printf("encryption weakness: %d\n", smallest+largest)
}

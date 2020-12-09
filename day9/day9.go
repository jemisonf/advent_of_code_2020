package main

import (
	"fmt"
	"log"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

func FindValidNums(nums []int, preambleSize int) []int {
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
	for index := range nums {
		sum := 0
		originalIndex := index
		for sum <= target {
			sum += nums[index]
			index++
		}

		if sum == target {
			return nums[originalIndex : index+1]
		}
	}

	return nil
}

func main() {
	arguments := args.ParseArgs()

	nums, err := io.ReadFileAsInts(arguments.File)

	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	badNums := FindValidNums(nums, 25)

	fmt.Println(badNums)

	p2Nums := FindContiguousSum(nums, badNums[0])
	fmt.Println(p2Nums)
}

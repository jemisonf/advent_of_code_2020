package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sort"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

func init() {

}

func ComputeDifferences(adapters []int) map[int]int {
	currentJolt := 0
	differences := map[int]int{}

	for _, adapter := range adapters {
		differences[adapter-currentJolt]++
		currentJolt = adapter
	}
	differences[3]++

	return differences
}

func ComputeArrangements(adapters []int) [][]int {
	arrangements := map[int][][]int{}

	arrangements[0] = [][]int{}

	for _, adapter := range adapters {
		adapterArrangements := [][]int{}
		if adapter-3 < 1 {
			adapterArrangements = append(adapterArrangements, []int{adapter})
		}

		for _, arrangement := range arrangements[adapter-1] {
			newArrangement := append(arrangement, adapter)
			adapterArrangements = append(adapterArrangements, newArrangement)
		}
		for _, arrangement := range arrangements[adapter-2] {
			newArrangement := append(arrangement, adapter)
			adapterArrangements = append(adapterArrangements, newArrangement)
		}
		for _, arrangement := range arrangements[adapter-3] {
			newArrangement := append(arrangement, adapter)
			adapterArrangements = append(adapterArrangements, newArrangement)
		}

		arrangements[adapter] = adapterArrangements
	}

	maxAdapter := adapters[len(adapters)-1] + 3

	maxAdapterArrangements := [][]int{}
	for _, arrangement := range arrangements[maxAdapter-1] {
		newArrangement := append(arrangement, maxAdapter)
		maxAdapterArrangements = append(maxAdapterArrangements, newArrangement)
	}
	for _, arrangement := range arrangements[maxAdapter-2] {
		newArrangement := append(arrangement, maxAdapter)
		maxAdapterArrangements = append(maxAdapterArrangements, newArrangement)
	}
	for _, arrangement := range arrangements[maxAdapter-3] {
		newArrangement := append(arrangement, maxAdapter)
		maxAdapterArrangements = append(maxAdapterArrangements, newArrangement)
	}

	return maxAdapterArrangements
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	arguments := args.ParseArgs()

	adapters, err := io.ReadFileAsInts(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	sort.Ints(adapters)

	if arguments.Part == 1 {
		differences := ComputeDifferences(adapters)
		fmt.Printf("product: %d\n", differences[1]*differences[3])
		os.Exit(0)
	}

	fmt.Printf("number of arrangements: %d\n", len(ComputeArrangements(adapters)))

}

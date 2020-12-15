package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type Bus struct {
	ID       int
	idOffset int
}

type byOffset []Bus

func (b byOffset) Len() int           { return len(b) }
func (b byOffset) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byOffset) Less(i, j int) bool { return b[i].idOffset < b[j].idOffset }

func ParseBuses(part int, time int, busSchedule string) ([]Bus, error) {
	rawBuses := strings.Split(busSchedule, ",")
	buses := []Bus{}

	for _, rawBus := range rawBuses {
		if rawBus == "x" {
			if part == 2 {
				buses = append(buses, Bus{
					ID: -1,
				})
			}
			continue
		}

		busID, err := strconv.Atoi(rawBus)

		if err != nil {
			return nil, err
		}

		buses = append(buses, Bus{
			ID:       busID,
			idOffset: (busID - time%busID) % busID,
		})
	}

	return buses, nil
}
func GCD(a, b int) int {
	fmt.Println(a, " ", b)
	if a == b {
		return a
	}
	return GCD(b, a%b)
}

func FindDeparture(buses []Bus) int {
	first := buses[0]
	last := buses[len(buses)-1]
	t := 0
	for {
		if (t+len(buses)-1)%last.ID != 0 {
			t += first.ID
			continue
		}
		if t%10000000 == 0 {
			fmt.Println(t)
		}

		valid := true
		for index, bus := range buses[0:] {
			if bus.ID == -1 {
				continue
			}

			if (t+(index))%bus.ID != 0 {
				valid = false
				break
			}
		}

		if valid {
			break
		}
		t += first.ID
	}

	return t
}
func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	time, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatalf("error parsing first line: %v", err)
	}
	buses, err := ParseBuses(arguments.Part, time, lines[1])

	if err != nil {
		log.Fatalf("error parsing bus schedule: %v", err)
	}

	if arguments.Part == 1 {
		sort.Sort(byOffset(buses))
		fmt.Printf("bus ID times wait time: %d X (%d) = %d\n", buses[0].ID, buses[0].idOffset, (buses[0].idOffset)*buses[0].ID)
		os.Exit(0)
	}

	t := FindDeparture(buses)

	fmt.Printf("first departure: %d\n", t)

}

package main

import (
	"fmt"
	"log"
	"math"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type Action string

var (
	North   Action = "N"
	South   Action = "S"
	East    Action = "E"
	West    Action = "W"
	Left    Action = "L"
	Right   Action = "R"
	Forward Action = "F"
)

type Direction struct {
	Action Action
	Amount int
}

type WayPoint struct {
	EastWestPosition   int
	NorthSouthPosition int
}

type Ferry struct {
	EastWestPosition   int
	NorthSouthPosition int
	Direction          int
	WayPoint           WayPoint
}

func (f *Ferry) FollowDirectionsPart1(d Direction) {
	switch d.Action {
	case North:
		f.NorthSouthPosition += d.Amount
	case South:
		f.NorthSouthPosition -= d.Amount
	case East:
		f.EastWestPosition += d.Amount
	case West:
		f.EastWestPosition -= d.Amount
	case Left:
		f.Turn(-1 * d.Amount)
	case Right:
		f.Turn(d.Amount)
	case Forward:
		f.GoForward(d.Amount)
	}
}

func (w *WayPoint) Rotate(amount int) {
	switch amount {
	case 90, -270:
		w.EastWestPosition, w.NorthSouthPosition = w.NorthSouthPosition, -1*w.EastWestPosition
	case 180, -180:
		w.NorthSouthPosition *= -1
		w.EastWestPosition *= -1
	case 270, -90:
		w.EastWestPosition, w.NorthSouthPosition = -1*w.NorthSouthPosition, w.EastWestPosition
	}
}

func (f *Ferry) ForwardPart2(amount int) {
	f.EastWestPosition += f.WayPoint.EastWestPosition * amount
	f.NorthSouthPosition += f.WayPoint.NorthSouthPosition * amount
}

func (f *Ferry) FollowDirectionsPart2(d Direction) {
	switch d.Action {
	case North:
		f.WayPoint.NorthSouthPosition += d.Amount
	case South:
		f.WayPoint.NorthSouthPosition -= d.Amount
	case East:
		f.WayPoint.EastWestPosition += d.Amount
	case West:
		f.WayPoint.EastWestPosition -= d.Amount
	case Left:
		f.WayPoint.Rotate(-1 * d.Amount)
	case Right:
		f.WayPoint.Rotate(d.Amount)
	case Forward:
		f.ForwardPart2(d.Amount)
	}
}

func (f *Ferry) FollowDirections(part int, d Direction) {
	if part == 1 {
		f.FollowDirectionsPart1(d)
	} else {
		f.FollowDirectionsPart2(d)
	}
}
func (f *Ferry) Turn(amount int) {
	f.Direction += amount
	f.Direction %= 360

	// keep direction positive
	// e.g. -90 == 270
	if f.Direction < 0 {
		f.Direction += 360
	}
}

func (f *Ferry) GoForward(amount int) {
	switch f.Direction {
	case 0:
		f.NorthSouthPosition += amount
	case 90:
		f.EastWestPosition += amount
	case 180:
		f.NorthSouthPosition -= amount
	case 270:
		f.EastWestPosition -= amount
	}
}

func ParseDirections(lines []string) []Direction {
	directions := []Direction{}

	for _, line := range lines {
		var a Action
		var am int
		fmt.Sscanf(line, "%1s%d\n", &a, &am)
		directions = append(directions, Direction{Action: a, Amount: am})
	}

	return directions
}

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	directions := ParseDirections(lines)

	f := Ferry{Direction: 90, WayPoint: WayPoint{EastWestPosition: 10, NorthSouthPosition: 1}}

	if arguments.Tick {
		fmt.Printf("%d %d (%d) %v %v\n", f.NorthSouthPosition, f.EastWestPosition, f.Direction, nil, f.WayPoint)
	}
	for _, dir := range directions {
		f.FollowDirections(arguments.Part, dir)
		if arguments.Tick {
			fmt.Printf("%d %d (%d) %v %v\n", f.NorthSouthPosition, f.EastWestPosition, f.Direction, dir, f.WayPoint)
		}
		io.WaitForKey()
	}

	fmt.Printf(
		"manhattan distance: abs(%d) + abs(%d) = %d\n",
		f.NorthSouthPosition,
		f.EastWestPosition,
		int(math.Abs(float64(f.NorthSouthPosition)))+int(math.Abs(float64(f.EastWestPosition))),
	)
}

package main

import (
	"fmt"
	"log"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type State string

var (
	Floor    State = "."
	Empty    State = "L"
	Occupied State = "#"
)

type Square struct {
	State State
}

type Grid struct {
	Positions [][]State
}

func BuildGrid(lines []string) Grid {
	g := Grid{
		Positions: make([][]State, len(lines)),
	}
	for i, line := range lines {
		row := make([]State, len(line))
		for j, position := range line {
			var newState State

			switch State(position) {
			case Floor:
				newState = Floor
			case Empty:
				newState = Empty
			case Occupied:
				newState = Occupied
			}

			row[j] = newState
		}

		g.Positions[i] = row
	}

	return g
}

func (g *Grid) GetPosition(x int, y int) *State {
	if x < 0 || x >= len(g.Positions) {
		return nil
	}

	if y < 0 || y >= len(g.Positions[x]) {
		return nil
	}

	return &g.Positions[x][y]
}

func (g *Grid) TickPart1() (numChanged int) {
	newPositions := make([][]State, len(g.Positions))
	for i, row := range g.Positions {
		newPositions[i] = make([]State, len(row))
		for j, position := range row {
			newPositions[i][j] = g.Positions[i][j]
			adjacentCount := 0
			for _, xOffset := range []int{-1, 0, 1} {
				for _, yOffset := range []int{-1, 0, 1} {
					if xOffset == 0 && yOffset == 0 {
						continue
					}

					s := g.GetPosition(i+xOffset, j+yOffset)

					if s == nil {
						continue
					}

					if *s == Occupied {
						adjacentCount++
					}
				}
			}

			switch position {
			case Empty:
				if adjacentCount == 0 {
					newPositions[i][j] = Occupied
					numChanged++
				}
			case Occupied:
				if adjacentCount >= 4 {
					newPositions[i][j] = Empty
					numChanged++
				}
			}
		}
	}

	g.Positions = newPositions

	return numChanged
}

func (g *Grid) TickPart2() (numChanged int) {
	newPositions := make([][]State, len(g.Positions))
	for i, row := range g.Positions {
		newPositions[i] = make([]State, len(row))
		for j, position := range row {
			newPositions[i][j] = g.Positions[i][j]
			adjacentCount := 0
			for _, xOffset := range []int{-1, 0, 1} {
				for _, yOffset := range []int{-1, 0, 1} {
					if xOffset == 0 && yOffset == 0 {
						continue
					}

					s := g.GetPosition(i+xOffset, j+yOffset)

					// keep looking until we're off the board or find a seat
					newXOffset := xOffset
					newYOffset := yOffset
					for s != nil && *s == Floor {
						newXOffset += xOffset
						newYOffset += yOffset
						s = g.GetPosition(i+newXOffset, j+newYOffset)
					}

					if s == nil {
						continue
					}

					if *s == Occupied {
						adjacentCount++
					}
				}
			}

			switch position {
			case Empty:
				if adjacentCount == 0 {
					newPositions[i][j] = Occupied
					numChanged++
				}
			case Occupied:
				if adjacentCount >= 5 {
					newPositions[i][j] = Empty
					numChanged++
				}
			}
		}
	}

	g.Positions = newPositions
	fmt.Println(numChanged)

	return numChanged
}

func (g *Grid) Tick(part int) int {
	if part == 1 {
		return g.TickPart1()
	}
	return g.TickPart2()
}

func (g *Grid) Print() {
	for _, row := range g.Positions {
		for _, position := range row {
			fmt.Print(position)
		}
		fmt.Print("\n")
	}
}

func (g *Grid) NumOccupied() (numOccuped int) {
	for _, row := range g.Positions {
		for _, position := range row {
			if position == Occupied {
				numOccuped++
			}
		}
	}

	return numOccuped
}

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	g := BuildGrid(lines)

	g.Print()

	io.WaitForKey()

	for g.Tick(arguments.Part) != 0 {
		g.Print()
		io.WaitForKey()
	}

	fmt.Printf("num occupied: %d\n", g.NumOccupied())

}

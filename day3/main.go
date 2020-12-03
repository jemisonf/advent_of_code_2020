package main

import (
	"fmt"
	"log"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type Tree bool

type Board [][]Tree

func buildBoard(lines []string, right int, down int) *Board {
	height := len(lines)

	numSteps := (height - 1) / down

	width := 0
	for numSteps*right >= width {
		width += len(lines[0])
	}

	var board Board = [][]Tree{}

	for _, line := range lines {
		row := []Tree{}
		for i := 0; i < width; i++ {
			row = append(row, line[i%len(line)] == '#')
		}
		board = append(board, row)
	}

	return &board
}

func (b *Board) Print() {
	for _, row := range *b {
		for _, tree := range row {
			if tree {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (b *Board) Walk(right, down int) int {
	xPos, yPos := 0, 0

	numTrees := 0
	for yPos != len(*b)-1 {
		yPos += down
		xPos += right
		if (*b)[yPos][xPos] {
			numTrees++
		}
	}
	fmt.Println(numTrees)

	return numTrees
}

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("%v", err)
	}

	var right, down int

	if arguments.Part == 1 {
		right, down = 3, 1

		board := buildBoard(lines, right, down)

		fmt.Printf("encountered %d trees\n", board.Walk(right, down))
	} else {

		product := 1

		board := buildBoard(lines, 7, 1)
		product *= board.Walk(1, 1)
		product *= board.Walk(3, 1)
		product *= board.Walk(5, 1)
		product *= board.Walk(7, 1)
		product *= board.Walk(1, 2)

		fmt.Printf("the product is %d\n", product)
	}

}

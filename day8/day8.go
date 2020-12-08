package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type OpCode string

var (
	Nop OpCode = "nop"
	Acc OpCode = "acc"
	Jmp OpCode = "jmp"
)

type Instruction struct {
	Op       OpCode
	Argument int
}

type Program struct {
	Instructions []Instruction
	Visited      []bool
	Accumulator  int
	Current      int
	ID           string
}

func (p *Program) Evaluate() int {
	if p.Current == len(p.Instructions) {
		return p.Accumulator
	}

	instr := p.Instructions[p.Current]

	p.Visited[p.Current] = true

	switch instr.Op {
	case Acc:
		p.Accumulator += instr.Argument
		p.Current++
	case Nop:
		p.Current++
	case Jmp:
		p.Current += instr.Argument
	}

	return p.Accumulator
}

func (p *Program) CurrentInstruction() (Instruction, bool, int) {
	if p.Current == len(p.Instructions) {
		return Instruction{}, false, p.Accumulator
	}

	return p.Instructions[p.Current], p.Visited[p.Current], p.Accumulator
}

func (p *Program) Copy() Program {
	p2 := Program{Visited: p.Visited, Accumulator: p.Accumulator, Current: p.Current}
	p2.Instructions = make([]Instruction, len(p.Instructions))
	copy(p2.Instructions, p.Instructions)
	return p2
}

func ParseInstructions(lines []string) ([]Instruction, error) {
	instructions := []Instruction{}

	for _, line := range lines {
		instr := Instruction{}
		arg := 0
		code := ""

		fmt.Sscanf(line, "%s %d", &code, &arg)

		switch OpCode(code) {
		case Nop:
			instr = Instruction{
				Op:       Nop,
				Argument: arg,
			}
		case Acc:
			instr = Instruction{
				Op:       Acc,
				Argument: arg,
			}
		case Jmp:
			instr = Instruction{
				Op:       Jmp,
				Argument: arg,
			}
		default:
			return nil, fmt.Errorf("unrecognized op code %s", code)
		}

		instructions = append(instructions, instr)
	}

	return instructions, nil
}

func EvaluateInstructionsPart1(p Program) (int, bool) {
	visited := false
	acc := 0
	for !visited && p.Current < len(p.Instructions) {
		p.Evaluate()
		_, visited, acc = p.CurrentInstruction()
	}

	return acc, visited
}

func EvaluateInstructionsPart2(instructions []Instruction) int {
	p := Program{
		Instructions: instructions,
		Visited:      make([]bool, len(instructions)),
		ID:           "p1",
	}

	accumulatorValue := 0

	canExit := false
	for !canExit {
		instr, _, _ := p.CurrentInstruction()

		current := p.Current
		if instr.Op == Nop || instr.Op == Jmp {

			newInstr := Instruction{}
			newInstr.Argument = instr.Argument

			if instr.Op == Nop {
				newInstr.Op = Jmp
			}

			if instr.Op == Jmp {
				newInstr.Op = Nop
			}

			p2 := p.Copy()
			p2.ID = "p2"

			p2.Instructions[current] = newInstr

			acc, visited := EvaluateInstructionsPart1(p2)
			if !visited {
				accumulatorValue = acc
				canExit = true
			}
		}

		p.Evaluate()
	}

	return accumulatorValue
}

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file to read: %v", err)
	}

	instructions, err := ParseInstructions(lines)

	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	if arguments.Part == 1 {

		p := Program{
			Instructions: instructions,
			Visited:      make([]bool, len(instructions)),
		}

		acc, _ := EvaluateInstructionsPart1(p)

		fmt.Printf("accumulator value: %d\n", acc)
		os.Exit(0)
	}

	index := EvaluateInstructionsPart2(instructions)

	fmt.Printf("the accumulator value is %d\n", index)

}

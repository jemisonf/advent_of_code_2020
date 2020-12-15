package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jemisonf/advent_of_code_2020/day14/parser"
	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type Process struct {
	Memory             map[int]int
	Program            parser.Program
	CurrentInstruction int
}

func ApplyMask(mask string, value int) int {
	zerosMask := 0
	onesMask := 0
	for index := range mask {
		bit := mask[len(mask)-index-1]
		if bit != '0' {
			zerosMask |= 1 << index
		}

		if bit == '1' {
			onesMask |= 1 << index
		}
	}
	value &= zerosMask
	value |= onesMask
	return value
}

func (p *Process) RunProgramPart1() {
	for _, step := range p.Program.Steps {
		mask := step.Bitmask.Mask

		for _, instruction := range step.Instruction {
			value := ApplyMask(mask, instruction.Value)

			p.Memory[instruction.Address] = value
		}
	}
}

func ApplyMaskToAddress(mask string, address int) []int {
	zerosMask := 0
	onesMask := 0
	floatingBits := []int{}

	for index := range mask {
		bit := mask[len(mask)-index-1]
		if bit == 'X' {
			floatingBits = append(floatingBits, index)
		}

		if bit == '1' {
			onesMask |= 1 << index
		}
	}
	baseAddress := address
	baseAddress &= zerosMask

	addresses := []int{baseAddress}

	for _, floatingBit := range floatingBits {
		newAddresses := []int{}
		zeroMask := 0
		for i := 0; i <= 36; i++ {
			if i != floatingBit {
				zeroMask |= 1 << i
			}
		}

		for _, address := range addresses {
			newAddresses = append(newAddresses, address|(1<<floatingBit), address&zeroMask)
		}

		addresses = newAddresses
	}

	/*
		for _, address := range addresses {
			fmt.Printf("%036b\n", address)
		}
	*/

	return addresses
}

func (p *Process) RunProgramPart2() {
	fmt.Println(len(p.Program.Steps))
	for _, step := range p.Program.Steps {
		mask := step.Bitmask.Mask

		for _, instruction := range step.Instruction {
			addresses := ApplyMaskToAddress(mask, instruction.Address)
			for _, addr := range addresses {
				p.Memory[addr] = instruction.Value
			}
		}
	}
}

func (p *Process) RunProgram(part int) {
	if part == 1 {
		p.RunProgramPart1()
	} else {
		p.RunProgramPart2()
	}
}

func (p *Process) SumMemory() int {
	sum := 0
	for _, m := range p.Memory {
		sum += m
	}
	return sum
}

func main() {
	arguments := args.ParseArgs()

	contents, err := io.ReadFileAsString(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	program, err := parser.ParseProgram(contents)

	if err != nil {
		log.Fatalf("error parsing program: %v", err)
	}

	process := Process{Program: program, Memory: map[int]int{}}

	process.RunProgram(arguments.Part)

	fmt.Printf("sum of memory values: %d\n", process.SumMemory())
	os.Exit(0)

}

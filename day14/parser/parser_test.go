package parser_test

import (
	"github.com/jemisonf/advent_of_code_2020/day14/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parser", func() {
	var (
		rule  parser.Program
		err   error
		input string
	)
	JustBeforeEach(func() {
		rule, err = parser.ParseProgram(input)
	})

	Context("example program", func() {
		BeforeEach(func() {
			input = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`
		})

		It("does not err", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("parses the program correctly", func() {
			Expect(rule.Steps).NotTo(BeNil())
			Expect(len(rule.Steps)).To(Equal(1))
			step := rule.Steps[0]

			Expect(step.Bitmask.Mask).To(Equal("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"))

			Expect(len(step.Instruction)).To(Equal(3))

			instructions := step.Instruction

			Expect(instructions[0].Value).To(Equal(11))
			Expect(instructions[1].Value).To(Equal(101))
			Expect(instructions[2].Value).To(Equal(0))
			Expect(instructions[0].Address).To(Equal(8))
			Expect(instructions[1].Address).To(Equal(7))
			Expect(instructions[2].Address).To(Equal(8))
		})
	})

	Context("multiple lines", func() {
		BeforeEach(func() {
			input = `mask = 0111X10100100X1111X10010X000X1000001
mem[50907] = 468673978
mem[22295] = 3337449
mem[58474] = 56418393
mem[15362] = 243184
mem[65089] = 110688658
mask = 010X010XX110X01X01X10X001001011X110X
mem[21952] = 950257
mem[44861] = 522064487
mem[38886] = 28536885
mask = 01X1X1010XX0011X110XX100101010X01011
mem[34148] = 165121
mem[25371] = 1910147
mem[4508] = 873`
		})

		It("does not err", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates multiple steps", func() {
			Expect(len(rule.Steps)).To(Equal(3))
		})

		It("parses each step correctly", func() {
			Expect(rule.Steps[0].Bitmask.Mask).To(Equal("0111X10100100X1111X10010X000X1000001"))
			Expect(rule.Steps[1].Bitmask.Mask).To(Equal("010X010XX110X01X01X10X001001011X110X"))
			Expect(rule.Steps[2].Bitmask.Mask).To(Equal("01X1X1010XX0011X110XX100101010X01011"))

			Expect(rule.Steps[0].Instruction).To(ContainElements(
				parser.Instruction{50907, 468673978},
				parser.Instruction{22295, 3337449},
				parser.Instruction{58474, 56418393},
				parser.Instruction{15362, 243184},
				parser.Instruction{65089, 110688658},
			))
			Expect(rule.Steps[1].Instruction).To(ContainElements(
				parser.Instruction{21952, 950257},
				parser.Instruction{44861, 522064487},
				parser.Instruction{38886, 28536885},
			))
			Expect(rule.Steps[2].Instruction).To(ContainElements(
				parser.Instruction{34148, 165121},
				parser.Instruction{25371, 1910147},
				parser.Instruction{4508, 873},
			))
		})
	})
})

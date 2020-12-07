package parser_test

import (
	"github.com/jemisonf/advent_of_code_2020/day7/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	It("handles a normal line correctly", func() {
		rule, err := parser.ParseLine("vibrant salmon bags contain 1 vibrant gold bag, 2 wavy aqua bags, 1 dotted crimson bag.")
		Expect(err).NotTo(HaveOccurred())
		Expect(rule).To(Equal(parser.LuggageRule{
			Color: &parser.Color{
				Adjective: "vibrant",
				Name:      "salmon",
			},
			EdgeList: &parser.EdgeList{
				Edges: &[]parser.Edge{
					{
						Count: 1,
						Color: &parser.Color{
							Adjective: "vibrant",
							Name:      "gold",
						},
					},
					{
						Count: 2,
						Color: &parser.Color{
							Adjective: "wavy",
							Name:      "aqua",
						},
					},
					{
						Count: 1,
						Color: &parser.Color{
							Adjective: "dotted",
							Name:      "crimson",
						},
					},
				},
			},
		}))
	})

	It("handles a line with no edges", func() {
		rule, err := parser.ParseLine("dotted black bags contain no other bags.")
		Expect(err).NotTo(HaveOccurred())
		Expect(rule).To(Equal(
			parser.LuggageRule{
				Color: &parser.Color{
					Adjective: "dotted",
					Name:      "black",
				},
				EdgeList: &parser.EdgeList{},
				Empty:    &parser.EmptyRule{Empty: true},
			},
		))
	})

	It("handles a line with a single edge", func() {
		rule, err := parser.ParseLine("bright white bags contain 1 shiny gold bag.")
		Expect(err).NotTo(HaveOccurred())
		Expect(rule).To(Equal(parser.LuggageRule{
			Color: &parser.Color{
				Adjective: "bright",
				Name:      "white",
			},
			EdgeList: &parser.EdgeList{
				Edges: &[]parser.Edge{
					{
						Count: 1,
						Color: &parser.Color{
							Adjective: "shiny",
							Name:      "gold",
						},
					},
				},
			},
		}))
	})
})

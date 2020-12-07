package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/jemisonf/advent_of_code_2020/day7/parser"
	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

type Node struct {
	ID string
}

type Edge struct {
	OuterNodeID string
	InnerNodeID string
	Count       int
}
type Graph struct {
	Nodes []Node
	Edges []Edge
}

func (g *Graph) LookUpNodeEdges(id string) (error, Node, []Edge) {
	foundNode := false
	var node Node
	for _, n := range g.Nodes {
		if n.ID == id {
			node = n
			foundNode = true
		}
	}

	if !foundNode {
		return errors.New("node not found"), Node{}, nil
	}

	edges := []Edge{}
	for _, e := range g.Edges {
		if e.OuterNodeID == id {
			edges = append(edges, e)
		}
	}

	return nil, node, edges
}

func HandleParsedLine(rule parser.LuggageRule) (Node, []Edge) {
	node := Node{
		ID: fmt.Sprintf("%s %s", rule.Color.Adjective, rule.Color.Name),
	}

	edges := []Edge{}

	if rule.Empty != nil {
		return node, edges
	}

	for _, edge := range *rule.EdgeList.Edges {
		edges = append(edges, Edge{
			OuterNodeID: node.ID,
			InnerNodeID: fmt.Sprintf("%s %s", edge.Color.Adjective, edge.Color.Name),
			Count:       edge.Count,
		})
	}

	return node, edges
}

func CanCarry(bagID string, g Graph) []string {
	ids := map[string]bool{}

	for _, e := range g.Edges {
		if e.InnerNodeID == bagID {
			ids[e.OuterNodeID] = true
		}
	}

	for id := range ids {
		res := CanCarry(id, g)
		for _, r := range res {
			ids[r] = true
		}
	}

	keys := make([]string, len(ids))
	for id := range ids {
		keys = append(keys, id)
	}

	return keys
}

func BuildGraph(lines []string) Graph {
	g := Graph{}

	for _, line := range lines {
		ast, err := parser.ParseLine(line)

		if err != nil {
			log.Fatalf("Error parsing line %s: %v", line, err)
		}

		node, edges := HandleParsedLine(ast)

		g.Edges = append(g.Edges, edges...)
		g.Nodes = append(g.Nodes, node)
	}

	return g
}

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	g := BuildGraph(lines)
	fmt.Printf("Valid bag colors: %d\n", len(CanCarry("shiny gold", g)))
}

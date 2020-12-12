package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate DAY",
	Short:   "generate a new AoC directory",
	Long:    `creates a new directory with a dayX.go file that can be executed to read in an input file`,
	Example: `aocutil generate 11`,
	Args:    cobra.ExactArgs(1),
	Run:     Generate,
}

var fileTemplate = `package main

import (
	"fmt"
	"log"

	"github.com/jemisonf/advent_of_code_2020/pkg/args"
	"github.com/jemisonf/advent_of_code_2020/pkg/io"
)

func main() {
	arguments := args.ParseArgs()

	lines, err := io.ReadFileAsLines(arguments.File)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	fmt.Println(lines)
}
`

func Generate(cmd *cobra.Command, args []string) {
	day, _ := strconv.Atoi(args[0])

	dayName := fmt.Sprintf("day%d", day)

	if err := os.Mkdir(dayName, 0755); err != nil {
		log.Fatalf("error creating dir: %v", err)
	}

	fmt.Printf("%s/%s.go", dayName, dayName)
	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.go", dayName, dayName), []byte(fileTemplate), 0644); err != nil {
		log.Fatalf("error creating file: %v", err)
	}
}

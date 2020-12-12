package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "aocutil",
	Short: "a collection of utilities",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error running command: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

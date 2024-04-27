/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates code with various templates",
	Long:  `A code generator command for generating code for cubit, blocs`,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

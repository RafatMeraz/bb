package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bb",
	Short: "Borobhai is a cli app with regular dev tools",
	Long: `BB - Borobhai contains various tools for development like
flutter code generator, golang template generator which can been use to fast
development`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

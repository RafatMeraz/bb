package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blocCmd = &cobra.Command{
	Use:   "bloc",
	Short: "Command for generate bloc",
	Long:  `Generate new bloc, state and event with given name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bloc called")
	},
}

func init() {
	generateCmd.AddCommand(blocCmd)
	// Flags
	blocCmd.PersistentFlags().String("name", "", "Name of the bloc")
	blocCmd.PersistentFlags().String("state", "", "Name of the state")
	blocCmd.PersistentFlags().String("event", "", "Name of the event")
}

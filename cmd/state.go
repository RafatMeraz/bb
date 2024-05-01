package cmd

import (
	"bb/code_templates"
	"bb/utility"
	"fmt"
	"github.com/spf13/cobra"
)

var stateCmd = &cobra.Command{
	Use:   "state [name]",
	Short: "Generate state class for cubit or bloc",
	Long:  `A generator command for state class of bloc or cubit`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		stateName := args[0]
		partOfFileName, _ := cmd.Flags().GetString("partOf")
		if partOfFileName == "" {
			writeCodeWithoutPartOf(stateName)
		} else {
			writeCodeWithPartOf(stateName, partOfFileName)
		}
	},
}

func writeCodeWithPartOf(stateName string, partOfFileName string) {
	fileName := utility.CreateFileName(stateName)
	file := utility.CreateFile(fileName)
	utility.WriteCodeOnFile(file, code_templates.StateCodeWithPartOf(stateName, partOfFileName))
	message := fmt.Sprintf("%v is generated", stateName)
	fmt.Println(message)
}

func writeCodeWithoutPartOf(stateName string) {
	fileName := utility.CreateFileName(stateName)
	file := utility.CreateFile(fileName)
	utility.WriteCodeOnFile(file, code_templates.StateCodeWithOutPartOf(stateName))
	message := fmt.Sprintf("%v is generated", stateName)
	fmt.Println(message)
}

func callStateCmd(stateName string, partOfFileName string) {
	cmd := &cobra.Command{}
	if partOfFileName != "" {
		cmd.Flags().String("partOf", partOfFileName, "")
	}
	stateCmd.Run(cmd, []string{stateName})
}

func init() {
	generateCmd.AddCommand(stateCmd)
	// FLAGS
	stateCmd.PersistentFlags().StringP("partOf", "p", "", "Define the file name if state is part of that cubit/bloc")
}

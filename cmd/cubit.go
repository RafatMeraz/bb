package cmd

import (
	"bb/code_templates"
	"bb/utility"
	"fmt"
	"github.com/spf13/cobra"
)

var cubitCmd = &cobra.Command{
	Use:   "cubit [name] [state]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		cubitName := args[0]

		if len(args) >= 2 {
			stateName := args[1]
			createState(stateName, cubitName)
			createCubitWithCustomState(cubitName, stateName)
		} else {
			primitiveType, _ := cmd.Flags().GetString("p")
			if primitiveType == "" {
				primitiveType = "int"
			}
			createCubitWithPrimitive(cubitName, primitiveType)
		}
	},
}

func createCubitWithPrimitive(cubitName string, primitiveType string) {
	cubitFileName := utility.CreateFileName(cubitName)
	file := utility.CreateFile(cubitFileName)
	utility.WriteCodeOnFile(file, code_templates.CubitCode(cubitName, primitiveType))
	message := fmt.Sprintf("%v is created", cubitName)
	fmt.Println(message)
}

func createCubitWithCustomState(cubitName string, stateName string) {
	cubitFileName := utility.CreateFileName(cubitName)
	file := utility.CreateFile(cubitFileName)
	stateFileName := utility.CreateFileName(stateName)
	utility.WriteCodeOnFile(file, code_templates.CubitCodeWithCustomState(stateFileName, cubitName, stateName))
	message := fmt.Sprintf("%v is created", cubitName)
	fmt.Println(message)
}

func createState(stateName string, cubitName string) {
	cubitFileName := utility.CreateFileName(cubitName)
	cmd := &cobra.Command{}
	cmd.Flags().String("partOf", cubitFileName, "")
	stateCmd.Run(cmd, []string{stateName})
}

func init() {
	generateCmd.AddCommand(cubitCmd)
	cubitCmd.PersistentFlags().String("p", "", "Set primitive type if state is dart primitive")
}

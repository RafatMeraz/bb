package cmd

import (
	"bb/code_templates"
	"bb/utility"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var eventCmd = &cobra.Command{
	Use:   "event [name]",
	Short: "Generate event class",
	Long:  `Command for generate the event class of Bloc`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		eventName := args[0]
		stateName, _ := cmd.Flags().GetString("partOf")
		generateEvent(eventName, stateName)
	},
}

func init() {
	generateCmd.AddCommand(eventCmd)
	//FLAGS
	eventCmd.PersistentFlags().StringP("partOf",
		"p", "",
		"Set if this event is part of any cubit/bloc")
}

func generateEvent(eventName string, partOfFileName string) {
	eventCode := getEventCode(eventName, partOfFileName)
	fileName := utility.CreateFileName(eventName)
	file := utility.CreateFile(fileName)
	utility.WriteCodeOnFile(file, eventCode)
	message := fmt.Sprintf("%v is generated!", eventName)
	fmt.Println(message)
}

func getEventCode(eventName string, partOfFileName string) string {
	var eventCode string
	if partOfFileName != "" {
		eventCode = code_templates.GetEventCodeWithPartOf(eventName, partOfFileName)
	} else {
		eventCode = code_templates.GetEventCodeWithoutPartOf(eventName)
	}
	return eventCode
}

func callEventCmd(eventName string, partOfFileName string) {
	cmd := &cobra.Command{}
	if partOfFileName != "" {
		cmd.Flags().String("partOf", partOfFileName, "")
	}
	eventCmd.Run(cmd, []string{eventName})
}

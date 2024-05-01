package cmd

import (
	"bb/code_templates"
	"bb/utility"
	"fmt"
	"github.com/spf13/cobra"
	"sync"
)

// blocCmd represents the bloc command
var blocCmd = &cobra.Command{
	Use:   "bloc [name]",
	Short: "Generate a new bloc",
	Long:  `Command for generate bloc, optionally you can generate state and event with it.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		blocName := args[0]
		stateName, _ := cmd.Flags().GetString("state")
		eventName, _ := cmd.Flags().GetString("event")
		if stateName == "" || eventName == "" {
			cmd.Help()
			return
		}
		generateBlocCode(blocName, eventName, stateName)
	},
}

func generateBlocCode(blocName string, eventName string, stateName string) {
	var wg sync.WaitGroup
	blocFileName := utility.CreateFileName(blocName)
	stateFileName := utility.CreateFileName(stateName)
	eventFileName := utility.CreateFileName(eventName)

	config := code_templates.BlocConfig{
		BlocName:      blocName,
		StateName:     stateName,
		StateFileName: stateFileName,
		EventName:     eventName,
		EventFileName: eventFileName,
	}
	wg.Add(3)
	go generateBlocEvent(&wg, blocFileName, eventName)
	go generateBlocState(&wg, blocFileName, stateName)
	go writeBlocCode(&wg, config)

	wg.Wait()
}

func writeBlocCode(w *sync.WaitGroup, config code_templates.BlocConfig) {
	defer w.Done()
	fileName := utility.CreateFileName(config.BlocName)
	file := utility.CreateFile(fileName)

	utility.WriteCodeOnFile(file, code_templates.GetBlocCode(config))
	message := fmt.Sprintf("%v is generated!", config.BlocName)
	fmt.Println(message)
}

func generateBlocEvent(w *sync.WaitGroup, blocFileName string, e string) {
	defer w.Done()
	callEventCmd(e, blocFileName)
}

func generateBlocState(w *sync.WaitGroup, blocFileName string, stateName string) {
	defer w.Done()
	callStateCmd(stateName, blocFileName)
}

func init() {
	generateCmd.AddCommand(blocCmd)
	//FLAGS
	blocCmd.PersistentFlags().StringP("state",
		"s", "",
		"Name of the state [Mandatory]")
	blocCmd.PersistentFlags().StringP("event",
		"e", "",
		"Name of the event [Mandatory]")
}

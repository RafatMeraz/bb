package cmd

import (
	"fmt"
	"github.com/fatih/camelcase"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var cubitCmd = &cobra.Command{
	Use:   "cubit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		state, _ := cmd.Flags().GetString("state")
		isPrimitive, _ := cmd.Flags().GetBool("p")
		if name == "" {
			name = "RandomCubit"
		}
		if state == "" {
			name = "int"
		}
		generateCubit(name, state, isPrimitive)
	},
}

func init() {
	generateCmd.AddCommand(cubitCmd)
	cubitCmd.PersistentFlags().String("name", "", "Set name of the cubit")
	cubitCmd.PersistentFlags().String("state", "", "Set name of the state")
	cubitCmd.PersistentFlags().Bool("p", false, "Set name of the state")
}

func generateCubit(cubitName string, stateName string, isPrimitive bool) {
	// generate cubit
	cubitFileName := getCubitFileName(cubitName)
	cubitFile := createFile(cubitFileName)
	defer cubitFile.Close()

	// generate state
	if isPrimitive == false {
		stateFileName := getCubitFileName(stateName)
		stateFile := createFile(stateFileName)
		defer stateFile.Close()
		writeStateCode(stateName, cubitFileName, stateFile)

		writeCubitCodeWithCustomState(stateFileName, cubitName, stateName, cubitFile)
		fmt.Println(fmt.Sprintf("%v has been generated", cubitName))
	} else {
		writeCubitCode(cubitName, stateName, cubitFile)
		fmt.Println(fmt.Sprintf("%v has been generated", cubitName))
	}
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	return file
}

func getCubitFileName(cubitName string) string {
	name := getFileName(cubitName)
	name = name + ".dart"
	return name
}

func getFileName(cubitName string) string {
	verbs := camelcase.Split(cubitName)
	for i, v := range verbs {
		verbs[i] = strings.ToLower(v)
	}
	fileName := strings.Join(verbs, "_")

	return fileName
}

func writeCubitCode(cubitName string, stateName string, file *os.File) {
	_, err := file.Write([]byte(cubitCode(cubitName, stateName)))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func writeCubitCodeWithCustomState(stateFileName string, cubitName string, stateName string, file *os.File) {
	_, err := file.Write([]byte(cubitCodeWithCustomState(stateFileName, cubitName, stateName)))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func writeStateCode(stateName string, cubitFileName string, file *os.File) {
	_, err := file.Write([]byte(stateCode(stateName, cubitFileName)))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func cubitCode(cubitName string, stateName string) string {
	return fmt.Sprintf(`import 'package:flutter_bloc/flutter_bloc.dart';

class %v extends Cubit<%v> {
			void doSomething() {}
}`, cubitName, stateName)
}

func cubitCodeWithCustomState(stateFileName string, cubitName string, stateName string) string {
	return fmt.Sprintf(`import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:equatable/equatable.dart';

part '%v';

class %v extends Cubit<%v> {
			void doSomething() {}
}`, stateFileName, cubitName, stateName)
}

func stateCode(stateName string, cubitFileName string) string {
	return fmt.Sprintf(`part of '%v';

sealed class %v extends Equatable {}
`, cubitFileName, stateName)
}

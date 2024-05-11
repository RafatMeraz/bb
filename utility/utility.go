package utility

import (
	"errors"
	"github.com/fatih/camelcase"
	"log"
	"os"
	"strings"
)

func CreateFileName(name string) string {
	verbs := camelcase.Split(name)
	for i, v := range verbs {
		verbs[i] = strings.ToLower(v)
	}
	fileName := strings.Join(verbs, "_")

	return fileName + ".dart"
}

func WriteCodeOnFile(file *os.File, code string) {
	_, err := file.Write([]byte(code))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func CreateInitialStateName(stateName string) string {
	res := strings.Split(stateName, "State")
	return res[0] + "InitialState"
}

func GetCurrentWorkingDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func ChangeDirectory(directory string) error {
	err := os.Chdir(directory)
	if err != nil {
		return errors.New("failed to change directory")
	}
	return nil
}

func CreateDirectory(directoryName string) error {
	err := os.Mkdir(directoryName, os.ModePerm)
	if err != nil {
		return errors.New("failed to create new directory")
	}
	return nil
}

func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	return file
}

package utility

import (
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

func CreateFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	return file
}

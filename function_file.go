package goduplicateword

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ProcessPrepareInputFile(inputpath string) {
	files, err := ioutil.ReadDir(inputpath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func CollectInputFileName(inputPath string) []string {
	files, err := ioutil.ReadDir(inputPath)
	var txtfile []string

	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		for _, file := range files {
			fullpath := inputPath + "/" + file.Name()
			fmt.Println(fullpath)
			txtfile = append(txtfile, fullpath)
		}
	}

	return txtfile
}

func CleanCreateDirPrepareAndOutput(pathPrepare string, pathOutput string) {
	CleanAndCreateDir(pathPrepare)
	CleanAndCreateDir(pathOutput)
}

func CleanAndCreateDir(pathTocreate string) {
	if PathExists(pathTocreate) {
		CleanDir(pathTocreate)
	}
	CreateDir(pathTocreate)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	var result bool = false
	if err == nil {
		result = true
	}
	if os.IsNotExist(err) {
		result = false
	}
	return result
}

func CreateDir(path string) {
	if !PathExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func CleanDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(path string) string {

	var result string = ""
	if PathExists(path) {
		body, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}

		result = string(body)
	}

	return result
}

func ReadWordConfig(path string) []string {

	var result string = ReadFile(path)
	result = strings.ReplaceAll(result, "\r", "")
	arr := strings.Split(result, "\n")
	return arr
}

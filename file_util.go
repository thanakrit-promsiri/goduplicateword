package goduplicateword

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CollectInputFileName(path string) []string {
	files, err := ioutil.ReadDir(path)
	var txtfile []string

	if err != nil {
		log.Fatal(err)
		panic(err)
	} else {
		for _, file := range files {
			filename := file.Name()
			txtfile = append(txtfile, filename)
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

func GetAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// Print the absolute path of the file
	return absPath
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

func WriteFile(pathfile string, filecontent string) {
	f, err := os.Create(pathfile)
	if err != nil {
		log.Fatalf("unable to Create file: %v", err)
		return
	}
	l, err := f.WriteString(filecontent)
	if err != nil {
		log.Fatalf("unable to WriteString file: %v", err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Fatalf("unable to WriteString file: %v", err)
		return
	}
}

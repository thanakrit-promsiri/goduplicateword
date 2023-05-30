package goduplicateword

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ProcessMapReduce(txt string) map[string]int {
	var wordSlice = strings.Split(txt, " ")
	reduceTxt := make(map[string]int)
	for _, word := range wordSlice {
		word = strings.ToLower(word)
		value, exists := reduceTxt[word]
		if exists {
			reduceTxt[word] = value + 1
		} else {
			reduceTxt[word] = 1
		}

	}
	return reduceTxt
}

func ProcessPrepareText(txt string) string {

	txt = textRemove(txt)
	txt = textSpace(txt)
	txt = textInsertSpace(txt)
	return txt
}

func processPrepareInputFile(inputpath string) string {
	files, err := ioutil.ReadDir(inputpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

func textRemove(txt string) string {
	return txt
}

func textSpace(txt string) string {
	return txt
}

func textInsertSpace(txt string) string {
	return txt
}

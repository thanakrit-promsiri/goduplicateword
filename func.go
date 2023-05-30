package goduplicateword

import (
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

func textRemove(txt string) string {
	return txt
}

func textSpace(txt string) string {
	return txt
}

func textInsertSpace(txt string) string {
	return txt
}

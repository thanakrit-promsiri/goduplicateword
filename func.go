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

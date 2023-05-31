package goduplicateword

import (
	"strings"
)

func ProcessMapReduce(txt string, reduceTxt map[string]int) map[string]int {

	txt = strings.Replace(txt, "\r\n", "\n", -1)
	txt = strings.Replace(txt, "\n", " ", -1)
	txt = strings.Replace(txt, "\t", " ", -1)

	var wordSlice = strings.Split(txt, " ")

	for _, word := range wordSlice {

		word = strings.ToLower(word)
		word = strings.TrimSpace(word)

		value, exists := reduceTxt[word]
		if exists {
			reduceTxt[word] = value + 1
		} else {
			reduceTxt[word] = 1
		}

	}
	return reduceTxt
}

func TextRemove(txt string, configList []string) string {
	for _, wordConfig := range configList {
		txt = strings.ReplaceAll(txt, wordConfig, "")
	}

	return txt
}

func TextSpace(txt string, configList []string) string {
	for _, wordConfig := range configList {
		txt = strings.ReplaceAll(txt, wordConfig, " ")
	}

	return txt
}

func TextInsertSpacebothSides(txt string, configList []string) string {
	for _, wordConfig := range configList {
		txt = strings.ReplaceAll(txt, wordConfig, " "+wordConfig+" ")
	}

	return txt
}

func TextSplitLine(txt string) []string {
	//text form windows
	txt = strings.Replace(txt, "\r\n", "\n", -1)
	txt = strings.TrimSpace(txt)
	result := strings.Split(txt, "\n")
	return result
}

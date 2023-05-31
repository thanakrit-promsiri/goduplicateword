package goduplicateword

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadWordConfig(path string) []string {
	var txt string = ReadFile(path)
	txt = strings.TrimSpace(txt)
	var arr []string
	if len(txt) > 0 {
		arr = TextSplitLine(txt)
	}

	return arr
}

func ProcessPrepareText(txt string, configRemove []string, configSpace []string, configInsertSpace []string) string {

	txt = TextRemove(txt, configRemove)
	txt = TextSpace(txt, configSpace)
	txt = TextInsertSpacebothSides(txt, configInsertSpace)
	return txt
}

func MapToCSVtext(mapReduce map[string]int, path string) {

	csvFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
	csvFile.Write(bomUtf8)
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	var wordSortedArr []string = SortKey(mapReduce)

	for _, wordSorted := range wordSortedArr {
		r := make([]string, 0, 2)
		r = append(r, wordSorted)
		r = append(r, strconv.Itoa(mapReduce[wordSorted]))

		err := csvWriter.Write(r)
		if err != nil {
			// handle error
		}
	}
	defer csvWriter.Flush()
}

func SortKey(mapReduce map[string]int) []string {
	keys := make([]string, 0, len(mapReduce))
	for k := range mapReduce {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
	//for _, k := range keys {
	//	fmt.Println(k, mapReduce[k])
	//}
}

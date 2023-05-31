package main

import (
	"fmt"

	engine "github.com/thanakrit-promsiri/goduplicateword"
	"github.com/thanakrit-promsiri/goduplicateword/internal/viperconf"
)

var conf viperconf.Configuration

var inputPath string
var preparePath string
var outputPath string

var configWordToSpace string
var configWordToRemove string
var configWordInsertSpace string

func initConfigAndDir() {
	conf = viperconf.InitConfig()

	rootPath := conf.Sourcefilerepository.RootDirectoryPath
	inputPath = rootPath + conf.Sourcefilerepository.DirectoryInput
	preparePath = rootPath + conf.Sourcefilerepository.DirectoryPrepare
	outputPath = rootPath + conf.Sourcefilerepository.DirectoryOutput

	configWordToSpace = conf.Sourcefilerepository.ConfigWordToSpace
	configWordToRemove = conf.Sourcefilerepository.ConfigWordToRemove
	configWordInsertSpace = conf.Sourcefilerepository.ConfigWordInsertSpace

	fmt.Println("Root Folder:", engine.GetAbsPath(rootPath))

	if engine.PathExists(rootPath) {
		fmt.Println("Check Root Folder is exists : OK")
	} else {
		panic("Root Folder is not exists")
	}

	fmt.Println("Input Folder:", engine.GetAbsPath(inputPath))
	fmt.Println()

	if engine.PathExists(inputPath) {
		files := engine.CollectInputFileName(inputPath)
		fmt.Println("Check Input Folder is exists : OK , Total files :", len(files))
		fmt.Println("Check Input Folder is exists : ", files)

	} else {
		panic("Input Folder is not exists")
	}
	fmt.Println()

	engine.CleanCreateDirPrepareAndOutput(preparePath, outputPath)

	fmt.Println("Clean And Create Prepare Folder:", engine.GetAbsPath(preparePath))
	fmt.Println("Clean And Create Output Folder :", engine.GetAbsPath(outputPath))

}

func prepareFileInput() {
	var wordToRemove []string = engine.ReadWordConfig(configWordToRemove)
	var wordToSpace []string = engine.ReadWordConfig(configWordToSpace)
	var WordInsertSpace []string = engine.ReadWordConfig(configWordInsertSpace)

	var txt string = ""

	files := engine.CollectInputFileName(inputPath)
	for _, file := range files {

		var pathInputReadFile string = inputPath + "/" + file
		var pathPrepareWriteFile string = preparePath + "/" + file

		txt = engine.ReadFile(pathInputReadFile)
		txt = engine.ProcessPrepareText(txt, wordToRemove, wordToSpace, WordInsertSpace)
		engine.WriteFile(pathPrepareWriteFile, txt)
	}

}

func processDuplication() {
	reduceTxtmap := make(map[string]int)
	files := engine.CollectInputFileName(preparePath)
	for _, file := range files {

		var pathPrepareFile string = preparePath + "/" + file

		txt := engine.ReadFile(pathPrepareFile)
		reduceTxtmap = engine.ProcessMapReduce(txt, reduceTxtmap)
		//engine.WriteFile(pathPrepareWriteFile, txt)

	}

	engine.MapToCSVtext(reduceTxtmap, outputPath+"/summary.csv")
	fmt.Println("Summary file Output successfully:", engine.GetAbsPath(outputPath+"/summary.csv"))

}

func main() {
	var any string
	initConfigAndDir()
	prepareFileInput()
	processDuplication()

	fmt.Println("press X and Enter to exit.....")
	fmt.Scan(&any)

}

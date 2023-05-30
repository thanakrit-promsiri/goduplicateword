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

func initConfig() {
	conf = viperconf.InitConfig()
	fmt.Println(viperconf.ViperEcho())

	rootPath := conf.Sourcefilerepository.RootDirectoryPath
	inputPath = rootPath + conf.Sourcefilerepository.DirectoryInput
	preparePath = rootPath + conf.Sourcefilerepository.DirectoryPrepare
	outputPath = rootPath + conf.Sourcefilerepository.DirectoryOutput

	configWordToSpace = conf.Sourcefilerepository.ConfigWordToSpace
	configWordToRemove = conf.Sourcefilerepository.ConfigWordToRemove

	fmt.Println()
	fmt.Println("inputPath :", inputPath)
	fmt.Println("preparepath :", preparePath)
	fmt.Println("outputpath :", outputPath)
	fmt.Println()
}

func main() {

	initConfig()

	engine.CleanCreateDirPrepareAndOutput(preparePath, outputPath)

	fmt.Println("xxx:", configWordToSpace)

	fmt.Println("xxx:", engine.ReadWordConfig(configWordToSpace))

	InputFileName := engine.CollectInputFileName(inputPath)

	fmt.Println("map:", InputFileName)

	processMapReduce := engine.ProcessMapReduce("xx vv xx")

	fmt.Println("map:", processMapReduce)
}

package main

import (
	"flag"

	"fmt"
	"log"
	"strings"

	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
	"github.com/zxjsdp/specimen-go/web"
)

func specimenInfo(markerDataFile, entryDataFile, outputDataFile string) {
	entryDataMatrix := files.GetDataMatrix(entryDataFile)
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)

	markerDataMatrix := files.GetDataMatrix(markerDataFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)

	speciesNames := converters.ExtractSpeciesNames(entryDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)

	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}

	files.SaveDataMatrix(outputDataFile, resultDataSlice)
}

// 解析命令行，并输出 query 文件名、data 文件名、output 文件名
func parseArgument() (string, string, string) {
	markerDataPtr := flag.String("query", "", "物种编号文件（query.xlsx）")
	entryDataPtr := flag.String("data", "", "物种记录及鉴定文件（data.xlsx）")
	outputDataPtr := flag.String("output", "", "输出文件（xlsx 格式）")

	flag.Parse()

	utils.CheckFileExists(*markerDataPtr, "-query", config.USAGE)
	utils.CheckFileExists(*entryDataPtr, "-data", config.USAGE)
	if len(strings.TrimSpace(*outputDataPtr)) == 0 {
		log.Fatal(fmt.Sprintf("ERROR! Blank argument for [ -output ].%s", config.USAGE))
	}

	return *markerDataPtr, *entryDataPtr, *outputDataPtr
}

func main() {
	markerDataFile, entryDataFile, outputDataFile := parseArgument()
	specimenInfo(markerDataFile, entryDataFile, outputDataFile)
}

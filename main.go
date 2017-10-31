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
	fmt.Printf("%%01  开始读取 entry 数据文件 ...\n")
	entryDataMatrix := files.GetDataMatrix(entryDataFile)
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)
	fmt.Printf("%%09  读取 entry 数据结束！\n")

	fmt.Printf("%%10  开始读取 marker 数据文件 ...\n")
	markerDataMatrix := files.GetDataMatrix(markerDataFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	fmt.Printf("%%19  读取 marker 数据结束！\n")

	fmt.Printf("%%20  开始提取网络信息 ...\n")
	speciesNames := converters.ExtractSpeciesNames(entryDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)
	fmt.Printf("%%90  提取网络信息结束！\n")

	fmt.Printf("%%91  开始整合本地数据及网络信息 ...\n")
	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}
	fmt.Printf("%%94  整合本地数据及网络信息结束！\n")

	fmt.Printf("%%95  开始将结果信息写入 xlsx 输出文件...\n")
	files.SaveDataMatrix(outputDataFile, resultDataSlice)

	fmt.Printf("%%100 任务完成！\n")
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

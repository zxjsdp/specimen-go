package specimen

import (
	"log"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/web"
)

func RunSpecimenInfo(markerDataFile, entryDataFile, outputDataFile string) {
	log.Printf("%%01  开始读取 entry 数据文件 ...\n")
	entryDataMatrix := files.GetDataMatrix(entryDataFile)
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)
	log.Printf("%%09  读取 entry 数据结束！\n")

	log.Printf("%%10  开始读取 marker 数据文件 ...\n")
	markerDataMatrix := files.GetDataMatrix(markerDataFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	log.Printf("%%19  读取 marker 数据结束！\n")

	log.Printf("%%20  开始提取网络信息 ...\n")
	speciesNames := converters.ExtractSpeciesNames(entryDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)
	log.Printf("%%90  提取网络信息结束！\n")

	log.Printf("%%91  开始整合本地数据及网络信息 ...\n")
	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}
	log.Printf("%%94  整合本地数据及网络信息结束！\n")

	log.Printf("%%95  开始将结果信息写入 xlsx 输出文件...\n")
	files.SaveDataMatrix(outputDataFile, resultDataSlice)

	log.Printf("%%100 任务完成！\n")
}

package specimen

import (
	"log"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
	"github.com/zxjsdp/specimen-go/web"
)

func RunSpecimenInfo(markerDataFile, entryDataFile, outputDataFile string, doesMarkerFileHasHeader bool) {
	log.Printf("开始读取 entry 数据文件 ...（进度 %%1）\n")
	entryDataMatrix := files.GetDataMatrix(entryDataFile)
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)
	log.Printf("读取 entry 数据结束！（进度 %%9）\n")

	log.Printf("开始读取 marker 数据文件 ...（进度 %%10）\n")
	markerDataMatrix := files.GetDataMatrix(markerDataFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	log.Printf("读取 marker 数据结束！（进度 %%19）\n")

	validationResult := utils.DataValidation(entryDataMatrix, markerDataMatrix)
	if !validationResult.Result {
		for i, failureInfo := range validationResult.FailureInfo {
			log.Printf("错误（%d）%s\n", i+1, failureInfo)
		}
		for i, warningInfo := range validationResult.WarningInfo {
			log.Printf("警告（%d）%s\n", i+1, warningInfo)
		}

		log.Printf("请解决上述错误后再重新运行。程序即将退出！\n")
		return
	} else {
		for i, warningInfo := range validationResult.WarningInfo {
			log.Printf("警告（%d）%s\n", i+1, warningInfo)
		}
	}

	log.Printf("开始提取网络信息 ...（进度 %%20）\n")
	speciesNames := converters.ExtractSpeciesNames(entryDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)
	log.Printf("提取网络信息结束！（进度 %%90）\n")

	log.Printf("开始整合本地数据及网络信息 ...（进度 %%91）\n")
	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	if doesMarkerFileHasHeader {
		markerDataSlice = markerDataSlice[1:] // 去除 marker 文件中的标题行
	}
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}
	log.Printf("整合本地数据及网络信息结束！（进度 %%94）\n")

	log.Printf("开始将结果信息写入 xlsx 输出文件...（进度 %%95）\n")
	files.SaveDataMatrix(outputDataFile, resultDataSlice)

	log.Printf("任务完成！（进度 %%100）\n")
}

package specimen

import (
	"log"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
	"github.com/zxjsdp/specimen-go/web"
)

func RunSpecimenInfo(markerDataFile, offlineDataFile, outputDataFile string, doesMarkerFileHasHeader bool) {
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 文件读取及解析
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("开始读取 offline 数据文件 ...（进度 %%1）\n")
	offlineDataMatrix := files.GetDataMatrix(offlineDataFile)
	offlineDataSlice := converters.ToOfflineDataSlice(offlineDataMatrix)
	offlineDataMap := converters.GenerateOfflineDataMap(offlineDataSlice)
	log.Printf("读取 offline 数据结束！（进度 %%9）\n")

	log.Printf("开始读取 marker 数据文件 ...（进度 %%10）\n")
	markerDataMatrix := files.GetDataMatrix(markerDataFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	log.Printf("读取 marker 数据结束！（进度 %%19）\n")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 数据校验
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	validationResult := utils.DataValidation(offlineDataMatrix, markerDataMatrix)
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

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 从网络获取信息
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("开始提取网络信息 ...（进度 %%20）\n")
	speciesNames := converters.ExtractSpeciesNames(offlineDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)
	log.Printf("提取网络信息结束！（进度 %%90）\n")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 整合数据信息及网络信息并生成结果
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("开始整合本地数据及网络信息 ...（进度 %%91）\n")
	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	if doesMarkerFileHasHeader {
		markerDataSlice = markerDataSlice[1:] // 去除 marker 文件中的标题行
	}
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, offlineDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}
	log.Printf("整合本地数据及网络信息结束！（进度 %%94）\n")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 将结果写入输出文件
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("开始将结果信息写入 xlsx 输出文件...（进度 %%95）\n")
	files.SaveDataMatrix(outputDataFile, resultDataSlice)

	log.Printf("任务完成！（进度 %%100）\n")
}

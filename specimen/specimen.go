package specimen

import (
	"log"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
	"github.com/zxjsdp/specimen-go/web"
)

func RunSpecimenInfo(snDataFile, offlineDataFile, outputDataFile string, doesSnFileHasHeader bool) {
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 文件读取及解析
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("[ 🟢 进度 %%1 ] 开始读取 “鉴定录入文件” 数据 ...\n")
	offlineDataMatrix := files.GetDataMatrix(offlineDataFile)
	offlineDataSlice := converters.ToOfflineDataSlice(offlineDataMatrix)
	offlineDataMap := converters.GenerateOfflineDataMap(offlineDataSlice)
	log.Printf("[ 🟢 进度 %%9 ] 读取 “鉴定录入文件” 数据结束！\n")

	log.Printf("[ 🟢 进度 %%10 ] 开始读取 “流水号文件” 数据 ...\n")
	snDataMatrix := files.GetDataMatrix(snDataFile)
	snDataSlice := converters.ToSnDataSlice(snDataMatrix)
	log.Printf("[ 🟢 进度 %%19 ] 读取 “流水号文件” 数据结束！\n")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 数据校验
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	validationResult := utils.DataValidation(offlineDataMatrix, snDataMatrix)
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
	log.Printf("[ 🟢 进度 %%20 ] 开始提取网络信息 ...\n")
	speciesNames := converters.ExtractSpeciesNames(offlineDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames)
	log.Printf("[ 🟢 进度 %%90 ] 提取网络信息结束！\n")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 整合数据信息及网络信息并生成结果
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("[ 🟢 进度 %%91 ] 开始整合本地数据及网络信息 ...\n")
	resultDataSlice := make([]entities.ResultData, len(snDataSlice))
	if doesSnFileHasHeader {
		snDataSlice = snDataSlice[1:] // 去除 “流水号文件” 中的标题行
	}
	for i, snData := range snDataSlice {
		resultData := converters.ToResultData(snData, offlineDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}
	log.Printf("[ 🟢 进度 %%94 ] 整合本地数据及网络信息结束！\n")

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 将结果写入输出文件
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	log.Printf("[ 🟢 进度 %%95 ] 开始将结果信息写入 xlsx 输出文件...\n")
	files.SaveResultDataToXlsx(outputDataFile, resultDataSlice)

	log.Printf("[ 🟢 进度 %%100 ] 任务完成！\n")
}

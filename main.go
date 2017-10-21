package main

import (
	"github.com/zxjsdp/specimen-go/constant"
	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
)

func generateResultData() string {
	return "Result"
}

func main() {
	entryDataMatrix := files.GetDataMatrix("data/data.xlsx")
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)

	markerDataMatrix := files.GetDataMatrix("data/query.xlsx")
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)

	webDataMap := make(map[string]entities.WebInfo)
	resultDataSlice := make([]entities.ResultData, 0)
	for _, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webDataMap)
		resultDataSlice = append(resultDataSlice, resultData)
	}

	files.SaveDataMatrix(constant.DefaultReslutXlsxName, resultDataSlice)
}

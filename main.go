package main

import (
	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/web"
)

func specimenInfo() {
	entryDataMatrix := files.GetDataMatrix("data/data.xlsx")
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)

	markerDataMatrix := files.GetDataMatrix("data/query.xlsx")
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)

	speciesNames := converters.ExtractSpeciesNames(entryDataSlice)
	webInfoMap := web.GenerateWebInfoMap(speciesNames[:3])

	resultDataSlice := make([]entities.ResultData, len(markerDataSlice))
	for i, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webInfoMap)
		resultDataSlice[i] = resultData
	}

	files.SaveDataMatrix(config.DefaultResultXlsxName, resultDataSlice)
}

func main() {
	specimenInfo()
}

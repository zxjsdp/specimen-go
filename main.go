package main

import (
	"fmt"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/filetype"
)

func generateResultData() string {
	return "Result"
}

func main() {
	entryDataMatrix := filetype.GetDataMatrix("data/data.xlsx")
	entryDatas := converters.ToEntryDatas(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDatas)

	markerDataMatrix := filetype.GetDataMatrix("data/query.xlsx")
	markerDatas := converters.ToMarkerDatas(markerDataMatrix)

	webDataMap := make(map[string]entities.WebInfo)
	for _, marker := range markerDatas {
		resultData := converters.ToResultData(marker, entryDataMap, webDataMap)
		fmt.Println(resultData)
	}
}

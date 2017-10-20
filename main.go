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
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	entryDataMap := converters.GenerateEntryDataMap(entryDataSlice)

	markerDataMatrix := filetype.GetDataMatrix("data/query.xlsx")
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)

	webDataMap := make(map[string]entities.WebInfo)
	for _, marker := range markerDataSlice {
		resultData := converters.ToResultData(marker, entryDataMap, webDataMap)
		fmt.Println(resultData)
	}
}

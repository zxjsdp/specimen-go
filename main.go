package main

import (
	"fmt"

	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/config/web"
	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/files"
)

func specimenInfo() {
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

	files.SaveDataMatrix(config.DefaultReslutXlsxName, resultDataSlice)
}

func main() {
	//specimenInfo()
	webInfo := web.GenerateWebInfo("Firmiana platanifolia")
	fmt.Println(webInfo.BodyHeight)
	fmt.Println(webInfo.DBH)
	fmt.Println(webInfo.Stem)
	fmt.Println(webInfo.Leaf)
	fmt.Println(webInfo.Flower)
	fmt.Println(webInfo.Fruit)
	fmt.Println(webInfo.Host)
}

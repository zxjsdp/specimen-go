package main

import (
	"fmt"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/filetype"
)

func generateResultData() string {
	resultData := entities.ResultData{}
	resultData.Altitude = ""
	return "Result"
}

func main() {
	entryDataMatrix := filetype.GetDataMatrix("data/data.xlsx")
	entryDatas := converters.ToEntryDatas(entryDataMatrix)
	for _, each := range entryDatas {
		fmt.Println(each.String())
	}

	markerDataMatrix := filetype.GetDataMatrix("data/query.xlsx")
	markerDatas := converters.ToMarkerDatas(markerDataMatrix)
	for _, each := range markerDatas {
		fmt.Println(each)
	}
}

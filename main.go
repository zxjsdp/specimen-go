package main

import (
	"fmt"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/filetype"
)

func generateResultData() string {
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

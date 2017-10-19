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
	//resultData := generateResultData()
	//fmt.Println(resultData)
	entryDataMatrix := filetype.GetDataMatrix("data.xlsx")
	entryDatas := converters.ToEntryDatas(entryDataMatrix)

	for _, each := range entryDatas {
		fmt.Println(each.String())
	}
	//fmt.Println(entryDatas)
}

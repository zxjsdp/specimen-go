package main

import (
	"fmt"

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
	fmt.Println(filetype.GetDataMatrix("data.xlsx"))
}

package test

import (
	"testing"

	"fmt"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/filetype"
	"github.com/zxjsdp/specimen-go/utils"
)

var testMarkerXlsxFile = "../data/query.xlsx"
var testEntryXlsxFile = "../data/data.xlsx"

func Test_GetDataMatrix(t *testing.T) {
	dataMatrix := filetype.GetDataMatrix(testMarkerXlsxFile)
	if dataMatrix.ColumnCount == 0 || dataMatrix.RowCount == 0 ||
		len(dataMatrix.Matrix) == 0 || len(dataMatrix.Matrix[0]) == 0 {
		t.Error("filetype.GetDataMatrix: failed to read xlsx file!")
	}
}

func Test_ToEntryDatas(t *testing.T) {
	entryDataMatrix := filetype.GetDataMatrix(testEntryXlsxFile)
	entryDatas := converters.ToEntryDatas(entryDataMatrix)
	if len(entryDataMatrix.Matrix) != len(entryDatas) ||
		len(entryDataMatrix.Matrix[0]) != 20 ||
		utils.GetNumberOfField(entryDatas[0]) != 4 {
		t.Error("converters.ToEntryDatas: failed to convert DataMatrix to EntryDatas")
	}
}

func Test_ToMarkerDatas(t *testing.T) {
	markerDataMatrix := filetype.GetDataMatrix(testMarkerXlsxFile)
	markerDatas := converters.ToMarkerDatas(markerDataMatrix)
	if len(markerDataMatrix.Matrix) != len(markerDatas) ||
		len(markerDataMatrix.Matrix[0]) != utils.GetNumberOfField(markerDatas[0]) {
		t.Error("converters.ToMarkerDatas: failed to convert DataMatrix to MarkerDatas")
		fmt.Println(len(markerDataMatrix.Matrix[0]), utils.GetNumberOfField(markerDatas[0]))
	}
}

package test

import (
	"testing"

	"fmt"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
)

const (
	testMarkerXlsxFile = "../data/query.xlsx"
	testEntryXlsxFile  = "../data/data.xlsx"
)

func TestGetDataMatrix(t *testing.T) {
	dataMatrix := files.GetDataMatrix(testMarkerXlsxFile)
	if dataMatrix.ColumnCount == 0 || dataMatrix.RowCount == 0 ||
		len(dataMatrix.Matrix) == 0 || len(dataMatrix.Matrix[0]) == 0 {
		t.Error("files.GetDataMatrix: failed to read xlsx file!")
	}
}

func TestToEntryDataSlice(t *testing.T) {
	entryDataMatrix := files.GetDataMatrix(testEntryXlsxFile)
	entryDataSlice := converters.ToEntryDataSlice(entryDataMatrix)
	if len(entryDataMatrix.Matrix) != len(entryDataSlice) ||
		len(entryDataMatrix.Matrix[0]) != 20 ||
		utils.GetNumberOfField(entryDataSlice[0]) != 4 {
		t.Error("converters.ToEntryDataSlice: failed to convert DataMatrix to EntryDataSlice")
	}
}

func TestToMarkerDataSlice(t *testing.T) {
	markerDataMatrix := files.GetDataMatrix(testMarkerXlsxFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	if len(markerDataMatrix.Matrix) != len(markerDataSlice) ||
		len(markerDataMatrix.Matrix[0]) != utils.GetNumberOfField(markerDataSlice[0]) {
		t.Error("converters.ToMarkerDatas: failed to convert DataMatrix to MarkerDataSlice")
		fmt.Println(len(markerDataMatrix.Matrix[0]), utils.GetNumberOfField(markerDataSlice[0]))
	}
}

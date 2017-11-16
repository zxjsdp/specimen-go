package test

import (
	"log"
	"testing"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
)

const (
	testSnXlsxFile      = "../data/query.xlsx"
	testOfflineXlsxFile = "../data/data.xlsx"
)

func TestGetDataMatrix(t *testing.T) {
	dataMatrix := files.GetDataMatrix(testSnXlsxFile)
	if dataMatrix.ColumnCount == 0 || dataMatrix.RowCount == 0 ||
		len(dataMatrix.Matrix) == 0 || len(dataMatrix.Matrix[0]) == 0 {
		t.Error("files.GetDataMatrix: failed to read xlsx file!")
	}
}

func TestToOfflineDataSlice(t *testing.T) {
	offlineDataMatrix := files.GetDataMatrix(testOfflineXlsxFile)
	offlineDataSlice := converters.ToOfflineDataSlice(offlineDataMatrix)
	if len(offlineDataMatrix.Matrix) != len(offlineDataSlice) ||
		len(offlineDataMatrix.Matrix[0]) != 20 ||
		utils.GetNumberOfField(offlineDataSlice[0]) != 4 {
		t.Error("converters.ToOfflineDataSlice: failed to convert DataMatrix to OfflineDataSlice")
	}
}

func TestToSnDataSlice(t *testing.T) {
	snDataMatrix := files.GetDataMatrix(testSnXlsxFile)
	snDataSlice := converters.ToSnDataSlice(snDataMatrix)
	if len(snDataMatrix.Matrix) != len(snDataSlice) ||
		len(snDataMatrix.Matrix[0]) != utils.GetNumberOfField(snDataSlice[0]) {
		t.Error("converters.ToSnDataSlice: failed to convert DataMatrix to SnDataSlice")
		log.Println(len(snDataMatrix.Matrix[0]), utils.GetNumberOfField(snDataSlice[0]))
	}
}

package test

import (
	"log"
	"testing"

	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/files"
	"github.com/zxjsdp/specimen-go/utils"
)

const (
	testMarkerXlsxFile  = "../data/query.xlsx"
	testOfflineXlsxFile = "../data/data.xlsx"
)

func TestGetDataMatrix(t *testing.T) {
	dataMatrix := files.GetDataMatrix(testMarkerXlsxFile)
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

func TestToMarkerDataSlice(t *testing.T) {
	markerDataMatrix := files.GetDataMatrix(testMarkerXlsxFile)
	markerDataSlice := converters.ToMarkerDataSlice(markerDataMatrix)
	if len(markerDataMatrix.Matrix) != len(markerDataSlice) ||
		len(markerDataMatrix.Matrix[0]) != utils.GetNumberOfField(markerDataSlice[0]) {
		t.Error("converters.ToMarkerDatas: failed to convert DataMatrix to MarkerDataSlice")
		log.Println(len(markerDataMatrix.Matrix[0]), utils.GetNumberOfField(markerDataSlice[0]))
	}
}

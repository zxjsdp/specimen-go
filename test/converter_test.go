package test

import (
	"testing"

	"github.com/zxjsdp/specimen-go/converters"
)

func TestGenerateColumnHeader(t *testing.T) {
	columnHeaderIndex := 40
	expectedResult := "AN"
	actualResult := converters.GenerateColumnHeader(columnHeaderIndex)
	if expectedResult != actualResult {
		t.Error("files.GenerateColumnHeader: result not match!", "expected:", expectedResult, "actual:", actualResult)
	}
}

func TestGenerateAxis(t *testing.T) {
	rowIndex := 5
	columnHeaderIndex := 38
	expectedResult := "AM6"
	actualResult := converters.GenerateAxis(rowIndex, columnHeaderIndex)
	if expectedResult != actualResult {
		t.Error("files.TestGenerateAxis: result not match!", "| expected:", expectedResult, "actual:", actualResult)
	}
}

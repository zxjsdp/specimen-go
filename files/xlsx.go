package files

import (
	"fmt"

	"log"

	"github.com/xuri/excelize"
	"github.com/zxjsdp/specimen-go/converters"
	"github.com/zxjsdp/specimen-go/entities"
)

// 读取 xlsx 文件并获取数据矩阵
func GetDataMatrix(xlsxFileName string) entities.DataMatrix {
	xlsx, err := excelize.OpenFile(xlsxFileName)
	if err != nil {
		panic(fmt.Sprintf("Error reading xlsx file: %s (%s)", xlsxFileName, err))
	}

	activeSheetIndex := xlsx.GetActiveSheetIndex()
	activeSheetName := xlsx.GetSheetName(activeSheetIndex)

	rows := xlsx.GetRows(activeSheetName)
	if len(rows) == 0 {
		panic("Empty sheet! " + activeSheetName)
	}

	matrix := make([][]string, 0)
	rowCount := len(rows)
	columnCount := len(rows[0])

	for _, row := range rows {
		cellsPerRow := make([]string, 0)
		for _, cell := range row {
			cellsPerRow = append(cellsPerRow, cell)
		}
		matrix = append(matrix, cellsPerRow)
	}

	return entities.DataMatrix{FileName: xlsxFileName, Matrix: matrix, RowCount: rowCount, ColumnCount: columnCount}
}

func SaveResultDataToXlsx(xlsxFileName string, resultDataSlice []entities.ResultData) {
	if len(resultDataSlice) == 0 {
		log.Fatal("Cannot save data matrix to xlsx file! Empty data matrix!")
	}

	resultDataMatrix := converters.FromResultDataSlice(resultDataSlice)
	SaveDataMatrixToXlsx(xlsxFileName, resultDataMatrix)
}

func SaveDataMatrixToXlsx(xlsxFileName string, dataMatrix entities.DataMatrix) {
	if dataMatrix.RowCount == 0 {
		log.Fatal("Cannot save data matrix to xlsx file! Empty data matrix!")
	}

	xlsx := excelize.NewFile()
	activeSheetIndex := xlsx.GetActiveSheetIndex()
	activeSheetName := xlsx.GetSheetName(activeSheetIndex)

	// Set value to cells.
	for columnIndex, eachHeader := range dataMatrix.Header {
		xlsx.SetCellValue(activeSheetName, converters.GenerateAxis(0, columnIndex), eachHeader)
	}
	for rowIndex, row := range dataMatrix.Matrix {
		for columnIndex, cell := range row {
			xlsx.SetCellValue(activeSheetName, converters.GenerateAxis(rowIndex+1, columnIndex), cell)
		}
	}

	// Save xlsx file by the given path.
	err := xlsx.SaveAs(xlsxFileName)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("已将结果写入文件：%s\n", xlsxFileName)
	}
}

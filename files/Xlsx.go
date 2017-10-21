package files

import (
	"fmt"

	"log"

	"github.com/xuri/excelize"
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
			//fmt.Printf(cell, "\t")
		}
		matrix = append(matrix, cellsPerRow)
		//fmt.Println()
	}

	return entities.DataMatrix{Matrix: matrix, RowCount: rowCount, ColumnCount: columnCount}
}

func SaveDataMatrix(xlsxFileName string, resultDataSlice []entities.ResultData) {
	if len(resultDataSlice) == 0 {
		log.Fatal("Cannot save data matrix to xlsx file! Empty data matrix!")
	}

	xlsx := excelize.NewFile()

	activeSheetIndex := xlsx.GetActiveSheetIndex()
	activeSheetName := xlsx.GetSheetName(activeSheetIndex)

	// Set value of a cell.
	xlsx.SetCellValue(activeSheetName, "A2", "Hello world.")

	// Save xlsx file by the given path.
	err := xlsx.SaveAs(xlsxFileName)
	if err != nil {
		fmt.Println(err)
	}
}

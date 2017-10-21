package converters

import "fmt"

// 根据 column index 生成 Excel 列标题 letter（A、B、...、AA、AB、...）
func GenerateColumnHeader(columnIndex int) string {
	if columnIndex < 0 {
		return GenerateColumnHeader(0)
	}

	div := columnIndex
	columnHeader := ""
	module := 0
	for div > 0 {
		module = (div - 1) % 26
		columnHeader = string(rune(65+module)) + columnHeader
		div = int((div - module) / 26)
	}
	return columnHeader
}

// 根据 row index 及 column index 生成 Excel axis 信息（A1、B2、...）
func GenerateAxis(rowIndex, columnIndex int) string {
	columnHeaderLetter := GenerateColumnHeader(columnIndex + 1)
	return fmt.Sprintf("%s%d", columnHeaderLetter, rowIndex+1)
}

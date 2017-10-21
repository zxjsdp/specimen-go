package entities

import "fmt"

// 数据矩阵
type DataMatrix struct {
	Matrix      [][]string // 数据矩阵内容
	Header      []string   // 标题行
	RowCount    int        //行数目
	ColumnCount int        //列数目
}

func (d DataMatrix) String() string {
	result := ""
	result += fmt.Sprintf("Total lines: %d, total columns: %d\n", d.RowCount, d.ColumnCount)
	for i, row := range d.Matrix {
		result += fmt.Sprintf("Line %3d | ", i+1)
		for _, cell := range row {
			result += fmt.Sprintf("%v\t", cell)
		}
		result += "\n"
	}

	return result
}

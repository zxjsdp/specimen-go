package utils

import (
	"fmt"

	"github.com/zxjsdp/specimen-go/entities"
)

// 数据校验结果
type ValidationResult struct {
	Result      bool
	FailureInfo []string
	WarningInfo []string
}

func DataValidation(offlineDataMatrix, snDataMatrix entities.DataMatrix) (validationResult ValidationResult) {
	allPass := true
	allFailureInfo := make([]string, 0)
	allWarningInfo := make([]string, 0)

	if offlineDataMatrix.RowCount <= 1 {
		allPass = false
		allFailureInfo = append(allFailureInfo, "“鉴定录入文件” 内容为空！")
	}

	if snDataMatrix.RowCount <= 1 {
		allPass = false
		allFailureInfo = append(allFailureInfo, "“流水号文件” 内容为空！")
	}

	offlineDataValidationResult := checkCellValueExistenceAndDuplication(entities.OfflineDataCellMap, offlineDataMatrix, "SpeciesNumber", false, false)
	if !offlineDataValidationResult.Result {
		allPass = false
	}
	for _, failureInfo := range offlineDataValidationResult.FailureInfo {
		allFailureInfo = append(allFailureInfo, failureInfo)
	}
	for _, warningInfo := range offlineDataValidationResult.WarningInfo {
		allWarningInfo = append(allWarningInfo, warningInfo)
	}

	snDataValidationResult := checkCellValueExistenceAndDuplication(entities.SnDataCellMap, snDataMatrix, "SpeciesNumber", false, true)
	if !snDataValidationResult.Result {
		allPass = false
	}
	for _, failureInfo := range snDataValidationResult.FailureInfo {
		allFailureInfo = append(allFailureInfo, failureInfo)
	}
	for _, warningInfo := range snDataValidationResult.WarningInfo {
		allWarningInfo = append(allWarningInfo, warningInfo)
	}

	return ValidationResult{Result: allPass, FailureInfo: allFailureInfo, WarningInfo: allWarningInfo}
}

func checkCellValueExistenceAndDuplication(cellMap [][]string, dataMatrix entities.DataMatrix, fieldName string, canBeNull bool, canDuplicate bool) (validationResult ValidationResult) {
	validationResult = ValidationResult{true, []string{}, []string{}}

	fieldNameToIndexMap := generateFieldNameToIndexMap(cellMap)
	fieldNameToChineseNameMap := generateFieldNameToChineseNameMap(cellMap)
	targetCellIndex := fieldNameToIndexMap[fieldName]
	chineseName := fieldNameToChineseNameMap[fieldName]

	seen := make(map[string]int, len(cellMap))

	for i, row := range dataMatrix.Matrix {
		if len(row) != len(cellMap) {
			isExtraCellsAllBlankCells := true
			for _, extraCell := range row[len(cellMap):] {
				if len(extraCell) > 0 {
					isExtraCellsAllBlankCells = false
				}
			}
			if !isExtraCellsAllBlankCells {
				validationResult.FailureInfo = append(validationResult.FailureInfo, fmt.Sprintf("数据列数不正确！应该为 %d 列，当前 %d 列。请参考示例数据！（文件：%s，行：%d）", len(cellMap), len(row), dataMatrix.FileName, i+1))
				validationResult.Result = false
			}
		}

		for j, cell := range row {
			if j == targetCellIndex {
				// Check cell value existence
				if len(cell) == 0 {
					if canBeNull {
						validationResult.WarningInfo = append(validationResult.WarningInfo, fmt.Sprintf("警告！Cell 值（%s）不能为空（文件：%s，行：%d，列：%d）", chineseName, dataMatrix.FileName, i+1, j+1))
					} else {
						validationResult.FailureInfo = append(validationResult.FailureInfo, fmt.Sprintf("错误！Cell 值（%s）不能为空（文件：%s，行：%d，列：%d）", chineseName, dataMatrix.FileName, i+1, j+1))
						validationResult.Result = false
					}
				}

				// Check cell value duplication
				if oldRowIndex, ok := seen[cell]; ok {
					if !canDuplicate {
						validationResult.FailureInfo = append(validationResult.FailureInfo, fmt.Sprintf("错误！Cell 值（%s）重复（文件：%s，行：%d，列：%d，已在第 %d 行出现过）", chineseName, dataMatrix.FileName, i+1, j+1, oldRowIndex+1))
						validationResult.Result = false
					}
				}

				seen[cell] = i
			}
		}
	}

	return
}

func generateFieldNameToIndexMap(cellMap [][]string) map[string]int {
	fieldNameToIndexMap := make(map[string]int)
	if len(cellMap) == 0 {
		return fieldNameToIndexMap
	}

	for i, nameTuple := range cellMap {
		fieldNameToIndexMap[nameTuple[0]] = i
	}

	return fieldNameToIndexMap
}

func generateFieldNameToChineseNameMap(cellMap [][]string) map[string]string {
	fieldNameToChineseNameMap := make(map[string]string)
	if len(cellMap) == 0 {
		return fieldNameToChineseNameMap
	}

	for _, nameTuple := range cellMap {
		fieldNameToChineseNameMap[nameTuple[0]] = nameTuple[1]
	}

	return fieldNameToChineseNameMap
}

package converters

import (
	"github.com/zxjsdp/specimen-go/entities"
)

// 由 ResultData slice 生成用于写入输出 Excel 文件的 DataMatrix
func FromResultDataSlice(resultDataSlice []entities.ResultData) entities.DataMatrix {
	dataMatrix := entities.DataMatrix{}
	if len(resultDataSlice) == 0 {
		return dataMatrix
	}

	// 获取标题信息
	headerNameSlice, _ := resultDataSlice[0].ToOrderedResultSlice()
	dataMatrix.Header = headerNameSlice

	// 获取内容信息
	for _, resultData := range resultDataSlice {
		_, orderedResultSlice := resultData.ToOrderedResultSlice()
		dataMatrix.Matrix = append(dataMatrix.Matrix, orderedResultSlice)
	}

	dataMatrix.RowCount = len(dataMatrix.Matrix)
	dataMatrix.ColumnCount = len(dataMatrix.Matrix[0])

	return dataMatrix
}

func FromTwoDimensionalSlice(twoDimensionalData [][]string, dataHeaderMap [][]string) entities.DataMatrix {
	dataMatrix := entities.DataMatrix{}
	if len(twoDimensionalData) == 0 {
		return dataMatrix
	}

	headerSlice := make([]string, 0)
	for _, tuple := range dataHeaderMap {
		headerSlice = append(headerSlice, tuple[1])
	}
	dataMatrix.Header = headerSlice

	for _, dataSlice := range twoDimensionalData {
		dataMatrix.Matrix = append(dataMatrix.Matrix, dataSlice)
	}

	dataMatrix.RowCount = len(twoDimensionalData)
	dataMatrix.ColumnCount = len(twoDimensionalData[0])

	return dataMatrix
}

// DataMatrix 转换至 SnData slice
func ToSnDataSlice(d entities.DataMatrix) []entities.SnData {
	snDataSlice := make([]entities.SnData, 0)
	for _, cells := range d.Matrix {
		snDataSlice = append(snDataSlice, entities.SnData{
			SpeciesNumber: cells[0],
			SerialNumber:  cells[1],
			Barcode:       cells[2],
			CopyNumber:    cells[3],
		})
	}

	return snDataSlice
}

// DataMatrix 转换至 OfflineData slice
func ToOfflineDataSlice(d entities.DataMatrix) []entities.OfflineData {
	offlineDataSlice := make([]entities.OfflineData, 0)
	for _, cells := range d.Matrix {
		speciesNumber := cells[0]
		chineseName := cells[1]
		fullLatinName := cells[2]
		familyChineseName := cells[3]
		familyLatinName := cells[4]
		province := cells[5]
		city := cells[6]
		detailedPlace := cells[7]
		latitude := cells[8]
		longitude := cells[9]
		altitude := cells[10]
		collectingDate := cells[11]
		inventory := cells[12]
		habit := cells[13]
		collector := cells[14]
		identifier := cells[15]
		identifyDate := cells[16]
		recordingPerson := cells[17]
		recordingDate := cells[18]

		offlineSpecimenMetaInfo := entities.OfflineSpecimenMetaInfo{SpeciesNumber: speciesNumber, Inventory: inventory}
		offlineCollectingInfo := entities.OfflineCollectingInfo{
			Province:       province,
			City:           city,
			DetailedPlace:  detailedPlace,
			Latitude:       latitude,
			Longitude:      longitude,
			Altitude:       altitude,
			CollectingDate: collectingDate,
			Collector:      collector,
		}
		offlineIdentificationInfo := entities.OfflineIdentificationInfo{
			Identifier:        identifier,
			IdentifyDate:      identifyDate,
			FullLatinName:     fullLatinName,
			ChineseName:       chineseName,
			FamilyChineseName: familyChineseName,
			FamilyLatinName:   familyLatinName,
			Habit:             habit,
		}

		offlineRecordingInfo := entities.OfflineRecordingInfo{
			RecordingPerson: recordingPerson,
			RecordingDate:   recordingDate,
		}

		offlineData := entities.OfflineData{
			OfflineSpecimenMetaInfo:   offlineSpecimenMetaInfo,
			OfflineCollectingInfo:     offlineCollectingInfo,
			OfflineIdentificationInfo: offlineIdentificationInfo,
			OfflineRecordingInfo:      offlineRecordingInfo,
		}

		offlineDataSlice = append(offlineDataSlice, offlineData)
	}

	return offlineDataSlice
}

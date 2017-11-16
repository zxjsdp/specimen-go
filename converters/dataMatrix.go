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

// DataMatrix 转换至 MarkerData slice
func ToMarkerDataSlice(d entities.DataMatrix) []entities.MarkerData {
	markerDataSlice := make([]entities.MarkerData, 0)
	for _, cells := range d.Matrix {
		markerDataSlice = append(markerDataSlice, entities.MarkerData{
			SpeciesNumber: cells[0],
			SerialNumber:  cells[1],
			Barcode:       cells[2],
			FullLatinName: cells[3],
			CopyNumber:    cells[4],
		})
	}

	return markerDataSlice
}

// DataMatrix 转换至 EntityData slice
func ToEntryDataSlice(d entities.DataMatrix) []entities.EntryData {
	entryDataSlice := make([]entities.EntryData, 0)
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

		entrySpecimenMetaInfo := entities.EntrySpecimenMetaInfo{SpeciesNumber: speciesNumber, Inventory: inventory}
		entryCollectingInfo := entities.EntryCollectingInfo{
			Province:       province,
			City:           city,
			DetailedPlace:  detailedPlace,
			Latitude:       latitude,
			Longitude:      longitude,
			Altitude:       altitude,
			CollectingDate: collectingDate,
			Collector:      collector,
		}
		entryIdentificationInfo := entities.EntryIdentificationInfo{
			Identifier:        identifier,
			IdentifyDate:      identifyDate,
			FullLatinName:     fullLatinName,
			ChineseName:       chineseName,
			FamilyChineseName: familyChineseName,
			FamilyLatinName:   familyLatinName,
			Habit:             habit,
		}

		entryRecordingInfo := entities.EntryRecordingInfo{
			RecordingPerson: recordingPerson,
			RecordingDate:   recordingDate,
		}

		entryData := entities.EntryData{
			EntrySpecimenMetaInfo:   entrySpecimenMetaInfo,
			EntryCollectingInfo:     entryCollectingInfo,
			EntryIdentificationInfo: entryIdentificationInfo,
			EntryRecordingInfo:      entryRecordingInfo,
		}

		entryDataSlice = append(entryDataSlice, entryData)
	}

	return entryDataSlice
}

package converters

import "github.com/zxjsdp/specimen-go/entities"

func ToMarkerDataSlice(d entities.DataMatrix) []entities.MarkerData {
	markerDataSlice := make([]entities.MarkerData, 0)
	for _, cells := range d.Matrix {
		speciesNumber := cells[0]
		fullLatinName := cells[1]
		serialNumber := cells[2]
		barcode := cells[3]
		copyNumber := cells[4]

		markerDataSlice = append(markerDataSlice, entities.MarkerData{
			speciesNumber,
			fullLatinName,
			serialNumber,
			barcode,
			copyNumber,
		})
	}

	return markerDataSlice
}

func ToEntryDataSlice(d entities.DataMatrix) []entities.EntryData {
	entryDataSlice := make([]entities.EntryData, 0)
	for _, cells := range d.Matrix {
		speciesNumber := cells[0]
		fullLatinName := cells[1]
		chineseName := cells[2]
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
		plantType := cells[13]
		collector := cells[14]
		identifier := cells[15]
		identifyDate := cells[16]
		recorder := cells[17]
		recordingDate := cells[18]

		entrySpecimenMetaInfo := entities.EntrySpecimenMetaInfo{speciesNumber, inventory}
		entryCollectingInfo := entities.EntryCollectingInfo{
			province,
			city,
			detailedPlace,
			latitude,
			longitude,
			altitude,
			collectingDate,
			collector,
		}
		entryIdentificationInfo := entities.EntryIdentificationInfo{
			identifier,
			identifyDate,
			fullLatinName,
			chineseName,
			familyChineseName,
			familyLatinName,
			plantType,
		}

		entryRecordingInfo := entities.EntryRecordingInfo{
			recorder,
			recordingDate,
		}

		entryData := entities.EntryData{
			entrySpecimenMetaInfo,
			entryCollectingInfo,
			entryIdentificationInfo,
			entryRecordingInfo,
		}

		entryDataSlice = append(entryDataSlice, entryData)
	}

	return entryDataSlice
}

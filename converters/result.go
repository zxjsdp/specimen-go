package converters

import (
	"fmt"
	"log"

	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/utils"
)

func ToResultData(
	marker entities.MarkerData,
	entryDataMap map[string]entities.EntryData,
	webInfoMap map[string]entities.WebInfo) entities.ResultData {

	var resultData entities.ResultData
	if entry, ok := entryDataMap[marker.SpeciesNumber]; ok {
		specimenMetaInfo := entities.SpecimenMetaInfo{
			LibraryCode:       config.LibraryCode,
			SerialNumber:      marker.SerialNumber,
			Barcode:           marker.Barcode,
			PatternType:       config.PatternType,
			SpecimenCondition: config.SpecimenCondition,
			Inventory:         config.Inventory,
		}

		collectingInfo := entities.CollectingInfo{
			Collector:        entry.Collector,
			CollectingNumber: entry.SpeciesNumber + "-" + marker.CopyNumber,
			CollectingDate:   entry.CollectingDate,
			Country:          config.Country,
			ProvinceAndCity:  entry.Province + "，" + entry.City,
			District:         config.District,
			Altitude:         entry.Altitude,
			NegativeAltitude: config.NegativeAltitude,
			DetailedPlace:    entry.DetailedPlace,
			Habitat:          config.DefaultHabitat,
			Longitude:        entry.Longitude,
			Latitude:         entry.Latitude,
			Remarks2:         config.Remarks2,
		}

		latinName := utils.ParseLatinName(marker.FullLatinName)
		identificationInfo := entities.IdentificationInfo{
			Family:       latinName.LatinNameString,
			Genus:        latinName.Genus,
			Species:      latinName.Species,
			NameGiver:    config.DefaultNameGiver,
			Level:        config.Level,
			ChineseName:  entry.ChineseName,
			Habit:        entry.Habit,
			Identifier:   entry.Identifier,
			IdentifyDate: entry.IdentifyDate,
			Remarks:      config.Remarks,
		}

		recordingInfo := entities.RecordingInfo{
			RecordingPerson: entry.RecordingPerson,
			RecordingDate:   entry.RecordingDate,
		}

		morphologyInfo := entities.Morphology{}

		// 若从网络上获取到了相关信息，则替换相应字段为网络信息
		if webInfo, ok := webInfoMap[marker.FullLatinName]; ok {
			collectingInfo.Habitat = webInfo.Habitat
			identificationInfo.NameGiver = webInfo.NameGiver

			morphologyInfo.BodyHeight = webInfo.BodyHeight
			morphologyInfo.DBH = webInfo.DBH
			morphologyInfo.Stem = webInfo.Stem
			morphologyInfo.Leaf = webInfo.Leaf
			morphologyInfo.Flower = webInfo.Flower
			morphologyInfo.Fruit = webInfo.Fruit
			morphologyInfo.Host = webInfo.Host
		}

		resultData = entities.ResultData{
			SpecimenMetaInfo:   specimenMetaInfo,
			CollectingInfo:     collectingInfo,
			IdentificationInfo: identificationInfo,
			RecordingInfo:      recordingInfo,
			Morphology:         morphologyInfo,
		}
	} else {
		log.Fatal(fmt.Sprintf("Entry 文件中缺失物种编号：%s", marker.SerialNumber))
	}
	return resultData
}

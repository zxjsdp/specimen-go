package converters

import (
	"fmt"
	"log"

	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/entities"
)

func ToResultData(
	marker entities.MarkerData,
	entryDataMap map[string]entities.EntryData,
	webInfoMap map[string]entities.WebInfo) entities.ResultData {

	var resultData entities.ResultData
	if entry, ok := entryDataMap[marker.SpeciesNumber]; ok {
		specimenMetaInfo := entities.SpecimenMetaInfo{
			config.LibraryCode,
			marker.SerialNumber,
			marker.Barcode,
			config.PatternType,
			config.SpecimenCondition,
			config.Inventory,
		}

		collectingInfo := entities.CollectingInfo{
			entry.Collector,
			entry.SpeciesNumber + "-" + marker.CopyNumber,
			entry.CollectingDate,
			config.Country,
			entry.Province + "，" + entry.City,
			config.District,
			entry.Altitude,
			config.NegativeAltitude,
			entry.DetailedPlace,
			config.DefaultHabitat,
			entry.Longitude,
			entry.Latitude,
			config.Remarks2,
		}

		identificationInfo := entities.IdentificationInfo{
			entry.FamilyLatinName,
			entry.FullLatinName,
			entry.FullLatinName,
			config.DefaultNameGiver,
			config.Level,
			entry.ChineseName,
			entry.Habit,
			entry.Identifier,
			entry.IdentifyDate,
			config.Remarks,
		}

		recordingInfo := entities.RecordingInfo{
			entry.RecordingPerson,
			entry.RecordingDate,
		}

		morphologyInfo := entities.Morphology{}

		// 若从网络上获取到了相关信息，则替换相应字段为网络信息
		log.Println(marker.FullLatinName)
		if webInfo, ok := webInfoMap[marker.FullLatinName]; ok {
			log.Println(marker.FullLatinName, ok)
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
			specimenMetaInfo,
			collectingInfo,
			identificationInfo,
			recordingInfo,
			morphologyInfo,
		}
	} else {
		log.Fatal(fmt.Sprintf("Entry 文件中缺失物种编号：%s", marker.SerialNumber))
	}
	return resultData
}

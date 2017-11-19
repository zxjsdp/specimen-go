package converters

import (
	"fmt"
	"log"

	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/utils"
)

func ToResultData(
	snData entities.SnData,
	offlineDataMap map[string]entities.OfflineData,
	webInfoMap map[string]entities.WebInfo) entities.ResultData {

	var resultData entities.ResultData
	if offlineData, ok := offlineDataMap[snData.SpeciesNumber]; ok {
		specimenMetaInfo := entities.SpecimenMetaInfo{
			LibraryCode:       config.LibraryCode,
			SerialNumber:      snData.SerialNumber,
			Barcode:           snData.Barcode,
			PatternType:       config.PatternType,
			SpecimenCondition: config.SpecimenCondition,
			Inventory:         config.Inventory,
		}

		collectingInfo := entities.CollectingInfo{
			Collector:        offlineData.Collector,
			CollectingNumber: offlineData.SpeciesNumber + "-" + snData.CopyNumber,
			CollectingDate:   offlineData.CollectingDate,
			Country:          config.Country,
			ProvinceAndCity:  offlineData.Province + "，" + offlineData.City,
			District:         config.District,
			Altitude:         offlineData.Altitude,
			NegativeAltitude: config.NegativeAltitude,
			DetailedPlace:    offlineData.DetailedPlace,
			Habitat:          config.DefaultHabitat,
			Longitude:        offlineData.Longitude,
			Latitude:         offlineData.Latitude,
			Remarks2:         config.Remarks2,
		}

		latinName := utils.ParseLatinName(offlineData.FullLatinName)
		identificationInfo := entities.IdentificationInfo{
			Family:        "", // TODO, 从网络获取到 “科” 的信息
			Genus:         latinName.Genus,
			Species:       latinName.Species,
			NamePublisher: config.DefaultNamePublisher,
			Level:         config.Level,
			ChineseName:   offlineData.ChineseName,
			Habit:         offlineData.Habit,
			Identifier:    offlineData.Identifier,
			IdentifyDate:  offlineData.IdentifyDate,
			Remarks:       config.Remarks,
		}

		recordingInfo := entities.RecordingInfo{
			RecordingPerson: offlineData.RecordingPerson,
			RecordingDate:   offlineData.RecordingDate,
		}

		morphologyInfo := entities.Morphology{}

		// 若从网络上获取到了相关信息，则替换相应字段为网络信息
		if webInfo, ok := webInfoMap[offlineData.FullLatinName]; ok {
			collectingInfo.Habitat = webInfo.Habitat
			identificationInfo.NamePublisher = webInfo.NamePublisher

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
		log.Fatal(fmt.Sprintf("“鉴定及录入文件” 中缺失物种编号：%s", snData.SerialNumber))
	}
	return resultData
}

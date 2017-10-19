package entities

import (
	"fmt"
	"reflect"
)

//// 个体录入信息
//type EntryData struct {
//	SpeciesNumber     string // 物种编号
//	FullLatinName     string // 种名（拉丁名）
//	ChineseName       string // 中文名
//	FamilyChineseName string // 科名（中文）
//	FamilyLatinName   string // 科名（拉丁名）
//	Province          string // 省
//	City              string // 市
//	DetailedPlace     string // 具体小地名
//	Latitude          string // 纬度
//	Longitude         string // 经度
//	Altitude          string // 海拔
//	CollectingDate    string // 采集日期
//	Copies            int32  // 份数
//	PlantType         string // 草灌
//	Collector         string // 采集人
//	Identifier        string // 鉴定人
//	IdentifyDate      string // 鉴定日期
//	Recorder          string // 录入人
//	RecordingDate     string // 录入日期
//}

// Entry 数据
type EntryData struct {
	EntrySpecimenMetaInfo   // Entry 标本基础信息
	EntryCollectingInfo     // Entry 采集信息
	EntryIdentificationInfo // Entry 鉴定信息
	EntryRecordingInfo      // Entry 录入信息
}

func (e EntryData) String() string {
	output := ""
	v := reflect.ValueOf(e)
	for i := 0; i < v.NumField(); i++ {
		output += fmt.Sprintf("%v\t", v.Field(i).Interface())
	}
	return output
}

// Entry 标本基础信息
type EntrySpecimenMetaInfo struct {
	SpeciesNumber string // 物种编号
	Copies        string // 份数
}

// Entry 采集信息
type EntryCollectingInfo struct {
	Province       string // 省
	City           string // 市
	DetailedPlace  string // 具体小地名
	Latitude       string // 纬度
	Longitude      string // 经度
	Altitude       string // 海拔
	CollectingDate string // 采集日期
	Collector      string // 采集人
}

// Entry 鉴定信息
type EntryIdentificationInfo struct {
	Identifier        string // 鉴定人
	IdentifyDate      string // 鉴定日期
	FullLatinName     string // 种名（拉丁名）
	ChineseName       string // 中文名
	FamilyChineseName string // 科名（中文）
	FamilyLatinName   string // 科名（拉丁名）
	PlantType         string // 习性（草灌）
}

// Entry 录入信息
type EntryRecordingInfo struct {
	Recorder      string // 录入人
	RecordingDate string // 录入日期
}

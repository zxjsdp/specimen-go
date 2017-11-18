package entities

import (
	"fmt"
	"reflect"
)

// 标本鉴定数据 map
var OfflineDataCellMap = [][]string{
	{"SpeciesNumber", "物种编号"},
	{"ChineseName", "中文名"},
	{"FullLatinName", "种名（拉丁名）"},
	{"FamilyChineseName", "科名（中文）"},
	{"FamilyLatinName", "科名（拉丁名）"},
	{"Province", "省"},
	{"City", "市"},
	{"DetailedPlace", "具体小地名"},
	{"Latitude", "纬度"},
	{"Longitude", "经度"},
	{"Altitude", "海拔"},
	{"CollectingDate", "采集日期"},
	{"Inventory", "库存"},
	{"Habit", "习性（草灌）"},
	{"Collector", "采集人"},
	{"Identifier", "鉴定人"},
	{"IdentifyDate", "鉴定日期"},
	{"RecordingPerson", "录入人"},
	{"RecordingDate", "录入日期"},
}

type OfflineDataModel struct {
	SpeciesNumber     string // 物种编号
	ChineseName       string // 中文名
	FullLatinName     string // 种名（拉丁名）
	FamilyChineseName string // 科名（中文）
	FamilyLatinName   string // 科名（拉丁名）
	Province          string // 省
	City              string // 市
	DetailedPlace     string // 具体小地名
	Latitude          string // 纬度
	Longitude         string // 经度
	Altitude          string // 海拔
	CollectingDate    string // 采集日期
	Inventory         string // 库存
	Habit             string // 习性（草灌）
	Collector         string // 采集人
	Identifier        string // 鉴定人
	IdentifyDate      string // 鉴定日期
	RecordingPerson   string // 录入人
	RecordingDate     string // 录入日期
}

// 标本鉴定数据
type OfflineData struct {
	OfflineSpecimenMetaInfo   // Offline 标本基础信息
	OfflineCollectingInfo     // Offline 标本采集信息
	OfflineIdentificationInfo // Offline 标本鉴定信息
	OfflineRecordingInfo      // Offline 标本录入信息
}

func (e OfflineData) String() string {
	output := ""
	v := reflect.ValueOf(e)
	for i := 0; i < v.NumField(); i++ {
		output += fmt.Sprintf("%v\t", v.Field(i).Interface())
	}
	return output
}

// Offline 标本基础信息
type OfflineSpecimenMetaInfo struct {
	SpeciesNumber string // 物种编号
	Inventory     string // 库存
}

// Offline 采集信息
type OfflineCollectingInfo struct {
	Province       string // 省
	City           string // 市
	DetailedPlace  string // 具体小地名
	Latitude       string // 纬度
	Longitude      string // 经度
	Altitude       string // 海拔
	CollectingDate string // 采集日期
	Collector      string // 采集人
}

// Offline 鉴定信息
type OfflineIdentificationInfo struct {
	Identifier        string // 鉴定人
	IdentifyDate      string // 鉴定日期
	FullLatinName     string // 种名（拉丁名）
	ChineseName       string // 中文名
	FamilyChineseName string // 科名（中文）
	FamilyLatinName   string // 科名（拉丁名）
	Habit             string // 习性（草灌）
}

// Offline 录入信息
type OfflineRecordingInfo struct {
	RecordingPerson string // 录入人
	RecordingDate   string // 录入日期
}

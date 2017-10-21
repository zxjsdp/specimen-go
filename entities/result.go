package entities

// 最终结果
type ResultData struct {
	SpecimenMetaInfo   // 标本基础信息
	CollectingInfo     // 采集信息
	IdentificationInfo // 鉴定信息
	RecordingInfo      // 录入信息
	Morphology         // 植物形态
}

// 将结果以输出文件中的顺序排序
func (rd ResultData) ToOrderedResultSlice() []string {
	orderedResultSlice := []string{
		rd.LibraryCode,       // 馆代码
		rd.SerialNumber,      // 流水号
		rd.Barcode,           // 条形码
		rd.PatternType,       // 模式类型
		rd.Inventory,         // 库存
		rd.SpecimenCondition, // 标本状态
		rd.Collector,         // 采集人
		rd.CollectingNumber,  // 采集号
		rd.CollectingDate,    // 采集日期
		rd.Country,           // 国家
		rd.ProvinceAndCity,   // 省市
		rd.District,          // 区县
		rd.Altitude,          // 海拔
		rd.NegativeAltitude,  // 负海拔
		rd.Family,            // 科
		rd.Genus,             // 属
		rd.Species,           // 种
		rd.NameGiver,         // 定名人
		rd.Level,             // 种下等级
		rd.ChineseName,       // 中文名
		rd.Identifier,        // 鉴定人
		rd.IdentifyDate,      // 鉴定日期
		rd.Remarks,           // 备注
		rd.DetailedPlace,     // 地名
		rd.Habitat,           // 生境
		rd.Longitude,         // 经度
		rd.Latitude,          // 纬度
		rd.Remarks2,          // 备注2
		rd.RecordingPerson,   // 录入员
		rd.RecordingDate,     // 录入日期
		rd.Habit,             // 习性（草灌）
		rd.BodyHeight,        // 体高
		rd.DBH,               // 胸径
		rd.Stem,              // 茎
		rd.Leaf,              // 叶
		rd.Flower,            // 花
		rd.Fruit,             // 果实
		rd.Host,              // 寄主
	}

	return orderedResultSlice
}

// 标本基础信息
type SpecimenMetaInfo struct {
	LibraryCode       string // 馆代码
	SerialNumber      string // 流水号
	Barcode           string // 条形码
	PatternType       string // 模式类型
	SpecimenCondition string // 标本状态
	Inventory         string // 库存
}

// 采集信息
type CollectingInfo struct {
	Collector        string // 采集人
	CollectingNumber string // 采集号
	CollectingDate   string // 采集日期
	Country          string // 国家
	ProvinceAndCity  string // 省市
	District         string // 区县
	Altitude         string // 海拔
	NegativeAltitude string // 负海拔
	DetailedPlace    string // 地名
	Habitat          string // 生境
	Longitude        string // 经度
	Latitude         string // 纬度
	Remarks2         string // 备注2
}

// 鉴定信息
type IdentificationInfo struct {
	Family       string // 科
	Genus        string // 属
	Species      string // 种
	NameGiver    string // 定名人
	Level        string // 种下等级
	ChineseName  string // 中文名
	Habit        string // 习性（草灌）
	Identifier   string // 鉴定人
	IdentifyDate string // 鉴定日期
	Remarks      string // 备注
}

// 录入信息
type RecordingInfo struct {
	RecordingPerson string // 录入员
	RecordingDate   string // 录入日期
}

// 植物形态
type Morphology struct {
	BodyHeight string // 体高
	DBH        string // 胸径
	Stem       string // 茎
	Leaf       string // 叶
	Flower     string // 花
	Fruit      string // 果实
	Host       string // 寄主
}

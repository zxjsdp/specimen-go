package entities

// 最终结果
type ResultData struct {
	SpecimenMetaInfo   // 标本基础信息
	CollectingInfo     // 采集信息
	IdentificationInfo // 鉴定信息
	RecordingInfo      // 录入信息
	Morphology         // 植物形态
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

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
func (r ResultData) ToOrderedResultSlice() ([]string, []string) {
	orderedResultPairs := [][]string{
		{r.LibraryCode, "馆代码"},
		{r.SerialNumber, " 流水号"},
		{r.Barcode, " 条形码"},
		{r.PatternType, " 模式类型"},
		{r.Inventory, " 库存"},
		{r.SpecimenCondition, " 标本状态"},
		{r.Collector, " 采集人"},
		{r.CollectingNumber, " 采集号"},
		{r.CollectingDate, " 采集日期"},
		{r.Country, " 国家"},
		{r.ProvinceAndCity, " 省市"},
		{r.District, " 区县"},
		{r.Altitude, " 海拔"},
		{r.NegativeAltitude, " 负海拔"},
		{r.Family, " 科"},
		{r.Genus, " 属"},
		{r.Species, " 种"},
		{r.NamePublisher, " 定名人"},
		{r.Level, " 种下等级"},
		{r.ChineseName, " 中文名"},
		{r.Identifier, " 鉴定人"},
		{r.IdentifyDate, " 鉴定日期"},
		{r.Remarks, " 备注"},
		{r.DetailedPlace, " 地名"},
		{r.Habitat, " 生境"},
		{r.Longitude, " 经度"},
		{r.Latitude, " 纬度"},
		{r.Remarks2, " 备注2"},
		{r.RecordingPerson, " 录入员"},
		{r.RecordingDate, " 录入日期"},
		{r.Habit, " 习性（草灌）"},
		{r.BodyHeight, " 体高"},
		{r.DBH, " 胸径"},
		{r.Stem, " 茎"},
		{r.Leaf, " 叶"},
		{r.Flower, " 花"},
		{r.Fruit, " 果实"},
		{r.Host, " 寄主"},
	}

	headerNamSlice := make([]string, len(orderedResultPairs))
	orderedResultSlice := make([]string, len(orderedResultPairs))
	for i, pair := range orderedResultPairs {
		headerNamSlice[i] = pair[1]
		orderedResultSlice[i] = pair[0]
	}

	return headerNamSlice, orderedResultSlice
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
	Family        string // 科
	Genus         string // 属
	Species       string // 种
	NamePublisher string // 定名人
	Level         string // 种下等级
	ChineseName   string // 中文名
	Habit         string // 习性（草灌）
	Identifier    string // 鉴定人
	IdentifyDate  string // 鉴定日期
	Remarks       string // 备注
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

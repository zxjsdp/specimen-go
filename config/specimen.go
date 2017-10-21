package config

import "regexp"

// Const: specimen info
const (
	LibraryCode       = "FUS" // 馆代码
	Country           = "中国"  // 国家
	PatternType       = ""    // 模式类型
	SpecimenCondition = ""    // 标本状态
	Inventory         = ""    // 库存
	District          = ""    //区

	NegativeAltitude = "" //
	Level            = ""
	Remarks          = ""
	Remarks2         = ""
)

// Const: default placeholder value
const (
	DefaultHabitat   = ""
	DefaultNameGiver = ""

	DefaultSeparator = " | "
)

// Const: io
const (
	DefaultResultXlsxName = "output.xlsx" // 默认输出 xlsx 文件名称
)

var (
	BodyHeightRegexp, _ = regexp.Compile("[^，。；]*高[^，。；]]*")   // 体高
	DBHRegexp, _        = regexp.Compile("[^，。；]]*胸径[^，。；]]*") // 胸径
	StemRegexp, _       = regexp.Compile("[^。；]]*茎[^。；]]*")    // 茎
	LeafRegexp, _       = regexp.Compile("[^。；]]*叶[^。；]]*")    // 叶
	FlowerRegexp, _     = regexp.Compile("[^。；]]*花[^。；]]*")    // 花
	FruitRegexp, _      = regexp.Compile("[^。；]]*果[^。；]]*")    // 果实
	HostRegexp, _       = regexp.Compile("[^。；]]*寄主[^。；]]*")   // 寄主

	NameGiverRegexpTemplate = "(?<=<b>%s</b> <b>%s</b>)[^><]*(?=<span)" // 命名人
)

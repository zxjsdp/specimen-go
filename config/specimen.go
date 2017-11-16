package config

const Version = "V1.3.0" // 请勿手动修改， 应使用 bumpversion 自动更新

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
	DefaultHabitat       = ""
	DefaultNamePublisher = ""

	DefaultSeparator = "；"
)

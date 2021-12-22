package config

import (
	"fmt"
)

const (
	Version = "v1.9.4" // 请勿手动修改， 应使用 bumpversion 自动更新
	Title   = "植物标本录入软件"
)

var About = fmt.Sprintf("%s\nspecimen-go GUI %s by zxjsdp\n复旦大学生科院 G417 实验室", Title, Version)

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

const (
	AmbiguousSpeciesName = "sp."
)

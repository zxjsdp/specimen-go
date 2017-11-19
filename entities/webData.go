package entities

// 网络信息
type WebInfo struct {
	FullLatinName string // 拉丁名（二名法）
	Morphology           // 植物形态
	NamePublisher string // 定名人
	Family        string
	Habitat       string //习性
}

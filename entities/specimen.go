package entities

// 拉丁名
type LatinName struct {
	LatinNameString string   // 拉丁名全名
	Genus           string   // 属名
	Species         string   // 种名
	Elements        []string // 拉丁名中所有部分
}

func (l LatinName) String() string {
	return l.LatinNameString
}

package entities

import "fmt"

// 拉丁名
type LatinName struct {
	LatinNameString string   // 拉丁名全名
	Genus           string   // 属名
	Species         string   // 种名
	Elements        []string // 拉丁名中所有部分
}

func (l LatinName) String() string {
	return fmt.Sprintf("%s (genus: %s, species: %s)", l.LatinNameString, l.Genus, l.Species)
}

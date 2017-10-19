package entities

// 物种信息
type SpeciesData struct {
	SpeciesNumber string // 物种编号
	FullLatinName string // 物种名（拉丁名，二名法）
	SerialNumber  string // 流水号
	Barcode       string // 条形码
	CopyNumber    string // 同一物种的个体份数编号（1、2、3、...）
}

package entities

// 物种信息
type MarkerData struct {
	SpeciesNumber string // 物种编号
	SerialNumber  string // 流水号
	Barcode       string // 条形码
	FullLatinName string // 物种名（拉丁名，二名法）
	CopyNumber    string // 同一物种的个体编号（1、2、3、...）
}

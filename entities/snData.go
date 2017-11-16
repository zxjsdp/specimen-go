package entities

var SnDataCellMap = [][]string{
	{"SpeciesNumber", "物种编号"},
	{"SerialNumber", "流水号"},
	{"Barcode", "条形码"},
	{"CopyNumber", "同一物种的个体编号（1、2、3、...）"},
}

// 物种信息
type SnData struct {
	SpeciesNumber string // 物种编号
	SerialNumber  string // 流水号
	Barcode       string // 条形码
	CopyNumber    string // 同一物种的个体编号（1、2、3、...）
}

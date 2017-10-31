package config

import (
	"regexp"
)

// web
const (
	// 中国植物志 URL
	URLPrefixEFLORA = "http://frps.eflora.cn/frps/"

	// URL 分隔符
	URLBlankSeparator = "%20"
)

// goroutine
const (
	WorkerPoolSize = 40 // worker poll size
)

// filtering keywords
var (
	AllKeywords      = []string{"高", "茎", "叶", "花", "果"}
	ModerateKeywords = []string{"叶", "花"}
	RelaxedKeywords  = []string{"茎", "叶"}
)

var (
	BodyHeightRegexp, _ = regexp.Compile("[^，。；]*高[^，。；]*")  // 体高
	DBHRegexp, _        = regexp.Compile("[^，。；]*胸径[^，。；]*") // 胸径
	StemRegexp, _       = regexp.Compile("[^。；]*茎[^。；]*")    // 茎
	LeafRegexp, _       = regexp.Compile("[^。；]*叶[^。；]*")    // 叶
	FlowerRegexp, _     = regexp.Compile("[^。；]*花[^。；]*")    // 花
	FruitRegexp, _      = regexp.Compile("[^。；]*果[^。；]*")    // 果实
	HostRegexp, _       = regexp.Compile("[^。；]*寄主[^。；]*")   // 寄主

	SpeciesInfoDiv          = "div#spinfoDiv" // 物种信息 div id: spinfoDiv
	NameGiverRegexpTemplate = "%s %s(.*)"     // 命名人
)

package config

import "regexp"

// goroutine
const (
	WorkerPoolSize = 30 // worker poll size
)

var (
	// filtering keywords
	AllKeywords      = []string{"高", "茎", "叶", "花", "果"}
	ModerateKeywords = []string{"叶", "花"}
	RelaxedKeywords  = []string{"茎", "叶"}
)

var (
	// filtering regexp
	BodyHeightRegexp, _ = regexp.Compile("[^，。；]*高[^，。；]*")  // 体高
	DBHRegexp, _        = regexp.Compile("[^，。；]*胸径[^，。；]*") // 胸径
	StemRegexp, _       = regexp.Compile("[^。；]*茎[^。；]*")    // 茎
	LeafRegexp, _       = regexp.Compile("[^。；]*叶[^。；]*")    // 叶
	FlowerRegexp, _     = regexp.Compile("[^。；]*花[^。；]*")    // 花
	FruitRegexp, _      = regexp.Compile("[^。；]*果[^。；]*")    // 果实
	HostRegexp, _       = regexp.Compile("[^。；]*寄主[^。；]*")   // 寄主
)

// web
const (
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 中国植物志静态页面 URL 信息（可提取到 spno）
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 中国植物志 URL 前缀
	URLPrefixEFLORA = "http://www.iplant.cn/info/"

	// 中国植物志 URL 后缀
	URLSuffix = "?t=z"

	// URL 分隔符
	URLBlankSeparator = "%20"

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 「详细描述」API
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 「详细描述」API URL 前缀
	//     - frpsspno: 可用于查询「纲目科属种」信息的 spno
	//     - frpsspclassid: 可用于查询「纲目科属种」信息的 spclassid
	//     - frpsdesc: 包含多个 <p> 的详细信息，可用于提取「茎」、「叶」、「花」、「果」等信息
	// 示例: http://www.iplant.cn/ashx/getfrps.ashx?key=Cephalotaxus+fortunei
	DetailedDescriptionApiURLPrefix = "http://www.iplant.cn/ashx/getfrps.ashx?key="
	FrpsspnoKeyInResponseMap        = "frpsspno"
	FrpsspclassidKeyInResponseMap   = "frpsspclassid"
	FrpsdescKeyInResponseMap        = "frpsdesc"

	// API URL 分隔符
	APIURLBlankSeparator = "+"

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 「物种分类（Texomony，界门纲目科属种）」信息查询 API
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 「物种分类（Texomony，界门纲目科属种）」信息查询 API URL
	DetailedCategoryApiUrl       = "http://www.iplant.cn/ashx/getfrpsclass.ashx"
	FrpsclasstxtKeyInResponseMap = "frpsclasstxt"
	PhylumKeyword                = "门 "
	ClassKeyword                 = "纲 "
	OrderKeyword                 = "目 "
	FamilyKeyword                = "科 "
	GenusKeyword                 = "属 "

	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 「拉丁名」信息查询 API，可解析命名人（NamePublisher）信息
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 「拉丁名」信息查询 API URL
	LatinApiUrl = "http://www.iplant.cn/ashx/getsplatin.ashx"

	// 命名人信息 API URL
	SpnoRegexpTemplate = "var spno = \"\\d+\"" // spno

	// 提取命名人的字符串分隔符
	// 示例: <span class='font20'><b>Cephalotaxus</b></span> <span class='font20'><b>fortunei</b></span> Hooker
	// 此 case 中需要 split 后取: Hooker
	NamePublisherSeparator = "</span>"
)

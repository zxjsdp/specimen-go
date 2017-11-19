package web

import (
	"strings"

	"fmt"

	"log"

	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/utils"
)

// （goroutine）根据拉丁名获取网络信息 map
func GenerateWebInfoMap(latinNames []string) map[string]entities.WebInfo {
	webInfoMap := make(map[string]entities.WebInfo)
	latinNames = utils.RemoveDuplicates(latinNames)
	size := len(latinNames)

	jobChannel := make(chan string, size)
	resultWebInfoChannel := make(chan entities.WebInfo, size)

	for i := 1; i <= config.WorkerPoolSize; i++ {
		go worker(i, jobChannel, resultWebInfoChannel)
	}

	for _, latinName := range latinNames {
		jobChannel <- latinName
	}

	close(jobChannel)

	for i := 1; i <= size; i++ {
		webInfo := <-resultWebInfoChannel
		webInfoMap[webInfo.FullLatinName] = webInfo
	}

	return webInfoMap
}

func worker(id int, jobs <-chan string, results chan<- entities.WebInfo) {
	for latinName := range jobs {
		//log.Println("worker", id, "started job", latinName)
		webInfo := GenerateWebInfo(latinName)
		//log.Println("worker", id, "finished job", latinName)
		results <- webInfo
	}
}

// （同步方法）根据拉丁名获取网络信息 map
func GenerateWebInfoMapSync(latinNames []string) map[string]entities.WebInfo {
	webInfoMap := make(map[string]entities.WebInfo)

	latinNames = utils.RemoveDuplicates(latinNames)
	for i, latinName := range latinNames {
		log.Println(i + 1)
		webInfoMap[latinName] = GenerateWebInfo(latinName)
	}

	return webInfoMap
}

// 获取并处理网络信息
func GenerateWebInfo(latinNameString string) entities.WebInfo {
	latinName := utils.ParseLatinName(latinNameString)
	if len(latinName.Species) == 0 || strings.HasSuffix(latinName.Species, config.AmbiguousSpeciesName) {
		// 若没有种名，或者种名为 sp.，则仅鉴定到属，无法从网络获取到比较精确的物种信息
		return entities.WebInfo{
			FullLatinName: latinNameString,
			Morphology:    entities.Morphology{},
			NamePublisher: "",
			Habitat:       "",
		}
	}

	fmt.Printf("    -> 开始从网络获取物种信息：%s\n", latinNameString)
	paragraphs, namePublisher, family := parseParagraphs(latinName)
	fmt.Printf("    <- 获取到物种信息：%s %s\n", latinNameString, namePublisher)

	fmt.Printf("    -> 开始寻找最匹配段落：%s\n", latinNameString)
	bestMatchParagraph := pickBestMatchedParagraph(latinNameString, paragraphs)
	fmt.Printf("    <- 找到最匹配段落：%s\n", latinNameString)

	fmt.Printf("    -> 开始从最匹配段落中提取形态描述信息：%s\n", latinNameString)
	morphology := getMorphologyFromMultipleParagraphs([]string{bestMatchParagraph})
	fmt.Printf("    <- 从最匹配段落中提取形态描述信息结束：%s\n", latinNameString)

	return entities.WebInfo{
		FullLatinName: latinNameString,
		Morphology:    morphology,
		NamePublisher: namePublisher,
		Family:        family,
		Habitat:       "",
	}
}

// 提取命名人信息
// TODO, 实现方式需要优化
func extractNamePublisher(latinName entities.LatinName, doc *goquery.Document) string {
	spinfoDiv := doc.Find(config.SpeciesInfoDiv)
	targetText := ""
	spinfoDiv.Find("div").Each(func(i int, div *goquery.Selection) {
		if i == 0 {
			targetText = div.Text()
		}
	})

	namePublisherRegexp, err := regexp.Compile(
		fmt.Sprintf(config.NamePublisherRegexpTemplate, latinName.Genus, latinName.Species))

	if err != nil {
		return ""
	}

	namePublisherSlice := namePublisherRegexp.FindAllString(targetText, -1)

	if len(namePublisherSlice) == 0 {
		return ""
	}

	namePublisher := namePublisherSlice[0]
	namePublisher = strings.Replace(namePublisher, latinName.Genus, "", -1)
	namePublisher = strings.Replace(namePublisher, latinName.Species, "", -1)

	return strings.TrimSpace(namePublisher)
}

// extractFamily: 从网页中提取物种的 “科”（family）信息
// TODO, 实现方式需要优化
func extractFamily(latinName entities.LatinName, doc *goquery.Document) (family string) {
	family = ""
	spinfoDiv := doc.Find(config.SpeciesInfoDiv)
	var contentDiv *goquery.Selection
	childrenDiv := spinfoDiv.Children()
	childrenDiv.Find("div").Each(func(i int, div *goquery.Selection) {
		if i == 5 {
			contentDiv = div
		}
	})

	if contentDiv == nil {
		return
	}

	var familyInfo string
	contentDiv.Find("a").Each(func(i int, div *goquery.Selection) {
		if i == 1 {
			familyInfo = div.Text()
		}
	})

	if len(familyInfo) > 0 && strings.Contains(familyInfo, "科") {
		elements := strings.Fields(familyInfo)
		if len(elements) == 2 {
			family = elements[1]
		}
	}

	return
}

// 选择最符合条件的段落
func pickBestMatchedParagraph(latinNameString string, paragraphs []string) string {
	// TODO, 需要实现：检查所有组合，找到第一个全部包含的段落
	// [A, B, C, D, E], ... 是否有全部包含的段落
	// [A, B, C, D], [A, B, C, E], ... 是否有全部包含的段落
	// [A, B, C], [A, B, D], ... } ... 是否有全部包含的段落
	// [A, B], [A, C], ... } ... 是否有全部包含的段落

	// 目前先使用已有的逻辑
	paragraphContainsAllKeywords := ""
	for _, paragraph := range paragraphs {
		paragraphContainsAllKeywords = checkIfParagraphThatContainsAllKeywords(paragraph, config.AllKeywords)
		if paragraphContainsAllKeywords != "" {
			return paragraph
		}
	}

	for _, paragraph := range paragraphs {
		paragraphContainsAllKeywords = checkIfParagraphThatContainsAllKeywords(paragraph, config.ModerateKeywords)
		if paragraphContainsAllKeywords != "" {
			return paragraph
		}
	}

	for _, paragraph := range paragraphs {
		paragraphContainsAllKeywords = checkIfParagraphThatContainsAllKeywords(paragraph, config.RelaxedKeywords)
		if paragraphContainsAllKeywords != "" {
			return paragraph
		}
	}

	return paragraphContainsAllKeywords
}

// 检测段落是否满足特定的关键字条件
func checkIfParagraphThatContainsAllKeywords(paragraph string, keywords []string) string {
	if strings.TrimSpace(paragraph) == "" {
		return ""
	}
	allKeywordInParagraph := true
	for _, keyword := range keywords {
		if !strings.Contains(paragraph, keyword) {
			allKeywordInParagraph = false
		}
		if !allKeywordInParagraph {
			break
		}
	}
	if allKeywordInParagraph {
		return paragraph
	}
	return ""
}

// 从所有段落里提取植物的各类形态信息
func getMorphologyFromMultipleParagraphs(paragraphs []string) entities.Morphology {
	morphologySlice := make([]entities.Morphology, 0)
	for _, paragraph := range paragraphs {
		morphologySlice = append(morphologySlice, parseMorphologyFromContent(paragraph))
	}

	bodyHeightInfoSlice := make([]string, 0)
	DBHInfoSlice := make([]string, 0)
	stemInfoSlice := make([]string, 0)
	leafInfoSlice := make([]string, 0)
	flowerInfoSlice := make([]string, 0)
	fruitInfoSlice := make([]string, 0)
	hostInfoSlice := make([]string, 0)

	for _, morphology := range morphologySlice {
		bodyHeightInfoSlice = append(bodyHeightInfoSlice, morphology.BodyHeight)
		DBHInfoSlice = append(DBHInfoSlice, morphology.DBH)
		stemInfoSlice = append(DBHInfoSlice, morphology.Stem)
		leafInfoSlice = append(leafInfoSlice, morphology.Leaf)
		flowerInfoSlice = append(flowerInfoSlice, morphology.Flower)
		fruitInfoSlice = append(fruitInfoSlice, morphology.Fruit)
		hostInfoSlice = append(hostInfoSlice, morphology.Host)
	}

	finalMorphology := entities.Morphology{
		BodyHeight: filterAndCombineMorphologyInfo(bodyHeightInfoSlice),
		DBH:        filterAndCombineMorphologyInfo(DBHInfoSlice),
		Stem:       filterAndCombineMorphologyInfo(stemInfoSlice),
		Leaf:       filterAndCombineMorphologyInfo(leafInfoSlice),
		Flower:     filterAndCombineMorphologyInfo(flowerInfoSlice),
		Fruit:      filterAndCombineMorphologyInfo(fruitInfoSlice),
		Host:       filterAndCombineMorphologyInfo(hostInfoSlice),
	}

	return finalMorphology
}

// 拼接形态信息
func filterAndCombineMorphologyInfo(infoSlice []string) string {
	resultSlice := make([]string, 0)
	for _, each := range infoSlice {
		each = strings.TrimSpace(each)
		if len(each) > 0 {
			resultSlice = append(resultSlice, each)
		}
	}

	resultInfo := strings.Join(resultSlice, config.DefaultSeparator)
	if len(resultInfo) > 0 {
		return resultInfo + "。"
	}
	return resultInfo
}

// 从段落中解析形态信息
func parseMorphologyFromContent(paragraph string) entities.Morphology {
	bodyHeightInfo := strings.Join(config.BodyHeightRegexp.FindAllString(paragraph, -1), config.DefaultSeparator) // 体高
	DBHInfo := strings.Join(config.DBHRegexp.FindAllString(paragraph, -1), config.DefaultSeparator)               // 胸径
	stemInfo := strings.Join(config.StemRegexp.FindAllString(paragraph, -1), config.DefaultSeparator)             // 茎
	leafInfo := strings.Join(config.LeafRegexp.FindAllString(paragraph, -1), config.DefaultSeparator)             // 叶
	flowerInfo := strings.Join(config.FlowerRegexp.FindAllString(paragraph, -1), config.DefaultSeparator)         // 花
	fruitInfo := strings.Join(config.FruitRegexp.FindAllString(paragraph, -1), config.DefaultSeparator)           // 果实
	hostInfo := strings.Join(config.HostRegexp.FindAllString(paragraph, -1), config.DefaultSeparator)             // 寄主

	return entities.Morphology{
		BodyHeight: bodyHeightInfo,
		DBH:        DBHInfo,
		Stem:       stemInfo,
		Leaf:       leafInfo,
		Flower:     flowerInfo,
		Fruit:      fruitInfo,
		Host:       hostInfo,
	}
}

// 从网络信息中提取段落及命名人信息
func parseParagraphs(latinName entities.LatinName) (paragraphs []string, namePublisher string, family string) {
	paragraphs = []string{}
	namePublisher = ""
	family = ""

	url := generateUrl(latinName)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	paragraphs = make([]string, 0)
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	namePublisher = extractNamePublisher(latinName, doc)
	family = extractFamily(latinName, doc)

	return paragraphs, namePublisher, family
}

// 根据拉丁名拼接 URL
func generateUrl(latinName entities.LatinName) string {
	return config.URLPrefixEFLORA + strings.Join(latinName.Elements, config.URLBlankSeparator)
}

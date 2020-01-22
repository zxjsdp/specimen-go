package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

	fmt.Printf("    🌐 [%s]        开始从网络获取「物种信息」\n", latinNameString)
	frpsspno, frpsspclassid, paragraphs := parseParagraphs(latinName)
	fmt.Printf("    💚 [%s]        获取到「物种信息」\n", latinNameString)

	fmt.Printf("    🌀 [%s]        开始寻找「最匹配段落」\n", latinNameString)
	bestMatchParagraph := pickBestMatchedParagraph(latinNameString, paragraphs)
	fmt.Printf("    💚 [%s]        寻找「最匹配段落」完成\n", latinNameString)

	fmt.Printf("    🌀 [%s]        开始从最匹配段落中「提取形态描述信息」\n", latinNameString)
	morphology := getMorphologyFromMultipleParagraphs([]string{bestMatchParagraph})
	fmt.Printf("    💚 [%s]        从最匹配段落中「提取形态描述信息」结束\n", latinNameString)

	fmt.Printf("    🌐 [%s]        开始从网络获取「命名人」信息\n", latinNameString)
	namePublisher := parseNamePublisher(latinName)
	fmt.Printf("    💚 [%s %s]        获取到「命名人」信息 \n", latinNameString, namePublisher)

	fmt.Printf("    🌐 [%s]        开始从网络获取「物种分类（Texomony，界门纲目科属种）」信息\n", latinNameString)
	phylum, class, order, family, genus := parseTaxonomyInfo(latinName, frpsspno, frpsspclassid)
	fmt.Printf("    💚 [%s %s]        获取到「物种分类（Texomony，界门纲目科属种）」信息: → 「%s %s %s %s %s」\n", latinNameString, namePublisher, phylum, class, order, family, genus)

	return entities.WebInfo{
		FullLatinName: latinNameString,
		Morphology:    morphology,
		NamePublisher: namePublisher,
		Family:        family,
		Habitat:       "",
	}
}

// 选择最符合条件的段落
func pickBestMatchedParagraph(latinNameString string, paragraphs []string) string {
	if len(paragraphs) == 0 {
		fmt.Printf("    💔 %s | %s | resp.Body 为空!\n", latinNameString, "pickBestMatchedParagraph")
		return ""
	}

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

// 「命名人」信息解析
func parseNamePublisher(latinName entities.LatinName) (namePublisher string) {
	baseUrl := config.URLPrefixEFLORA + strings.Join(latinName.Elements, config.URLBlankSeparator) + config.URLSuffix
	resp, err := http.Get(baseUrl)
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseNamePublisher http.Get", err)
	}
	defer resp.Body.Close()

	spnoRegexp, err := regexp.Compile(fmt.Sprintf(config.SpnoRegexpTemplate))
	if err != nil {
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	spnoSlice := spnoRegexp.FindAllString(string(body), -1)
	if len(spnoSlice) == 0 {
		return ""
	}
	spnoInfo := spnoSlice[0]
	spnoInfoSlice := strings.Split(spnoInfo, "\"")
	if len(spnoInfoSlice) != 3 {
		return ""
	}

	spno := spnoInfoSlice[1]

	// 「拉丁名」信息查询 data 格式，注意此 spno 与 frpsspno 为不同的 ID
	//     - spno: 从「http://www.iplant.cn/ashx/getfrps.ashx?key=Cephalotaxus+fortunei」API 返回的结果中取 var spno = "10726"
	// 示例：spno=10726
	resp, err = http.PostForm(config.LatinApiUrl, url.Values{"spno": {spno}})
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseNamePublisher http.PostForm", err)
	}
	defer resp.Body.Close()

	// 示例: <span class='font20'><b>Cephalotaxus</b></span> <span class='font20'><b>fortunei</b></span> Hooker
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseNamePublisher http.PostForm", err)
		return ""
	}

	contentSlice := strings.Split(string(body), config.NamePublisherSeparator)
	if len(contentSlice) > 0 {
		namePublisher = strings.TrimSpace(contentSlice[len(contentSlice)-1])
	}

	return namePublisher
}

// 从网络信息中提取「生物分类（门纲目科属）」信息
// 界（Kingdom）、门（Phylum）、纲（Class）、目（Order）、科（Family）、属（Genus）、种（Species）
func parseTaxonomyInfo(latinName entities.LatinName, frpsspno string, frpsspclassid string) (phylum string, class string, order string, family string, genus string) {
	if len(frpsspno) == 0 || len(frpsspclassid) == 0 {
		fmt.Printf("    💔 %s | %s \n", latinName.LatinNameString, "parseTaxonomyInfo frpsspno 或 frpsspclassid 为空! "+"frpsspno: "+frpsspno+", frpsspclassid: "+frpsspclassid)
		return "", "", "", "", ""
	}

	// 「物种分类（Texomony，界门纲目科属种）」信息查询 data 格式，
	//     - spno: 从「详细描述」API 返回的结果中取 frpsspno
	//     - spclassid: 从「详细描述」API 返回的结果中取 frpsspclassid
	// 示例：spno=52&spclassid=24
	resp, err := http.PostForm(config.DetailedCategoryApiUrl, url.Values{"spno": {frpsspno}, "spclassid": {frpsspclassid}})
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseTaxonomyInfo http.PostForm", err)
	}
	defer resp.Body.Close()

	var responseMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseTaxonomyInfo json.NewDecoder", err)
	}

	frpsclasstxt := responseMap[config.FrpsclasstxtKeyInResponseMap]

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(frpsclasstxt))
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseTaxonomyInfo goquery.NewDocumentFromReader", err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, config.PhylumKeyword) {
			phylum = strings.TrimSpace(text)
		} else if strings.Contains(text, config.ClassKeyword) {
			class = strings.TrimSpace(text)
		} else if strings.Contains(text, config.OrderKeyword) {
			order = strings.TrimSpace(text)
		} else if strings.Contains(text, config.FamilyKeyword) {
			family = strings.TrimSpace(text)
		} else if strings.Contains(text, config.GenusKeyword) {
			genus = strings.TrimSpace(text)
		}
	})

	return phylum, class, order, family, genus
}

// 从网络 API 返回的详细段落信息中，提取「茎」、「叶」、「花」、「果」等信息
func parseParagraphs(latinName entities.LatinName) (frpsspno string, frpsspclassid string, paragraphs []string) {
	paragraphs = []string{}

	apiUrl := config.DetailedDescriptionApiURLPrefix + strings.Join(latinName.Elements, config.APIURLBlankSeparator)
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseParagraphs http.Get", err)
	}

	var responseMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		fmt.Printf("    💔 %s | %s %v \n", latinName.LatinNameString, "parseParagraphs json.NewDecoder", err)
		return "", "", []string{}
	}
	defer resp.Body.Close()

	frpsspno = responseMap[config.FrpsspnoKeyInResponseMap]
	frpsspclassid = responseMap[config.FrpsspclassidKeyInResponseMap]
	frpsdesc := responseMap[config.FrpsdescKeyInResponseMap]

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(frpsdesc))
	if err != nil {
		fmt.Printf("    💔 %s | %s: %v \n", latinName.LatinNameString, "parseParagraphs goquery.NewDocumentFromReader", err)
	}

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	return frpsspno, frpsspclassid, paragraphs
}

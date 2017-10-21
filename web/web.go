package web

import (
	"log"
	"strings"

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

	for i := 1; i <= config.WORKER_POOL_SIZE; i++ {
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
		log.Println("worker", id, "started job", latinName)
		webInfo := GenerateWebInfo(latinName)
		log.Println("worker", id, "finished job", latinName)
		results <- webInfo
	}
}

// （同步方法）根据拉丁名获取网络信息 map
func GenerateWebInfoMapSync(latinNames []string) map[string]entities.WebInfo {
	webInfoMap := make(map[string]entities.WebInfo)

	latinNames = utils.RemoveDuplicates(latinNames)
	for _, latinName := range latinNames {
		webInfoMap[latinName] = GenerateWebInfo(latinName)
	}

	return webInfoMap
}

func GenerateWebInfo(latinNameString string) entities.WebInfo {
	latinName := utils.ParseLatinName(latinNameString)
	log.Println(latinName)
	url := generateUrl(latinName)
	paragraphs := parseParagraphs(url)
	morphology := getMorphologyFromMultipleParagraphs(paragraphs)

	return entities.WebInfo{
		FullLatinName: latinNameString,
		Morphology:    morphology,
		NameGiver:     "default",
		Habitat:       "default",
	}
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

func parseParagraphs(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	paragraphs := make([]string, 0)
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	return paragraphs
}

func generateUrl(latinName entities.LatinName) string {
	return config.URL_PREFIX_EFLORA + strings.Join(latinName.Elements, config.URL_BLANK_SEPERATOR)
}

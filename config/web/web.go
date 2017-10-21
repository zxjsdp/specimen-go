package web

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zxjsdp/specimen-go/config"
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/utils"
)

func GenerateWebInfo(latinNameString string) entities.WebInfo {
	latinName := utils.ParseLatinName(latinNameString)
	url := generateUrl(latinName)
	paragraphs := parseParagraphs(url)
	morphology := getMorphologyFromMultipleParagraphs(paragraphs)

	return entities.WebInfo{
		Morphology: morphology,
		NameGiver:  "default",
		Habitat:    "default",
	}
}

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
		BodyHeight: strings.Join(bodyHeightInfoSlice, config.DefaultSeparator),
		DBH:        strings.Join(DBHInfoSlice, config.DefaultSeparator),
		Stem:       strings.Join(stemInfoSlice, config.DefaultSeparator),
		Leaf:       strings.Join(leafInfoSlice, config.DefaultSeparator),
		Flower:     strings.Join(flowerInfoSlice, config.DefaultSeparator),
		Fruit:      strings.Join(fruitInfoSlice, config.DefaultSeparator),
		Host:       strings.Join(hostInfoSlice, config.DefaultSeparator),
	}

	return finalMorphology
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

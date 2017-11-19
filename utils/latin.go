package utils

import (
	"regexp"

	"strings"

	"github.com/zxjsdp/specimen-go/entities"
)

// 从字符串中解析拉丁名，以及属名及种名
func ParseLatinName(latinNameString string) entities.LatinName {
	nonBlankSubStringRegexp, _ := regexp.Compile("\\S+")
	elements := nonBlankSubStringRegexp.FindAllString(latinNameString, -1)

	latinNameString = strings.Join(elements, " ")

	genus := ""
	species := ""
	if len(elements) >= 2 {
		genus = elements[0]   // 第一个元素认为是属
		species = elements[1] // 第二个元素认为是种。若种为 sp. (config.AmbiguousSpeciesName) 则认为是属中不确定的一种
	} else if len(elements) == 1 {
		genus = elements[0]
	}

	return entities.LatinName{LatinNameString: latinNameString, Genus: genus, Species: species, Elements: elements}
}

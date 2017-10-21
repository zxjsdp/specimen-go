package converters

import (
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/utils"
)

// 生成 Entry info map
func GenerateEntryDataMap(entryDataSlice []entities.EntryData) map[string]entities.EntryData {
	entryDataMap := make(map[string]entities.EntryData)
	if len(entryDataSlice) == 0 {
		return entryDataMap
	}

	for _, entry := range entryDataSlice {
		entryDataMap[entry.SpeciesNumber] = entry
	}

	return entryDataMap
}

// 从 entryData 中提取 species slice
func ExtractSpeciesNames(entryDataSlice []entities.EntryData) []string {
	speciesNames := make([]string, len(entryDataSlice))
	index := 0
	for _, entryData := range entryDataSlice {
		speciesNames[index] = entryData.FullLatinName
		index++
	}
	return utils.RemoveDuplicates(speciesNames[:index])
}

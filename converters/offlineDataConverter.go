package converters

import (
	"github.com/zxjsdp/specimen-go/entities"
	"github.com/zxjsdp/specimen-go/utils"
)

// 生成 Offline info map
func GenerateOfflineDataMap(offlineDataSlice []entities.OfflineData) map[string]entities.OfflineData {
	offlineDataMap := make(map[string]entities.OfflineData)
	if len(offlineDataSlice) == 0 {
		return offlineDataMap
	}

	for _, offlineData := range offlineDataSlice {
		offlineDataMap[offlineData.SpeciesNumber] = offlineData
	}

	return offlineDataMap
}

// 从 offlineData 中提取 species slice
func ExtractSpeciesNames(offlineDataSlice []entities.OfflineData) []string {
	speciesNames := make([]string, len(offlineDataSlice))
	index := 0
	for _, offlineData := range offlineDataSlice {
		speciesNames[index] = offlineData.FullLatinName
		index++
	}
	return utils.RemoveDuplicates(speciesNames[:index])
}

package converters

import "github.com/zxjsdp/specimen-go/entities"

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

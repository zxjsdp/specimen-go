package converters

import "github.com/zxjsdp/specimen-go/entities"

// 生成 Entry info map
func GenerateEntryDataMap(entryDatas []entities.EntryData) map[string]entities.EntryData {
	entryDataMap := make(map[string]entities.EntryData)
	if len(entryDatas) == 0 {
		return entryDataMap
	}

	for _, entry := range entryDatas {
		entryDataMap[entry.SpeciesNumber] = entry
	}

	return entryDataMap
}

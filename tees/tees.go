package tees

import (
	"tcas-cli/collectors"
	"tcas-cli/tees/csv"
)

func GetCollectors() map[string]collectors.EvidenceCollector {
	collectorMap := make(map[string]collectors.EvidenceCollector, 0)
	csvCollector := csv.NewCollector()
	collectorMap[csvCollector.Name()] = csvCollector

	return collectorMap

}

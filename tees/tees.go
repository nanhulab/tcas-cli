package tees

import (
	"tcas-cli/collectors"
	"tcas-cli/tees/csv"
	"tcas-cli/tees/virtcca"
)

func GetCollectors() map[string]collectors.EvidenceCollector {
	collectorMap := make(map[string]collectors.EvidenceCollector, 0)
	csvCollector := csv.NewCollector()
	collectorMap[csvCollector.Name()] = csvCollector

	vccaCollector := virtcca.NewCollector()
	collectorMap[vccaCollector.Name()] = vccaCollector

	return collectorMap

}

package tees

import (
	"github.com/nanhulab/tcas-cli/collectors"
	"github.com/nanhulab/tcas-cli/tees/csv"
	nvidia "github.com/nanhulab/tcas-cli/tees/nvidiamock"
	"github.com/nanhulab/tcas-cli/tees/virtcca"
)

func GetCollectors() map[string]collectors.EvidenceCollector {
	collectorMap := make(map[string]collectors.EvidenceCollector, 0)
	csvCollector := csv.NewCollector()
	collectorMap[csvCollector.Name()] = csvCollector

	vccaCollector := virtcca.NewCollector()
	collectorMap[vccaCollector.Name()] = vccaCollector

	nvidiaCollector := nvidia.NewCollector()
	collectorMap[nvidiaCollector.Name()] = nvidiaCollector

	return collectorMap

}

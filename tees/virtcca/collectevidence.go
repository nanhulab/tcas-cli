package vcca

import (
	"tcas-cli/collectors"
)

type Collector struct {
}

func (c *Collector) CollectEvidence(userdata []byte) (*collectors.Evidence, error) {
	return nil, nil
}

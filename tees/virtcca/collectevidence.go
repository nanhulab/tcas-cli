package vcca

import "tcas-cli/manager"

type Collector struct {
}

func (c *Collector) CollectEvidence(userdata []byte) (*manager.Evidence, error) {
	return nil, nil
}

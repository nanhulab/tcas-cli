package manager

type Evidence struct {
	Tee         string      `json:"tee"`
	TeeReport   string      `json:"tee_report"`
	Parameter   interface{} `json:"parameter"`
	RuntimeData string      `json:"runtime_data"`
}

type EvidenceCollector interface {
	CollectEvidence(userdata []byte) (*Evidence, error)
}

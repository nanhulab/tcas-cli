package collectors

type Evidence struct {
	TeeType   string      `json:"tee"`
	TeeReport string      `json:"tee_report"`
	Parameter interface{} `json:"parameter"`
}

type EvidenceCollector interface {
	CollectEvidence(userdata []byte) (*Evidence, error)
	Name() string
}

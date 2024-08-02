package manager

type SetPolicyReq struct {
	Name            string `json:"policy_name"`
	Policy          string `json:"policy_rego"`
	AttestationType string `json:"attestation_type"`
}

type TrustDeviceReport struct {
	DeviceType   string      `json:"type"`
	DeviceReport string      `json:"device_report"`
	Parameter    interface{} `json:"parameter"`
}

type NodeEvidence struct {
	Tee         string               `json:"tee"`
	TeeReport   string               `json:"tee_report"`
	Parameter   interface{}          `json:"parameter"`
	TrustDevice []*TrustDeviceReport `json:"trust_devices"`
	RuntimeData string               `json:"runtime_data"`
	InitData    string               `json:"init_data"`
	EventLog    string               `json:"event_log"`
}

type NodeAttestInfoReq struct {
	Report    *NodeEvidence `json:"report"`
	Nonce     string        `json:"nonce"`
	PolicyIds []string      `json:"policy_ids"`
}

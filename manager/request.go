package manager

type SetPolicyReq struct {
	Name            string `json:"policy_name"`
	Policy          string `json:"policy_rego"`
	AttestationType string `json:"attestation_type"`
}

/*
 * @Author: jffan
 * @Date: 2024-08-15 09:16:45
 * @LastEditTime: 2024-08-15 10:56:35
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\manager\request.go
 * @Description: The type of request params
 */
package manager

type SetPolicyReq struct {
	Name            string `json:"policy_name"`
	Policy          string `json:"policy_rego"`
	AttestationType string `json:"attestation_type"`
}

type SetSecretReq struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

type UpdateSecretReq struct {
	Id     string `json:"id"`
	Secret string `json:"secret"`
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

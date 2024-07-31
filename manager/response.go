package manager

type HttpBaseResponse struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

type PolicySetResponse struct {
	HttpBaseResponse
	PolicyID string `json:"policy_id"`
}

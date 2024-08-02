package manager

import "time"

type HttpBaseResponse struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

type PolicySetResponse struct {
	HttpBaseResponse
	PolicyID string `json:"policy_id"`
}

type Nonce struct {
	Nonce       string    `json:"nonce,omitempty"`
	ExpiredTime time.Time `json:"expired,omitempty"`
}
type NonceResponse struct {
	HttpBaseResponse
	Data *Nonce `json:"data"`
}

type TokenResponse struct {
	HttpBaseResponse
	Token string `json:"token"`
}

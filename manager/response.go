/*
 * @Author: jffan
 * @Date: 2024-07-31 15:01:17
 * @LastEditTime: 2024-08-02 15:33:34
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\manager\response.go
 * @Description: 🎉🎉🎉
 */
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

type PolicyDeleteResponse struct {
	HttpBaseResponse
	PolicyID string `json:"policy_id"`
}

type PolicyListResData struct {
	No              int    `json:"no"`
	PolicyId        string `json:"policy_id"`
	PolicyRego      string `json:"policy_rego"`
	PolicyName      string `json:"policy_name"`
	AttestationType string `json:"attestation_type"`
	PolicyHash      string `json:"policy_hash"`
	Version         int    `json:"version"`
	CreateTime      string `json:"createTime"`
	UpdateTime      string `json:"updateTime"`
}

type PolicyListResponse struct {
	HttpBaseResponse
	Data []PolicyListResData `json:"data"`
}

type PolicyListJsonFormat struct {
	Policies []PolicyListResData `json:"policies"`
}

type SecretListResData struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type SecretListResponse struct {
	HttpBaseResponse
	Data []SecretListResData `json:"data"`
}

type SecretListJsonFormat struct {
	Secrets []SecretListResData `json:"secrets"`
}

type SecretDeleteResponse struct {
	HttpBaseResponse
	SecretID string `json:"secret_id"`
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

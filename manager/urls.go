/*
 * @Author: jffan
 * @Date: 2024-07-31 15:01:17
 * @LastEditTime: 2024-08-14 14:33:15
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\manager\urls.go
 * @Description: Define the constant for the request
 */
package manager

// The backend interface address
const (
	NonceGetInterface = "/v1/nonce"
	PolicyUrl         = "/v1/policy"
	SecretUrl         = "/v1/secret"
	SecretListUrl     = "/v1/secret/list"
	NonceUrl          = "/v1/nonce"
	AttestUrl         = "/v1/attest"
	CaUrl             = "/v1/pki/ca"
)

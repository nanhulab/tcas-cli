package manager

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"os"
	"strings"
)

type Manager struct {
	APIEndpoint string
	TlsConfig   *tls.Config
}

func New(apiEndpoint, caPath string) (*Manager, error) {
	tc := new(tls.Config)
	if caPath != "" {
		certBytes, err := os.ReadFile(caPath)
		if err != nil {
			return nil, err
		}

		CaCertPool := x509.NewCertPool()
		ok := CaCertPool.AppendCertsFromPEM(certBytes)
		if !ok {
			return nil, fmt.Errorf("add ca to pool failed")
		}
		tc.RootCAs = CaCertPool
	}

	return &Manager{
		APIEndpoint: apiEndpoint,
		TlsConfig:   tc,
	}, nil
}

func (m *Manager) newClient(method string, url string) *httplib.BeegoHTTPRequest {
	var client *httplib.BeegoHTTPRequest
	me := strings.ToUpper(method)
	client = httplib.NewBeegoRequest(m.APIEndpoint+url, me)

	if m.TlsConfig != nil {
		client.SetTLSClientConfig(m.TlsConfig)
	}
	return client
}

func (m *Manager) SetPolicy(name, policy, attestationType string) (*PolicySetResponse, error) {

	if name == "" || policy == "" {
		return nil, fmt.Errorf("name or policy is null")
	}
	client := m.newClient("post", PolicyUrl)
	if client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	req := SetPolicyReq{
		Name:            name,
		Policy:          policy,
		AttestationType: attestationType,
	}

	fmt.Println(req)
	client, err := client.JSONBody(req)
	if err != nil {
		return nil, err
	}

	res := new(PolicySetResponse)
	err = client.ToJSON(res)
	if err != nil {
		errMesg := fmt.Sprintf("requst set policy api failed, , error: %s ", err)
		return nil, fmt.Errorf(errMesg)
	}

	return res, nil
}

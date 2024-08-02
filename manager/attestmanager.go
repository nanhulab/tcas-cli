package manager

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"tcas-cli/collectors"
	"tcas-cli/tees"
)

type Manager struct {
	APIEndpoint string
	TlsConfig   *tls.Config
	Collectors  map[string]collectors.EvidenceCollector
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

	c := tees.GetCollectors()

	return &Manager{
		APIEndpoint: apiEndpoint,
		TlsConfig:   tc,
		Collectors:  c,
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

func (m *Manager) GetNonce() (*NonceResponse, error) {
	client := m.newClient("get", NonceUrl)
	res := new(NonceResponse)
	err := client.ToJSON(res)
	if err != nil {
		return nil, fmt.Errorf("request get nonce failed, error: %s", err)
	}

	return res, nil

}

func (m *Manager) getNodeAttestInfo(tee, runtimedata, devices, policies string) (*NodeAttestInfoReq, error) {
	if tee == "" {
		return nil, fmt.Errorf("tee is null")
	}

	nonceRes, err := m.GetNonce()
	if err != nil {
		return nil, fmt.Errorf("get nonce failed, error: %s", err)
	}
	if nonceRes.Code != 200 {
		return nil, fmt.Errorf("get nonce failed, error: %s", nonceRes.Message)
	}

	logrus.Debugf("nonce is %s", nonceRes.Data.Nonce)
	userdata := fmt.Sprintf("%x", sha256.Sum256([]byte(nonceRes.Data.Nonce+runtimedata)))

	c, ok := m.Collectors["csv"]
	if !ok {
		return nil, fmt.Errorf("tee: %s not support yet", tee)
	}
	teeReport, err := c.CollectEvidence([]byte(userdata))
	if err != nil {
		return nil, fmt.Errorf("collect evidence from %s failed, error: %s", c.Name(), err)
	}

	var devicelist []string
	if devices != "" {
		devicelist = strings.Split(devices, ",")
	}
	deviceReports := make([]*TrustDeviceReport, 0)
	for _, d := range devicelist {
		c, ok := m.Collectors[d]
		if !ok {
			return nil, fmt.Errorf("device tee: %s not support yet", tee)
		}
		deviceEvidence, err := c.CollectEvidence([]byte(userdata))
		if err != nil {
			return nil, fmt.Errorf("collect evidence from %s failed, error: %s", c.Name(), err)
		}

		deviceReports = append(deviceReports, &TrustDeviceReport{
			DeviceType:   d,
			DeviceReport: deviceEvidence.TeeReport,
			Parameter:    deviceEvidence.Parameter,
		})
	}

	var policyIds []string
	if policies != "" {
		policyIds = strings.Split(policies, ",")
	}
	return &NodeAttestInfoReq{
		Report: &NodeEvidence{
			Tee:         teeReport.TeeType,
			TeeReport:   teeReport.TeeReport,
			Parameter:   teeReport.Parameter,
			RuntimeData: runtimedata,
			TrustDevice: deviceReports,
		},
		Nonce:     nonceRes.Data.Nonce,
		PolicyIds: policyIds,
	}, nil
}

func (m *Manager) AttestForToken(tee, runtimedata, devices, policies string) (*TokenResponse, error) {
	attestReq, err := m.getNodeAttestInfo(tee, runtimedata, devices, policies)
	if err != nil {
		return nil, fmt.Errorf("get node attestInfo failed, error: %s", err)
	}

	client := m.newClient("post", AttestUrl)
	client, err = client.JSONBody(attestReq)
	if err != nil {
		return nil, fmt.Errorf("set request body failed, error: %s", err)
	}

	tokenRes := new(TokenResponse)
	err = client.ToJSON(tokenRes)
	if err != nil {
		return nil, fmt.Errorf("do request to attest api failed, error: %s", err)
	}

	if tokenRes.Code != 200 {
		return nil, fmt.Errorf("response error: %s", tokenRes.Message)
	}

	return tokenRes, nil
}

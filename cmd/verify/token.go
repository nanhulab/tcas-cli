/*
 * @Author: jffan
 * @Date: 2024-08-09 09:29:01
 * @LastEditTime: 2024-08-14 16:37:52
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\cmd\verify\token.go
 * @Description: The CA certificate is used to verify the token, which is divided into online verification and local verification
 */
package verify

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"os"
	"strings"
	consts "tcas-cli/constants"
	"tcas-cli/manager"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type TokenHeader struct {
	Alg string `json:"alg"`
	Jku string `json:"jku"`
	Kid string `json:"kid"`
	Typ string `json:"typ"`
}

var verifyTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "verify token",
	Long:  `verify token`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenString, _ := cmd.Flags().GetString("token")
		caFilePath, _ := cmd.Flags().GetString("file")
		if caFilePath == "" {
			//online verify
			logrus.Debugf(consts.ColorYellow + "Verify Token Online" + consts.OutReset)
			url, _ := cmd.Flags().GetString("url")
			m, err := manager.New(url, "")
			if err != nil {
				logrus.Errorf("create attest manager failed, error: %s", err)
				return
			}
			res, err := m.GetRootCert()
			if err != nil {
				logrus.Errorf("Request policy list failed: %v", err)
				return
			}
			// parse token header
			parts := strings.Split(tokenString, ".")
			header, err := base64.RawURLEncoding.DecodeString(parts[0])
			if err != nil {
				logrus.Errorf(consts.ColorRed+"Decoding token header failed:%v"+consts.OutReset, err)
				return
			}
			var headerToken TokenHeader
			if err := json.Unmarshal(header, &headerToken); err != nil {
				logrus.Errorf(consts.ColorRed+"Unmarshalling header failed:%v"+consts.OutReset, err)
				return
			}
			for _, key := range res.Keys {
				if key.Kid == headerToken.Kid {
					pemData, err := manager.X5cToCertPem(key.X5c)
					if err != nil {
						logrus.Errorf(err.Error())
						return
					}
					block, _ := pem.Decode(pemData.Bytes())
					if block == nil {
						logrus.Errorf("Failed to decode PEM block containing the CA certificate")
						return
					}
					caCert, err := x509.ParseCertificate(block.Bytes)
					if err != nil {
						logrus.Errorf("Failed to parse CA certificate: %v", err)
						return
					}
					publicKey := caCert.PublicKey
					token, err := manager.ParseTokenByPk(publicKey, tokenString)
					if err != nil {
						logrus.Errorf(err.Error())
						return
					}
					logrus.Infof(consts.ColorGreen + "Verify token successful" + consts.OutReset)
					err = manager.PrintFormatToken(token)
					if err != nil {
						logrus.Errorf(err.Error())
						return
					}
					return
				}
			}
			logrus.Errorf(consts.ColorRed + "Verify Token Failed" + consts.OutReset)
		} else {
			//local verify
			logrus.Debugf(consts.ColorYellow + "Verify Token Local" + consts.OutReset)
			caCertPEM, err := os.ReadFile(caFilePath)
			if err != nil {
				logrus.Errorf("read ca file failed, error: %s", err)
				return
			}
			block, _ := pem.Decode(caCertPEM)
			if block == nil {
				logrus.Errorf("Failed to decode PEM block containing the CA certificate")
				return
			}
			caCert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				logrus.Errorf("Failed to parse CA certificate: %v", err)
				return
			}
			publicKey := caCert.PublicKey
			token, err := manager.ParseTokenByPk(publicKey, tokenString)
			if err != nil {
				logrus.Errorf(err.Error())
				return
			}
			logrus.Infof(consts.ColorGreen + "Verify token successful" + consts.OutReset)
			err = manager.PrintFormatToken(token)
			if err != nil {
				logrus.Errorf(err.Error())
				return
			}
		}

	},
}

func init() {
	Cmd.AddCommand(verifyTokenCmd)
	//set parameter for secret set
	verifyTokenCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
	verifyTokenCmd.Flags().StringP("token", "t", "", "must, tcas's token")
	verifyTokenCmd.Flags().StringP("file", "f", "", "optionally, the path of the CA certificate. If not, it will be verified online")
	verifyTokenCmd.MarkFlagRequired("token")
}

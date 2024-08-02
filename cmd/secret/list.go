/*
 * @Author: jffan
 * @Date: 2024-08-01 15:36:12
 * @LastEditTime: 2024-08-02 17:00:39
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\secret\list.go
 * @Description: 🎉🎉🎉
 */
package secret

import (
	"encoding/json"
	consts "tcas-cli/constants"
	"tcas-cli/manager"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var secretListCmd = &cobra.Command{
	Use:   "list",
	Short: "get the secret base info list",
	Long:  `get the secret base info list`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		logrus.Debugf("secret url: " + consts.ColorYellow + url + consts.OutReset)
		m, err := manager.New(url, "")
		if err != nil {
			logrus.Errorf("create attest manager failed, error: %s", err)
		}
		res, err := m.ListSecret()
		if err != nil {
			logrus.Errorf("Request secret list failed: %v", err)
		}
		if res.Code == 200 {
			jsonData, err := json.MarshalIndent(manager.SecretListJsonFormat{Secrets: res.Data}, "", "  ")
			if err != nil {
				logrus.Errorf("Error marshaling JSON:", err)
			} else {
				logrus.Debugf("------------------secret list start------------------")
				logrus.Debugf(consts.ColorYellow + string(jsonData) + consts.OutReset)
				logrus.Debugf("------------------secret list end--------------------")
			}
		} else {
			logrus.Errorf(consts.ColorRed + "request secret list failed:" + res.Message + consts.OutReset)
		}
	},
}

func init() {
	Cmd.AddCommand(secretListCmd)
	//set parameter for secret list
	secretListCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
}

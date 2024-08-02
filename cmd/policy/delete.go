/*
 * @Author: jffan
 * @Date: 2024-07-31 14:46:14
 * @LastEditTime: 2024-08-02 17:02:49
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\policy\delete.go
 * @Description: 🎉🎉🎉
 */
package policy

import (
	consts "tcas-cli/constants"
	"tcas-cli/manager"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var policyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete policy",
	Long:  `delete policy`,
	Run: func(cmd *cobra.Command, args []string) {
		ploicyId, _ := cmd.Flags().GetString("id")
		if ploicyId == "" {
			logrus.Errorf(consts.ColorRed + "policy id is required ! use `--id <policy_id>`" + consts.OutReset)
			return
		}

		url, _ := cmd.Flags().GetString("url")
		logrus.Debugf("policy delete url: " + consts.ColorYellow + url + consts.OutReset)

		m, err := manager.New(url, "")
		if err != nil {
			logrus.Errorf("create attest manager failed, error: %s", err)
		}
		res, err := m.DeletePolicy(ploicyId)
		if err != nil {
			logrus.Errorf("Request failed: %v", err)
		}
		if res.Code == 200 {
			logrus.Debugf(consts.ColorGreen + "delete policy successful, the policy id is " + res.PolicyID + consts.OutReset)
		} else {
			logrus.Errorf(consts.ColorRed + "delete policy failed:" + res.Message + consts.OutReset)
		}
	},
}

func init() {
	Cmd.AddCommand(policyDeleteCmd)
	//set parameter for policy delete
	policyDeleteCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
	policyDeleteCmd.Flags().StringP("id", "i", "", "must the id of policy")
	policyDeleteCmd.MarkFlagRequired("id")
}

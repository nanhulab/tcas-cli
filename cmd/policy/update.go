/*
 * @Author: jffan
 * @Date: 2024-07-31 16:41:04
 * @LastEditTime: 2024-08-02 16:57:13
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\policy\update.go
 * @Description: 🎉🎉🎉
 */
package policy

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var policyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update policy",
	Long:  `update policy`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debugf("policy update called")
	},
}

func init() {
	Cmd.AddCommand(policyUpdateCmd)
	//set parameter for policy delete
	policyUpdateCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
}

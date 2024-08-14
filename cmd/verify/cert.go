/*
 * @Author: jffan
 * @Date: 2024-08-14 09:33:52
 * @LastEditTime: 2024-08-14 15:30:54
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\cmd\verify\cert.go
 * @Description:
 */
package verify

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verfifyCertCmd = &cobra.Command{
	Use:   "cert",
	Short: "verify cert",
	Long:  `verify cert`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("verify cert called")
	},
}

func init() {
	Cmd.AddCommand(verfifyCertCmd)
	//set parameter for verify cert
	verfifyCertCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
	verfifyCertCmd.Flags().StringP("file", "f", "", "must, the path of the cert to be verified")
}

/*
 * @Author: jffan
 * @Date: 2024-08-01 15:15:30
 * @LastEditTime: 2024-08-02 17:01:31
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\secret\set.go
 * @Description: 🎉🎉🎉
 */
package secret

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var secretSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set secret",
	Long:  `set secret`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debugf("secret set called")
	},
}

func init() {
	Cmd.AddCommand(secretSetCmd)
	//set parameter for secret set
	secretSetCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
}

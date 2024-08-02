/*
 * @Author: jffan
 * @Date: 2024-08-01 15:18:37
 * @LastEditTime: 2024-08-02 17:01:52
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\secret\update.go
 * @Description: 🎉🎉🎉
 */
package secret

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var secretUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update secret",
	Long:  `update secret`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debugf("secret update called ")
	},
}

func init() {
	Cmd.AddCommand(secretUpdateCmd)
	//set parameter for secret update
	secretUpdateCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
}

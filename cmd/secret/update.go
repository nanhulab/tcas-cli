/*
 * @Author: jffan
 * @Date: 2024-08-01 15:18:37
 * @LastEditTime: 2024-08-14 14:35:46
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\cmd\secret\update.go
 * @Description: update secret
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
		logrus.Infof("secret update called ")
	},
}

func init() {
	Cmd.AddCommand(secretUpdateCmd)
	//set parameter for secret update
	secretUpdateCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
}

package attest

import (
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "attest for getting token",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand(tokenCmd)
	//set parameter for policy set
	tokenCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
	tokenCmd.Flags().StringP("tee", "t", "", "must, tee type")
	tokenCmd.MarkFlagRequired("tee")
	tokenCmd.Flags().StringP("devices", "v", "", "optional, the trust devices")

	tokenCmd.Flags().StringP("userdata", "d", "", "optional, the base64 encoded userdata")
	tokenCmd.Flags().StringP("policies", "p", "", "optional, the ids of the policy needed matching")
}

package attest

import (
	"github.com/spf13/cobra"
	consts "tcas-cli/constants"
)

const TEE = "tee"
const USERDATA = "userdata"
const POLICIES = "policies"
const DEVICES = "DEVICES"

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
	tokenCmd.Flags().StringP(consts.UrlParam, consts.UrlShotParam, "https://api.trustcluster.cn", "optional, tcas's api url")
	tokenCmd.Flags().StringP(TEE, "t", "", "must, tee type")
	tokenCmd.MarkFlagRequired(TEE)
	tokenCmd.Flags().StringP(DEVICES, "v", "", "optional, the trust devices")

	tokenCmd.Flags().StringP(USERDATA, "d", "", "optional, the base64 encoded userdata")
	tokenCmd.Flags().StringP(POLICIES, "p", "", "optional, the ids of the policy needed matching")
}

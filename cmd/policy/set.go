package policy

import (
	"fmt"
	"github.com/spf13/cobra"
	consts "tcas-cli/constants"
	"tcas-cli/manager"
	"tcas-cli/utils/file"
	"tcas-cli/utils/tools"
)

var policySetCmd = &cobra.Command{
	Use:   "set",
	Short: "set policy",
	Long:  `set policy`,
	Run: func(cmd *cobra.Command, args []string) {
		filepath, _ := cmd.Flags().GetString("rego-file")
		if filepath == "" {
			fmt.Println(consts.ColorRed + "Error: Please set the rego policy file path first! use `--file /path/to/file`" + consts.OutReset)
			return
		}
		var fileBase64 string
		var err error
		if file.IsExists(filepath) {
			fileBase64, err = file.FileToBase64(filepath)
			if err != nil {
				fmt.Println("Error: transfer base64 failed：", err)
				return
			}
		} else {
			fmt.Println(consts.ColorRed + "Error: The file path you set is not exist! Please check it!" + consts.OutReset)
			return
		}

		name, _ := cmd.Flags().GetString("name")
		fmt.Println(name)
		if name == "" {
			name = tools.GenerateName("policy")
			fmt.Println("There is no name set So We generate a policy name (" + consts.ColorYellow + name + consts.OutReset + ") for you! ")
		}
		url, _ := cmd.Flags().GetString("url")
		fmt.Println("policy url:" + consts.ColorYellow + url + consts.OutReset)
		attestationType, _ := cmd.Flags().GetString("type")
		fmt.Println("policy type:" + consts.ColorYellow + attestationType + consts.OutReset)

		m, err := manager.New(url, "")
		if err != nil {
			fmt.Printf("create attest manager failed, error: %s", err)
		}
		res, err := m.SetPolicy(name, fileBase64, attestationType)
		if err != nil {
			fmt.Printf("Request failed: %v", err)
		}
		if res.Code == 200 {
			fmt.Println(consts.ColorGreen + "set policy successful, policy id: " + res.PolicyID + consts.OutReset)
		} else {
			fmt.Println(consts.ColorRed + "set policy failed:" + res.Message + consts.OutReset)
		}
	},
}

func init() {
	Cmd.AddCommand(policySetCmd)
	//set parameter for policy set
	policySetCmd.Flags().StringP("url", "u", "https://api.trustcluster.cn", "optional, tcas's api url")
	policySetCmd.Flags().StringP("name", "n", "", "must, policy name")
	policySetCmd.MarkFlagRequired("name")

	policySetCmd.Flags().StringP("type", "t", "trust_node", "optional, the attestation-type of policy, support `trust_node` or `trust_cluster`")
	policySetCmd.Flags().StringP("rego-file", "f", "", "must, the path of policy file in rego format")
	policySetCmd.MarkFlagRequired("rego-file")
}

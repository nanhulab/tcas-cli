/*
 * @Author: jffan
 * @Date: 2024-07-29 09:48:09
 * @LastEditTime: 2024-08-14 14:30:04
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\cmd\policy\policy.go
 * @Description:
 */
package policy

import (
	"github.com/spf13/cobra"
)

// policyCmd represents the policy command
var Cmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "policy",
	Short: "manager policy",
	Long:  "",
	//SilenceUsage:               true,
	SuggestionsMinimumDistance: 1,
	DisableSuggestions:         false,
}

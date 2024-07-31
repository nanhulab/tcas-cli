/*
 * @Author: jffan
 * @Date: 2024-07-29 09:48:09
 * @LastEditTime: 2024-07-30 16:53:42
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\cmd\policy.go
 * @Description: Copyright © 2024 <jffan@nanhulab.ac.cn>
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

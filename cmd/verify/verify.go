/*
 * @Author: jffan
 * @Date: 2024-08-02 14:59:33
 * @LastEditTime: 2024-08-02 15:25:28
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\verify\verify.go
 * @Description: 🎉🎉🎉
 */
package verify

import (
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var Cmd = &cobra.Command{
	Args:                       cobra.NoArgs,
	Use:                        "verify",
	Short:                      "Verify the legitimacy of the information",
	Long:                       "",
	SuggestionsMinimumDistance: 1,
	DisableSuggestions:         false,
}

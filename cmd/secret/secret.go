/*
 * @Author: jffan
 * @Date: 2024-07-31 15:01:17
 * @LastEditTime: 2024-08-02 15:00:48
 * @LastEditors: jffan
 * @FilePath: \gitee-tcas\cmd\secret\secret.go
 * @Description:
 */
package secret

import (
	"github.com/spf13/cobra"
)

// secretCmd represents the secret command
var Cmd = &cobra.Command{
	Args:                       cobra.NoArgs,
	Use:                        "secret",
	Short:                      "manager secret",
	Long:                       "",
	SuggestionsMinimumDistance: 1,
	DisableSuggestions:         false,
}

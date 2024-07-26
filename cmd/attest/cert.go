package attest

import (
	"fmt"
	"github.com/spf13/cobra"
)

// certCmd represents the cert command
var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cert called")
	},
}

func init() {
	Cmd.AddCommand(certCmd)
}

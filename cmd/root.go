package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"tcas-cli/cmd/attest"
	"tcas-cli/cmd/policy"
	"tcas-cli/cmd/secret"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tcasctl",
	Short: "tcasctl is a client to manage trust cluster attestation server",
	Args:  cobra.NoArgs,
	Long:  "",
	//SilenceUsage:               true,
	SuggestionsMinimumDistance: 1,
	DisableSuggestions:         false,
}

func Execute() {
	RootCmd.AddCommand(attest.Cmd)
	RootCmd.AddCommand(policy.Cmd)
	RootCmd.AddCommand(secret.Cmd)
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

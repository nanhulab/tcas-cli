package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
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
	loglevel := os.Getenv("LogLevel")
	if loglevel != "" {
		level, err := logrus.ParseLevel(loglevel)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse logging level: %s\n", loglevel)
			os.Exit(1)
		}
		logrus.SetLevel(level)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})

	RootCmd.AddCommand(attest.Cmd)
	RootCmd.AddCommand(policy.Cmd)
	RootCmd.AddCommand(secret.Cmd)
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

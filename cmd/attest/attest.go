package attest

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Args:  cobra.NoArgs,
	Use:   "attest",
	Short: "do trust node attest for get token, cert or secret",
	Long:  "",
	//SilenceUsage:               true,
	SuggestionsMinimumDistance: 1,
	DisableSuggestions:         false,
}

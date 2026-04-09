package config

import (
	"tsacl/internal/tailscale"

	"github.com/spf13/cobra"
)


var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "List Current Tailscale ACL rules",
	Long: `
	Applies your settings to the Tailnet. With no flags, the command takes the local state and applies it
	`,
	Run: func(cmd *cobra.Command, args []string) {
		context := tailscale.CurrentTailnetContext("")
		println(context)
	},
}


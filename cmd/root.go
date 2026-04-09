package cmd

import (
	"tsacl/cmd/apply"
	"tsacl/cmd/list"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "tsacl",
	Short: "TSACL is helping you configure Tailscale ACL.",
	Long: `
					* TSACL *

	TSACL is managing configuration and ACLs for different Tailscale contexts (Tailnets).
	You just supply the API key on login and the application will store it in an encrypted and safe manner.
	Each contextreferences both a Tailnet and an API key.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init(){
	rootCmd.AddCommand(apply.ApplyCmd)
	rootCmd.AddCommand((list.ListCmd))
}

func Execute() error {
	return rootCmd.Execute()
}
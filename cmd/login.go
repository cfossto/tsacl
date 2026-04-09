package cmd

import "github.com/spf13/cobra"

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Tailscale",
	Run: func(cmd *cobra.Command, args []string) {
		// logic
	},
}

func init() {
	// Add it to the root command here!
	rootCmd.AddCommand(loginCmd)
}
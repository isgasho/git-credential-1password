package cmd

import "github.com/spf13/cobra"

// eraseCmd represents the store operation.
var eraseCmd = &cobra.Command{
	Use:              "erase",
	TraverseChildren: true,
	Args:             cobra.NoArgs,
	Short:            "Remove a matching credential, if any, from the helper’s storage",
	Run: func(cmd *cobra.Command, args []string) {
		// not supported, ignore
	},
}

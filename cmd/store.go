package cmd

import "github.com/spf13/cobra"

// storeCmd represents the store operation.
var storeCmd = &cobra.Command{
	Use:              "store",
	TraverseChildren: true,
	Args:             cobra.NoArgs,
	Short:            "Store the credential",
	Run: func(cmd *cobra.Command, args []string) {
		// not supported, ignore
	},
}

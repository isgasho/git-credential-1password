package cmd

import (
	"github.com/spf13/cobra"
)

var account string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:               "git-credential-1password [operation]",
	DisableAutoGenTag: true,
	Version:           getVersion(),
	Short:             "Helper to store credentials inside 1password",
	Long: `This command stores credentials inside 1password for use by future Git programs.

You probably don't want to invoke this command directly; it is meant to be used
as as credential helper by other parts of Git.`,
	Example: `The point of this helper is to reduce the number of time you must type your
username and password and to sync your credentials across multiple systems.
For example:

$ git config credential.helper !git-credential-1password
$ git push http://example.com/repo.git
Enter your 1password master password: <type your master password>
[your credentials are used automatically]`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// increase default mousetrap display duration
		cobra.MousetrapDisplayDuration = 0

		if len(args) > 1 {
			// arguments given, disable Mousetrap handler
			cobra.MousetrapHelpText = ""
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute the git-credential-1password cli.
func Execute() error {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(storeCmd)
	rootCmd.AddCommand(eraseCmd)

	getCmd.Flags().SortFlags = true
	storeCmd.Flags().SortFlags = true
	eraseCmd.Flags().SortFlags = true

	rootCmd.PersistentFlags().StringVarP(&account, "account", "a", "my",
		"the shorthand of the 1password account you want to use")

	return rootCmd.Execute()
}

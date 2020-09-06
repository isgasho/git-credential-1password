package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/develerik/git-credential-1password/git"
	"github.com/develerik/git-credential-1password/onepassword"
)

// debugInfoCmd represents the debug-info command.
var debugInfoCmd = &cobra.Command{
	Use:              "debug-info",
	TraverseChildren: true,
	Args:             cobra.NoArgs,
	Hidden:           true,
	Short:            "Print debug information (useful for issue submitting)",
	Run: func(cmd *cobra.Command, args []string) {
		info := map[string]string{}

		info["Version"] = getVersion()

		var err error

		info["Git Version"], err = git.GetVersion()

		if err != nil {
			info["Git Version"] = err.Error()
		}

		info["1Password CLI Version"], err = onepassword.GetVersion()

		if err != nil {
			info["1Password CLI Version"] = err.Error()
		}

		info["Architecture"] = runtime.GOARCH
		info["Operating System"] = runtime.GOOS

		for k, v := range info {
			fmt.Printf("%s: %s\n", k, v)
		}
	},
}

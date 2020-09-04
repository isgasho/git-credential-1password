package main

import (
	"github.com/develerik/git-credential-1password/cmd"
)

func main() {
	_ = cmd.Execute() // nolint:errcheck // error is displayed by cobra
}

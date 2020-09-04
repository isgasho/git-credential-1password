package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/develerik/git-credential-1password/onepassword"
)

// getCmd represents the get operation.
var getCmd = &cobra.Command{
	Use:              "get",
	TraverseChildren: true,
	Args:             cobra.NoArgs,
	Short:            "Return a matching credential, if any exists",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getCredentials(os.Stdin, os.Stdout); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func getCredentials(reader io.Reader, writer io.Writer) error { // nolint:gocognit // TODO: refactor
	c := onepassword.Client{}

	scanner := bufio.NewScanner(reader)

	data := map[string]string{}
	buffer := new(bytes.Buffer)

	for scanner.Scan() {
		keyAndValue := bytes.SplitN(scanner.Bytes(), []byte("="), 2)
		if len(keyAndValue) > 1 {
			data[string(keyAndValue[0])] = string(keyAndValue[1])
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		return err
	}

	if err := c.Login(); err != nil {
		return err
	}

	host, exist := data["host"]

	if !exist {
		return fmt.Errorf("missing host to check credentials: %v", data) // nolint:goerr113 // TODO: refactor
	}

	creds, err := c.GetCredentials(host)

	if err != nil {
		return err
	}

	data["username"] = creds.Username
	data["password"] = creds.Password

	buffer.Reset()

	for key, value := range data {
		if _, err = fmt.Fprintf(buffer, "%s=%s\n", key, value); err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(writer, buffer.String())

	return err
}

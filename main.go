package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/develerik/git-credential-1password/onepassword"
)

const (
	gitHostKey     = "host"
	gitUsernameKey = "username"
	gitPasswordKey = "password"
)

var c = onepassword.Client{}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <get> \n", os.Args[0])
	}

	if err := handleCommand(os.Args[1], os.Stdin, os.Stdout); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleCommand(key string, in io.Reader, out io.Writer) error {
	switch key {
	case "get":
		return getCredentials(in, out)
	}
	return fmt.Errorf("unsupported operation: %s", key)
}

func getCredentials(reader io.Reader, writer io.Writer) error {
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

	host, exist := data[gitHostKey]

	if !exist {
		return fmt.Errorf("missing host to check credentials: %v", data)
	}

	creds, err := c.GetCredentials(host)

	if err != nil {
		return err
	}

	data[gitUsernameKey] = creds.Username
	data[gitPasswordKey] = creds.Password

	buffer.Reset()

	for key, value := range data {
		fmt.Fprintf(buffer, "%s=%s\n", key, value)
	}

	fmt.Fprint(writer, buffer.String())
	return nil
}

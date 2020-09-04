package onepassword

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
)

// Credentials defines git credentials.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Details struct {
		Fields []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"details"`
}

// Client defines a 1password client.
type Client struct {
	token string
}

// GetCredentials loads credentials from 1password.
func (c *Client) GetCredentials(host string) (*Credentials, error) {
	var stdout bytes.Buffer

	var stderr bytes.Buffer

	cmd := exec.Command("op", "--session", c.token, "get", "item", host) // nolint:gosec // TODO: validate
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return nil, errors.New(stderr.String()) // nolint:goerr113 // TODO: refactor
	}

	var res response
	if err := json.Unmarshal(stdout.Bytes(), &res); err != nil {
		return nil, err
	}

	var username string

	var password string

	for _, field := range res.Details.Fields {
		if field.Name == "username" {
			username = field.Value
		}

		if field.Name == "password" {
			password = field.Value
		}
	}

	return &Credentials{
		Username: username,
		Password: password,
	}, nil
}

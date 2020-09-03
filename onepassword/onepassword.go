package onepassword

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"golang.org/x/crypto/ssh/terminal"
)

// Credentials defines git credentials
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

// Client defines a 1password client
type Client struct {
	token string
}

// Login to 1password
func (c *Client) Login() error {
	s := os.Getenv("OP_SESSION_my")

	if s == "" {
		fmt.Fprint(os.Stderr, "Enter your 1password master password: ")

		tty, err := os.Open("/dev/tty")

		if err != nil {
			return err
		}

		defer tty.Close()
		pass, err := terminal.ReadPassword(int(tty.Fd()))

		if err != nil {
			return err
		}

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		var b bytes.Buffer
		b.Write([]byte(fmt.Sprintf("%s\n", pass)))

		cmd := exec.Command("op", "signin", "my")
		cmd.Stdout = &stdout
		cmd.Stdin = &b
		cmd.Stderr = &stderr

		err = cmd.Run()

		if err != nil {
			return errors.New(stderr.String())
		}

		r, err := regexp.Compile("export OP_SESSION_my=\"([a-zA-Z0-9-_]+)\".*")
		res := r.FindStringSubmatch(stdout.String())

		if err != nil {
			return err
		}

		if len(res) < 2 {
			return errors.New("no session token found")
		}

		c.token = res[1]
	} else {
		c.token = s
	}

	os.Setenv("OP_SESSION_my", c.token)

	return nil
}

// GetCredentials loads credentials from 1password
func (c *Client) GetCredentials(host string) (*Credentials, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("op", "--session", c.token, "get", "item", host)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return nil, errors.New(stderr.String())
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

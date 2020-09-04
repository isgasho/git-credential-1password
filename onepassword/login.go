package onepassword

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	errNoSessionToken = errors.New("no session token found")
)

// Login to 1password.
func (c *Client) Login() error {
	c.token = os.Getenv("OP_SESSION_my")

	if c.token != "" {
		return nil
	}

	if _, err := fmt.Fprint(os.Stderr, "Enter your 1password master password: "); err != nil {
		return err
	}

	tty, err := os.Open("/dev/tty")

	if err != nil {
		return err
	}

	defer func() {
		_ = tty.Close()
	}()

	fd := int(tty.Fd())
	pass, err := terminal.ReadPassword(fd)

	if err != nil {
		return err
	}

	var stdout, stderr, stdin bytes.Buffer

	stdin.Write([]byte(fmt.Sprintf("%s\n", pass)))

	cmd := exec.Command("op", "signin", "my")
	cmd.Stdout = &stdout
	cmd.Stdin = &stdin
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return errors.New(stderr.String()) // nolint:goerr113 // TODO: correctly handle error
	}

	m := regexp.
		MustCompile("export OP_SESSION_my=\"([a-zA-Z0-9-_]+)\".*").
		FindStringSubmatch(stdout.String())

	if len(m) < 2 { // nolint:gomnd // see regex
		return errNoSessionToken
	}

	c.token = m[1]

	return nil
}

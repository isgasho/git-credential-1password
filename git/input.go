package git

import (
	"bufio"
	"bytes"
	"io"
)

// ReadInput reads key-value pairs from git output api.
func ReadInput(r io.Reader) (map[string]string, error) {
	scan := bufio.NewScanner(r)

	data := map[string]string{}

	for scan.Scan() {
		kv := bytes.SplitN(scan.Bytes(), []byte("="), 2)
		if len(kv) > 1 {
			data[string(kv[0])] = string(kv[1])
		}
	}

	if err := scan.Err(); err != nil && err != io.EOF {
		return nil, err
	}

	return data, nil
}

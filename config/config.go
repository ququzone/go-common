package config

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

type Config struct {
	data map[string]string
}

func (c *Config) read(buf *bufio.Reader) error {
	for {
		l, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		r := strings.TrimSpace(string(l))
		switch {
		case len(r) == 0: // empty line
			continue

		case r[0] == '#': // comment
			continue

		case r[0] == '=': // error row
			continue

		default:
			vals := bytes.SplitN([]byte(r), []byte{'='}, 2)
			c.data[strings.TrimSpace(string(vals[0]))] = strings.TrimSpace(string(vals[1]))
		}
	}
	return nil
}

func (c *Config) String(key string) string {
	return c.data[key]
}

func NewConfig(fname string) (*Config, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	c := &Config{make(map[string]string)}

	if err := c.read(bufio.NewReader(file)); err != nil {
		return nil, err
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return c, nil
}

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	Path string
}

func (c *Config) SetConfig(path string) {
	c.Path = path
}

func TestSetConfig(t *testing.T) {
	cfg := Config{}
	cfg.SetConfig("/path/to/config")
	fmt.Printf("Config path: %s\n", cfg.Path)

	assert.Equal(t, "/path/to/config", cfg.Path)
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	fmt.Println("Hello ", buf.String())
	Process()

	assert.Contains(t, buf.String(), "Processing data...")
}

package hellonow_test

import (
	"bytes"
	"testing"

	"github.com/stakkato95/hellownow2/hellonow"
	"github.com/stretchr/testify/assert"
)

func TestPrintTime(t *testing.T) {
	writer := new(bytes.Buffer)

	hellonow.PrintTime(writer)

	out := writer.String()
	assert.Contains(t, out, "current time: ")
	assert.Contains(t, out, "exact time: ")
}

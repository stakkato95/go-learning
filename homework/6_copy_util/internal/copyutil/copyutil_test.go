package copyutil

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintErrors(t *testing.T) {
	var buffer bytes.Buffer

	errs := ValidateConfig(CopyConfig{})
	PrintErrors(errs, bufio.NewWriter(&buffer))

	errsString := buffer.String()

	assert.Contains(t, errsString, ErrorMessageHeader)
	assert.Contains(t, errsString, ErrFromIsEmpty.Error())
	assert.Contains(t, errsString, ErrToIsEmpty.Error())
}

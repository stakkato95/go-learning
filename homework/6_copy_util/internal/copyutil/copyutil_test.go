package copyutil

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
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

func TestCopyInternal(t *testing.T) {
	t.Run("test copying even number of bytes", func(t *testing.T) {
		testCopyInternal(t, CopyConfig{
			ChunkSize:    2,
			CopyBytes:    4,
			Offset:       2,
			ShowProgress: false,
			From:         "",
			To:           "",
		})
	})

	t.Run("test copying odd number of bytes", func(t *testing.T) {
		testCopyInternal(t, CopyConfig{
			ChunkSize:    2,
			CopyBytes:    5,
			Offset:       2,
			ShowProgress: false,
			From:         "",
			To:           "",
		})
	})
}

func testCopyInternal(t *testing.T, cfg CopyConfig) {
	inFile, err := os.CreateTemp("", "in")
	if err != nil {
		assert.FailNow(t, "can not create 'in' file")
	}
	defer os.Remove(inFile.Name())

	inFileContent := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	_, err = inFile.Write(inFileContent)
	inFile.Close()
	if err != nil {
		assert.FailNow(t, "can not write to 'in' file")
	}

	outFile, err := os.CreateTemp("", "out")
	if err != nil {
		assert.FailNow(t, "can not create 'out' file")
	}
	defer os.Remove(outFile.Name())

	cfg.From = inFile.Name()
	cfg.To = outFile.Name()
	CopyInternal(cfg)

	actual, err := ioutil.ReadAll(outFile)
	if err != nil {
		assert.FailNow(t, "can not read from 'out' file")
	}

	expected := inFileContent[cfg.Offset : cfg.Offset+cfg.CopyBytes]
	assert.EqualValues(t, expected, actual)
}

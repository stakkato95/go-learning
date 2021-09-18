package copyutil

import (
	"os"
	"testing"

	"github.com/stakkato95/copyutil/internal/copyutil/slices"
	"github.com/stretchr/testify/assert"
)

const tempContentSize = 10

func TestEmptyFromAndTo(t *testing.T) {
	errs := ValidateConfig(CopyConfig{})
	assert.True(t, slices.ContainsAll(errs, ErrFromIsEmpty, ErrToIsEmpty))
}

func TestFromDoesNotExist(t *testing.T) {
	errs := ValidateConfig(CopyConfig{From: "nonExistingFromFile", To: "toFile"})
	assert.True(t, slices.ContainsAll(errs, ErrFfomIsNotExisting))
}

func TestFromIsDir(t *testing.T) {
	if dir, err := os.MkdirTemp("", "tempDir"); err != nil {
		assert.FailNow(t, "can not create tempDir")
	} else {
		errs := ValidateConfig(CopyConfig{From: dir, To: "toFile"})
		assert.True(t, slices.ContainsAll(errs, ErrFromIsDir))
	}
}

func TestFromSizeIsEmpty(t *testing.T) {
	if file, err := os.CreateTemp("", "tempFile"); err != nil {
		assert.FailNow(t, "can not create tempFile")
	} else {
		errs := ValidateConfig(CopyConfig{From: file.Name(), To: "toFile"})
		assert.True(t, slices.ContainsAll(errs, ErrFromSizeIsEmpty))
	}
}

func TestOffsetIsBiggerThanSize(t *testing.T) {
	startTestWithOffsetAndLimit(t, tempContentSize+1, 0, ErrOffsetIsBiggerThanSize)
}

func TestOffsetAndLimitAreBiggerThanSize(t *testing.T) {
	startTestWithOffsetAndLimit(t, tempContentSize-5, 7, ErrOffsetAndLimitAreBiggerThanSize)
}

func startTestWithOffsetAndLimit(t *testing.T, offset, limit int64, expectedErr error) {
	if file, err := os.CreateTemp("", "tempFile"); err != nil {
		assert.FailNow(t, "can not create tempFile")
	} else {
		var contentSize int64 = 10
		file.Write(make([]byte, contentSize))
		file.Close()

		if file, err = os.OpenFile(file.Name(), os.O_RDONLY, 0); err != nil {
			assert.FailNow(t, "can not open tempFile")
		}

		errs := ValidateConfig(CopyConfig{From: file.Name(), To: "toFile", Offset: offset, CopyBytes: limit})
		assert.True(t, slices.ContainsAll(errs, expectedErr))
	}
}

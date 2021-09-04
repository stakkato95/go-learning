package hellownow_test

import (
	"bytes"
	"testing"

	"github.com/stakkato95/hellownow/hellownow"
)

func TestWriteTime(t *testing.T) {
	writer1 := new(bytes.Buffer)
	hellownow.WriteTime(writer1)

	if writer1.String() != "3-9-2021" {
		t.Fatal("wrong output")
	}
}

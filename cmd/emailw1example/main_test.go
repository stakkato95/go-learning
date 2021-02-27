package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
2
1
2`

var testOkResult = `1
2
3
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)

	if err := uniqFromUnsortedInput(in, out); err != nil {
		t.Errorf("test for OK failed")
	}

	if out.String() != testOkResult {
		t.Errorf("test for OK failed - result doesn't match")
	}
}

var testFail = `1
2
3
2
1
2
a`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)

	if err := uniqFromUnsortedInput(in, out); err == nil {
		t.Errorf("test for Error failed")
	}
}

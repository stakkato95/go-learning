package unpack_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stakkato95/unpack_string/unpack"
	"github.com/stretchr/testify/assert"
)

const crasher = "CRASHER"
const data = "DATA"

func TestUnpack(t *testing.T) {
	cases := []struct {
		Intput    string
		Expected  string
		IsCorrect bool
	}{
		{"a4bc2d5e", "aaaabccddddde", true},
		{"abcd", "abcd", true},
		{"3abc", "", false},
		{"45", "", false},
		{"aaa10b", "", false},
		{"aaa0b", "aab", true},
		{"", "", true},
		{"d\n5abc", "d\n\n\n\n\nabc", true},
		{`qwe\4\5`, "qwe45", true},
		{`qwe\45`, "qwe44444", true},
		{`qwe\\5a`, `qwe\\\\\a`, true},
		{`qw\\ne`, "", false},
	}

	for i, caze := range cases {
		t.Run(fmt.Sprintf("case %d. %s => %s", i+1, caze.Intput, caze.Expected), func(t *testing.T) {
			actual, err := unpack.Unpack(caze.Intput)
			assert.Equal(t, caze.Expected, actual)
			if caze.IsCorrect {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	if os.Getenv(crasher) == "1" {
		r := []rune(os.Getenv(data))
		unpack.ParseInt(r[0])
		return
	}

	cases := []struct {
		Input     string
		IsCorrect bool
	}{
		{"a", false},
		{"1", true},
	}

	for i, caze := range cases {
		t.Run(fmt.Sprintf("case %d. ParseInt(%v)", i+1, caze.Input), func(t *testing.T) {
			cmd := exec.Command(os.Args[0], "-test.run=TestParseInt")
			cmd.Env = append(os.Environ(), crasher+"=1", data+"="+caze.Input)

			err := cmd.Run()

			if e, ok := err.(*exec.ExitError); ok {
				//process finished with error
				assert.True(t, !e.Success() && !caze.IsCorrect)
			} else {
				//process finished without error
				assert.True(t, caze.IsCorrect)
			}
		})
	}
}

package unpack_test

import (
	"fmt"
	"testing"

	"github.com/stakkato95/unpack_string/unpack"
	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	cases := [...][3]interface{}{
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
		input := caze[0].(string)
		expected := caze[1].(string)
		isCorrect := caze[2].(bool)
		t.Run(fmt.Sprintf("case %d. %s => %s", i+1, input, expected), func(t *testing.T) {
			actual, err := unpack.Unpack(input)
			assert.Equal(t, expected, actual)
			if isCorrect {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

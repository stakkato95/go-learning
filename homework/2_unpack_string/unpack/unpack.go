package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const errorString = ""
const slash = "\\"

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return s, nil
	}

	var b strings.Builder

	var anyCharBefore bool
	var slashCharBefore bool
	var slashBefore bool
	for i, r := range s {
		if unicode.IsDigit(r) {
			if slashBefore {
				slashBefore = false
				anyCharBefore = true
				continue
			}

			if slashCharBefore {
				slashCharBefore = false
				count, _ := strconv.Atoi(string(r))
				b.WriteString(strings.Repeat(slash, count))
				continue
			}

			if anyCharBefore {
				anyCharBefore = false
				count, _ := strconv.Atoi(string(r))
				b.WriteString(strings.Repeat(string(s[i-1]), count))
			} else {
				return errorString, errors.New("numbers are fobbiden, only figures are allowed")
			}
		} else {
			if anyCharBefore {
				b.WriteByte(s[i-1])
			}

			if string(r) == slash {
				if slashBefore {
					slashCharBefore = true
					slashBefore = false
				} else {
					slashBefore = true
					anyCharBefore = false
				}
			} else {
				if slashBefore || slashCharBefore {
					return errorString, errors.New("only numbers can be slashed")
				}
				anyCharBefore = true
			}
		}
	}

	if anyCharBefore {
		b.WriteByte(s[len(s)-1])
	}

	return b.String(), nil
}

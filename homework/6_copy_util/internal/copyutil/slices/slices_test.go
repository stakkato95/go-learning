package slices

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteFirstElement(t *testing.T) {
	slc := []error{errors.New("a"), errors.New("b"), errors.New("c")}
	Delete(&slc, errors.New("a"))
	assert.EqualValues(t, []error{errors.New("c"), errors.New("b")}, slc)
}

func TestDeleteLastElement(t *testing.T) {
	slc := []error{errors.New("a"), errors.New("b"), errors.New("c")}
	Delete(&slc, errors.New("a"))
	assert.EqualValues(t, []error{errors.New("c"), errors.New("b")}, slc)
}

func TestDeleteMiddleElement(t *testing.T) {
	slc := []error{errors.New("a"), errors.New("b"), errors.New("c")}
	Delete(&slc, errors.New("b"))
	assert.EqualValues(t, []error{errors.New("a"), errors.New("c")}, slc)
}

func TestDeleteMultipleElements(t *testing.T) {
	slc := []error{errors.New("a"), errors.New("b"), errors.New("c"), errors.New("d")}
	Delete(&slc, errors.New("b"))
	Delete(&slc, errors.New("c"))
	assert.EqualValues(t, []error{errors.New("a"), errors.New("d")}, slc)
}

func TestContainsAll(t *testing.T) {
	errs1 := []error{errors.New("a"), errors.New("b"), errors.New("c")}
	errs2 := []error{errors.New("b"), errors.New("a"), errors.New("c")}
	assert.True(t, ContainsAll(errs1, errs2...))
}

func TestContainsNotAll(t *testing.T) {
	errs1 := []error{errors.New("a"), errors.New("b"), errors.New("c")}
	errs2 := []error{errors.New("b"), errors.New("c")}
	assert.False(t, ContainsAll(errs1, errs2...))
}

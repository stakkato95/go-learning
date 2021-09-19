package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringLengthValidation(t *testing.T) {
	person := []struct {
		Name string `validation:"maxLength:4"`
	}{
		{Name: "Jack"},
		{Name: "Anton"},
	}

	t.Run("string length is correct", func(t *testing.T) {
		assert.Empty(t, Validate(&person[0]))
	})

	t.Run("string length is incorrect", func(t *testing.T) {
		errs := Validate(&person[1])
		assert.ErrorIs(t, errs[0], ErrStringFieldLength)
	})
}

func TestStringEndingValidation(t *testing.T) {
	streets := []struct {
		Address string `validation:"endsWith:street"`
	}{
		{Address: "Baker street"},
		{Address: "Streetwall"},
		{Address: "Strasse"},
	}

	t.Run("string ending is found", func(t *testing.T) {
		assert.Empty(t, Validate(&streets[0]))
	})

	t.Run("string ending is not found when there is partial match", func(t *testing.T) {
		errs := Validate(&streets[1])
		assert.ErrorIs(t, errs[0], ErrStringUnexpectedEnding)
	})

	t.Run("string ending is not found when there is no match at all", func(t *testing.T) {
		errs := Validate(&streets[2])
		assert.ErrorIs(t, errs[0], ErrStringUnexpectedEnding)
	})
}

func TestStringMultipleErrorsValidation(t *testing.T) {
	streets := []struct {
		Address string `validation:"maxLength:12|endsWith:street"`
	}{
		{Address: "Baker street"},
		{Address: "Streetwall"},
	}

	t.Run("string ending is found", func(t *testing.T) {
		assert.Empty(t, Validate(&streets[0]))
	})

	t.Run("string ending is not found when there is partial match", func(t *testing.T) {
		errs := Validate(&streets[1])
		assert.ErrorIs(t, errs[0], ErrStringUnexpectedEnding)
	})
}

func TestIntMinValidation(t *testing.T) {
	person := []struct {
		Age int `validation:"min:18"`
	}{
		{Age: 18},
		{Age: 17},
	}

	t.Run("min constraint is not violated", func(t *testing.T) {
		assert.Empty(t, Validate(&person[0]))
	})

	t.Run("min constraint is violated", func(t *testing.T) {
		errs := Validate(&person[1])
		assert.ErrorIs(t, errs[0], ErrIntMinViolation)
	})
}

func TestIntMaxValidation(t *testing.T) {
	person := []struct {
		Age int `validation:"max:100"`
	}{
		{Age: 100},
		{Age: 101},
	}

	t.Run("max constraint is not violated", func(t *testing.T) {
		assert.Empty(t, Validate(&person[0]))
	})

	t.Run("max constraint is violated", func(t *testing.T) {
		errs := Validate(&person[1])
		assert.ErrorIs(t, errs[0], ErrIntMaxViolation)
	})
}

func TestIntMinAndMaxValidation(t *testing.T) {
	person := []struct {
		Age int `validation:"min:18|max:100"`
	}{
		{Age: 50},
		{Age: 17},
	}

	t.Run("min and max constraints are not violated", func(t *testing.T) {
		assert.Empty(t, Validate(&person[0]))
	})

	t.Run("one of min or max constraints is violated", func(t *testing.T) {
		errs := Validate(&person[1])
		assert.ErrorIs(t, errs[0], ErrIntMinViolation)
	})
}

func TestStringFieldWithIntTagValidation(t *testing.T) {
	street := struct {
		Street string `validation:"min:18|max:100"`
	}{Street: "Baker street"}

	errs := Validate(&street)
	assert.Len(t, errs, 2)
	assert.ErrorIs(t, errs[0], ErrConversion)
	assert.ErrorIs(t, errs[1], ErrConversion)
}

func TestIntFieldWithStringTagValidation(t *testing.T) {
	street := struct {
		Age int `validation:"maxLength:18|endsWith:street"`
	}{Age: 10}

	errs := Validate(&street)
	assert.Len(t, errs, 2)
	assert.ErrorIs(t, errs[0], ErrConversion)
	assert.ErrorIs(t, errs[1], ErrConversion)
}

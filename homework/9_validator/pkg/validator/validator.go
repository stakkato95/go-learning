package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrNil                 = errors.New("passed value is nil")
	ErrIsNotPtr            = errors.New("passed value is not pointer")
	ErrPtrValueIsNotStruct = errors.New("pointer of the passed value is not struct")

	ErrConversion = errors.New("type conversion error")

	ErrStringFieldLength      = errors.New("string is too long")
	ErrStringUnexpectedEnding = errors.New("string does not have expected ending")

	ErrIntMinViolation = errors.New("min value violated")
	ErrIntMaxViolation = errors.New("max value violated")
)

const tagName = "validation"
const tagSeparator = "|"
const tagPartsSeparator = ":"
const notPresent = -1

type validatorFunc func(string, interface{}, string) error

var validators = map[string]validatorFunc{
	"maxLength": func(fieldName string, value interface{}, parameter string) error {
		stringVal, ok := value.(string)
		if !ok {
			return fmt.Errorf("value of field '%s' should be of type string: %w", fieldName, ErrConversion)
		}

		lengthLimit, err := strconv.Atoi(parameter)
		if err != nil {
			return fmt.Errorf("parameter for 'maxLength' should be of type int: %w", ErrConversion)
		}

		if len(stringVal) > lengthLimit {
			return fmt.Errorf("len(%s) > %d: %w", fieldName, lengthLimit, ErrStringFieldLength)
		}
		return nil
	},
	"endsWith": func(fieldName string, value interface{}, parameter string) error {
		stringVal, ok := value.(string)
		if !ok {
			return fmt.Errorf("value of field '%s' should be of type string: %w", fieldName, ErrConversion)
		}

		index := strings.LastIndex(stringVal, parameter)

		if index == notPresent || len(stringVal)-len(parameter) != index {
			return fmt.Errorf("field '%s' with value '%s' does not end with '%s': %w", fieldName, stringVal, parameter, ErrStringUnexpectedEnding)
		}
		return nil
	},
	"min": func(fieldName string, value interface{}, parameter string) error {
		intVal, intParameter, err := minMaxParamConversion(fieldName, value, parameter, "min")
		if err != nil {
			return err
		}
		if intVal < intParameter {
			return fmt.Errorf("field '%s' with value '%d' > '%d': %w", fieldName, intVal, intParameter, ErrIntMinViolation)
		}
		return nil
	},
	"max": func(fieldName string, value interface{}, parameter string) error {
		intVal, intParameter, err := minMaxParamConversion(fieldName, value, parameter, "max")
		if err != nil {
			return err
		}
		if intVal > intParameter {
			return fmt.Errorf("field '%s' with value '%d' > '%d': %w", fieldName, intVal, intParameter, ErrIntMaxViolation)
		}
		return nil
	},
}

func Validate(s interface{}) []error {
	if s == nil {
		return []error{ErrNil}
	}

	reflectType := reflect.TypeOf(s)
	if reflectType.Kind() != reflect.Ptr {
		return []error{ErrIsNotPtr}
	}

	dereferencedType := reflectType.Elem()
	if dereferencedType.Kind() != reflect.Struct {
		return []error{ErrPtrValueIsNotStruct}
	}

	errs := []error{}
	structValue := reflect.ValueOf(s).Elem()
	for i := 0; i < dereferencedType.NumField(); i++ {
		fieldInfo := dereferencedType.Field(i)
		fieldValue := structValue.Field(i)

		var genericValue interface{}
		switch fieldValue.Kind() {
		case reflect.String:
			genericValue = fieldValue.String()
		case reflect.Int:
			genericValue = fieldValue.Int()
		}

		tag, ok := fieldInfo.Tag.Lookup(tagName)
		if ok && tag != "" {
			errs = append(errs, handleString(genericValue, fieldInfo.Name, tag)...)
		}
	}

	return errs
}

func handleString(value interface{}, fieldName, tag string) []error {
	parts := strings.Split(tag, tagSeparator)

	errs := []error{}
	for _, part := range parts {
		keyValuePair := strings.Split(part, tagPartsSeparator)
		validatorFun, ok := validators[keyValuePair[0]]
		if ok {
			err := validatorFun(fieldName, value, keyValuePair[1])
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}

func minMaxParamConversion(fieldName string, value interface{}, parameter, parameterName string) (int64, int64, error) {
	intVal, ok := value.(int64)
	if !ok {
		return 0, 0, fmt.Errorf("value of field '%s' should be of type int: %w", fieldName, ErrConversion)
	}
	intParameter, err := strconv.ParseInt(parameter, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parameter for '%s' should be of type int: %w", parameterName, ErrConversion)
	}

	return intVal, intParameter, nil
}

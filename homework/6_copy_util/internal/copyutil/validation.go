package copyutil

import (
	"errors"
	"os"
)

const (
	from                            = "-from"
	to                              = "-to"
	offset                          = "-offset"
	empty                           = "is empty"
	fromEmpty                       = from + " " + empty
	toEmpty                         = to + " " + empty
	fromNotExisting                 = from + " does not exist"
	fromIsDir                       = from + " is a directory"
	fromSizeEmpty                   = from + " " + empty
	offsetIsBiggerThanSize          = offset + " is bigger than size of input file"
	offsetAndLimitAreBiggerThanSize = "offset+limit > 'from' file size"
)

var (
	ErrFromIsEmpty                     = errors.New(fromEmpty)
	ErrToIsEmpty                       = errors.New(toEmpty)
	ErrFfomIsNotExisting               = errors.New(fromNotExisting)
	ErrFromIsDir                       = errors.New(fromIsDir)
	ErrFromSizeIsEmpty                 = errors.New(fromSizeEmpty)
	ErrOffsetIsBiggerThanSize          = errors.New(offsetIsBiggerThanSize)
	ErrOffsetAndLimitAreBiggerThanSize = errors.New(offsetAndLimitAreBiggerThanSize)
)

func ValidateConfig(config CopyConfig) []error {
	errs := []error{}

	var isFromEmpty bool
	if config.From == "" {
		errs = append(errs, ErrFromIsEmpty)
		isFromEmpty = true
	}
	if config.To == "" {
		errs = append(errs, ErrToIsEmpty)
	}
	if isFromEmpty {
		return errs
	}

	if stat, err := os.Stat(config.From); err != nil {
		if os.IsNotExist(err) {
			errs = append(errs, ErrFfomIsNotExisting)
		}
	} else {
		if stat.IsDir() {
			errs = append(errs, ErrFromIsDir)
		} else if stat.Size() == 0 {
			errs = append(errs, ErrFromSizeIsEmpty)
		} else if stat.Size() <= int64(config.Offset) {
			errs = append(errs, ErrOffsetIsBiggerThanSize)
		} else if stat.Size() <= config.Offset+config.CopyBytes {
			errs = append(errs, ErrOffsetAndLimitAreBiggerThanSize)
		}
	}

	return errs
}

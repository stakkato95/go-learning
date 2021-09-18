package slices

func Contains(slc []interface{}, element interface{}) bool {
	for _, item := range slc {
		if item == element {
			return true
		}
	}
	return false
}

func ContainsAll(slc []error, slcDelete ...error) bool {
	if len(slc) != len(slcDelete) {
		return false
	}

	for _, item := range slc {
		var matchItemFound bool

		for _, itemToDelete := range slcDelete {
			if item.Error() == itemToDelete.Error() {
				matchItemFound = true
				break
			}
		}

		if !matchItemFound {
			return false
		}
	}
	return true
}

func ToSliceOfInterfaces(errs []error) []interface{} {
	result := make([]interface{}, len(errs))
	for i, err := range errs {
		result[i] = err
	}
	return result
}

func Delete(slc *[]error, err error) {
	length := len(*slc)

	for i := 0; i < length; i++ {
		if (*slc)[i].Error() == err.Error() {
			(*slc)[i] = (*slc)[length-1]
			*slc = (*slc)[:length-1]
			length--
		}
	}
}

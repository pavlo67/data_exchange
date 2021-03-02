package structures

import "github.com/pavlo67/common/common/errors"

type ErrorsMap map[int]map[string]errors.Error

func (errorsMap ErrorsMap) Stat() ErrorsStat {
	var errorsStat ErrorsStat

	if errorsMap == nil {
		return errorsStat
	}

	var errKeys []errors.Key
	errorsStat.Fields = map[string]int{}

	for _, errs := range errorsMap {
		errorsStat.Total += len(errs)
	ITEM_ERRORS:
		for field, err := range errs {
			errorsStat.Fields[field]++

			errKey := err.Key()
			for _, errKeyAlready := range errKeys {
				if errKeyAlready == errKey {
					continue ITEM_ERRORS
				}
			}
			errKeys = append(errKeys, errKey)
		}

	}
	errorsStat.Distinct = len(errKeys)

	return errorsStat
}

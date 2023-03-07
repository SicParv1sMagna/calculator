package uniq

import "errors"

func UncountFields(num_fields int, fields *[]string, regiPtr bool) error {
	if num_fields > len(*fields) {
		return errors.New("fields len less then number of fields you want to decrease")
	}

	if num_fields < 0 {
		return errors.New("negative number of fields")
	}
	tempArr := *fields
	*fields = tempArr[num_fields:]
	UniqStrings(fields, regiPtr)
	return nil
}

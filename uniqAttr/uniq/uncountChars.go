package uniq

import (
	"errors"
	"strings"
)

func UncountChars(num_chars int, fields *[]string, regiPtr bool) error {
	tempArr := strings.Split(strings.Join(*fields, "\n"), "")
	if num_chars > len(tempArr) {
		return errors.New("you want to decrease too much chars")
	}

	if num_chars < 0 {
		return errors.New("you want to decrease negative number")
	}
	tempArr = tempArr[num_chars:]
	tempArr = strings.Split(strings.Join(tempArr, ""), "\n")
	*fields = tempArr
	UniqStrings(fields, regiPtr)
	return nil
}

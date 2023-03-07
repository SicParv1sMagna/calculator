package uniq

import (
	"strings"
)

func UniqStrings(fields *[]string, regiPtr bool) {
	var uniq []string
	tempArr := *fields

	if regiPtr {
		for i := 0; i < len(tempArr)-1; i++ {
			count := 0
			for j := i + 1; j < len(tempArr); j++ {
				if strings.ToUpper(tempArr[i]) != strings.ToUpper(tempArr[j]) {
					uniq = append(uniq, tempArr[i])
					i += count
					break
				}
				count++
			}
		}

		if strings.ToUpper(tempArr[len(tempArr)-1]) != strings.ToUpper(uniq[len(uniq)-1]) {
			for i := len(tempArr) - 2; i > 0; i-- {
				if strings.ToUpper(tempArr[len(tempArr)-1]) != strings.ToUpper(tempArr[i]) {
					uniq = append(uniq, tempArr[i+1])
					break
				}
			}
		}
	} else {
		for i := 0; i < len(tempArr)-1; i++ {
			count := 0
			for j := i + 1; j < len(tempArr); j++ {
				if tempArr[i] != tempArr[j] {
					uniq = append(uniq, tempArr[i])
					i += count
					break
				}
				count++
			}
		}

		if tempArr[len(tempArr)-1] != uniq[len(uniq)-1] {
			uniq = append(uniq, tempArr[len(tempArr)-1])
		}
	}

	*fields = uniq
}

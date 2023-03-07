package uniq

import (
	"strings"
)

func RegisterLines(fields *[]string) {
	var uniq []string
	tempArr := *fields

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
		uniq = append(uniq, tempArr[len(tempArr)-1])
	}

	*fields = uniq
}

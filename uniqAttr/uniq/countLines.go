package uniq

import (
	"strings"
)

func CountLines(fields *[]string, regiPtr bool) []int {
	var num_of_coincedence []int
	tempArr := *fields

	if regiPtr {
		for i := 0; i < len(tempArr)-1; i++ {
			count := 0
			for j := i + 1; j < len(tempArr); j++ {
				if strings.ToUpper(tempArr[i]) != strings.ToUpper(tempArr[j]) {
					num_of_coincedence = append(num_of_coincedence, count+1)
					i += count
					break
				}
				count++
			}
		}

		if strings.ToUpper(tempArr[len(tempArr)-1]) == strings.ToUpper(tempArr[len(tempArr)-1]) {
			count := 1
			for i := len(tempArr) - 2; i != 0; i-- {
				if strings.ToUpper(tempArr[i]) != strings.ToUpper(tempArr[len(tempArr)-1]) {
					num_of_coincedence = append(num_of_coincedence, count)
					break
				}
				count++
			}
		}
	} else {
		for i := 0; i < len(tempArr)-1; i++ {
			count := 0
			for j := i + 1; j < len(tempArr); j++ {
				if tempArr[i] != tempArr[j] {
					num_of_coincedence = append(num_of_coincedence, count+1)
					i += count
					break
				}
				count++
			}
		}

		if tempArr[len(tempArr)-1] == tempArr[len(tempArr)-1] {
			count := 1
			for i := len(tempArr) - 2; i != 0; i-- {
				if tempArr[i] != tempArr[len(tempArr)-1] {
					num_of_coincedence = append(num_of_coincedence, count)
					break
				}
				count++
			}
		}
	}

	return num_of_coincedence
}

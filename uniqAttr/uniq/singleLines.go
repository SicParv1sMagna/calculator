package uniq

func SingleLines(fields *[]string, regiPtr bool) {
	var only []string
	tempArr := *fields

	num_of_coincedence := CountLines(&tempArr, regiPtr)
	UniqStrings(&tempArr, regiPtr)

	for i := 0; i < len(num_of_coincedence); i++ {
		if num_of_coincedence[i] == 1 {
			only = append(only, tempArr[i])
		}
	}

	*fields = only
}

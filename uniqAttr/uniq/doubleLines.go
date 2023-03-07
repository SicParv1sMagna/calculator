package uniq

func DoubleLines(fields *[]string, regiPtr bool) {
	var double []string
	tempArr := *fields

	num_of_coincedence := CountLines(&tempArr, regiPtr)
	UniqStrings(&tempArr, regiPtr)

	for i := 0; i < len(num_of_coincedence); i++ {
		if num_of_coincedence[i] > 1 {
			double = append(double, tempArr[i])
		}
	}

	*fields = double
}

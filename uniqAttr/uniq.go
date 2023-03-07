package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"uniq/uniq"
)

func main() {
	num_fields := flag.Int("f", 0, "Не учитывать первые n полей в строке")
	num_chars := flag.Int("s", 0, "Не учитывать первые n символов в строке")
	copyPtr := flag.Bool("c", false, "Подсчитать количество встреч строки во входных данных")
	doublePtr := flag.Bool("d", false, "Вывод тех строк, которые повторились во входных данных")
	uniqPtr := flag.Bool("u", false, "Вывод тех строк, которые не повторились во входных данных")
	regiPtr := flag.Bool("i", false, " Не учитывать регистр букв")
	flag.Parse()

	input, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	if *copyPtr && *doublePtr || *copyPtr && *uniqPtr || *doublePtr && *uniqPtr {
		fmt.Println("Параллельно параметры -c, -d, -u не имеют никакого смысла")
		os.Exit(1)
	}

	fields := strings.Split(string(input), "\n")

	if *num_fields != 0 {
		uniq.UncountFields(*num_fields, &fields, *regiPtr)
	}

	if *num_chars != 0 {
		uniq.UncountChars(*num_chars, &fields, *regiPtr)
	}

	if *copyPtr {
		num_of_coincedence := uniq.CountLines(&fields, *regiPtr)
		uniq.UniqStrings(&fields, *regiPtr)
		for i := 0; i < len(num_of_coincedence); i++ {
			fields[i] = strconv.Itoa(num_of_coincedence[i]) + " " + fields[i]
		}
	}

	if *doublePtr {
		uniq.DoubleLines(&fields, *regiPtr)
	}

	if *uniqPtr {
		uniq.SingleLines(&fields, *regiPtr)
	}

	if *regiPtr && (!*doublePtr || !*uniqPtr || !*copyPtr || *num_fields == 0 || *num_chars == 0) {
		uniq.RegisterLines(&fields)
	}

	if !*regiPtr && !*doublePtr && !*uniqPtr && !*copyPtr && *num_fields == 0 && *num_chars == 0 {
		uniq.UniqStrings(&fields, *regiPtr)
	}

	if flag.Arg(1) == "" {
		for i := 0; i < len(fields); i++ {
			fmt.Println(fields[i])
		}
	} else {
		f, err := os.Create("output.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		for i := 0; i < len(fields); i++ {
			f.WriteString(fields[i] + "\n")
		}
	}
}

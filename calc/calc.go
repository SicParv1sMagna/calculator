package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrCalculate = errors.New("unexpected literal")

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return "error: Stack is empty"
	} else {
		element := (*s)[len(*s)-1]
		return element
	}
}

func main() {
	var expr string
	fmt.Scanln(&expr)

	RPN, _ := Get_RPN(expr)
	fmt.Println(RPN)
	Calculate(RPN)
}

func Get_RPN(src string) ([]string, error) {

	var dst []string
	var op Stack

	var j int
	var t string
	for i := 0; i < len(src); {
		if i < len(src) && (string(src[i]) >= "0" && string(src[i]) <= "9") {
			for t = ""; i < len(src) && (string(src[i])) >= "0" && string(src[i]) <= "9"; i++ {
				t = t + string(src[i])
			}
			dst = append(dst, t)
			j++
		}

		if i < len(src) && string(src[i]) == "(" {
			op.Push(string(src[i]))
			i++
		} else if i < len(src) && string(src[i]) == ")" {
			for !op.IsEmpty() && op.Top() != "(" {
				dst = append(dst, op.Pop())
				j++
			}
			op.Pop()
			i++
		} else if i < len(src) && op.IsEmpty() {
			op.Push(string(src[i]))
			i++
		} else if i < len(src) && op.Top() == "(" {
			op.Push(string(src[i]))
			i++
		} else if i < len(src) && (string(src[i]) == "*" || string(src[i]) == "/") && (op.Top() == "+" || op.Top() == "-") {
			op.Push(string(src[i]))
			i++
		} else if i < len(src) {
			dst = append(dst, op.Pop())
			j++
		}
	}
	for !op.IsEmpty() {
		dst = append(dst, op.Pop())
		j++
	}
	return dst, nil
}

func Calculate(RPN []string) error {
	var num Stack

	for i := 0; i < len(RPN); i++ {
		if RPN[i] != "" {
			if string(RPN[i]) >= "0" && string(RPN[i]) <= "9" {
				num.Push(RPN[i])
			} else if RPN[i] == "+" {
				x, _ := strconv.Atoi(num.Pop())
				y, _ := strconv.Atoi(num.Pop())
				res := x + y
				num.Push(strconv.Itoa(res))
			} else if RPN[i] == "-" {
				x, _ := strconv.Atoi(num.Pop())
				y, _ := strconv.Atoi(num.Pop())
				res := y - x
				num.Push(strconv.Itoa(res))
			} else if RPN[i] == "*" {
				x, _ := strconv.Atoi(num.Pop())
				y, _ := strconv.Atoi(num.Pop())
				res := y * x
				num.Push(strconv.Itoa(res))
			} else if RPN[i] == "/" {
				x, _ := strconv.Atoi(num.Pop())
				y, _ := strconv.Atoi(num.Pop())
				res := y / x
				num.Push(strconv.Itoa(res))
			} else {
				return ErrCalculate
			}
		} else {
			return nil
		}
	}
	fmt.Println(num.Pop())
	return nil
}

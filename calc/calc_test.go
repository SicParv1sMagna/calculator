package main

import (
	"testing"
)

func TestGetRPN_1(t *testing.T) {
	var src = "30*(17+13)"

	_, err := Get_RPN(src)
	if err != nil {
		t.Fatalf("Check failed on get RPN #1: %s", err)
	}
}

func TestGetRPN_2(t *testing.T) {
	var src = "758*217/(1+5)*4/3"

	_, err := Get_RPN(src)
	if err != nil {
		t.Fatalf("Check failed on get RPN #2: %s", err)
	}
}

func TestGetRPN_3(t *testing.T) {
	var src = "(496+874)/12*3-(328+1296)*16"

	_, err := Get_RPN(src)
	if err != nil {
		t.Fatalf("Check failed on get RPN #3: %s", err)
	}
}

func TestCalculate_1(t *testing.T) {
	var src = []string{"1", "3", "+"}

	err := Calculate(src)
	if err != nil {
		t.Fatalf("Check failed on Calculating #1: %s", err)
	}
}

func TestCalculate_2(t *testing.T) {
	var src = []string{"30", "1", "2", "+", "*"}

	err := Calculate(src)
	if err != nil {
		t.Fatalf("Check failed on Calculating #1: %s", err)
	}
}

func TestCalculate_3(t *testing.T) {
	var src = []string{"375", "375", "/", "375", "*"}

	err := Calculate(src)
	if err != nil {
		t.Fatalf("Check failed on Calculating #1: %s", err)
	}
}

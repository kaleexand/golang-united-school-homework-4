package string_sum

//package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func pase_num(s string, i *int) (num string) {
	for {
		if len(s) == *i || string(s[*i]) == "+" || string(s[*i]) == "-" {
			return num
		}
		num += string(s[*i])
		*i++
	}
}

func parseOperand(str string, cursor *int) (firstNum int, err error) {
	firstNum = 1
	if string(str[*cursor]) == "-" {
		firstNum = -1
		*cursor++
	}
	if string(str[*cursor]) == "+" {
		*cursor++
	}
	if *cursor == len(str) {
		return 0, fmt.Errorf("error while calculating sum: %w", errorNotTwoOperands)
	}
	first, err := strconv.Atoi(pase_num(str, cursor))
	if err != nil {
		e := err.(*strconv.NumError)
		return 0, fmt.Errorf("error while calculating sum: %w", e)
	}
	firstNum *= first
	return
}

func StringSum(input string) (output string, err error) {
	cursor := 0
	if len(input) == 0 {
		return "", fmt.Errorf("string is empty: %w", errorEmptyInput)
	}
	trimResult := strings.ReplaceAll(input, " ", "")
	//fmt.Println(trimResult)
	firstNum, err := parseOperand(trimResult, &cursor)
	secondNum, err := parseOperand(trimResult, &cursor)
	if cursor != len(trimResult) {
		return "", fmt.Errorf("error while calculating sum: %w", errorNotTwoOperands)
	}
	if err != nil {
		return "", fmt.Errorf("error while calculating sum: %w", errorNotTwoOperands)
	}
	output = strconv.Itoa(firstNum + secondNum)
	return output, nil
}

//
//func main() {
//	in := "30+5"
//	fmt.Println(StringSum(in))
//}

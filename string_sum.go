//package string_sum

package main

import (
	"errors"
	"fmt"
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

func pase_num(str string) (num string) {
	for i := 0; i < len(str); i++ {
		if str[i] != 43 || string(str[i]) != "-" {
			num += string(str[i])
			//fmt.Println(str[i])
		} else {
			break
		}
	}
	return num
}
func StringSum(input string) (output string, err error) {

	if len(input) == 0 {
		return "", fmt.Errorf("string is empty: %w", errorEmptyInput)
	}
	trimResult := strings.TrimSpace(input)
	if string(trimResult[0]) == "+" {
		trimResult = trimResult[1 : len(trimResult)-1]
	}
	first := ""
	if string(trimResult[0]) == "-" {
		first = "-"
	}
	first += pase_num(trimResult)
	fmt.Println(first)
	return output, nil
}

func main() {
	in := " +30 + 5"
	fmt.Println(StringSum(in))
}

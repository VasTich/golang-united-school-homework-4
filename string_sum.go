package string_sum

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

func StringSum(input string) (output string, err error) {
	if len(strings.TrimSpace(input)) == 0 {
		return "", fmt.Errorf("%s", errorEmptyInput.Error())
	}

	pos := 0
	var strNumbers []string
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' {
			strNumbers = append(strNumbers, input[pos:i])
			pos = i
		}
	}

	strNumbers = append(strNumbers, input[pos:len(input)])

	var noEmptyStrNumbers []string
	for i := 0; i < len(strNumbers); i++ {
		if len(strings.TrimSpace(strNumbers[i])) != 0 {
			noEmptyStrNumbers = append(noEmptyStrNumbers, strNumbers[i])
		}
	}

	var numberCount int = 2
	if len(noEmptyStrNumbers) > numberCount {
		return "", fmt.Errorf("%s", errorNotTwoOperands.Error())
	}

	var sum int = 0
	for i := 0; i < len(noEmptyStrNumbers); i++ {
		num, err := strconv.Atoi(strings.ReplaceAll(strings.TrimSpace(noEmptyStrNumbers[i]), " ", ""))
		if err != nil {
			return "", fmt.Errorf("%s", err.Error())
		}

		sum = sum + num

	}

	output = strconv.Itoa(sum)
	return output, nil
}

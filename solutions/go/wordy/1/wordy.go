package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func calculate(firstnumber, secondnumber int, operand string) (int, bool) {
	var val int
	var expectError bool

	switch operand {
	case "+":
		val = firstnumber + secondnumber
		expectError = true
	case "-":
		val = firstnumber - secondnumber
		expectError = true
	case "*":
		val = firstnumber * secondnumber
		expectError = true
	case "/":
		val = firstnumber / secondnumber
		expectError = true
	}

	return val, expectError
}

func Answer(question string) (int, bool) {
	var expectError bool
	var result int

	re := regexp.MustCompile(`plus|minus|multiplied|divided|\+|\-|\*|\\`)

	r := strings.NewReplacer("What is ", "", "?", "", "plus", "+", "minus", "-", "multiplied by", "*", "divided by", "/")

	a := strings.Split(r.Replace(question), " ")

	if len(a) < 1 {
		return 0, false
	} else if len(a) == 1 {
		result, _ = strconv.Atoi(a[0])
		return result, true
	} else if len(a) <= 2 {
		return 0, false
	} else if len(re.FindAllString(strings.Join(a, ""), -1)) <= 1 && len(a) > 3 {
		return 0, false
	} else if !re.MatchString(question) {
		return 0, false
	}

	firstnumber, err := strconv.Atoi(a[0])
	if err != nil {
		return 0, false
	}

	secondnumber, err := strconv.Atoi(a[2])
	if err != nil {
		return 0, false
	}

	operand := a[1]

	result, expectError = calculate(firstnumber, secondnumber, operand)

	if len(a) > 3 && len(a) >= 5 {
		thirdnumber, err := strconv.Atoi(a[4])
		if err != nil {
			return 0, false
		}

		operand = a[3]
		result, expectError = calculate(result, thirdnumber, operand)

		return result, expectError
	}

	return result, expectError
}

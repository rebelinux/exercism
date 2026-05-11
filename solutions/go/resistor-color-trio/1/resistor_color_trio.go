package resistorcolortrio

import (
	"fmt"
	"os"
	"strconv"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

var resistor = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

var multiplier = map[string]int{
	"ohms":     0,
	"kiloohms": 1000,
	"megaohms": 1000000,
	"gigaohms": 1000000000,
}

var cero = map[int]string{
	0: "0",
	1: "0",
	2: "00",
	3: "000",
	4: "0000",
	5: "00000",
	6: "000000",
	7: "0000000",
	8: "00000000",
	9: "000000000",
}

func ohmsmetric(i int) string {

	var metric string

	switch {
	case i <= 3:
		metric = "ohms"
	case i <= 7 && i > 3:
		metric = "kiloohms"
	case i <= 9 && i > 7:
		metric = "megaohms"
	case i >= 10:
		metric = "gigaohms"
	}

	return metric
}

func convert(i []int) string {

	var a string

	if i[2] == 0 {
		a = fmt.Sprintf("%d%d", i[0], i[1])
	} else {
		a = fmt.Sprintf("%d%d%s", i[0], i[1], cero[i[2]])
	}

	return a
}

func Label(colors []string) string {

	var code []int
	var i int
	var output string

	for _, v := range colors {
		code = append(code, resistor[v])
	}

	i, err := strconv.Atoi(convert(code))

	if err != nil {
		fmt.Println("Error during conversion")
		os.Exit(1)
	}

	o := ohmsmetric(len(strconv.Itoa(i)))

	if multiplier[o] == 0 {
		output = fmt.Sprintf("%d %s", i, o)
	} else {
		output = fmt.Sprintf("%d %s", i/multiplier[o], o)
	}

	return output
}

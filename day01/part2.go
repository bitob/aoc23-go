package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type digit struct {
	text  string
	value int
}

var digits = []digit{
	digit{
		text:  "one",
		value: 1,
	},
	digit{
		text:  "two",
		value: 2,
	},
	digit{
		text:  "three",
		value: 3,
	},
	digit{
		text:  "four",
		value: 4,
	},
	digit{
		text:  "five",
		value: 5,
	},
	digit{
		text:  "six",
		value: 6,
	},
	digit{
		text:  "seven",
		value: 7,
	},
	digit{
		text:  "eight",
		value: 8,
	},
	digit{
		text:  "nine",
		value: 9,
	},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculate_calibration_value(file_name string) uint32 {
	var calibration_value uint32 = 0

	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// TODO: Find out if go has an optional type concept
		var first_digit int32 = -1
		var second_digit int32 = -1

		line := scanner.Text()

		for i, c := range line {
			if unicode.IsDigit(c) {
				if first_digit == -1 {
					first_digit = int32(c) - int32('0')
				}
				second_digit = int32(c) - int32('0')
			} else {
				for _, d := range digits {
					if len(line) >= (i + len(d.text)) {
						slice := line[i : i+len(d.text)]

						if d.text == slice {
							if first_digit == -1 {
								first_digit = int32(d.value)
							}
							second_digit = int32(d.value)
						}
					}
				}
			}
		}
		fmt.Println("First: ", first_digit, "Second: ", second_digit)
		calibration_value += uint32(first_digit)*10 + uint32(second_digit)
	}

	return calibration_value
}

func main() {
	// calibration_value := calculate_calibration_value("./test_input2.txt")
	calibration_value := calculate_calibration_value("./input.txt")

	fmt.Println(calibration_value)
}

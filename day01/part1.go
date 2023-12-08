package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

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

		for _, c := range line {
			if unicode.IsDigit(c) {
				if first_digit == -1 {
					first_digit = int32(c) - int32('0')
				}
				second_digit = int32(c) - int32('0')
			}
		}
		calibration_value += uint32(first_digit)*10 + uint32(second_digit)
	}

	return calibration_value
}

func main() {
	// calibration_value := calculate_calibration_value("./test_input1.txt")
	calibration_value := calculate_calibration_value("./input.txt")

	fmt.Println(calibration_value)
}

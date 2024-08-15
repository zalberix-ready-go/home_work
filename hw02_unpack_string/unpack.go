package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var repeatChar string
	var unpucked strings.Builder

	for _, v := range input {
		char := string(v)

		if repeatChar == "" {
			if _, err := strconv.Atoi(char); err == nil {
				return "", ErrInvalidString
			}
			repeatChar = char
			continue
		}

		multiplier, err := strconv.Atoi(char)
		if err != nil {
			unpucked.WriteString(repeatChar)
			repeatChar = char
			continue
		}

		repeatString := strings.Repeat(repeatChar, multiplier)

		unpucked.WriteString(repeatString)

		repeatChar = ""
	}

	// last char
	unpucked.WriteString(repeatChar)

	return unpucked.String(), nil
}

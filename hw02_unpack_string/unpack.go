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

	nextCharEscape := false

	for _, v := range input {
		char := string(v)

		if char == `\` {
			if !nextCharEscape {
				nextCharEscape = true
				continue
			}

			// Если есть символ до экранированного, то его надо сохранить
			if repeatChar != "" {
				unpucked.WriteString(repeatChar)
			}

			repeatChar = char

			nextCharEscape = false

			continue
		}

		if nextCharEscape {
			// Если есть символ до экранированного, то его надо сохранить
			if repeatChar != "" {
				unpucked.WriteString(repeatChar)
			}

			// Если хотят экранировать не число, то это ошибка
			_, err := strconv.Atoi(char)
			if err != nil {
				return "", ErrInvalidString
			}

			repeatChar = char

			nextCharEscape = false

			continue
		}

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

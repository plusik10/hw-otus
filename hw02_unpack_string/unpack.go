package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if IsDigit(rune(str[0])) {
		return "", ErrInvalidString
	}
	sb := strings.Builder{}
	r := []rune(str)
	for i := 1; i <= len(r)-1; i++ {
		cur := r[i]
		prev := r[i-1]
		switch {
		case IsDigit(cur) && !IsDigit(prev):
			num, _ := strconv.Atoi(string(cur))
			text := strings.Repeat(string(prev), num)
			sb.WriteString(text)
		case IsDigit(cur) && IsDigit(prev):
			return "", ErrInvalidString
		case IsDigit(prev) && !IsDigit(cur):
			continue
		case !IsDigit(cur) && !IsDigit(prev):
			sb.WriteRune(prev)
		}
	}
	if !IsDigit(r[len(r)-1]) {
		sb.WriteRune(r[len(r)-1])
	}
	return sb.String(), nil
}

func IsDigit(r rune) bool {
	if r > 47 && r < 58 {
		return true
	}
	return false
}

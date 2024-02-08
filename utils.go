package bankverification

import (
	"strconv"
	"strings"
)

const (
	AlgorithmMod10 = "mod10"
	AlgorithmMod11 = "mod11"
)

func padLeft(s string, length int, padChar rune) string {
	if len(s) > length {
		return s
	}

	padding := strings.Repeat(string(padChar), length-len(s))

	return padding + s
}

func mod10(value string) bool {
	runes := []rune(value)
	sum := 0

	for i := len(runes) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(runes[i]))
		weight := 1
		if (len(runes)-i)%2 == 0 {
			weight = 2
		}
		tmp := digit * weight
		if tmp > 9 {
			tmp -= 9
		}
		sum += tmp
	}

	return sum%10 == 0
}

func mod11(value string) bool {
	if len(value) > 11 {
		return false
	}

	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	runes := []rune(value)
	sum := 0

	for i := len(runes) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(runes[i]))
		sum += digit * weights[len(runes)-1-i]
	}

	return sum%11 == 0
}

func filter(slice []string, predicate func(string) bool) []string {
	var result []string
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

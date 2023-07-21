package inputparser

import (
	"strconv"
)

func ParseInput(input []string) (ints []int, floats []float64, strings []string) {
	for _, str := range input {
		if i, err := strconv.Atoi(str); err == nil {
			ints = append(ints, i)
		} else if f, err := strconv.ParseFloat(str, 64); err == nil {
			floats = append(floats, f)
		} else {
			strings = append(strings, str)
		}
	}
	return ints, floats, strings
}
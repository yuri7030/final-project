package sorter

import (
    "sort"
)

type MixedSorter []interface{}

func SortInts(arr []int) {
    sort.Ints(arr)
}

func SortFloats(arr []float64) {
    sort.Float64s(arr)
}

func SortStrings(arr []string) {
    sort.Strings(arr)
}

func Less(a, b interface{}) bool {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return a.(int) < b.(int)
		case float64:
			return float64(a.(int)) < b.(float64)
		}
	case float64:
		switch b.(type) {
		case int:
			return a.(float64) < float64(b.(int))
		case float64:
			return a.(float64) < b.(float64)
		}
	}
	panic("unsupported type")
}

func SortMixedNumber(ints []int, floats []float64) MixedSorter {
	var mixed MixedSorter
	for _, i := range ints {
		mixed = append(mixed, i)
	}
	for _, f := range floats {
		mixed = append(mixed, f)
	}

	sort.Slice(mixed, func(i, j int) bool {
		return Less(mixed[i], mixed[j])
	})

	return mixed
}

func SortMixedArray(arr MixedSorter) {
	var ints []int
	var floats []float64
	var strings []string

	for _, v := range arr {
		switch v.(type) {
		case int:
			ints = append(ints, v.(int))
		case float64:
			floats = append(floats, v.(float64))
		case string:
			strings = append(strings, v.(string))
		}
	}

	SortInts(ints)
	SortFloats(floats)
	SortStrings(strings)

	var sorted MixedSorter
	sortedMixed := SortMixedNumber(ints, floats)
	for _, v := range sortedMixed {
		sorted = append(sorted, v)
	}
	for _, v := range strings {
		sorted = append(sorted, v)
	}

	copy(arr, sorted)
}
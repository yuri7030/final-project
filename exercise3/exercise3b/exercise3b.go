package main

import (
	"fmt"
	"strings"
	"unicode"
)

func numDifferentIntegers(word string) int {
	uniqueIntegers := make(map[string]bool)

	runes := []rune(word)
	for i, r := range runes {
		if !unicode.IsDigit(r) {
			runes[i] = ' '
		}
	}

	tokens := splitBySpace(string(runes))

	for _, token := range tokens {
		num := removeLeadingZeros(token)
		if num != "" {
			uniqueIntegers[num] = true
		}
	}

	return len(uniqueIntegers)
}

func splitBySpace(s string) []string {
	result := []string{}
	tokens := strings.Split(s, " ")
	for _, token := range tokens {
		if token != "" {
			result = append(result, token)
		}
	}
	return result
}

func removeLeadingZeros(s string) string {
	i := 0
	for i < len(s) && s[i] == '0' {
		i++
	}
	if i == len(s) {
		return "0"
	}
	return s[i:]
}

func main() {
	word := "a1b01c001"
	count := numDifferentIntegers(word)
	fmt.Printf("%v", count)
}

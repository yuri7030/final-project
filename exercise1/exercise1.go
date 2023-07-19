package main

import (
	"fmt"
	"os"
	"strings"
)

func reorderName(firstName, lastName, middleName, countryCode string) string {
	switch countryCode {
	case "VN", "CN", "JP", "KR", "IN", "TH", "ID", "PH", "MY", "SG":
		return fmt.Sprintf("%s %s", lastName, firstName)
	case "US", "CA", "GB", "FR", "DE", "IT", "ES", "AU", "NZ", "BR":
		if middleName != "" {
			return fmt.Sprintf("%s %s %s", firstName, lastName, middleName)
		}
		return fmt.Sprintf("%s %s", firstName, lastName)
	default:
		return fmt.Sprintf("%s %s %s", firstName, middleName, lastName)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("Not enough arguments.")
		fmt.Println("Usage: go run reorder_names.go <first_name> <last_name> <country_code> [middle_name]")
		return
	}

	firstName := args[0]
	lastName := args[1]
	countryCode := strings.ToUpper(args[2])
	middleName := ""
	if len(args) == 4 {
		middleName = args[2]
		countryCode = strings.ToUpper(args[3])
	}

	reorderedName := reorderName(firstName, lastName, middleName, countryCode)
	fmt.Println(reorderedName)
}

package cmd

import (
	"fmt"
	"strings"
	"os"
)

func ReorderName(firstName, lastName, middleName, countryCode string) string {
	switch countryCode {
	case "VN", "CN", "JP", "KR", "IN", "TH", "ID", "PH", "MY", "SG":
		if middleName != "" {
			return fmt.Sprintf("%s %s %s", lastName, middleName, firstName)
		}
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

func Execute() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments.")
		fmt.Println("Usage: go run main.go <first_name> <last_name> [middle_name] <country_code>.")
		return
	}

	firstName := os.Args[1]
	lastName := os.Args[2]
	countryCode := strings.ToUpper(os.Args[len(os.Args)-1])
	middleName := strings.Join(os.Args[3:len(os.Args)-1], " ")

	reorderedName := ReorderName(firstName, lastName, middleName, countryCode)
	fmt.Println(reorderedName)
}
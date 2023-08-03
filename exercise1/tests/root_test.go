package tests

import (
	"testing"
	
	"github.com/diegovanne/go23/exercise1/cmd"
)

func TestReorderName(t *testing.T) {
	tests := []struct {
		firstName   string
		lastName    string
		middleName  string
		countryCode string
		expected    string
	}{
		// Test cases for VN, CN, JP, KR, IN, TH, ID, PH, MY, SG
		{"John", "Doe", "Michael", "VN", "Doe Michael John"},
		{"John", "Doe", "", "CN", "Doe John"},
		{"John", "Doe", "Michael", "JP", "Doe Michael John"},

		// Test cases for US, CA, GB, FR, DE, IT, ES, AU, NZ, BR
		{"John", "Doe", "Michael", "US", "John Doe Michael"},
		{"John", "Doe", "", "CA", "John Doe"},

		// Test case for other country codes
		{"John", "Doe", "Michael", "XX", "John Michael Doe"},
	}

	for _, test := range tests {
		actual := cmd.ReorderName(test.firstName, test.lastName, test.middleName, test.countryCode)
		if actual != test.expected {
			t.Errorf("reorderName(%s, %s, %s, %s) = %s; expected %s", test.firstName, test.lastName, test.middleName, test.countryCode, actual, test.expected)
		}
	}
}
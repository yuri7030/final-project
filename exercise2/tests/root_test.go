package tests

import (
	"testing"
	"math"
	
	"github.com/diegovanne/go23/exercise2/cmd"
)

func TestParseIntArray(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected []int
		wantErr  bool
	}{
		// Test cases with valid inputs
		{"case1", []string{"1", "2", "3"}, []int{1, 2, 3}, false},
		{"case2", []string{"-5", "10", "20"}, []int{-5, 10, 20}, false},
		{"case3", []string{"100", "200", "300"}, []int{100, 200, 300}, false},

		// Test cases with invalid inputs
		{"case4", []string{"1", "2", "a"}, nil, true},      // "a" is not a valid integer
		{"case5", []string{"1", "2", "9223372036854775808"}, nil, true}, // 9223372036854775808 is out of int range on most systems
		{"case6", []string{}, nil, false},                 // Empty input should return an empty array
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.ParseIntArray(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseIntArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the result array elements match the expected array elements
			for i := range tt.expected {
				if got[i] != tt.expected[i] {
					t.Errorf("parseIntArray() got = %v, expected %v", got, tt.expected)
					return
				}
			}
		})
	}
}

func TestParseFloatArray(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected []float64
		wantErr  bool
	}{
		// Test cases with valid inputs
		{"case1", []string{"1.5", "2.75", "3.99"}, []float64{1.5, 2.75, 3.99}, false},
		{"case2", []string{"-0.5", "10.25", "-20.75"}, []float64{-0.5, 10.25, -20.75}, false},
		{"case3", []string{"100.1", "200.2", "300.3"}, []float64{100.1, 200.2, 300.3}, false},

		// Test cases with invalid inputs
		{"case4", []string{"1.5", "2.75", "abc"}, nil, true}, // "abc" is not a valid float
		{"case5", []string{"1.5", "2.75", "1.7976931348623159e+308"}, nil, true}, // float64 limit exceeded
		{"case6", []string{}, nil, false}, // Empty input should return an empty array
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.ParseFloatArray(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFloatArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if the result array elements match the expected array elements
			for i := range tt.expected {
				if !floatEquals(got[i], tt.expected[i]) {
					t.Errorf("parseFloatArray() got = %v, expected %v", got, tt.expected)
					return
				}
			}
		})
	}
}

func TestParseMixedArray(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expected      []interface{}
		expectedError bool
	}{
		// Test cases with valid inputs
		{"case1", []string{"1", "2.75", "apple"}, []interface{}{1, 2.75, "apple"}, false},
		{"case2", []string{"-5", "10", "20.5"}, []interface{}{-5, 10, 20.5}, false},
		{"case3", []string{"100.1", "200", "300"}, []interface{}{100.1, 200, 300}, false},

		// Test cases with invalid inputs
		{"case4", []string{}, nil, false}, // Empty input should return an empty array
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.ParseMixedArray(tt.args)
			if (err != nil) != tt.expectedError {
				t.Errorf("parseMixedArray() error = %v, expectedError %v", err, tt.expectedError)
				return
			}

			if !compareMixedArray(got, tt.expected) {
				t.Errorf("parseMixedArray() got = %v, expected %v", got, tt.expected)
				return
			}
		})
	}
}

// floatEquals checks if two floats are equal within a tolerance (epsilon)
func floatEquals(a, b float64) bool {
	epsilon := 1e-9 // Adjust this tolerance as per your needs
	return math.Abs(a-b) < epsilon
}

// compareMixedArray compares two mixed arrays for equality
func compareMixedArray(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
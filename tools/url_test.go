package tools

import "testing"

// AI
// TestIsValidCypherhunterURL tests the IsValidCypherhunterURL function
func TestIsValidCypherhunterURL(t *testing.T) {
	// Define some test cases with inputs and expected outputs
	testCases := []struct {
		url      string
		expected bool
	}{
		{"https://www.cypherhunter.com/en/p/ethereum/", true},
		{"https://www.cypherhunter.com/en/p/bitcoin", true},
		{"https://www.cypherhunter.com/en/p/ethereum/defi", false},
		{"https://www.cypherhunter.com/ru/p/ethereum", true},
		{"https://www.google.com", false},
	}
	// Loop over the test cases
	for _, tc := range testCases {
		// Call the function with the input and get the output
		actual := IsValidCypherhunterURL(tc.url)
		// Compare the output with the expected output
		if actual != tc.expected {
			// Report an error if they don't match
			t.Errorf("IsValidCypherhunterURL(%q) = %v; want %v", tc.url, actual, tc.expected)
		}
	}
}

// TestIsValidURL tests the IsValidURL function
func TestIsValidURL(t *testing.T) {
	// Define some test cases with inputs and expected outputs
	testCases := []struct {
		url      string
		expected bool
	}{
		{"https://www.google.com", true},
		{"http://", false},
		{"/foo/bar", false},
		{"invalid", false},
		{"https://www.google.com:badport", false},
	}
	// Loop over the test cases
	for _, tc := range testCases {
		// Call the function with the input and get the output
		actual := IsValidURL(tc.url)
		// Compare the output with the expected output
		if actual != tc.expected {
			// Report an error if they don't match
			t.Errorf("IsValidURL(%q) = %v; want %v", tc.url, actual, tc.expected)
		}
	}
}

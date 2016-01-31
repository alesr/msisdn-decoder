package msisdn

import "testing"

// CASES
var sanitizeCases = []struct {
	input    string
	expected error
}{
	{"a", ErrSanitizeError},
	{"D3", ErrSanitizeError},
	{"111111111111", nil},
	{"+111111111111", nil},
	{"00111111111111", nil},
	{"_9283777dsd", ErrSanitizeError},
	{"223", ErrSanitizeError},
	{"34344456677777777755544", ErrSanitizeError},
	{"4", ErrSanitizeError},
	{"343f566723r4", ErrSanitizeError},
	{"6544455454", nil},
	{"++54445555", ErrSanitizeError},
	{"00999_99", ErrSanitizeError},
	{"+9999999999", nil},
	{"009999999999", nil},
}

var countryCodeCases = []struct {
	input string
	//expectedInfo []countryData
	expectedError error
}{
	{"35196234887", nil},
}

// TESTS
func TestSanitize(t *testing.T) {
	n := new(Msisdn)
	for _, test := range sanitizeCases {
		observed := n.sanitize(test.input)
		if observed != test.expected {
			t.Errorf("For input: %s\nExpected: %s\nObserved: %s",
				test.input, test.expected, observed)
		}
	}
}

func TestCountryCode(t *testing.T) {
	n := new(Msisdn)
	LoadJSON("../data/country-code.json", n)
	for _, test := range countryCodeCases {
		n.input = test.input
		_, err := n.countryCode()
		if err != nil && err != test.expectedError {
			t.Errorf("should get nil here")
		}
	}
}

package msisdn

import (
	"fmt"
	"testing"
)

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
	input                string
	name, code, dialCode []string
	expectedError        error
}{
	{"35196234887", []string{"Portugal"}, []string{"PT"}, []string{"351"}, nil},
	{"09196234887", []string{}, []string{}, []string{}, ErrCodeCountryError},
	{"55196234887", []string{"Brazil"}, []string{"BR"}, []string{"55"}, nil},
	{"3862121212", []string{"Slovenia"}, []string{"SI"}, []string{"386"}, nil},
	{"44655446212",
		[]string{
			"Guernsey",
			"Isle of Man",
			"Jersey",
			"United Kingdom",
		},
		[]string{
			"GG",
			"IM",
			"JE",
			"GB",
		},
		[]string{
			"44",
			"44",
			"44",
			"44",
		},
		nil,
	},
	{"11554545455",
		[]string{
			"Canada",
			"United States",
		},
		[]string{
			"CA",
			"US",
		},
		[]string{
			"1",
			"1",
		},
		nil,
	},
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

	errorMsg := fmt.Sprint("For input: %s, expected %s. Got %s")

	n := new(Msisdn)
	LoadJSON("../data/country-code.json", n)

	for _, test := range countryCodeCases {
		n.input = test.input
		observed, err := n.countryCode()

		if err != nil && err != test.expectedError {
			t.Errorf("should get nil here")
		}

		for i, obs := range observed {
			if obs.Name != test.name[i] {
				t.Errorf(errorMsg, n.input, test.name[i], obs.Name)
			}
			if obs.Code != test.code[i] {
				t.Errorf(errorMsg, n.input, test.code[i], obs.Code)
			}
			if obs.DialCode != test.dialCode[i] {
				t.Errorf(errorMsg, n.input, test.dialCode[i], obs.DialCode)
			}
		}
	}
}

package msisdn

import "testing"

var errSanitizeError string = "MSISDN must have between 8 and 15 digits"
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

func TestSanitize(t *testing.T) {
	for _, test := range sanitizeCases {
		n := new(Msisdn)
		observed := n.sanitize(test.input)
		if observed != test.expected {
			t.Errorf("For input: %s\nExpected: %s\nObserved: %s",
				test.input, test.expected, observed)
		}
	}
}

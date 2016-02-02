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

var decoderCases = []struct {
	input, expectedCC, expectedNDC, expectedMNO, expectedSN string
}{
	{"386016400500", "SI", "1", "T-2", "6400500"},
	{"386024000411", "SI", "2", "SI.MOBIL d.d.", "4000411"},
	{"386037111500", "SI", "3", "Telekom Slovenije", "7111500"},
	{"386047000375", "SI", "4", "Telemach", "7000375"},
	{"386056800497", "SI", "5", "SI.MOBIL d.d.", "6800497"},
	{"386076400500", "SI", "7", "T-2", "6400500"},
	{"351962348810", "PT", "", "", ""},
	{"551962348810", "BR", "", "", ""},
	{"672962348810", "AQ", "", "", ""},
}

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

	coutryDataFile = "../data/country-code.json"
	ndcDataFile = "../data/slovenia-ndc.json"
	mnoDataFile = "../data/slovenia-mno.json"

	// possible formatting directive in Sprint call go vet will call
	//but keep calm and test
	errorMsg := fmt.Sprint("For input: %s, expected %s. Got %s")

	n := new(Msisdn)
	LoadData(n)

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

// Partial test. Does not test CC w/ many areas and errors
func TestDecoder(t *testing.T) {
	errMsgFmt := fmt.Sprint("in: %s. expected: %s. got: %s")
	n := new(Msisdn)

	LoadData(n)
	for _, test := range decoderCases {
		r := new(Response)
		err := n.Decode(test.input, r)
		if err != nil {
			t.Error(err)
		}

		switch {
		case r.CC[0].Code != test.expectedCC:
			t.Errorf(errMsgFmt, test.input, test.expectedCC, r.CC[0].Code)
		case r.NDC.Code != test.expectedNDC:
			t.Errorf(errMsgFmt, test.input, test.expectedNDC, r.NDC.Code)
		case r.MNO.Operator != test.expectedMNO:
			t.Errorf(errMsgFmt, test.input, test.expectedMNO, r.MNO.Operator)
		case r.SN != test.expectedSN:
			t.Errorf(errMsgFmt, test.input, test.expectedSN, r.SN)
		}
	}
}

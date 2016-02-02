package msisdn

import (
	"errors"
	"regexp"
	"strings"
)

// good practice to list all package errors here
var (
	ErrSanitizeError    = errors.New("only of digits and optional prefixes (+, 00), 8-15 characters")
	ErrCodeCountryError = errors.New("sorry, didn't find any code country for this msisdn")
	ErrNotSInumberError = errors.New("the number provided is not a valid slovenian msisdn: wrong length")
	ErrUnknownNDCError  = errors.New("the number provided is not a valid slovenian msisdn: unknown ndc")
	ErrUnknownMNOError  = errors.New("the number provided is not a valid slovenian msisdn: unknown mno")
)

// Msisdn is kinda a Oracle that knows everything
// it knows about the questions the user do
// the data to look for the answer and
// all the methods to convert one thing into other
type Msisdn struct {
	input       string
	CountryData []cc
	NdcData     []ndc
	MnoData     []mno
}

type cc struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	DialCode string `json:"dial_code"`
}

type ndc struct {
	Code     string   `json:"code"`
	Locality []string `json:"locality"`
}

type mno struct {
	Operator string   `json:"operator"`
	Code     []string `json:"code"`
}

// Decode is our guy. Our contact with the client.
// he's responsible to get the question, call some tough guys to work on it
// and put the answer on the paper.
func (n *Msisdn) Decode(s string, reply *Response) error {

	// let's take the user input to quarantine
	if err := n.sanitize(s); err != nil {
		return err
	}

	// get CC
	cc, err := n.countryCode()
	if err != nil {
		return err
	}
	reply.CC = cc

	// Due to our restriction on data. We will only consider NDC, MNO and SN for Slovenia.
	if cc[0].Name == "Slovenia" {

		// Before start look for a NDC, MNO and SN.
		// Let's check if the input is a valid Slovenian number
		// or just some weird number starting with 386.
		// For that, we need to get rid of the CC and the potential zero following it.
		// After that, we see how many digits we have are left.
		// We need to count eight digits (NDC + (MNO + SN)) - initial 0.
		// Thanks Wikipedia. I hope you're right.
		_, err := n.isValidSInumber(cc[0].DialCode)
		if err != nil {
			return err
		}

		// get NDC
		ndcCode, ndcLocality, err := n.nationalDestCode()
		if err != nil {
			return err
		}
		reply.NDC.Code = ndcCode
		reply.NDC.Locality = ndcLocality

		// keep removing elements from the msisdn
		n.input = strings.TrimPrefix(n.input, ndcCode)

		// MSISDN = CC + NDC + SN
		// SN = MSISDN - CC - NDC
		reply.SN = n.input

		// get MNO
		mnoOp, mnoCode, err := n.mobileNetworkOp()
		reply.MNO.Operator = mnoOp
		reply.MNO.Code = mnoCode

		return nil
	}
	return nil
}

// this guy is cleaning obsessed. believe me.
func (n *Msisdn) sanitize(s string) error {
	// if i knew regular expression well enough
	// i would solve the whole story with our guy down there
	// but at least we're working on pointers. so, no copies!
	sPtr := &s
	*sPtr = strings.Replace(*sPtr, " ", "", -1)
	*sPtr = strings.Replace(*sPtr, "\t", "", -1)
	*sPtr = strings.TrimPrefix(*sPtr, "+")
	*sPtr = strings.TrimPrefix(*sPtr, "00")

	// only digits between 8(so far) and 15
	r, err := regexp.Compile("^[0-9]{8,15}$")
	if err != nil {
		return err
	}

	// after all the cleaning and preparation we can say.
	// that's a valid input
	if !r.MatchString(*sPtr) {
		return ErrSanitizeError
	}

	// we accept you, one of us
	n.input = *sPtr
	return nil
}

func (n *Msisdn) countryCode() ([]cc, error) {

	// hold all matchs for that msisdn
	countries := []cc{}

	// for each country in the whole world
	// if dial code is equal to the slice with same length (starting at index 0)
	// of the input data. then, we have a fellow cc.
	for _, c := range n.CountryData {
		if c.DialCode == n.input[:len(c.DialCode)] {
			match := cc{c.Name, c.Code, c.DialCode}
			countries = append(countries, match)
		}
	}

	// if empty slice, there's no match for this msisdn
	if len(countries) == 0 {
		return nil, ErrCodeCountryError
	}
	return countries, nil
}

func (n *Msisdn) nationalDestCode() (string, []string, error) {

	var ndcCode []string
	var ndcLocality [][]string

	// for each zone in JSON we compare the zone value to
	// the slice[:len(zone value)] on input
	for _, zone := range n.NdcData {
		if zone.Code == n.input[:len(zone.Code)] {
			ndcCode = append(ndcCode, zone.Code)
			ndcLocality = append(ndcLocality, zone.Locality)
		}
	}

	// no match found. so, this is not a valid SI number
	if len(ndcCode) == 0 {
		return "", nil, ErrUnknownNDCError
	}
	return ndcCode[0], ndcLocality[0], nil
}

func (n *Msisdn) mobileNetworkOp() (string, []string, error) {

	var mnoOp string
	var mnoCode []string

	// more or less the same thing above but nested loops
	for _, data := range n.MnoData {
		for _, code := range data.Code {
			if code == n.input[:len(code)] {
				mnoOp = data.Operator
				mnoCode = append(mnoCode, code)
			}
		}
	}
	// if no MNO were found...
	if len(mnoOp) == 0 {
		return "", nil, ErrUnknownMNOError
	}
	return mnoOp, mnoCode, nil
}

func (n *Msisdn) isValidSInumber(cc string) (bool, error) {
	n.input = strings.TrimPrefix(n.input, cc)

	// remove dispensable digit zero before area code
	// see! that's only* for slovenia! how do you guys do that?
	// i've read something about this zero be optional or not necessary when
	// you make phone call to the same area...
	n.input = strings.TrimPrefix(n.input, "0")

	// I know that the maximum length for any MSISDN is 15
	// but what is the minimum SN in slovenia? and other countries?
	// I assume 8 for slovenia with help of wikipedia
	// but wikipedia SHOULDN'T be used as documentation for writing software!
	if len(n.input) != 8 {
		return false, ErrNotSInumberError
	}
	return true, nil
}

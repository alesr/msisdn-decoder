package msisdn

import (
	"errors"
	"regexp"
	"strings"
)

// really good practice to list all package errors here
// ErrSanitizeError - when user type something nonsense that's what we say
var (
	ErrSanitizeError    = errors.New("only of digits and optional prefixes (+, 00), 8-15 characters")
	ErrCodeCountryError = errors.New("sorry, didn't find any code country for this msisdn")
)

// Msisdn is kinda a Oracle that knows everything
// it knows about the questions the user do
// the data to look for the answer and
// all the methods to convert one thing into other
type Msisdn struct {
	input string
	data  []Data
}

// Decode is our guy. Our contact with the client.
// he's responsible to get the question, call some tough guys to work on it
// and put the answer on paper.
func (n *Msisdn) Decode(s string, reply *Response) error {

	// let's take the user input to quarantine
	if err := n.sanitize(s); err != nil {
		return err
	}

	cc, err := n.countryCode()
	if err != nil {
		return (err)
	}

	reply.CC = cc

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

func (n *Msisdn) countryCode() ([]string, error) {
	cc := []string{}

	// for each country in the whole world
	// if dial code is equal to the slice with same length
	// of the input data. then, we have a fellow cc.
	for _, country := range n.data {
		if country.DialCode == n.input[:len(country.DialCode)] {
			cc = append(cc, country.Code)
		}
	}

	// if empty slice, there's no match for this msisdn
	if len(cc) == 0 {
		return nil, ErrCodeCountryError
	}

	return cc, nil
}

package msisdn

import (
	"errors"
	"regexp"
	"strings"
)

// really good practice to list all package erros here
// ErrSanitizeError - when user type something nonsense that's what we say
var ErrSanitizeError = errors.New("only of digits and optional prefixes (+, 00), 8-15 characters")

// Msisdn is kinda a Oracle that knows everything
// it knows about the questions the user do
// the data to look for the answer and
// all the methods to convert one thing into other
type Msisdn struct {
	input string
	data  []Data
}

// Decode is our guy. Our contact with the client side.
// he's responsible to get the question, call some tough guys to work on it
// and send put the answer on the paper.
func (n *Msisdn) Decode(s string, reply *Response) error {

	// let's take the user input to quarantine
	if err := n.sanitize(s); err != nil {
		return err
	}

	*reply = Response{"cc", "ndc", "mno"}
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

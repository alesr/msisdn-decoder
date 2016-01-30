package msisdn

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrSanitizeError error = errors.New("only of digits and optional prefixes (+, 00), 8-15 characters")

type Msisdn struct {
	input string
	data  []Data
}

func (n *Msisdn) Decode(s string, reply *Response) error {

	if err := n.sanitize(s); err != nil {
		return err
	}

	*reply = Response{"cc", "ndc", "mno"}
	return nil
}

func (n *Msisdn) sanitize(s string) error {
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

	if !r.MatchString(*sPtr) {
		return ErrSanitizeError
	}

	n.input = *sPtr
	return nil
}

// Response holds your data to be sent to client
type Response struct {
	CC, NDC, MNO string
}

// Implements Stringer interface
func (r *Response) String() string {
	return fmt.Sprintf("CC: %s  |  NDC: %s  |  MNO: %s",
		r.CC, r.NDC, r.MNO)
}

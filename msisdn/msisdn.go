package msisdn

import (
	"fmt"
	"regexp"
	"strings"
)

type Msisdn struct {
	input string
}

func (n *Msisdn) Decode(s string, reply *Response) error {

	if err := n.sanitize(&s); err != nil {
		return err
	}

	*reply = Response{"cc", "ndc", "mno"}
	return nil
}

func (n *Msisdn) sanitize(s *string) error {
	*s = strings.Replace(*s, " ", "", -1)
	*s = strings.Replace(*s, "\t", "", -1)
	*s = strings.TrimPrefix(*s, "+")
	*s = strings.TrimPrefix(*s, "00")

	// only digits between 8(so far) and 15
	r, err := regexp.Compile("^[0-9]{8,15}$")
	if err != nil {
		return err
	}

	if !r.MatchString(*s) {
		return fmt.Errorf("MSISDN must have between 8 and 15 digits")
	}

	n.input = *s
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

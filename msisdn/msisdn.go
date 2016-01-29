package msisdn

import (
	"fmt"
	"log"
	"strings"
)

type Msisdn struct {
	input string
}

func (n *Msisdn) Decode(s string, reply *Response) error {

	if err := n.sanitize(s); err != nil {
		log.Fatal(err)
	}

	*reply = Response{"cc", "ndc", "mno"}
	return nil
}

func (n *Msisdn) sanitize(s *string) error {
	*s = strings.Replace(*s, " ", "", -1)
	*s = strings.Replace(*s, "\t", "", -1)
	*s = strings.TrimPrefix(*s, "+")
	*s = strings.TrimPrefix(*s, "00")
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

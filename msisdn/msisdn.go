package msisdn

import "fmt"

type Msisdn struct {
	input string
}

func (n *Msisdn) Decode(s string, reply *Response) error {
	n.input = s
	*reply = Response{"cc", "ndc", "mno"}
	return nil
}

// func (n *Msisdn) sanitize(s *string)

// Response holds your data to be sent to client
type Response struct {
	CC, NDC, MNO string
}

// Implements Stringer interface
func (r *Response) String() string {
	return fmt.Sprintf("CC: %s  |  NDC: %s  |  MNO: %s",
		r.CC, r.NDC, r.MNO)
}

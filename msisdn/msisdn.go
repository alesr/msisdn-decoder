package msisdn

import "fmt"

type Msisdn string

func (n *Msisdn) Decode(s string, reply *Response) error {
	*reply = Response{"cc", "ndc", "mno"}
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

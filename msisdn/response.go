package msisdn

import "fmt"

// Response holds our data to be sent to client.
type Response struct {
	CC       []string // worth to say that some countries share the same CC (like USA and Canada)
	NDC, MNO string
}

// Implements Stringer interface o/
func (r *Response) String() string {
	return fmt.Sprintf("CC: %s  |  NDC: %s  |  MNO: %s",
		r.CC, r.NDC, r.MNO)
}

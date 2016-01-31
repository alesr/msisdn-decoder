package msisdn

import "fmt"

// Response holds our data to be sent to client.
type Response struct {
	CC       map[string][]string // worth to say that some countries share the same CC (eg.: USA and Canada)
	NDC, MNO string
}

// Implements Stringer interface o/
func (r *Response) String() string {
	return fmt.Sprintf("\nCC: %s  |  NDC: %s  |  MNO: %s\n",
		r.CC, r.NDC, r.MNO)
}

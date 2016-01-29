package msisdn

import "strings"

type Msisdn struct{}

type Response struct {
	CC, NDC, MNO string
}

func (n *Msisdn) Decode(s string, reply *Response) error {
	*reply = Response{strings.ToUpper(s), "ndc", "mno"}
	return nil
}

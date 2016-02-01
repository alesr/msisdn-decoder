package msisdn

import "fmt"

// Response holds our data to be sent to client.
type Response struct {
	CC  []country // worth to say that some countries share the same CC (eg.: USA and Canada)
	NDC ndc
	MNO mno
}

// just formats the answer in a readable way
// Implements Stringer interface o/
func (r *Response) String() string {

	// we can have multiples results for the same dial code
	// in such case we add this info to countryInfo
	countryInfo := []string{}
	for _, cc := range r.CC {
		data := fmt.Sprintf("\n    name: %s\n    code: %s\n    dial code: %s\n", cc.Name, cc.Code, cc.DialCode)
		countryInfo = append(countryInfo, data)
	}

	return fmt.Sprintf("\ncountry info: \n%s\n\nNDC: %s\nMNO: %s\n",
		countryInfo, r.NDC, r.MNO)
}

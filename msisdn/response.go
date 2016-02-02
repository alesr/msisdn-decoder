package msisdn

import "fmt"

// Response holds our data to be sent to client.
type Response struct {
	CC  []cc // worth to say that some countries share the same CC (eg.: USA and Canada)
	NDC ndc
	MNO mno
	SN  string
}

// format and print the answer in a readable way
func (r *Response) PrintReply() {

	// we can have multiples results for the same dial code
	// in such case we add this info to countryInfo
	fmt.Println("\nDial code: ", r.CC[0].DialCode)
	for _, country := range r.CC {
		fmt.Printf("  %s - %s\n", country.Name, country.Code)
	}

	if len(r.NDC.Code) != 0 {
		fmt.Println("\nNDC: ", r.NDC.Code)
		for _, locality := range r.NDC.Locality {
			fmt.Println("  ", locality)
		}
	}

	if len(r.MNO.Code) != 0 {
		fmt.Println("\nMNO: ", r.MNO.Code[0])
		fmt.Println("  ", r.MNO.Operator)
	}

	if len(r.SN) != 0 {
		fmt.Println("\nSN: ", r.SN)
	}
	fmt.Println()
}

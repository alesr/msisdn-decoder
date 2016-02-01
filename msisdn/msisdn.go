package msisdn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// good practice to list all package errors here
// ErrSanitizeError - when user type something nonsense that's what we say
var (
	ErrSanitizeError    = errors.New("only of digits and optional prefixes (+, 00), 8-15 characters")
	ErrCodeCountryError = errors.New("sorry, didn't find any code country for this msisdn")
	ErrNotSInumberError = errors.New("the number provided is not a valid slovenian msisdn: wrong length")
	ErrUnknownNDCError  = errors.New("the number provided is not a valid slovenian msisdn: unknown ndc")
)

// Msisdn is kinda a Oracle that knows everything
// it knows about the questions the user do
// the data to look for the answer and
// all the methods to convert one thing into other
type Msisdn struct {
	input       string
	CountryData []country
	NdcData     []ndc
}

// this dear buddy hold the country data we get from JSON
type country struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	DialCode string `json:"dial_code"`
}

type ndc struct {
	Code     string   `json:"code"`
	Locality []string `json:"locality"`
}

// Decode is our guy. Our contact with the client.
// he's responsible to get the question, call some tough guys to work on it
// and put the answer on the paper.
func (n *Msisdn) Decode(s string, reply *Response) error {

	// let's take the user input to quarantine
	if err := n.sanitize(s); err != nil {
		return err
	}

	// get CC
	cc, err := n.countryCode()
	if err != nil {
		return err
	}

	// Due to our restriction on data. We will only consider NDC, MNO and SN for Slovenia.
	if cc[0].Name == "Slovenia" {

		_, err := isValidSInumber(&n.input, cc[0].DialCode)
		if err != nil {
			return err
		}

		ndcCode, ndcLocality, err := n.nationalDestCode()
		if err != nil {
			return err
		}

		fmt.Println(ndcCode, ndcLocality)

		reply.NDC.Code = ndcCode
		reply.NDC.Locality = ndcLocality
		// n.mobileNetworkOp()
	}

	reply.CC = cc

	return nil
}

// this guy is cleaning obsessed. believe me.
func (n *Msisdn) sanitize(s string) error {

	// if i knew regular expression well enough
	// i would solve the whole story with our guy down there
	// but at least we're working on pointers. so, no copies!
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

	// after all the cleaning and preparation we can say.
	// that's a valid input
	if !r.MatchString(*sPtr) {
		return ErrSanitizeError
	}

	// we accept you, one of us
	n.input = *sPtr
	return nil
}

func (n *Msisdn) countryCode() ([]country, error) {

	// hold all matchs for that msisdn
	countries := []country{}

	// for each country in the whole world
	// if dial code is equal to the slice with same length
	// of the input data. then, we have a fellow cc.
	for _, c := range n.CountryData {
		if c.DialCode == n.input[:len(c.DialCode)] {
			match := country{c.Name, c.Code, c.DialCode}
			countries = append(countries, match)
		}
	}

	// if empty slice, there's no match for this msisdn
	if len(countries) == 0 {
		return nil, ErrCodeCountryError
	}
	return countries, nil
}

func (n *Msisdn) nationalDestCode() (string, []string, error) {

	var ndcCode []string
	var ndcLocality [][]string

	for _, zone := range n.NdcData {
		if zone.Code == n.input[:len(zone.Code)] {
			ndcCode = append(ndcCode, zone.Code)
			ndcLocality = append(ndcLocality, zone.Locality)
		}
	}

	if len(ndcCode) == 0 {
		return "", nil, ErrUnknownNDCError
	}
	return ndcCode[0], ndcLocality[0], nil
}

// Before start look for a NDC, MNO and SN.
// Let's check if the input is a valid Slovenian number
// or just some weird number starting with 386.
// For that, we need to get rid of the CC and the potential zero following it.
// After that, we see how many digits we have are left.
// We need to count eight digits (NDC + (MNO + SN)) - initial 0.
// Thanks Wikipedia. I hope you're right.
func isValidSInumber(input *string, cc string) (bool, error) {
	*input = strings.TrimPrefix(*input, cc)

	// remove dispensable digit zero before area code
	*input = strings.TrimPrefix(*input, "0")

	if len(*input) != 8 {
		return false, ErrNotSInumberError
	}
	return true, nil
}

// LoadData guess what. Loads data from JSON files into msisdn structs
func LoadData(n *Msisdn) {
	countryJSON, err := handleFile("data/country-code.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(countryJSON, &n.CountryData); err != nil {
		log.Fatal(err)
	}

	ndcJSON, err := handleFile("data/slovenia-ndc.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(ndcJSON, &n.NdcData); err != nil {
		log.Fatal(err)
	}
}

// handleFile checks if file exists, open and load it
func handleFile(filepath string) ([]byte, error) {

	// checks if the file is still there =]
	_, err := checkFile(filepath)
	if err != nil {
		return nil, err
	}

	// well, we open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// and we load the whole []byte
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, err
}

// Check if file exist in directory.
func checkFile(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

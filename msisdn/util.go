package msisdn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	coutryDataFile = "data/country-code.json"
	ndcDataFile    = "data/slovenia-ndc.json"
	mnoDataFile    = "data/slovenia-mno.json"
)

// LoadData guess what. Loads data from JSON files into msisdn structs
func LoadData(n *Msisdn) {
	countryJSON, err := handleFile(coutryDataFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(countryJSON, &n.CountryData); err != nil {
		log.Fatal(err)
	}

	ndcJSON, err := handleFile(ndcDataFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(ndcJSON, &n.NdcData); err != nil {
		log.Fatal(err)
	}

	mnoJSON, err := handleFile(mnoDataFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(mnoJSON, &n.MnoData); err != nil {
		fmt.Println("uhaeh")
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

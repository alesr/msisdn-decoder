package msisdn

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// LoadJSON data into msisdn struct so we can use it for search matchs later.
func LoadJSON(filepath string, n *Msisdn) {

	// checks if the file is still there =]
	_, err := checkFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// well, we open the file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// and we load the whole []byte
	content, err := ioutil.ReadAll(file)

	// to conclude we take all this content and unmarchal to a struct
	if err := json.Unmarshal(content, &n.countryData); err != nil {
		log.Fatal(err)
	}
}

// Check if file exist in directory.
func checkFile(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

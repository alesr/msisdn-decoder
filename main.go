package main

import (
	"fmt"
	"time"

	"github.com/alesr/msisdn-decoder/msisdn"
	"github.com/alesr/msisdn-decoder/rpc"
)

const countryCodeFilepath string = "data/country-code.json"

func main() {
	n := new(msisdn.Msisdn)
	go msisdn.LoadJSON(countryCodeFilepath, n)
	go rpc.Server(n)
	time.Sleep(1 * time.Second)
	fmt.Printf("\n*** lauching client... ***\n\n")
	time.Sleep(1 * time.Second)
	rpc.Client()
}

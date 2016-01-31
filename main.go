package main

import (
	"fmt"
	"time"

	"github.com/alesr/msisdn-decoder/msisdn"
	"github.com/alesr/msisdn-decoder/rpc"
)

// the pot under the rainbow for code country and other data
const countryCodeFilepath string = "data/country-code.json"

func main() {
	n := new(msisdn.Msisdn)

	// go load that data for me. gonna use it soon
	go msisdn.LoadJSON(countryCodeFilepath, n)

	// hey server! follow that taxi! kidding, just start
	go rpc.Server(n)

	// play: drum-roll-sound-effect.midi
	time.Sleep(1 * time.Second)
	fmt.Printf("\n*** lauching client... ***\n\n")
	time.Sleep(1 * time.Second)

	// starts the client
	rpc.Client()
}

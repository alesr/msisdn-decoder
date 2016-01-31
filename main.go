package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/alesr/msisdn-decoder/msisdn"
	"github.com/alesr/msisdn-decoder/rpc"
)

func main() {
	n := new(msisdn.Msisdn)

	b, err := msisdn.LoadFile("data/country-code.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, &n.CountryData); err != nil {
		log.Fatal(err)
	}

	// hey server! follow that taxi! kidding, just start
	go rpc.Server(n)

	// play: drum-roll-sound-effect.midi
	time.Sleep(1 * time.Second)
	fmt.Printf("\n*** lauching client... ***\n\n")
	time.Sleep(1 * time.Second)

	// starts the client
	rpc.Client()
}

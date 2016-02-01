package main

import (
	"fmt"
	"time"

	"github.com/alesr/msisdn-decoder/msisdn"
	"github.com/alesr/msisdn-decoder/rpc"
)

func main() {

	n := new(msisdn.Msisdn)

	// load the all the data we gonna need
	go msisdn.LoadData(n)

	// starts server
	go rpc.Server(n)

	// play: drum-roll-sound-effect.midi
	time.Sleep(1 * time.Second)
	fmt.Printf("\n*** launching client ***\n\n")
	time.Sleep(1 * time.Second)

	// starts client
	rpc.Client()
}

package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/alesr/msisdn-decoder/msisdn"
)

// Server tpc for our buddies
func Server() {

	// here's our address
	addr := "0.0.0.0:80"

	// honestly, never happened to understand what resolve address means
	rAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		log.Fatal(err)
	}
	// please, close the door before you leave.
	defer l.Close()

	fmt.Print("\n\n*** RPC server up and running... (ctrl-c to exit) ***\n")

	// make our buddy visible
	rpc.Register(new(msisdn.Msisdn))

	// do you accept everything you listen?
	rpc.Accept(l)
}

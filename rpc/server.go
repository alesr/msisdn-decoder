package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/alesr/msisdn-decoder/msisdn"
)

func Server() {

	addr := "0.0.0.0:80"

	rAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Print("\n\n*** RPC server up and running... (ctrl-c to exit) ***\n")

	rpc.Register(new(msisdn.Msisdn))
	rpc.Accept(l)
}

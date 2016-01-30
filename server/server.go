package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/alesr/msisdn-decoder/msisdn"
)

func main() {

	addr := "0.0.0.0:8080"

	rAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	list, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer list.Close()

	log.Println("RPC server up and running (ctrl-c to exit)")

	rpc.Register(new(msisdn.Msisdn))
	rpc.Accept(list)
}

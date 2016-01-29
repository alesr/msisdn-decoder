package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/alesr/msisdn-decoder3000/msisdn"
)

func main() {

	addr := "0.0.0.0:80"

	rAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	list, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer list.Close()

	n := new(msisdn.Msisdn)
	rpc.Register(n)
	rpc.Accept(list)

}

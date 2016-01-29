package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/alesr/msisdn-decoder3000/number"
)

func main() {

	*addr = "0.0.0.0:80"

	rAddr, err := net.ResolveTCPAddr("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	list, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		log.Fatal(err)
	}

	n := new(number.Number)
	rpc.Register(n)
	rpc.Accept(list)

}

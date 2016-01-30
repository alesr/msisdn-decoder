package main

import (
	"github.com/alesr/msisdn-decoder/rpc"
)

func main() {
	go rpc.Server()
	rpc.Client()
}

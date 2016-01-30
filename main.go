package main

import (
	"fmt"
	"time"

	"github.com/alesr/msisdn-decoder/rpc"
)

func main() {
	go rpc.Server()
	time.Sleep(1 * time.Second)
	fmt.Printf("\n*** lauching client... ***\n\n")
	time.Sleep(1 * time.Second)
	rpc.Client()
}

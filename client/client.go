package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/alesr/msisdn-decoder3000/msisdn"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)

	fmt.Print("msisdn: ")
	line, _, err := in.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	reply := new(msisdn.Response)
	if err = client.Call("Msisdn.Decode", string(line), &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(reply.String())

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
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

	var reply string
	if err = client.Call("Number.Translate", string(line), &reply); err != nil {
		log.Fatal(err)
	}

	log.Println(reply)

}

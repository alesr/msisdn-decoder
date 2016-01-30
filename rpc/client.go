package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/alesr/msisdn-decoder/msisdn"
)

func Client() {

	client, err := rpc.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}
	getRequest(client)
}

func getRequest(c *rpc.Client) {

	input, err := askInput(c)
	if err != nil {
		log.Fatal(err)
	}

	reply := new(msisdn.Response)
	if err = c.Call("Msisdn.Decode", input, &reply); err != nil {
		if err == msisdn.ErrSanitizeError {
			fmt.Println(err)
			getRequest(c)
		}
		fmt.Println(err)
		getRequest(c)
	}

	fmt.Printf("%s\n", reply.String())
	getRequest(c)
}

// askInput - interacts with the user asking a msisdn number
func askInput(c *rpc.Client) (string, error) {

	var input string
	fmt.Print("msisdn: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}

	switch input {
	case "exit":
		fmt.Println("\n*** exiting client ***")
		os.Exit(0)
	case "help":
		fmt.Println("enter a MSISDN composed only of digits and optional prefixes (+, 00), 8-15 characters")
		getRequest(c)
	}

	return input, nil
}

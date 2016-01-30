package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/alesr/msisdn-decoder/msisdn"
)

// Client - Here we deal with the connection to the server to ask him some serious stuff...
func Client() {

	// "hey i just met you. and this is crazy, but here's my number. so call me maybe?"
	client, err := rpc.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatal(err)
	}

	// let's check with user what he want to ask to server
	getRequest(client)
}

// it's all about good questions
func getRequest(c *rpc.Client) {

	// so, let's get some good ones from the client
	input, err := askInput(c)
	if err != nil {
		log.Fatal(err)
	}

	// hey server, please answer me back on this paper!
	reply := new(msisdn.Response)

	// Mr. Server i'm sending a paper &reply followed by a question (input)
	// so you can take it to Mrs. Msisdn.Decode give me some good news
	if err = c.Call("Msisdn.Decode", input, &reply); err != nil {

		// if the reply is an error saying that we ask the wrong question...
		if err == msisdn.ErrSanitizeError {
			fmt.Println(err)
		} else {
			// gee! not good!
			log.Fatal(err)
		}

		// we ask request again... (but what now?).
		getRequest(c)
	}

	// let's  announce the good news to user
	fmt.Printf("%s\n", reply.String())
	getRequest(c) // and ask again... shiiii..
}

// askInput - interacts with the user asking a msisdn number
func askInput(c *rpc.Client) (string, error) {

	// input will hold... well, the input
	var input string
	fmt.Print("msisdn: ")
	// this thing you typed, give to me
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}

	// too tired to explain
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

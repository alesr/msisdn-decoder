package rpc

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/alesr/msisdn-decoder/msisdn"
)

// Client - Here we deal with the connection client connection to the Server
// ask for input, outpur errors and good news
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

		// if the reply is an error saying that the input msisdn didn't
		// match with our rules. or that we didn't find any correspondent
		// result for that number. let's just print it out as a message and
		// allow the program to continue.
		// otherwise, it's an error coming from std library. in this case, exit.
		//
		// NOTE:
		// Here, the err is returned by decoder method VIA Server.
		// Because of that, the error is of type rpc.ServerError (or something like that)
		// Said that, we need to compare the VALUE of this error to the string
		// representation of our error.
		switch err.Error() {
		case msisdn.ErrSanitizeError.Error():
			fmt.Println(err)
		case msisdn.ErrCodeCountryError.Error():
			fmt.Println(err)
		case msisdn.ErrNotSInumberError.Error():
			fmt.Println(err)
		case msisdn.ErrUnknownNDCError.Error():
			fmt.Println(err)
		case msisdn.ErrUnknownMNOError.Error():
			fmt.Println(err)
		default:
			log.Fatal(err)
		}

		// come back to user and wait for another request
		getRequest(c)
	}

	// let's  announce the good news to user
	reply.String()
	getRequest(c) // and ask again...
}

// askInput - interacts with the user asking a msisdn number
func askInput(c *rpc.Client) (string, error) {

	var input string
	fmt.Print("msisdn: ")
	// this thing you typed, give to me
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}

	switch input {
	case "exit":
		fmt.Println("\n*** exit ***")
		os.Exit(0)
	case "help":
		fmt.Println("enter a MSISDN composed only by digits and optional prefixes (+, 00), 8-15 characters")
		getRequest(c)
	}
	return input, nil
}

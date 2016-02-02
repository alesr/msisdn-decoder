![msisdn-decoder](https://github.com/alesr/msisdn-decoder/blob/master/msisdn-decoder.gif)

### how to run it

If you have GO installed:
```
git clone github.com/alesr/msisdn-decoder
cd msisdn-decoder/
go run main.go
```

If you have docker:
```
docker pull alesr/msisdn-decoder
docker run -it --rm --name msisdn alesr/msisdn-decoder
```

### how it works

This is a simple MSISDN decoder using the Go net/rpc package to server on a tcp address and expose a single method API which can decode a MSISDN in some useful data.

After the successful connection to the server running at **127.0.0.1:80**, the client can then make requests by calling the only available method Decoder, pass a MSISDN string and receive a Response struct containing the result for that query.

The conversation happens over a simple I/O comunication interface, where the user provide an MSISDN number on stdin and receive your response or errors via stdout and stderr.

### input

For the input, the user should type a valid MSISDN number composed by digits, spaces, or optional the prefixes + and 00. The mininum length is 8 and the maximum 15.

By providing an invalid input, the user should receive helpful error message corresponding to that particular issue via stderr.

Also, the user can interact with the program by typing **help** to get a message hint and **exit** to close de connection and leave the application.

### output

After a valid submission, the following behavior is expected:

For any MSISDN number starting with a valid country dial code (CC) eg.: +55, 351, 0044. The program will map the corresponding country code in the ISO 3166-1-alpha-2 format, output the result and ask the user for a new MSISDN.

For Slovenian MSISDN an enhanced output is expected:
After verify that the country code is from Slovenia, the program will look to the next numbers to get the National Destination Code (NDC), which if valid, will point to the corresponding Slovenian region for that code. Although the NDC has not been asked in the task. This step is important to separate the CC from the Subscriber Number (SN) by following the definition:

**MSISDN = CC + NDC + SN**

After find the NDC number, it is already possible to determine the Subscriber Number (SN) which are the remaining digits after the NDC.

The next step is to map the Mobile Network Operator (MNO). The MNO is expected to be found as the first 2 or 3 digits at the Subscriber number.

For each part of the analysis there is a corresponding error message. So the user can understand the reason of a failure.

At the and of the proccess the result will be printed to the user console and the program will ask for a new MSISDN.

### considerations

[1]
The program scope and implementation of the code, follow my superficial analysis of the MSISDN definition found on the following links:

http://www.msisdn.org/

https://en.wikipedia.org/wiki/MSISDN

https://en.wikipedia.org/wiki/Telephone_numbers_in_Slovenia

http://www.akos-rs.si/numbering-space

Particularly to the scope of the program.
After find the CC, to get the SN from the MSISDN, is necessary to identify the NDC. And either to find the CC, the NDC or to map the MNO names to their codes it is only possible by using a database. That is the reason to the enhanced behavior be restrict only for Slovenians MSISDN so far.

[2]
At main.go, the server is called on its own goroutine and followed by a call to the client function for the
merely convenience of execute the program in the same terminal window. In any way the implementation of the server and client should be considered a single entity.

[3]
More about the implementation can be found in the code commentaries.

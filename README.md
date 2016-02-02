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

It's a simple MSISDN decoder using the Go net/rpc package to listen on a tcp address serving exposing a single method to the client.

After the successful connection to the server on **127.0.0.1:80**, the client can
then make requests to the server calling the only available method Decoder to
pass a MSISDN string and receive a Response struct containing the result for that query.

The conversation happens on a simple I/O comunication where the user input provide an MSISDN number and receive your response via stdout.

### input

For the input the user should type a valid MSISDN number composed by only digits, spaces, or optional the prefixes + and 00. The mininum length is 8 and the maximum 15.

By providing an invalid input, the user should receive helpfull error message corresponding to that particular issue via stder.

Also, the user can interact with the program by typing **help** to get a hint and **exit** to leave the application.

### output

After a valid submission, the following behavior is expected:

For any MSISDN number starting with a valid country dial code (CC) eg.: +55, 351, 0044. The program will match the corresponding country code in the ISO 3166-1-alpha-2 format, output the result and ask the user for a new MSISDN.

For Slovenian MSISDN an enhanced output is expected:
After verify that the country code is from Slovenia, the program will look to the next numbers to get the National Destination Code (NDC) which if valid will point to the corresponding Slovenian region. Although this code has not been asked in the task. This step is important to separate the the CC from the Subscriber Number which follows the definition:

**MSISDN = CC + NDC + SN**

After find the NDC number, it is already possible to find the Subscriber Number (SN) which are the remaining digits after the NDC.

The next step is to follow the previous principles used to find the CC and NDC to map the Mobile Network Operator (MNO). The MNO is expected to be found as the first 2 or 3 digits at the Subscriber number.

For each analysis if there is a corresponding error message. So that the user can understand the reason for a failure.

### considerations

[1]
The program scope and implementation of the code, follow my superficial analysis of the MSISDN definition found on the following links:

http://www.msisdn.org/
<<<<<<< HEAD
https://en.wikipedia.org/wiki/MSISDN
https://en.wikipedia.org/wiki/Telephone_numbers_in_Slovenia
=======

https://en.wikipedia.org/wiki/MSISDN

https://en.wikipedia.org/wiki/Telephone_numbers_in_Slovenia

>>>>>>> dev
http://www.akos-rs.si/numbering-space

Particularly to the scope of the program.
To get the the SN from the MSISDN, after find the CC is necessary to identify the NDC. And either to find the CC, the NDC and either to map the MNO names to their codes. It is only possible by using a database.


[2]
At main.go the server is called on its on goroutine being followed by the call
to the client for the pure convenience of execute the program in the same terminal window.

[3]
More about the implementation must be found in the code commentaries.   

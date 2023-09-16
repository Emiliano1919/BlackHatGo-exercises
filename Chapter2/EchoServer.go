package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echoSlower(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 512)
	for {
		size, err := conn.Read(buffer[0:])
		if err == io.EOF {
			log.Print("Client disconnected \n")
			break
		}
		if err != nil {
			log.Print("Unexpected error \n")
			break
		}
		log.Printf("Received %d bytes:%s\n", size, string(buffer))

		log.Println("Writing data")
		if _, err := conn.Write(buffer[0:size]); err != nil {
			log.Fatalln("Unable to write data")
			//This method exits after printing
		}
	}
}
func echoSlow(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n') //delimitates how much to read
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)
	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func echoFast(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Print("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Print("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}

//How to test this program:
/*
This program will listen on the port listed on net.Listen
in this case that port is 20080, but you won't be able to actually taste this program
at least in a macos environment using the method described on the book, since we don't have telnet.
Type:
echo -n "Can you hear me?" | nc localhost 20080
in another window of the terminal
this will send "Can you hear me" to the localhost in the port we opened by passing this with a pipeline to netat.

This part of the netcat manual might be useful:
TALKING TO SERVERS
     It is sometimes useful to talk to servers “by hand” rather than through a user interface.  It can aid in troubleshooting, when it might be necessary to verify what data a server is sending in
     response to commands issued by the client.  For example, to retrieve the home page of a web site:

           $ echo -n "GET / HTTP/1.0\r\n\r\n" | nc host.example.com 80

     Note that this also displays the headers sent by the web server.  They can be filtered, using a tool such as sed(1), if necessary.

     More complicated examples can be built up when the user knows the format of requests required by the server.  As another example, an email may be submitted to an SMTP server using:

           $ nc localhost 25 << EOF
           HELO host.example.com
           MAIL FROM: <user@host.example.com>
           RCPT TO: <user2@host.example.com>
           DATA
           Body of email.
           .
           QUIT
           EOF
Source: Netcat manual:*Hobbit* ⟨hobbit@avian.org⟩ and Eric Jackson ⟨ericj@monkey.org⟩.

Additional source: https://dev.to/hgsgtk/how-go-handles-network-and-system-calls-when-tcp-server-1nbdß
*/

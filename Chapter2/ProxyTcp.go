func handle(src net.Conn) {
	//Dial connects to a server
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	//dst is the implementation of a Conn interface
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
		//these elements are needed and conventional, just to detect errors
	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
func main() {
	//Listen creates a server
	listener, err := net.Listen("tcp:", "80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn) //This function will serve as our client receiver on our server
	}
}

/* Conclusion:
This server can be used with netcat or curl commands.
How can I connect two computers? For instance run this in one computer and send or receive information by connecting to this computer?
*/

/*Source:
Black Hat Go book
https://dev.to/hgsgtk/how-go-handles-network-and-system-calls-when-tcp-server-1nbd   : Blog post
https://pkg.go.dev/net    : Net Package documentation
https://medium.com/@adityapathak1189/tcp-server-in-golang-3c75766a8b08      : Blog post
*/
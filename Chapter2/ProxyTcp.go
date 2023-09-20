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
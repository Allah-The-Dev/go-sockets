package main

import (
	"log"
	"net"
)

func main() {
	//Empty host --> INADDR_ANY: listen on host's available unicast + anycast IP's
	protocol, port := "tcp", ":8080"
	//encapsulates socket(), bind(), listen() syscalls and return net.TCPListener
	ln, err := net.Listen(protocol, port)
	handleErr(err)
	log.Printf("Started listening (%s) on port %s", protocol, port)
	for {
		//ln.Accept() waits upto next request
		//and returns next connection which wraps a socket() for that individual connection
		//encapsulates accept() syscall and returns net.TCPConn
		conn, err := ln.Accept()
		handleErr(err)
		log.Printf("created connection for remote addr : %s with network : %s", conn.RemoteAddr().String(), conn.RemoteAddr().Network())
		go handleConnection(conn)
	}
}

func handleErr(err error) {
	log.Print(err)
}

func handleConnection(conn net.Conn) {
	//read data from connection and echo it back
	for {
		buf := make([]byte, 1024)
		//encapsulates read() syscall with no direct filedescriptor logic
		size, err := conn.Read(buf)
		if err != nil {
			handleErr(err)
			break
		}
		//just need to send back written portion of the buffer
		//encapsulates write() syscall
		data := buf[:size]
		conn.Write(append([]byte("TCP echo: "), data...))
	}
	//encapsulates close() syscall
	//close() conn on error, free connection file descriptor
	conn.Close()
}

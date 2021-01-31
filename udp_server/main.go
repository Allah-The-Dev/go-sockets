package main

import (
	"log"
	"net"
)

func main() {
	//Empty host --> INADDR_ANY: listen on host's available unicast + anycast IP's
	protocol, port := "udp", ":8080"
	//encapsulates socket(), bind(), returns net.UDPConn
	pc, err := net.ListenPacket(protocol, port)
	handleErr(err)
	log.Printf("Started listening (%s) on port %s", protocol, port)
	handlePacket(pc)
}

func handleErr(err error) {
	log.Print(err)
}

func handlePacket(pc net.PacketConn) {
	//read data from connection and echo it back
	for {
		buf := make([]byte, 1024)
		//encapsulates recvfrom() syscall with no direct file descriptor logic
		size, addr, err := pc.ReadFrom(buf)
		if err != nil {
			break
		}
		//just need to send back written portion of the buffer
		data := buf[:size]
		//encapsulates sendto() with no direct fd logic
		pc.WriteTo(append([]byte("UDP echo: "), data...), addr)
	}
	//close conn on error, free socket file descriptor
	pc.Close()
}

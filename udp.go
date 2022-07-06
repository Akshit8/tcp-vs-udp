package main

import (
	"log"
	"net"
)

func handleIncomingData(udp *net.UDPConn, n int, addr *net.UDPAddr, buffer []byte) {
	log.Printf("new data from %v, length: %v, data: %s\n", addr, n, buffer)

	n, err := udp.WriteToUDP(buffer, addr)
	if err != nil {
		log.Fatalf("write error: %v\n", err)
	}

	log.Printf("sent %v bytes\n", n)
}

func udp() {
	udp, err := net.ListenUDP("udp4", &net.UDPAddr{
		Port: 8081,
	})
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	defer udp.Close()

	for {
		buffer := make([]byte, 1024)
		n, addr, err := udp.ReadFromUDP(buffer)
		if err != nil {
			log.Fatalf("read error: %v\n", err)
		}

		go handleIncomingData(udp, n, addr, buffer)
	}
}

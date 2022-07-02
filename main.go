package main

import (
	"flag"
	"log"
)

var protocol = flag.String("p", "tcp", "protocol")

func main() {
	flag.Parse()
	log.Printf("protocol: %v\n", *protocol)

	switch *protocol {
	case "tcp":
		tcp()
	case "udp":
		panic("udp not implemented")
	default:
		log.Fatalf("unknown protocol: %v\n", *protocol)
	}
}

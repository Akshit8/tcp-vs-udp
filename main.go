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
		udp()
	default:
		log.Fatalf("unknown protocol: %v\n", *protocol)
	}
}

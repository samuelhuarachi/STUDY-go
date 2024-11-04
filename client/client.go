// Client
package main

import (
	"context"
	"log"
	"time"

	"github.com/plgd-dev/go-coap/v3/udp"
)

// See /examples/simpler/client/main.go
func main() {
	co, err := udp.Dial("[aaaa:7558:9102:293:212:4b00:2b4d:da78]:5683")

	// for tcp
	// co, err := tcp.Dial("localhost:5688")

	// for tcp-tls
	// co, err := tcp.Dial("localhost:5688", tcp.WithTLS(&tls.Config{...}))

	// for dtls
	// co, err := dtls.Dial("localhost:5688", &dtls.Config{...}))

	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	//resp, err := co.Get(ctx, "/ep/rf/txpw")
	resp, err := co.Get(ctx, "/meter/energ_ativa_total_direta")
	
	if err != nil {
		log.Fatalf("Cannot get response: %v", err)
		return
	}
	log.Printf("Response: %+v", resp)

}
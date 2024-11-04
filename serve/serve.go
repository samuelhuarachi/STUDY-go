// Server
package main

import (
	"bytes"
	"log"

	"github.com/plgd-dev/go-coap/v3"
	"github.com/plgd-dev/go-coap/v3/message"
	"github.com/plgd-dev/go-coap/v3/message/codes"
	"github.com/plgd-dev/go-coap/v3/mux"
)

// Middleware function, which will be called for each request.
func loggingMiddleware(next mux.Handler) mux.Handler {
	return mux.HandlerFunc(func(w mux.ResponseWriter, r *mux.Message) {
		log.Printf("ClientAddress %v, %v\n", w.Conn().RemoteAddr(), r.String())
		next.ServeCOAP(w, r)
	})
}

// See /examples/simple/server/main.go
func handleA(w mux.ResponseWriter, req *mux.Message) {
	err := w.SetResponse(codes.GET, message.TextPlain, bytes.NewReader([]byte("hello world")))
	if err != nil {
		log.Printf("cannot set response: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Handle("/a", mux.HandlerFunc(handleA))
	//r.Handle("/b", mux.HandlerFunc(handleB))

	// log.Fatal(coap.ListenAndServe("udp", ":5688", r))


	// for tcp
	log.Fatal(coap.ListenAndServe("tcp", ":5688",  r))

	// for tcp-tls
	// log.Fatal(coap.ListenAndServeTLS("tcp", ":5688", &tls.Config{...}, r))

	// for udp-dtls
	// log.Fatal(coap.ListenAndServeDTLS("udp", ":5688", &dtls.Config{...}, r))
}
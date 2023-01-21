package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//
//var tlsCertFile = "./certbundle.pem"
//var tlsKeyFile = "./server.key"

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working... at... "+r.Host+"\n")
}

func Server() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	fmt.Println(listenAddr)
	if len(listenAddr) == 0 {
		fmt.Println("Starting New Server at :8443")
		listenAddr = ":8443"
	}
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServeTLS(listenAddr, "certbundle.pem", "server.key", nil))
}

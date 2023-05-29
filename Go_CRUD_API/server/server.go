package server

import (
	"log"

	"fmt"
	"net/http"
)

func StartServer(port string) {
	fmt.Println("Starting application at port : ", port)

	handler, _ := New()

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

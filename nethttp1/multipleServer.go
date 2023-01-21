package main

//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func testFunc(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "testFunc at "+r.Host+"\n\n")
//}
//
//func main() {
//
//	http.HandleFunc("/test", testFunc)
//	go func() {
//		fmt.Println("In go routine")
//		log.Fatal(http.ListenAndServe(":9000", nil))
//	}()
//
//	//log.Fatal(http.ListenAndServe(":9001", nil))
//}

package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func clientReq(w http.ResponseWriter) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: transport}
	url := "http://localhost:8085/users"
	method := "GET"

	payloadStr := fmt.Sprintf("test the req")
	payload := strings.NewReader(payloadStr)
	req, err := http.NewRequest(method, url, payload)

	fmt.Println("Request : ", req)
	req.Header.Add("Content-Type", "text/plain")
	res, err := client.Do(req)
	fmt.Println("Response : ", res)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(body))
}

func clientReq2(w http.ResponseWriter, url string) {
	res, err := http.Get(url)
	errorHandler(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	errorHandler(err)
	fmt.Fprintf(w, string(body))
}

func handFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit proxy........")

	//clientReq(w)
	clientReq2(w, "http://localhost:8085/users")
}

func main() {
	fmt.Println("Starting proxy server....")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handFunc)
	http.ListenAndServe(":8088", r)
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

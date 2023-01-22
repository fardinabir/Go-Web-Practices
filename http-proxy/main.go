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

func clientReq(w http.ResponseWriter, url string) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: transport}
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
	urlParam := r.URL.Query().Get("url")
	decodedPath := decryptUrl(urlParam)
	url := "http://localhost:8085" + decodedPath // decodedPath -> "/users/3"

	if IsURL(url) == false {
		fmt.Fprintf(w, "URL is not valid")
		return
	}
	//clientReq(w, url)
	clientReq2(w, url)
}

// demo get req at http://localhost:8088/?url=63SSTsNXVSjHVmApEfMkLHHVDkmZ2Nujor1ZbpWR3kuX5g8fCE5n31hYVmABgL1KxjEo1Ad8Ctq7V5JxyCeiUSR5x49zx9ZxPAFUSAjJSCLfnh1pyETcGoXtFqgQqwt8GJ5m244xsYDEApXUVJDHrwF3eF8qzZTnqKdxdZyfZYocsMX

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

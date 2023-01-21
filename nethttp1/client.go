package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func Client() {
	var tlsConfig *tls.Config

	url := "https://localhost:8443/test"
	method := "GET"

	switch os.Getenv("AUTH_TYPE") {
	case "mtls":
		fmt.Println("AUTH_TYPE set as mtls")
		// Load public/private key pair from a pair of files. The files must contain PEM encoded data.
		cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
		if err != nil {
			panic(err)
		}
		// Load CA cert
		caCert, err := ioutil.ReadFile("cacert.pem")
		if err != nil {
			panic(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		}

	case "tls":
		fmt.Println("AUTH_TYPE set as tls")
		// Load CA cert
		caCert, err := ioutil.ReadFile("cacert.pem")
		if err != nil {
			panic(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig = &tls.Config{
			RootCAs: caCertPool,
		}

	default:
		fmt.Println("Insecure communication selected, skipping server verification")
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: transport}

	for i := 1; i < 5; i++ {
		payloadStr := fmt.Sprintf("test the %v time", i)
		payload := strings.NewReader(payloadStr)
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "text/plain")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(body))
	}
}

package service

import (
	"fmt"
	"net/http"
)

func NativeApiSecurity(h http.Handler) http.Handler {
	fmt.Println("Middleware initilized....")
	fn := func(w http.ResponseWriter, r *http.Request) {
		//reqID := r.Header.Get("reqID")
		//authKey := r.Header.Get("auth_key")
		//authSecret := r.Header.Get("auth_secret")
		fmt.Println("Middleware func called....")
		fmt.Println(r.Header)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func NativeApiSecurity2(h http.Handler) http.Handler {
	fmt.Println("Middleware initilized2222222....")
	fn := func(w http.ResponseWriter, r *http.Request) {
		//reqID := r.Header.Get("reqID")
		//authKey := r.Header.Get("auth_key")
		//authSecret := r.Header.Get("auth_secret")
		fmt.Println("Middleware func called2222222222222222222222....")
		fmt.Println(r.Header)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

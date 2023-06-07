package service

import (
	"net/http"
)

func GetHeaderValue(r *http.Request, header string) string {
	return r.Header.Get(header)
}

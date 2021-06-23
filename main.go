package main

import (
	"crypto/tls"
	"net/http"
	"vdart/config"
	"vdart/handler"
)

func main() {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	http.ListenAndServe(":"+config.Port, handler.New())
}

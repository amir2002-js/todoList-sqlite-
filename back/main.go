package main

import (
	"back/Handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler.Handler)
}

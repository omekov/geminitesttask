package parsesite

import (
	"log"
	"net/http"
)

const (
	addr = ":8080"
)

// Run ...
func Run() {
	log.Fatal(http.ListenAndServe(addr, server()))
}

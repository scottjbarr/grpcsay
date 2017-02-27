package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/scottjbarr/grpcsay"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.FormValue("m")
	w.Write([]byte(grpcsay.Say(message)))
}

func main() {
	bind := os.Getenv("BIND")
	if len(bind) == 0 {
		fmt.Printf("Usage: BIND=:9000 bovine_http")
		os.Exit(1)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(bind, nil)
}

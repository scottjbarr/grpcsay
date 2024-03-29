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
		fmt.Printf("Usage: BIND=:9000 grpcsay_http\n")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	if err := http.ListenAndServe(bind, mux); err != nil {
		panic(err)
	}
}

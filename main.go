package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(r.RemoteAddr)) }).Methods(http.MethodGet)
	s := http.Server{Addr: "0.0.0.0:1000", Handler: r}

	fmt.Println(s.ListenAndServe())
}

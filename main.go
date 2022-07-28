package main

import (
	"fmt"
	"net/http"
	//"os"
	//"strconv"

	"github.com/gorilla/mux"
)

func main() {
	port := 3000
	/*if p := os.Getenv("PORT"); len(p) != 0 {
		port, _ = strconv.Atoi(p)
	}*/

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(r.RemoteAddr)) }).Methods(http.MethodGet)
	s := http.Server{Addr: fmt.Sprintf("0.0.0.0:%d", port), Handler: r}

	fmt.Println(s.ListenAndServe())
}

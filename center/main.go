package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		path := vars["path"]
		path = fmt.Sprintf("/data/%s", path)

		fileBytes, e := ioutil.ReadFile(path)
		if e != nil {
			perr, ok := e.(*fs.PathError)
			if !ok {
				return
			}
			if perr.Err.Error() == "is a directory" {
				files, err := ioutil.ReadDir(path)
				if err != nil {
					return
				}
				for _, f := range files {
					w.Write([]byte(fmt.Sprintf("<p><a href=\"%s\"> <span style=\"margin-right: 30px;\">%s</span> <span>size:%d</span> </a></p>", f.Name(), f.Name(), f.Size())))
				}

			} else {
				return
			}
		}

		w.Write(fileBytes)
	}
}

func upload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.ParseMultipartForm(33554432)

		file, fileHeader, err := r.FormFile("file")
		fmt.Println(file, fileHeader, err)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f, err := os.OpenFile(fmt.Sprintf("/data/%s", fileHeader.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		http.Redirect(w, r, r.Header["Referer"][0], http.StatusSeeOther)
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/upload", upload()).Methods(http.MethodPost)
	r.HandleFunc("/{path:[a-zA-Z0-9/.]*}", get()).Methods(http.MethodGet)
	s := http.Server{Addr: "0.0.0.0:80", Handler: r}
	s.ListenAndServe()
}

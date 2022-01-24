package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"study1/errhanding/filelistingserver/filelisting"
)

type appHandler func(rw http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic: %v", rec)
				http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(rw, r)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			if useErr, ok := err.(userError); ok {
				http.Error(rw, useErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(rw, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

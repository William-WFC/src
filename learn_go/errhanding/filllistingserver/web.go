package main

import (
	"learnGo/errhanding/filllistingserver/filllisting"
	"log"
	"net/http"
	"os"
)

type apphandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func errorWrapper(handler apphandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("error: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("error on handling: %s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			var code = http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}

func main() {
	http.HandleFunc("/", errorWrapper(filllisting.Handler))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

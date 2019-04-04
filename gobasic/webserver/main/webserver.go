package main

import (
	"github.com/dlstonedl/leet-code-go/gobasic/webserver/filehandle"
	"net/http"
	"os"
)

type errorHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errorWrapper(handler errorHandler) func(writer http.ResponseWriter,
	request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := handler(writer, request); err != nil {
			code := http.StatusOK
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
	http.HandleFunc("/list/", errorWrapper(filehandle.FileHandler))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

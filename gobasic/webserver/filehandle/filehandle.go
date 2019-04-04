package filehandle

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
)

func FileHandler(writer http.ResponseWriter, request *http.Request) error {
	fileName := request.URL.Path[len("/list/"):]
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	writer.Write(bytes)
	return nil
}

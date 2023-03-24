package filllisting

import (
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func Handler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("url must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		// panic(err)
		// fmt.Println(http.StatusInternalServerError) // 500
		// http.Error(writer, err.Error(), 200)
		return err
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}
	writer.Write(all)
	return nil
}

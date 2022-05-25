package utils

import (
	"bytes"
	"io"
	"net/http"
)

func ReceiveFileContent(r *http.Request) string {
	r.ParseMultipartForm(32 << 20)

	var buffer bytes.Buffer

	file, _, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	io.Copy(&buffer, file)

	return buffer.String()
}

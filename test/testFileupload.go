package main

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"bytes"
	"github.com/julienschmidt/httprouter"
	"log"
)

func uploadBigFile(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {

	mr, err := r.MultipartReader()
	if err != nil {
		fmt.Sprintln(err)
		fmt.Fprintln(w, err)

		return
	}

	values := make(map[string][]string, 0)
	maxValueBytes := int64(10 << 20)
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}

		name := part.FormName()
		if name == "" {
			continue
		}

		fileName := part.FileName()

		var b bytes.Buffer

		if fileName == "" {
			n, err := io.CopyN(&b, part, maxValueBytes)
			if err != nil && err != io.EOF {
				fmt.Sprintln(err)
				fmt.Fprintln(w, err)

				return
			}

			maxValueBytes -= n
			if maxValueBytes <= 0 {
				msg := "multipart message too large"
				fmt.Fprint(w, msg)
				return
			}

			values[name] = append(values[name], b.String())
		}

		dst, err := os.Create("/tmp/upload/" + fileName)
		defer dst.Close()

		for {
			buffer := make([]byte, 100000)
			cBytes, err := part.Read(buffer)
			if err == io.EOF {
				break
			}
			dst.Write(buffer[0:cBytes])
		}

	}
}

func main() {
	router:=httprouter.New()
	router.GET("/upload/:file",uploadBigFile)
	log.Fatal(http.ListenAndServe("8080",router))
}
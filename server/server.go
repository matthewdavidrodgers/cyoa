package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/matthewdavidrodgers/cyoa/adventure"
)

func writeError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "500 Internal Server Error")
}

func writeNotFound(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "404 Not Found")
}

func writeData(rw http.ResponseWriter, contentType string, data []byte) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", contentType)
	rw.Write(data)
}

func SetupServer(story adventure.Story) error {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handling request for %s %s\n", r.Method, r.URL.Path)
		nodeKey := r.URL.Path[1:]
		if r.Method == http.MethodGet && nodeKey == "" {
			page, err := adventure.RenderPage(story["intro"])
			if err != nil {
				fmt.Println(err)
				writeError(rw)
				return
			}
			writeData(rw, "text/html", page)
			return
		} else if _, ok := story[nodeKey]; ok && r.Method == http.MethodGet {
			node := story[nodeKey]
			page, err := adventure.RenderPage(node)
			if err != nil {
				fmt.Println(err)
				writeError(rw)
				return
			}
			writeData(rw, "text/html", page)
			return
		}
		writeNotFound(rw)
		return
	})

	fmt.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", nil)

	return nil
}

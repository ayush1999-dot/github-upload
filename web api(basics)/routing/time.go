package main

import (
	"io"
	"net/http"
	"time"
)

func Printtime(w http.ResponseWriter, r *http.Request) {
	timenow := time.Now()
	io.WriteString(w, timenow.String())

}
func main() {
	http.HandleFunc("/", Printtime)
	http.ListenAndServe(":8080", nil)
}

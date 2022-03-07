package main

import (
	"fmt"
	"net/http"
)

func main() {
	originalhandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalhandler))
	http.ListenAndServe(":8000", nil)

}
func handle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("execting manhandler.....")
	w.Write([]byte("OK"))

}
func middleware(originalhandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("execting request phase")
		//passing control back to original handler
		originalhandler.ServeHTTP(w, r)
		fmt.Println("execcuting response phase")

	})
}

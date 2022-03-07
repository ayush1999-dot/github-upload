package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string `json:"name"`
	Area int    `json:"area"`
}

func main() {
	//handlerr :=http.handlefunc(check)
	http.Handle("/dog", filter(servercookie(http.HandlerFunc(check))))
	http.ListenAndServe(":8080", nil)

}
func check(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in check fun")

	if r.Method == "POST" {
		var cityA city
		json.NewDecoder(r.Body).Decode(&cityA)
		fmt.Println("the city name is ", cityA.Name, "and the area is ", cityA.Area)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")

	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "not ok")
	}
}
func filter(servercookie http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "in filter")

		servercookie.ServeHTTP(w, r)

	})
}
func servercookie(check http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		fmt.Fprintln(w, "Currently in the set server time middleware")
		check.ServeHTTP(w, r)
		// Setting cookie to every API response
		cookie := http.Cookie{Name: "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)

	})
}

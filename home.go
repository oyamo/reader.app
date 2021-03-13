package main

import (
	"fmt"
	"net/http"
	auth "reader.app/auth"
)

func testConnection(w http.ResponseWriter, r *http.Request) {
	var authorization auth.Authorization
	auth.VerifyAuthentication(w, r, &authorization)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	_, _ = fmt.Fprintf(w, "OK")
}

func ApiHome(w http.ResponseWriter, r *http.Request) {
	var authorization auth.Authorization
	auth.VerifyAuthentication(w, r, &authorization)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	response := "{}"
	_, err := fmt.Fprint(w, response)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func HtmlHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "./templates/index.html")
}

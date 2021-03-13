package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"path/filepath"
	"reader.app/const"
)

func main() {
	router := mux.NewRouter()
	server := http.Server{
		Addr: _const.Address,
	}


	router.HandleFunc("/", HtmlHome).Methods("GET")
	router.HandleFunc("/test", testConnection).Methods("GET")
	router.HandleFunc("/api", ApiHome).Methods("GET")
	router.HandleFunc("/api/books", listBooks).Methods("GET")
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Printf("%s %s", "Server Running on", _const.Address)

	http.Handle("/", router)
	dir, err := filepath.Abs(filepath.Dir(".."))
	if err != nil {
		log.Fatal(err)
	}

	// Load .env file
	if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
		fmt.Print("No .env file found")
		panic(err)
	}
	println(filepath.Join(dir, ".env"))

	// Connect to the database
	ConnectToDB()

	// Start Serving
	err = server.ListenAndServe()

	// Listen for start-up errors
	if err != nil {
		fmt.Printf("%v", err)
	}
}

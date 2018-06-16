package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

func GetHealtEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "not implemented yet!")
}

func GetAllMailsEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "not implemented yet!")
}

func PostMailEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "not implemented yet!")
}

func main() {
	// Configure Router and Routes
	router := mux.NewRouter()
	router.HandleFunc("/health", GetHealtEndPoint).Methods("GET")
	router.HandleFunc("/mailstore", GetAllMailsEndPoint).Methods("GET")
	router.HandleFunc("/mailstore", PostMailEndPoint).Methods("POST")

	// Serve
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
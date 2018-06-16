package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mailStorageService/models"
	"encoding/json"
	"github.com/joho/godotenv"
	"strconv"
	"os"
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
	// load configuration
	err := godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	privateKey := os.Getenv("PRIVATE_KEY")
	publicKey := os.Getenv("PUBLIC_KEY")
	if err != nil {
		log.Fatal("Unable to read configuration file.")
	}


	// Configure Router and Routes
	router := mux.NewRouter()
	router.HandleFunc("/health", GetHealthEndPoint).Methods("GET")
	router.HandleFunc("/mailstore", GetQueryMailsEndPoint).Methods("GET")
	router.HandleFunc("/mailstore", PostMailEndPoint).Methods("POST")

	// Serve
	fmt.Printf("MailStorageService ist listening on port %d. \n", port)
	err = http.ListenAndServeTLS(":" + strconv.Itoa(port), publicKey, privateKey, router)
	if err != nil {
		log.Fatal(err)
	}
}
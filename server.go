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
	config2 "mailStorageService/config"
)

var config = config2.Config{}

func GetHealthEndPoint(writer http.ResponseWriter, _ *http.Request) {
	respondWithJson(writer, http.StatusOK, map[string]string{"status": "200", "message": "Service operational"})
}

func GetQueryMailsEndPoint(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "not implemented yet!")
}

func PostMailEndPoint(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var mail models.Mail

	err := json.NewDecoder(request.Body).Decode(&mail)
	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	err = mail.IsValid()
	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	// TODO: Add model storing
	respondWithJson(writer, http.StatusCreated, mail)
}

func main() {
	// load configuration
	err := godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	privateKey := os.Getenv("PRIVATE_KEY")
	publicKey := os.Getenv("PUBLIC_KEY")
func init() {
	// read configuration
	err := config.Read()
	if err != nil {
		log.Fatal("Unable to read configuration. Error: " + err.Error())
	}

}

func main() {
	// Configure Router and Routes
	router := mux.NewRouter()
	router.HandleFunc("/health", GetHealthEndPoint).Methods("GET")
	router.HandleFunc("/mailstore", GetQueryMailsEndPoint).Methods("GET")
	router.HandleFunc("/mailstore", PostMailEndPoint).Methods("POST")

	// Serve
	fmt.Printf("MailStorageService ist listening on port %d. \n", config.Port)
	err := http.ListenAndServeTLS(":" + strconv.Itoa(config.Port), config.PublicKey, config.PrivateKey, router)
	if err != nil {
		log.Fatal(err)
	}
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
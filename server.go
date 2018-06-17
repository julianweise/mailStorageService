package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mailStorageService/models"
	"encoding/json"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	dao2 "mailStorageService/dao"
	config2 "mailStorageService/config"
)

var dao = dao2.MailsDAO{}
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

	// decode request
	err := json.NewDecoder(request.Body).Decode(&mail)
	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// check completeness
	err = mail.IsValid()
	if err != nil {
		respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// insert into database
	mail.Id = bson.NewObjectId()
	err = dao.Insert(mail)
	if err != nil {
		respondWithJson(writer, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	respondWithJson(writer, http.StatusCreated, mail)
}

func init() {
	// read configuration
	err := config.Read()
	if err != nil {
		log.Fatal("Unable to read configuration. Error: " + err.Error())
	}

	// setup database connection
	dao.Server = config.Server
	dao.Database = config.Database
	err = dao.Connect()
	if err != nil {
		log.Fatal("Unable to establish a database connection. Error: " + err.Error())
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
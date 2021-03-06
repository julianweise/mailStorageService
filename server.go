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
	"strings"
	"time"
)

type MailListResponse struct {
	MailList	[]models.Mail	`json:"mail_list"`
}

func NewMailListResponse() MailListResponse {
	response := MailListResponse{}
	response.MailList = make([]models.Mail, 0)
	return response
}

func (this *MailListResponse) SetMailList(mails []models.Mail) {
	if mails == nil || len(mails) <= 0 {
		return
	}

	this.MailList = mails
}

var dao = dao2.MailsDAO{}
var config = config2.Config{}

func GetHealthEndPoint(writer http.ResponseWriter, _ *http.Request) {
	respondWithJson(writer, http.StatusOK, map[string]string{"status": "200", "message": "Service operational"})
}

func GetQueryMailsEndPoint(writer http.ResponseWriter, request *http.Request) {
	if len(request.URL.Query().Encode()) <= 0 {
		GetAllMailsEndPoint(writer, request)
		return
	}

	var err error
	responseLimit := 1024
	responseOffset := 0
	queryAttributes := bson.M{}

	ids := parseStringList(request.URL.Query().Get("id"))
	if len(ids) > 0 {
		queryAttributes["_id"] = bson.M{"$in": ids}
	}

	/*
	receivedBeforeTime := time.Now()
	receivedAfterTime := time.Time{}

	receivedBefore := request.URL.Query().Get("received_before")
	if len(receivedBefore) > 0 {
		receivedBeforeTime, err = parseTime(receivedBefore)

		if err != nil {
			respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": "unable to parse 'received_before': " + err.Error()})
			return
		}
	}

	receivedAfter := request.URL.Query().Get("received_after")
	if len(receivedAfter) > 0 {
		receivedAfterTime, err = parseTime(receivedAfter)

		if err != nil {
			respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": "unable to parse 'received_before': " + err.Error()})
			return
		}
	}

	log.Printf("received after '%s'   received before '%s'", receivedAfterTime, receivedBeforeTime)

	queryAttributes["received"] = bson.M{"received": bson.M{"$and":
		[]bson.M{
			bson.M{"received": bson.M{"$gt": receivedAfterTime}},
			bson.M{"received": bson.M{"$lt": receivedBeforeTime}},
	}}}
	*/

	senders := parseStringList(request.URL.Query().Get("mail_from"))
	if len(senders) > 0 {
		queryAttributes["mail_from"] = bson.M{"$in": senders}
	}

	receivers := parseStringList(request.URL.Query().Get("rcpt_to"))
	if len(receivers) > 0 {
		queryAttributes["rcpt_to"] = bson.M{"$in": receivers}
	}

	limitQueryParameter := request.URL.Query().Get("limitQueryParameter")
	if len(limitQueryParameter) > 0 {
		responseLimit, err = strconv.Atoi(limitQueryParameter)

		if err != nil {
			respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": "unable to parse 'limitQueryParameter': " + err.Error()})
			return
		}

		responseLimit = max(responseLimit, 0)
	}

	offset := request.URL.Query().Get("offset")
	if len(offset) > 0 {
		responseOffset, err = strconv.Atoi(offset)

		if err != nil {
			respondWithJson(writer, http.StatusBadRequest, map[string]string{"message": "unable to parse 'offset': " + err.Error()})
			return
		}

		responseOffset = max(responseOffset, 0)
	}

	log.Printf("query: '%s'", queryAttributes)
	mails, err := dao.Select(queryAttributes)
	if err != nil {
		respondWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "unable to fetch from database: " + err.Error()})
		return
	}

	response := NewMailListResponse()

	startIndex := limit(0, len(mails), responseOffset)
	endIndex := limit(0, len(mails), min(len(mails), responseLimit))

	if endIndex - startIndex > 0 {
		response.SetMailList(mails[startIndex:endIndex])
	}

	respondWithJson(writer, http.StatusOK, response)
}

func GetAllMailsEndPoint(writer http.ResponseWriter, request *http.Request) {
	mails, err := dao.SelectAll()
	if err != nil {
		respondWithJson(writer, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	response := MailListResponse{}
	response.SetMailList(mails)

	respondWithJson(writer, http.StatusOK, response)
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
	mail.Id = bson.NewObjectId().Hex()
	err = dao.Insert(mail)
	if err != nil {
		respondWithJson(writer, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	log.Printf("stored mail with id = '%s'", mail.Id)
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
	router.HandleFunc("/mailstore", GetAllMailsEndPoint).Methods("GET")
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

func parseTime(timeString string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeString)
}

func parseStringList(input string) []string {
	if len(input) <= 0 {
		return []string{}
	}

	return strings.Split(input, "|")
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func limit(lowerBound int, upperBound int, value int) int {
	return max(lowerBound, min(upperBound, value))
}
package main

import (
	"fmt"
	"io"
	"net/http"
	. "os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	file, err := OpenFile("access.log", O_RDWR|O_CREATE|O_APPEND, 0666)

	if err != nil {
		fmt.Println("Could not open file with error: " + err.Error())
	}

	log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting TODO API")

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheckHandler).Methods("GET")
	http.ListenAndServe(":8001", router)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	io.WriteString(w, `{"alive":true}`)
}
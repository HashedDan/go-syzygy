package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/bitly/go-simplejson"
)

func startApi() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HomeHandler).Methods("GET")

	go http.ListenAndServe(":8000", r)
	fmt.Println("started")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("key", "value")

	payload, err := json.MarshalJSON()

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}


package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-led-strip/services"
	"log"
	"net/http"
)

type LedStripController struct {
	ledStripService services.LedStripService
}

type StripPayload struct {
	Active bool `json:"active"`
}

func NewLedStripController(ledStripService *services.LedStripService) *LedStripController {
	l := LedStripController{ledStripService: *ledStripService}
	l.initialize()
	return &l
}

func (ledStripController *LedStripController) toggle(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	strip := StripPayload{Active: ledStripController.ledStripService.Toggle()}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&strip)
}

func (ledStripController *LedStripController) status(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	strip := StripPayload{Active: ledStripController.ledStripService.Status()}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&strip)
}

func (ledStripController *LedStripController) turnOff(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	strip := StripPayload{Active: ledStripController.ledStripService.Stop()}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&strip)
}

func (ledStripController *LedStripController) turnOn(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ledStripController.ledStripService.Start())
}

func (ledStripController *LedStripController) handle(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.Method == http.MethodPost {
		var payload StripPayload
		err := json.NewDecoder(request.Body).Decode(&payload)
		fmt.Println(payload)
		if err != nil {
			log.Fatalln("There was an error decoding the request body into the struct")
		}
		if payload.Active {
			ledStripController.ledStripService.Start()
		} else {
			ledStripController.ledStripService.Stop()
		}
	}
	strip := StripPayload{Active: ledStripController.ledStripService.Status()}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&strip)
}

func (ledStripController *LedStripController) initialize() {
	r := mux.NewRouter()
	r.HandleFunc("/ledstrip", ledStripController.handle).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/status", ledStripController.status).Methods(http.MethodGet)
	r.HandleFunc("/toggle", ledStripController.toggle).Methods(http.MethodGet)
	r.HandleFunc("/off", ledStripController.turnOff).Methods(http.MethodGet)
	r.HandleFunc("/on", ledStripController.turnOn).Methods(http.MethodGet)

	fmt.Println("Launching http and listening..")
	log.Fatal(http.ListenAndServe(":8081", r))
}

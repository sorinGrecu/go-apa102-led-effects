package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-led-strip/services"
	"golang.org/x/image/colornames"
	"log"
	"net/http"
)

type LedStripController struct {
	ledStripService services.LedStripService
}

func NewLedStripController(ledStripService *services.LedStripService) *LedStripController {
	l := LedStripController{ledStripService: *ledStripService}
	l.initialize()
	return &l
}

func (ledStripController *LedStripController) toggle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ledStripController.ledStripService.Toggle()
	json.NewEncoder(w).Encode("TOGGLED")
}

func (ledStripController *LedStripController) turnOff(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ledStripController.ledStripService.Stop()
	ledStripController.ledStripService.LedStrip.Fill(colornames.Black)
	json.NewEncoder(w).Encode("OFF")
}

func (ledStripController *LedStripController) turnOn(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ledStripController.ledStripService.Start()
	json.NewEncoder(w).Encode("ON")
}

func (ledStripController *LedStripController) initialize() {
	r := mux.NewRouter()
	r.HandleFunc("/toggle", ledStripController.toggle).Methods("GET")
	r.HandleFunc("/off", ledStripController.turnOff).Methods("GET")
	r.HandleFunc("/on", ledStripController.turnOn).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}

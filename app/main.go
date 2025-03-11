package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	mux "github.com/gorilla/mux"
)

func Start() {
	// Multiplexor

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet_handler).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet, http.MethodPatch, http.MethodDelete)

	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getTime(writer http.ResponseWriter, request *http.Request) {
	// Read query parameters
	query := request.URL.Query()
	tz := query.Get("tz")

	if tz == "" {
		tz = "UTC"
	}

	loc, err := time.LoadLocation(tz) // Cargar zona horaria
	if err != nil {
		http.Error(writer, "Error cargando la zona horaria", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(map[string]string{
		"current_time": time.Now().In(loc).Format("2006-01-02 15:04:05 MST"),
		"timezone":     loc.String(),
	})
}

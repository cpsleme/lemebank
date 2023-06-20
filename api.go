package main

import (
	"encoding/json"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func api(system *actor.ActorSystem, bankingService *actor.PID) error {

	router := mux.NewRouter()

	router.HandleFunc("/register-account", func(w http.ResponseWriter, r *http.Request) {
		var request RegisterAccount
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newAccountNumber := request.AccountNumber
		newActorAccount := system.Root.Spawn(actor.PropsFromProducer(func() actor.Actor { return &AccountActor{account: newAccountNumber, balance: 0} }))
		request.AccountActor = newActorAccount

		response, err := system.Root.RequestFuture(bankingService, &request, 5000*time.Second).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("POST")

	router.HandleFunc("/debit", func(w http.ResponseWriter, r *http.Request) {
		var request DebitRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := system.Root.RequestFuture(bankingService, &request, 5*time.Second).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("POST")

	router.HandleFunc("/credit", func(w http.ResponseWriter, r *http.Request) {
		var request CreditRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := system.Root.RequestFuture(bankingService, &request, 5000*time.Second).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("POST")

	router.HandleFunc("/daily-summary", func(w http.ResponseWriter, r *http.Request) {
		accountNumber := r.URL.Query().Get("accountnumber")
		date := r.URL.Query().Get("date")

		if accountNumber == "" || date == "" {
			http.Error(w, "Missing 'accountnumber' or 'date' parameter", http.StatusBadRequest)
			return
		}

		request := &GetDailySummaryRequest{Date: date, AccountNumber: accountNumber}

		response, err := system.Root.RequestFuture(bankingService, request, 5*time.Second).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	log.Info("Starting HTTP server on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

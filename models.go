package main

import "github.com/asynkron/protoactor-go/actor"

type DebitRequest struct {
	AccountNumber string
	Amount        float64
	TransactionID string
}

type DebitResponse struct {
	TransactionID string
	Success       bool
}

type CreditRequest struct {
	AccountNumber string
	Amount        float64
	TransactionID string
}

type CreditResponse struct {
	TransactionID string
	Success       bool
}

type GetDailySummaryRequest struct {
	AccountNumber string
	Date          string
}

type DailySummaryResponse struct {
	Date        string
	TotalDebit  float64
	TotalCredit float64
	Balance     float64
}

type PersistenceResponse struct {
	Success bool
	Error   string
}

type PersistStateRequest struct {
	AccountNumber string
}

type LoadStateRequest struct {
	AccountNumber string
}

type RegistrationResponse struct {
	Success bool
	Error   string
}

type RegisterAccount struct {
	AccountNumber string
	AccountActor  *actor.PID
}

type Error struct {
	AccountNumber string
	Success       bool
}

package main

import (
	"github.com/asynkron/protoactor-go/actor"
	log "github.com/sirupsen/logrus"
)

type BankingServiceActor struct {
	accounts map[string]*actor.PID
}

func (state *BankingServiceActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		log.Info("Banking service started.")
	case *RegisterAccount:
		state.accounts[msg.AccountNumber] = msg.AccountActor
		account, ok := state.accounts[msg.AccountNumber]
		if ok {
			context.Forward(account)
		} else {
			context.Respond(&RegistrationResponse{Success: false})
		}
	case *DebitRequest:
		account, ok := state.accounts[msg.AccountNumber]
		if ok {
			context.Forward(account)
		} else {
			context.Respond(&DebitResponse{TransactionID: msg.TransactionID, Success: false})
		}
	case *CreditRequest:
		account, ok := state.accounts[msg.AccountNumber]
		if ok {
			context.Forward(account)
		} else {
			context.Respond(&CreditResponse{TransactionID: msg.TransactionID, Success: false})
		}
	case *GetDailySummaryRequest:
		account, ok := state.accounts[msg.AccountNumber]
		if ok {
			context.Forward(account)
		} else {
			context.Respond(&Error{AccountNumber: msg.AccountNumber, Success: false})
		}
	case *PersistStateRequest:
		account, ok := state.accounts[msg.AccountNumber]
		if ok {
			context.Forward(account)
		} else {
			context.Respond(&Error{AccountNumber: msg.AccountNumber, Success: false})
		}
	}
}

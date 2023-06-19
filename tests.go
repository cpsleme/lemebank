package main

import (
	"github.com/asynkron/protoactor-go/actor"
	log "github.com/sirupsen/logrus"
	"time"
)

func tests(system *actor.ActorSystem, bankingService *actor.PID) error {

	// Cria algumas contas de exemplo
	account1 := system.Root.Spawn(actor.PropsFromProducer(func() actor.Actor { return &AccountActor{account: "123", balance: 0} }))
	account2 := system.Root.Spawn(actor.PropsFromProducer(func() actor.Actor { return &AccountActor{account: "456", balance: 0} }))

	system.Root.Send(bankingService, &RegisterAccount{AccountNumber: "123", AccountActor: account1})
	system.Root.Send(bankingService, &RegisterAccount{AccountNumber: "456", AccountActor: account2})

	// Lançamento de débito
	debitRequest1 := &DebitRequest{
		AccountNumber: "123",
		Amount:        50,
		TransactionID: "debit-1",
	}
	debitResponse1, err := system.Root.RequestFuture(bankingService, debitRequest1, 5*time.Second).Result()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info("Debit response:", debitResponse1)

	// Lançamento de crédito
	creditRequest1 := &CreditRequest{
		AccountNumber: "456",
		Amount:        200,
		TransactionID: "credit-1",
	}
	creditResponse1, err := system.Root.RequestFuture(bankingService, creditRequest1, 5*time.Second).Result()
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Info("Credit response:", creditResponse1)

	// Lançamento de débito
	debitRequest2 := &DebitRequest{
		AccountNumber: "456",
		Amount:        150,
		TransactionID: "credit-1",
	}
	debitResponse2, err := system.Root.RequestFuture(bankingService, debitRequest2, 5*time.Second).Result()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info("Credit response:", debitResponse2)

	// Recuperação do resumo diário
	summaryRequest := &GetDailySummaryRequest{
		AccountNumber: "456",
		Date:          "2023-06-18",
	}
	summaryResponse, err := system.Root.RequestFuture(bankingService, summaryRequest, 500*time.Second).Result()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info("Daily summary:", summaryResponse)

	return nil
}

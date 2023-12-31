package main

import (
	"encoding/json"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/boltdb/bolt"
	log "github.com/sirupsen/logrus"

	"time"
)

type AccountActor struct {
	account      string
	balance      float64
	transactions []Transaction
	db           *bolt.DB
}

type Transaction struct {
	ID     string
	Date   string
	Type   string
	Amount float64
}

func (state *AccountActor) Receive(context actor.Context) {

	switch msg := context.Message().(type) {
	case *RegisterAccount:
		dbFile := "./db/" + state.account + ".db"
		db, err := bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Error("Can't create or open a database file.")
			context.Respond(&RegistrationResponse{Success: false})
		} else {
			state.db = db
			err = state.loadState()
			if err != nil {
				log.Error()
				context.Respond(&RegistrationResponse{Success: false})
			}
		}
		context.Respond(&RegistrationResponse{Success: true})

	case *DebitRequest:
		if state.balance >= msg.Amount {
			state.balance -= msg.Amount
			t := Transaction{ID: msg.TransactionID, Date: time.Now().Format("2006-01-02"), Amount: msg.Amount, Type: "debit"}
			err := state.saveState(t)
			if err != nil {
				context.Respond(&DebitResponse{TransactionID: msg.TransactionID, Success: false})
			} else {
				state.transactions = append(state.transactions, t)
				context.Respond(&CreditResponse{TransactionID: msg.TransactionID, Success: true})
			}
		} else {
			context.Respond(&DebitResponse{TransactionID: msg.TransactionID, Success: false})
		}

	case *CreditRequest:
		state.balance += msg.Amount
		t := Transaction{ID: msg.TransactionID, Date: time.Now().Format("2006-01-02"), Amount: msg.Amount, Type: "credit"}
		err := state.saveState(t)
		if err != nil {
			context.Respond(&CreditResponse{TransactionID: msg.TransactionID, Success: false})
		} else {
			state.transactions = append(state.transactions, t)
			context.Respond(&CreditResponse{TransactionID: msg.TransactionID, Success: true})
		}

	case *GetDailySummaryRequest:
		totalDebit := 0.0
		totalCredit := 0.0
		for _, t := range state.transactions {
			if t.Date == msg.Date {
				if t.Type == "debit" {
					totalDebit += t.Amount
				} else {
					totalCredit += t.Amount
				}
			}
		}
		context.Respond(
			&DailySummaryResponse{Date: msg.Date,
				TotalDebit:  totalDebit,
				TotalCredit: totalCredit,
				Balance:     totalCredit - totalDebit})

	case *LoadStateRequest:
		err := state.loadState()
		if err != nil {
			context.Respond(&PersistenceResponse{Success: false, Error: err.Error()})
		} else {
			context.Respond(&PersistenceResponse{Success: true})
		}
	}
}

func (state *AccountActor) saveState(t Transaction) error {
	err := state.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("account"))
		if err != nil {
			return err
		}

		data, err := json.Marshal(t)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(t.ID), data)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (state *AccountActor) loadState() error {
	err := state.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("account"))
		if bucket == nil {
			return nil // Bucket doesn't exist yet, no state to load
		}

		bucket.ForEach(func(k, v []byte) error {
			var t Transaction
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}
			state.transactions = append(state.transactions, t)
			return nil
		})

		return nil
	})

	return err
}

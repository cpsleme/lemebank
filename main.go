package main

import (
	"github.com/asynkron/protoactor-go/actor"
	log "github.com/sirupsen/logrus"
)

func main() {
	system := actor.NewActorSystem()

	props := actor.
		PropsFromProducer(func() actor.Actor { return &BankingServiceActor{accounts: make(map[string]*actor.PID)} })

	bankingService := system.Root.Spawn(props)

	// Start API
	err := tests(system, bankingService)
	if err != nil {
		log.Panic(err)
	}

	// Start API
	err = api(system, bankingService)
	if err != nil {
		log.Fatal(err)
	}

}

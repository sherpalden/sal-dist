package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sherpalden/sal-dist/appConfig"
	"github.com/sherpalden/sal-dist/entity"
	"github.com/sherpalden/sal-dist/repository"
	"github.com/sherpalden/sal-dist/service"
	"github.com/sherpalden/sal-dist/store"
)

func entityValidator(val string) error {
	if val != "sender" && val != "employee" {
		errMsg := "invalid entity value: should be either 'sender' or 'employee'"
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func main() {
	var object string

	for {
		fmt.Print("Check balance of sender or employee: ")
		fmt.Scanln(&object)
		if entityValidator(object) == nil {
			break
		}
	}

	env := appConfig.GetEnv()

	client, err := ethclient.DialContext(context.Background(), env.NetworkURl)
	if err != nil {
		log.Fatalf("failed to initialize network client: %v", err)
	}

	if object == "employee" {
		empStore := store.NewJSONFileStore("./employees.json")
		empRepo := repository.NewEmployeeRepository(empStore)
		empService := service.NewEmployeeService(empRepo)
		employees, err := empService.GetAll()
		if err != nil {
			log.Fatalf("failed to get employees: %v", err)
		}
		if len(employees) == 0 {
			fmt.Println("no employees created")
		}

		for _, emp := range employees {
			fmt.Println()
			empBalance, err := emp.GetBalance(client)
			if err != nil {
				log.Fatalf("failed to employee balance: %v", err)
			}
			fmt.Printf("Employee %v balance: %v ETH\n", emp.Name, empBalance.Eth())
		}

	} else {
		sender, err := entity.NewSender(env.OwnerPrivateKey, env.OwnerPublicAddress)
		if err != nil {
			log.Fatalf("failed to get sender: %v", err)
		}
		senderBalance, err := sender.GetBalance(client)
		fmt.Printf("\nSender balance: %v ETH\n", senderBalance.Eth())
	}
}

package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sherpalden/sal-dist/entity"
	"github.com/sherpalden/sal-dist/goCli"
	"github.com/sherpalden/sal-dist/repository"
	"github.com/sherpalden/sal-dist/service"
	"github.com/sherpalden/sal-dist/store"
)

func salaryValidator(val string) error {
	_, ok := new(big.Int).SetString(val, 10) // base 10 for decimal numbers
	if !ok {
		log.Fatal("invalid salary value")
	}
	return nil
}

func walletAddressValidator(address string) error {
	if !common.IsHexAddress(address) {
		log.Fatal("invalid wallet address")
	}

	return nil
}

func main() {
	argConfig := []goCli.ArgConfig{
		{Name: "id", Required: true},
		{Name: "name", Required: true},
		{Name: "salary", Required: true, Validate: salaryValidator},
		{Name: "walletaddress", Required: true, Validate: walletAddressValidator},
	}

	args, err := goCli.GetArgs(argConfig)
	if err != nil {
		log.Fatalf("failed to get arguments for employee: %v", err)
	}

	salary, _ := new(big.Int).SetString(args["salary"].(string), 10)

	emp := entity.Employee{
		ID:            args["id"].(string),
		Name:          args["name"].(string),
		WalletAddress: args["walletaddress"].(string),
		Salary:        salary,
	}

	empStore := store.NewJSONFileStore("./employees.json")
	empRepo := repository.NewEmployeeRepository(empStore)
	empService := service.NewEmployeeService(empRepo)
	if err := empService.Create(emp); err != nil {
		log.Fatalf("failed to create an employee: %v", err)
	}
	fmt.Println("employee creation successful")
}

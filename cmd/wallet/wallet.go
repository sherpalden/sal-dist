package main

import (
	"fmt"
	"log"

	"github.com/sherpalden/sal-dist/goCli"
	"github.com/sherpalden/sal-dist/wallet"
)

func main() {
	argConfig := []goCli.ArgConfig{
		{Name: "passphrase", Required: true},
		{Name: "keypath", Required: true},
	}

	args, err := goCli.GetArgs(argConfig)
	if err != nil {
		log.Fatal("failed to get arguments for wallet: passphrase and keypath args required")
	}

	password := args["passphrase"].(string)
	keyPath := args["keypath"].(string)

	if err := wallet.CreateWallet(password, keyPath); err != nil {
		log.Fatal("failed to create wallet")
	}
	fmt.Println("wallet creation successful")
}

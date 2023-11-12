package appConfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	NetworkURl         string
	OwnerPrivateKey    string
	OwnerPublicAddress string
}

func GetEnv() Env {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return Env{
		NetworkURl:         os.Getenv("NETWORK_URL"),
		OwnerPrivateKey:    os.Getenv("OWNER_PRIVATE_KEY"),
		OwnerPublicAddress: os.Getenv("OWNER_PUBLIC_ADDRESS"),
	}
}

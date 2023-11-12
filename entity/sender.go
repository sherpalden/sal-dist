package entity

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Sender struct {
	privateKey ecdsa.PrivateKey
	Account
}

func NewSender(pk string, address string) (*Sender, error) {
	privateKeyBytes, err := hex.DecodeString(pk)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	sender := Sender{
		privateKey: *privateKey,
		Account: Account{
			Address: common.HexToAddress(address),
		},
	}

	return &sender, nil
}

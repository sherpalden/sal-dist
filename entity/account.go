package entity

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	Address common.Address
}

func (ac Account) GetBalance(cl *ethclient.Client) (*WeiCurrency, error) {
	bigIntVal, err := cl.BalanceAt(context.Background(), ac.Address, nil)
	wc := WeiCurrency(*bigIntVal)
	return &wc, err
}

func (ac Account) GetPendingNonce(cl *ethclient.Client) (uint64, error) {
	return cl.PendingNonceAt(context.Background(), ac.Address)
}

package entity

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Employee struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	WalletAddress string   `json:"wallet_address"`
	Salary        *big.Int `json:"salary"`
}

func (emp Employee) GetBalance(cl *ethclient.Client) (*WeiCurrency, error) {
	bigIntVal, err := cl.BalanceAt(context.Background(), common.HexToAddress(emp.WalletAddress), nil)
	wc := WeiCurrency(*bigIntVal)
	return &wc, err
}

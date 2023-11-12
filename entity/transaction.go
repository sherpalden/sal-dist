package entity

import (
	"context"
	"math/big"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Transaction struct {
	Sender    *Sender
	ToAccount Account
	Value     *big.Int
}

func (tx Transaction) Execute(cl *ethclient.Client) error {
	nonce, err := tx.Sender.GetPendingNonce(cl)
	if err != nil {
		return err
	}

	gasPrice, err := cl.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainId, err := cl.NetworkID(context.Background())
	if err != nil {
		return err
	}

	txn := ethTypes.NewTransaction(nonce, tx.ToAccount.Address, tx.Value, 21000, gasPrice, nil)

	txn, err = ethTypes.SignTx(txn, ethTypes.NewEIP155Signer(chainId), &tx.Sender.privateKey)
	if err != nil {
		return err
	}

	return cl.SendTransaction(context.Background(), txn)
}

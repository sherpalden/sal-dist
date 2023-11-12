package entity

import (
	"math/big"
)

const (
	EthToWeiMultiplier int64 = 1e18
)

type WeiCurrency big.Int

func (wc *WeiCurrency) Eth() *big.Float {
	ether := new(big.Float)
	ether.SetInt(new(big.Int).Set((*big.Int)(wc)))
	ether.Quo(ether, big.NewFloat(float64(EthToWeiMultiplier)))
	return ether
}

func (wc *WeiCurrency) BigInt() *big.Int {
	return new(big.Int).Set((*big.Int)(wc))
}

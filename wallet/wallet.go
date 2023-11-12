package wallet

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateWallet(passphrase string, destPath string) error {
	key := keystore.NewKeyStore(destPath, keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := key.NewAccount(passphrase)
	return err
}

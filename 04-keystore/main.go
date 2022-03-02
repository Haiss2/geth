package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password123"
	// account, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print(account)

	b, err := ioutil.ReadFile("./wallet/UTC--2022-03-01T08-54-24.994780000Z--f85a475e344f7b263b6b92ff7376c5f7541afcc0")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("PrivateKey", hexutil.Encode(pData))

	puData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub", hexutil.Encode(puData))

	fmt.Println("Address", crypto.PubkeyToAddress(key.PrivateKey.PublicKey))

}

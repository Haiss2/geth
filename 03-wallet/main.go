package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Fail to create new private key %v", err)
	}
	// Private key is the most important part and should always be stored in a secure place
	// The person who has this key owns the account
	// It is used to encrypt a transaction
	// The private key is generated using Elliptic Curve Digital Signature Algorithm
	pData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pData))

	// Public key can be generated from private key, it's one way
	// PubKey is used to check and validate a transaction
	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	// Public address is the key that used to make transaction
	address := crypto.PubkeyToAddress(pvk.PublicKey).Hex()

	fmt.Println(address)
}

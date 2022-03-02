/*
From: {
	from: "Address_1",
	to: "Address_2",
	value: "Amount of Ether that we send from W1 to W2, value set to Wei",
	gasLimit: "Default Value = 21000",
	gasPrice: "Important value, if too low -> fail, high -> spend more Ether"
	nonce: Number that we use once, each time we do transaction, we += 1
		we get this value from Ethereum network
}

*/

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	mainUrl  = "https://mainnet.infura.io/v3/5117d8719cc94d49a2d2af843cd35309"
	kovanUrl = "https://kovan.infura.io/v3/5117d8719cc94d49a2d2af843cd35309"
	palmUrl  = "https://palm-testnet.infura.io/v3/5117d8719cc94d49a2d2af843cd35309"

	addr_1 = "feec37b179f2c97d6f16c43300557ed12254b190"
	addr_2 = "3a4b7979d740049434f645c59f39883e512fb811"
)

func main() {
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	// _, err := ks.NewAccount("password123")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = ks.NewAccount("password123")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client, err := ethclient.Dial(palmUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	a1 := common.HexToAddress(addr_1)
	a2 := common.HexToAddress(addr_2)

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("B1:", b1)
	fmt.Println("B2:", b2)

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}

	amount := big.NewInt(100000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)

	chanID, err := client.NetworkID(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2022-03-02T04-00-19.649694000Z--feec37b179f2c97d6f16c43300557ed12254b190")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password123")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP155Signer(chanID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx send : %s", tx.Hash().Hex())
}

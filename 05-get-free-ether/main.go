package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
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
}

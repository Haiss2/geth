package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/5117d8719cc94d49a2d2af843cd35309"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}

	fmt.Println("The block number:", block.Number())

	addr := "0xdfde315ab0aaf8abf6d22a9991c535718d6211b3"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance: %v", err)
	}
	fmt.Println("The balance:", balance)

	// 1 ether = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())

	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("Balance of account in ETH:", value)

}

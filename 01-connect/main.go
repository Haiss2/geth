package main

import (
	"context"
	"fmt"
	"log"

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

}

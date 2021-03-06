package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	LP "./contracts"
)

func main() {

	var contractAddress = "0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc";
	var url = "wss://jolly-brattain:class-awning-speech-rule-mangy-gating@ws-nd-780-530-825.p2pify.com";

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := LP.NewLiquidityPool(address, client)

	if err != nil {
		log.Fatal(err)
	}

	reserves, err := instance.GetReserves(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Reserves of reserve0: ", reserves.Reserve0)
	fmt.Println("Reserves of reserve1: ", reserves.Reserve1)
	fmt.Println("Timestamp is: ", reserves.BlockTimestampLast)
}

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SmartContractAddr = "0x27695e09149adc738a978e9a678f99e4c39e9eb9"
)

func main() {
	client, err := ethclient.Dial("mainnet.infura.io/v3/8f81eca071e245e6934ad24b23c722d0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	// Loading smart contract
	address := common.HexToAddress(SmartContractAddr)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance

}

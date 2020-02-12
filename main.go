package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/emikhalev/kickico/contracts/cstoken"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractABI       = "contracts/CSToken.json"
	SmartContractAddr = "0x27695e09149adc738a978e9a678f99e4c39e9eb9"
)

func main() {
	listMethods()

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/8f81eca071e245e6934ad24b23c722d0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	// Loading smart contract:
	address := common.HexToAddress(SmartContractAddr)
	instance, err := cstoken.NewCstoken(address, client)
	if err != nil {
		log.Fatal(err)
	}
	if ca, err := instance.CountAddresses(nil); err==nil {
		fmt.Printf("ca: %v", *ca)
	} else {
		log.Fatal(err)
	}
}

func listMethods() {
	abiContent, err := ioutil.ReadFile(contractABI)
	if err!=nil {
		log.Fatal(err)
	}

	cABI, err := abi.JSON(bytes.NewReader(abiContent))
	if err!=nil {
		log.Fatal(err)
	}
	for mName := range cABI.Methods {
		println(mName)
	}
}

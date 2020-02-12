package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/emikhalev/kickio/contracts/cstoken"
)

const (
	contractABI = "contracts/CSToken.json"
	SmartContractAddr = "0x27695e09149adc738a978e9a678f99e4c39e9eb9"
)

func main() {
	/*_, err := ethclient.Dial("mainnet.infura.io/v3/8f81eca071e245e6934ad24b23c722d0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")*/

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
	// Loading smart contract
	//address := common.HexToAddress(SmartContractAddr)
	//instance, err := store.NewStore(address, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = instance

}

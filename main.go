package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/emikhalev/kickico/abi"
)

const (
	contractABI       = "contracts/CSToken.json"
	smartContractAddr = "0x27695e09149adc738a978e9a678f99e4c39e9eb9"
	infuraURL         = "https://mainnet.infura.io/v3/8f81eca071e245e6934ad24b23c722d0"
	callMethodTimeout = 30 * time.Second
)

type params struct {
	list   bool
	method string
	params string
}

var (
	execParams params = params{}
)

func init() {
	flag.BoolVar(&execParams.list, "list", false, "List contract methods")
	flag.StringVar(&execParams.method, "method", "", "Contract method name to call")
	flag.StringVar(&execParams.params, "p", "", "Input parameters to call function (if any)")
}

func main() {

	flag.Parse()

	inst, err := abi.NewABI(contractABI, infuraURL, smartContractAddr)
	if err != nil {
		log.Fatalf("cannot load ABI: %v", err)
	}

	if execParams.list {
		inst.PrintMethods()
		return
	}

	method := strings.TrimSpace(execParams.method)
	if method != "" {
		params, err := inst.ParseParameters(method, execParams.params)
		if err != nil {
			log.Fatalf("cannot read parameters: %v", err)
		}
		flag.Parse()

		ctx, cancelFn := context.WithTimeout(context.Background(), callMethodTimeout)
		defer cancelFn()

		result, err := inst.CallMethod(ctx, method, params...)
		if err != nil {
			log.Fatalf("cannot call ABI: %v", err)
		}

		fmt.Printf("Call method %s return\n", method)
		for k, v := range result {
			if k == "" {
				k = "<unnamed>"
			}
			fmt.Printf("%s: %v\n", k, v)
		}

		return
	}

	printUsage()
}

func printUsage() {
	println("")
	println("Usage:")
	println("  -list - list contract methods")
	println("  -method=<methodName> -p=<parameters separated with comma> - call contract method with parameters (i.e. -method=disableManuallyBurnTokens -p=false)")
	println("")
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kickico/abi"
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
	params []interface{}
}

var (
	execParams params = params{}
)

func init() {
	flag.BoolVar(&execParams.list, "list", false, "List contract methods")
	flag.StringVar(&execParams.method, "method", "", "Contract method name to call")
}

func main() {

	flag.Parse()

	inst, err := abi.NewABI(contractABI, infuraURL)
	if err != nil {
		log.Fatalf("cannot load ABI: %v", err)
	}

	if execParams.list {
		inst.PrintMethods()
		return
	}

	method := strings.TrimSpace(execParams.method)
	if method != "" {
		execParams.params, err = inst.FlagMethodParams(method)
		if err!=nil {
			log.Fatalf("cannot read parameters: %v", err)
		}
		flag.Parse()

		ctx, cancelFn := context.WithTimeout(context.Background(), callMethodTimeout)
		defer cancelFn()

		result, err := inst.CallMethod(ctx, execParams.method, execParams.params)
		if err != nil {
			log.Fatalf("cannot call ABI: %v", err)
		}
		fmt.Printf("Call method %s returns: %#v", method, result)
		return
	}

	printUsage()
}

func printUsage() {
	println("")
	println("Usage:")
	println("  -list - list contract methods")
	println("  -method=<methodName> -<paramName1>=<paramValue1> -<paramName2>=<paramValue2> ... - call contract method with parameters (i.e. -method=disableManuallyBurnTokens -_disable=false)")
	println("")
}

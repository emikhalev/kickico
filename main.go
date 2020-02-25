package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/emikhalev/kickico/internal/abi"
	"github.com/emikhalev/kickico/internal/config"
)

type params struct {
	configPath string
	list       bool
	method     string
	params     string
}

var (
	execParams params = params{}
)

func init() {
	flag.StringVar(&execParams.configPath, "config", "configs/config.yaml", "config file (default: configs/config.yaml")
	flag.BoolVar(&execParams.list, "list", false, "List contract methods")
	flag.StringVar(&execParams.method, "method", "", "Contract method name to call")
	flag.StringVar(&execParams.params, "p", "", "Input parameters to call function (if any)")
}

func main() {
	flag.Parse()

	if execParams.list {
		inst := loadABI(loadConfig(execParams.configPath))
		inst.PrintMethods()
		return
	}

	method := strings.TrimSpace(execParams.method)
	if method != "" {
		cfg := loadConfig(execParams.configPath)
		inst := loadABI(cfg)

		params, err := inst.ParseParameters(method, execParams.params)
		if err != nil {
			log.Fatalf("cannot read parameters: %v", err)
		}

		ctx, cancelFn := context.WithTimeout(context.Background(), cfg.CallMethodTimeout)
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

func loadABI(cfg *config.Config) *abi.ABI {
	inst, err := abi.NewABI(cfg.ContractABI, cfg.InfuraURL, cfg.SmartContractAddress)
	if err != nil {
		log.Fatalf("cannot load ABI: %v", err)
	}
	return inst
}

func loadConfig(path string) *config.Config {
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	return cfg
}

func printUsage() {
	println("")
	println("Usage:")
	println("  -list - list contract methods")
	println("  -method=<methodName> -p=<parameters separated with comma> - call contract method with parameters (i.e. -method=disableManuallyBurnTokens -p=false)")
	println("  -config=<path to config file> - to set alternative config file (default: configs/config.yaml_")
	println("")
}

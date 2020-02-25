package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ContractABI          string
	SmartContractAddress string
	InfuraURL            string
	CallMethodTimeout    time.Duration
}

const (
	contractABI          = "ContractABI"
	smartContractAddress = "SmartContractAddress"
	infuraURL            = "InfuraURL"
	callMethodTimeout    = "CallMethodTimeout"
)

func init() {
	viper.SetDefault(contractABI, "contracts/CSToken.json")
	viper.SetDefault(smartContractAddress, "0x27695e09149adc738a978e9a678f99e4c39e9eb9")
	viper.SetDefault(infuraURL, "https://mainnet.infura.io/v3/8f81eca071e245e6934ad24b23c722d0")
	viper.SetDefault(callMethodTimeout, 30*time.Second)
}

func NewConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config file: %s \n", err)
	}
	c := new(Config)
	c.readValues()
	return c, nil
}

func (c *Config) readValues() {
	c.ContractABI = viper.GetString(contractABI)
	c.SmartContractAddress = viper.GetString(smartContractAddress)
	c.InfuraURL = viper.GetString(infuraURL)
	c.CallMethodTimeout = viper.GetDuration(callMethodTimeout)
}

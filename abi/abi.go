package abi

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ABI struct {
	abi *abi.ABI
	url string
}

func NewABI(contractABI string, url string) (*ABI, error) {
	abiContent, err := ioutil.ReadFile(contractABI)
	if err != nil {
		return nil, err
	}

	cABI, err := abi.JSON(bytes.NewReader(abiContent))
	if err != nil {
		return nil, err
	}

	return &ABI{
		abi: &cABI,
		url: url,
	}, nil
}

func (a *ABI) PrintMethods() {
	for _, m := range a.abi.Methods {
		fmt.Printf("NAME: %s\nINPUTS: %s\nOUTPUTS: %s\n\n", m.Name,
			a.sprintArguments(m.Inputs), a.sprintArguments(m.Outputs))
	}
}

func (a *ABI) FlagMethodParams(method string) ([]interface{}, error) {
	m, ok := a.abi.Methods[method]
	if !ok {
		return nil, errors.New("method not found")
	}

	p := make([]interface{}, len(m.Inputs))

	for idx, o := range m.Inputs {
		switch o.Type.T {
		case abi.IntTy:
			p[idx] = flag.Int64(o.Name, 0, "")
		case abi.UintTy:
			p[idx] = flag.Uint64(o.Name, 0, "")
		case abi.BoolTy:
			p[idx] = flag.Bool(o.Name, false, "")
		case abi.SliceTy:
			p[idx] = flag.String(o.Name, "", "")
		case abi.ArrayTy:
			p[idx] = flag.String(o.Name, "", "")
		case abi.TupleTy: // TODO: not implemented
		case abi.AddressTy:
			p[idx] = flag.String(o.Name, "", "")
		case abi.FixedBytesTy:
			p[idx] = flag.String(o.Name, "", "")
		case abi.BytesTy: // TODO: not implemented
		case abi.HashTy: // TODO: not implemented
		case abi.FixedPointTy: // TODO: not implemented
		case abi.FunctionTy: // TODO: not implemented
		}
	}

	return p, nil
}

func (a *ABI) CallMethod(ctx context.Context, method string, params ...interface{}) (map[string]interface{}, error) {
	client, err := ethclient.Dial(a.url)
	if err != nil {
		return nil, fmt.Errorf("cannot dial: %v", err)
	}

	// Prepare data for eth_call
	address := common.HexToAddress(smartContractAddr)

	methodData, err := a.packCallMethodData(method, params)
	if err != nil {
		return nil, errors.Errorf("packCallMethodData error: %v", err)
	}

	// eth_call
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &address,
		Data: methodData,
	}, nil)

	if err != nil {
		return nil, fmt.Errorf("call method returns error: %v", err)
	}

	r, err := a.unpackCallMethodData(method, result)
	if err != nil {
		return nil, fmt.Errorf("cannot unpack method return result: %v", err)
	}

	return r, nil
}

func (a *ABI) sprintArguments(args abi.Arguments) string {
	res := ""
	for idx := range args {
		res += fmt.Sprintf("%s: %s[%d]", args[idx].Name, args[idx].Type, args[idx].Type.T)
		if idx != len(args)-1 {
			res += ", "
		}
	}
	return res
}

func (a *ABI) packCallMethodData(method string, args ...interface{}) ([]byte, error) {
	data, err := a.abi.Pack(method, args)
	if err != nil {
		return nil, fmt.Errorf("cannot pack method call data: %v", err)
	}

	return data, nil
}

func (a *ABI) unpackCallMethodData(method string, data []byte) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	if err := a.abi.UnpackIntoMap(v, method, data); err != nil {
		return nil, fmt.Errorf("cannot unpack into map: %v", err)
	}

	return v, nil
}

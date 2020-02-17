package abi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

type ABI struct {
	abi               *abi.ABI
	url               string
	smartContractAddr string
}

func NewABI(contractABI, url, smartContractAddr string) (*ABI, error) {
	abiContent, err := ioutil.ReadFile(contractABI)
	if err != nil {
		return nil, err
	}

	cABI, err := abi.JSON(bytes.NewReader(abiContent))
	if err != nil {
		return nil, err
	}

	return &ABI{
		abi:               &cABI,
		url:               url,
		smartContractAddr: smartContractAddr,
	}, nil
}

func (a *ABI) PrintMethods() {
	for _, m := range a.abi.Methods {
		fmt.Printf("NAME: %s\nINPUTS: %s\nOUTPUTS: %s\n\n", m.Name,
			a.sprintArguments(m.Inputs), a.sprintArguments(m.Outputs))
	}
}

func (a *ABI) CallMethod(ctx context.Context, method string, params ...interface{}) (map[string]interface{}, error) {
	client, err := ethclient.Dial(a.url)
	if err != nil {
		return nil, fmt.Errorf("cannot dial: %v", err)
	}

	// Prepare data for eth_call
	address := common.HexToAddress(a.smartContractAddr)

	methodData, err := a.packCallMethodData(method, params...)
	if err != nil {
		return nil, errors.Errorf("packCallMethodData error: %v", err)
	}

	// eth_call
	result, err := client.CallContract(ctx, ethereum.CallMsg{
		To:   &address,
		Data: methodData,
	}, nil)

	if err == nil && len(result) == 0 {
		if code, err := client.CodeAt(ctx, address, nil); err != nil {
			return nil, err
		} else if len(code) == 0 {
			return nil, errors.New("contract code empty")
		}
	}
	if err != nil {
		return nil, fmt.Errorf("call method returns error: %v", err)
	}

	r, err := a.unpackCallMethodData(method, result)
	if err != nil {
		return nil, fmt.Errorf("cannot unpack method return result: %v", err)
	}

	return r, nil
}

func (a *ABI) ParseParameters(method, params string) ([]interface{}, error) {
	m, ok := a.abi.Methods[method]
	if !ok {
		return nil, errors.New("method not found")
	}

	p := make([]interface{}, len(m.Inputs))
	iP := strings.Split(params, ",")

	if len(iP) < len(m.Inputs) {
		return nil, errors.New("not all parameters set")
	}

	for idx, o := range m.Inputs {
		switch o.Type.T {
		case abi.IntTy, abi.UintTy:
			pv, err := strconv.ParseUint(iP[idx], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("cannot parse parametes: %s (type: %v, index: %d)", o.Name, o.Type, idx)
			}
			p[idx] = big.NewInt(int64(pv))
		case abi.BoolTy: //TODO: Check
			v, err := strconv.ParseBool(iP[idx])
			if err != nil {
				return nil, fmt.Errorf("cannot parse parametes: %s (type: %v, index: %d)", o.Name, o.Type, idx)
			}
			p[idx] = v
		case abi.SliceTy: //TODO: slice
		case abi.ArrayTy: //TODO: array
		case abi.TupleTy: // TODO: struct
		case abi.AddressTy: //TODO: Check
			var v [20]byte
			copy(v[:], iP[idx])
			p[idx] = v
		case abi.FixedBytesTy: // TODO: Array
		case abi.BytesTy: //TODO: Check
			v := make([]byte, len(iP[idx]))
			copy(v, iP[idx])
			p[idx] = v
		case abi.HashTy: // TODO: not implemented
		case abi.FixedPointTy: // TODO: not implemented
		case abi.FunctionTy: // TODO: not implemented
		}
	}

	return p, nil
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
	data, err := a.abi.Pack(method, args...)
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

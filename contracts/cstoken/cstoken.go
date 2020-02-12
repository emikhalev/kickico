// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cstoken

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CstokenABI is the input ABI used to generate the binding from.
const CstokenABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"allAgingTimesAdded\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disable\",\"type\":\"bool\"}],\"name\":\"disableTransfers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\"},{\"name\":\"agingTime\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"calculateDividends\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowManuallyBurnTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disable\",\"type\":\"bool\"}],\"name\":\"disableManuallyBurnTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveDividends\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"addAgingTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfersEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"accountBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"agingBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"poolAddress\",\"type\":\"address\"},{\"name\":\"agingTime\",\"type\":\"uint256\"}],\"name\":\"addAgingTimesForPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countAddresses\",\"outputs\":[{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"agingTime\",\"type\":\"uint256\"}],\"name\":\"AgingTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Issuance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Destruction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"NewSmartToken\",\"type\":\"event\"}]"

// Cstoken is an auto generated Go binding around an Ethereum contract.
type Cstoken struct {
	CstokenCaller     // Read-only binding to the contract
	CstokenTransactor // Write-only binding to the contract
	CstokenFilterer   // Log filterer for contract events
}

// CstokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CstokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CstokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CstokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CstokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CstokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CstokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CstokenSession struct {
	Contract     *Cstoken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CstokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CstokenCallerSession struct {
	Contract *CstokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CstokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CstokenTransactorSession struct {
	Contract     *CstokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CstokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CstokenRaw struct {
	Contract *Cstoken // Generic contract binding to access the raw methods on
}

// CstokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CstokenCallerRaw struct {
	Contract *CstokenCaller // Generic read-only contract binding to access the raw methods on
}

// CstokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CstokenTransactorRaw struct {
	Contract *CstokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCstoken creates a new instance of Cstoken, bound to a specific deployed contract.
func NewCstoken(address common.Address, backend bind.ContractBackend) (*Cstoken, error) {
	contract, err := bindCstoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cstoken{CstokenCaller: CstokenCaller{contract: contract}, CstokenTransactor: CstokenTransactor{contract: contract}, CstokenFilterer: CstokenFilterer{contract: contract}}, nil
}

// NewCstokenCaller creates a new read-only instance of Cstoken, bound to a specific deployed contract.
func NewCstokenCaller(address common.Address, caller bind.ContractCaller) (*CstokenCaller, error) {
	contract, err := bindCstoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CstokenCaller{contract: contract}, nil
}

// NewCstokenTransactor creates a new write-only instance of Cstoken, bound to a specific deployed contract.
func NewCstokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CstokenTransactor, error) {
	contract, err := bindCstoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CstokenTransactor{contract: contract}, nil
}

// NewCstokenFilterer creates a new log filterer instance of Cstoken, bound to a specific deployed contract.
func NewCstokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CstokenFilterer, error) {
	contract, err := bindCstoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CstokenFilterer{contract: contract}, nil
}

// bindCstoken binds a generic wrapper to an already deployed contract.
func bindCstoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CstokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cstoken *CstokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cstoken.Contract.CstokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cstoken *CstokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cstoken.Contract.CstokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cstoken *CstokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cstoken.Contract.CstokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cstoken *CstokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cstoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cstoken *CstokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cstoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cstoken *CstokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cstoken.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xd294cb0f.
//
// Solidity: function accountBalance(address _address) constant returns(uint256 balance)
func (_Cstoken *CstokenCaller) AccountBalance(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "accountBalance", _address)
	return *ret0, err
}

// AccountBalance is a free data retrieval call binding the contract method 0xd294cb0f.
//
// Solidity: function accountBalance(address _address) constant returns(uint256 balance)
func (_Cstoken *CstokenSession) AccountBalance(_address common.Address) (*big.Int, error) {
	return _Cstoken.Contract.AccountBalance(&_Cstoken.CallOpts, _address)
}

// AccountBalance is a free data retrieval call binding the contract method 0xd294cb0f.
//
// Solidity: function accountBalance(address _address) constant returns(uint256 balance)
func (_Cstoken *CstokenCallerSession) AccountBalance(_address common.Address) (*big.Int, error) {
	return _Cstoken.Contract.AccountBalance(&_Cstoken.CallOpts, _address)
}

// AddressByIndex is a free data retrieval call binding the contract method 0xd8ab9208.
//
// Solidity: function addressByIndex(uint256 ) constant returns(address)
func (_Cstoken *CstokenCaller) AddressByIndex(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "addressByIndex", arg0)
	return *ret0, err
}

// AddressByIndex is a free data retrieval call binding the contract method 0xd8ab9208.
//
// Solidity: function addressByIndex(uint256 ) constant returns(address)
func (_Cstoken *CstokenSession) AddressByIndex(arg0 *big.Int) (common.Address, error) {
	return _Cstoken.Contract.AddressByIndex(&_Cstoken.CallOpts, arg0)
}

// AddressByIndex is a free data retrieval call binding the contract method 0xd8ab9208.
//
// Solidity: function addressByIndex(uint256 ) constant returns(address)
func (_Cstoken *CstokenCallerSession) AddressByIndex(arg0 *big.Int) (common.Address, error) {
	return _Cstoken.Contract.AddressByIndex(&_Cstoken.CallOpts, arg0)
}

// AgingBalanceOf is a free data retrieval call binding the contract method 0xe27f0236.
//
// Solidity: function agingBalanceOf(address , uint256 ) constant returns(uint256)
func (_Cstoken *CstokenCaller) AgingBalanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "agingBalanceOf", arg0, arg1)
	return *ret0, err
}

// AgingBalanceOf is a free data retrieval call binding the contract method 0xe27f0236.
//
// Solidity: function agingBalanceOf(address , uint256 ) constant returns(uint256)
func (_Cstoken *CstokenSession) AgingBalanceOf(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Cstoken.Contract.AgingBalanceOf(&_Cstoken.CallOpts, arg0, arg1)
}

// AgingBalanceOf is a free data retrieval call binding the contract method 0xe27f0236.
//
// Solidity: function agingBalanceOf(address , uint256 ) constant returns(uint256)
func (_Cstoken *CstokenCallerSession) AgingBalanceOf(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Cstoken.Contract.AgingBalanceOf(&_Cstoken.CallOpts, arg0, arg1)
}

// AllowManuallyBurnTokens is a free data retrieval call binding the contract method 0x30762e2e.
//
// Solidity: function allowManuallyBurnTokens() constant returns(bool)
func (_Cstoken *CstokenCaller) AllowManuallyBurnTokens(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "allowManuallyBurnTokens")
	return *ret0, err
}

// AllowManuallyBurnTokens is a free data retrieval call binding the contract method 0x30762e2e.
//
// Solidity: function allowManuallyBurnTokens() constant returns(bool)
func (_Cstoken *CstokenSession) AllowManuallyBurnTokens() (bool, error) {
	return _Cstoken.Contract.AllowManuallyBurnTokens(&_Cstoken.CallOpts)
}

// AllowManuallyBurnTokens is a free data retrieval call binding the contract method 0x30762e2e.
//
// Solidity: function allowManuallyBurnTokens() constant returns(bool)
func (_Cstoken *CstokenCallerSession) AllowManuallyBurnTokens() (bool, error) {
	return _Cstoken.Contract.AllowManuallyBurnTokens(&_Cstoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_Cstoken *CstokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_Cstoken *CstokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Cstoken.Contract.Allowance(&_Cstoken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) constant returns(uint256 remaining)
func (_Cstoken *CstokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Cstoken.Contract.Allowance(&_Cstoken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Cstoken *CstokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Cstoken *CstokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Cstoken.Contract.BalanceOf(&_Cstoken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256 balance)
func (_Cstoken *CstokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Cstoken.Contract.BalanceOf(&_Cstoken.CallOpts, _owner)
}

// CountAddresses is a free data retrieval call binding the contract method 0xec530de6.
//
// Solidity: function countAddresses() constant returns(uint256 length)
func (_Cstoken *CstokenCaller) CountAddresses(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "countAddresses")
	return *ret0, err
}

// CountAddresses is a free data retrieval call binding the contract method 0xec530de6.
//
// Solidity: function countAddresses() constant returns(uint256 length)
func (_Cstoken *CstokenSession) CountAddresses() (*big.Int, error) {
	return _Cstoken.Contract.CountAddresses(&_Cstoken.CallOpts)
}

// CountAddresses is a free data retrieval call binding the contract method 0xec530de6.
//
// Solidity: function countAddresses() constant returns(uint256 length)
func (_Cstoken *CstokenCallerSession) CountAddresses() (*big.Int, error) {
	return _Cstoken.Contract.CountAddresses(&_Cstoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Cstoken *CstokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Cstoken *CstokenSession) Decimals() (uint8, error) {
	return _Cstoken.Contract.Decimals(&_Cstoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Cstoken *CstokenCallerSession) Decimals() (uint8, error) {
	return _Cstoken.Contract.Decimals(&_Cstoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Cstoken *CstokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Cstoken *CstokenSession) Name() (string, error) {
	return _Cstoken.Contract.Name(&_Cstoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Cstoken *CstokenCallerSession) Name() (string, error) {
	return _Cstoken.Contract.Name(&_Cstoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Cstoken *CstokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Cstoken *CstokenSession) Owner() (common.Address, error) {
	return _Cstoken.Contract.Owner(&_Cstoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Cstoken *CstokenCallerSession) Owner() (common.Address, error) {
	return _Cstoken.Contract.Owner(&_Cstoken.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_Cstoken *CstokenCaller) Standard(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "standard")
	return *ret0, err
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_Cstoken *CstokenSession) Standard() (string, error) {
	return _Cstoken.Contract.Standard(&_Cstoken.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_Cstoken *CstokenCallerSession) Standard() (string, error) {
	return _Cstoken.Contract.Standard(&_Cstoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Cstoken *CstokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Cstoken *CstokenSession) Symbol() (string, error) {
	return _Cstoken.Contract.Symbol(&_Cstoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Cstoken *CstokenCallerSession) Symbol() (string, error) {
	return _Cstoken.Contract.Symbol(&_Cstoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 totalSupply)
func (_Cstoken *CstokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 totalSupply)
func (_Cstoken *CstokenSession) TotalSupply() (*big.Int, error) {
	return _Cstoken.Contract.TotalSupply(&_Cstoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256 totalSupply)
func (_Cstoken *CstokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Cstoken.Contract.TotalSupply(&_Cstoken.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() constant returns(bool)
func (_Cstoken *CstokenCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cstoken.contract.Call(opts, out, "transfersEnabled")
	return *ret0, err
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() constant returns(bool)
func (_Cstoken *CstokenSession) TransfersEnabled() (bool, error) {
	return _Cstoken.Contract.TransfersEnabled(&_Cstoken.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() constant returns(bool)
func (_Cstoken *CstokenCallerSession) TransfersEnabled() (bool, error) {
	return _Cstoken.Contract.TransfersEnabled(&_Cstoken.CallOpts)
}

// AddAgingTime is a paid mutator transaction binding the contract method 0x8d37f52c.
//
// Solidity: function addAgingTime(uint256 time) returns()
func (_Cstoken *CstokenTransactor) AddAgingTime(opts *bind.TransactOpts, time *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "addAgingTime", time)
}

// AddAgingTime is a paid mutator transaction binding the contract method 0x8d37f52c.
//
// Solidity: function addAgingTime(uint256 time) returns()
func (_Cstoken *CstokenSession) AddAgingTime(time *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.AddAgingTime(&_Cstoken.TransactOpts, time)
}

// AddAgingTime is a paid mutator transaction binding the contract method 0x8d37f52c.
//
// Solidity: function addAgingTime(uint256 time) returns()
func (_Cstoken *CstokenTransactorSession) AddAgingTime(time *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.AddAgingTime(&_Cstoken.TransactOpts, time)
}

// AddAgingTimesForPool is a paid mutator transaction binding the contract method 0xea6ca182.
//
// Solidity: function addAgingTimesForPool(address poolAddress, uint256 agingTime) returns()
func (_Cstoken *CstokenTransactor) AddAgingTimesForPool(opts *bind.TransactOpts, poolAddress common.Address, agingTime *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "addAgingTimesForPool", poolAddress, agingTime)
}

// AddAgingTimesForPool is a paid mutator transaction binding the contract method 0xea6ca182.
//
// Solidity: function addAgingTimesForPool(address poolAddress, uint256 agingTime) returns()
func (_Cstoken *CstokenSession) AddAgingTimesForPool(poolAddress common.Address, agingTime *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.AddAgingTimesForPool(&_Cstoken.TransactOpts, poolAddress, agingTime)
}

// AddAgingTimesForPool is a paid mutator transaction binding the contract method 0xea6ca182.
//
// Solidity: function addAgingTimesForPool(address poolAddress, uint256 agingTime) returns()
func (_Cstoken *CstokenTransactorSession) AddAgingTimesForPool(poolAddress common.Address, agingTime *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.AddAgingTimesForPool(&_Cstoken.TransactOpts, poolAddress, agingTime)
}

// AllAgingTimesAdded is a paid mutator transaction binding the contract method 0x037ca6c4.
//
// Solidity: function allAgingTimesAdded() returns()
func (_Cstoken *CstokenTransactor) AllAgingTimesAdded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "allAgingTimesAdded")
}

// AllAgingTimesAdded is a paid mutator transaction binding the contract method 0x037ca6c4.
//
// Solidity: function allAgingTimesAdded() returns()
func (_Cstoken *CstokenSession) AllAgingTimesAdded() (*types.Transaction, error) {
	return _Cstoken.Contract.AllAgingTimesAdded(&_Cstoken.TransactOpts)
}

// AllAgingTimesAdded is a paid mutator transaction binding the contract method 0x037ca6c4.
//
// Solidity: function allAgingTimesAdded() returns()
func (_Cstoken *CstokenTransactorSession) AllAgingTimesAdded() (*types.Transaction, error) {
	return _Cstoken.Contract.AllAgingTimesAdded(&_Cstoken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Cstoken *CstokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Approve(&_Cstoken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Approve(&_Cstoken.TransactOpts, _spender, _value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool success)
func (_Cstoken *CstokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool success)
func (_Cstoken *CstokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Cstoken.Contract.ApproveAndCall(&_Cstoken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address _spender, uint256 _value, bytes _extraData) returns(bool success)
func (_Cstoken *CstokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _Cstoken.Contract.ApproveAndCall(&_Cstoken.TransactOpts, _spender, _value, _extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Cstoken *CstokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Burn(&_Cstoken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Burn(&_Cstoken.TransactOpts, _value)
}

// CalculateDividends is a paid mutator transaction binding the contract method 0x2cf86006.
//
// Solidity: function calculateDividends(uint256 limit) returns()
func (_Cstoken *CstokenTransactor) CalculateDividends(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "calculateDividends", limit)
}

// CalculateDividends is a paid mutator transaction binding the contract method 0x2cf86006.
//
// Solidity: function calculateDividends(uint256 limit) returns()
func (_Cstoken *CstokenSession) CalculateDividends(limit *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.CalculateDividends(&_Cstoken.TransactOpts, limit)
}

// CalculateDividends is a paid mutator transaction binding the contract method 0x2cf86006.
//
// Solidity: function calculateDividends(uint256 limit) returns()
func (_Cstoken *CstokenTransactorSession) CalculateDividends(limit *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.CalculateDividends(&_Cstoken.TransactOpts, limit)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Cstoken *CstokenTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Cstoken *CstokenSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Cstoken.Contract.ChangeOwner(&_Cstoken.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Cstoken *CstokenTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Cstoken.Contract.ChangeOwner(&_Cstoken.TransactOpts, newOwner)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(address _from, uint256 _amount) returns()
func (_Cstoken *CstokenTransactor) Destroy(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "destroy", _from, _amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(address _from, uint256 _amount) returns()
func (_Cstoken *CstokenSession) Destroy(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Destroy(&_Cstoken.TransactOpts, _from, _amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(address _from, uint256 _amount) returns()
func (_Cstoken *CstokenTransactorSession) Destroy(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Destroy(&_Cstoken.TransactOpts, _from, _amount)
}

// DisableManuallyBurnTokens is a paid mutator transaction binding the contract method 0x71766ae3.
//
// Solidity: function disableManuallyBurnTokens(bool _disable) returns()
func (_Cstoken *CstokenTransactor) DisableManuallyBurnTokens(opts *bind.TransactOpts, _disable bool) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "disableManuallyBurnTokens", _disable)
}

// DisableManuallyBurnTokens is a paid mutator transaction binding the contract method 0x71766ae3.
//
// Solidity: function disableManuallyBurnTokens(bool _disable) returns()
func (_Cstoken *CstokenSession) DisableManuallyBurnTokens(_disable bool) (*types.Transaction, error) {
	return _Cstoken.Contract.DisableManuallyBurnTokens(&_Cstoken.TransactOpts, _disable)
}

// DisableManuallyBurnTokens is a paid mutator transaction binding the contract method 0x71766ae3.
//
// Solidity: function disableManuallyBurnTokens(bool _disable) returns()
func (_Cstoken *CstokenTransactorSession) DisableManuallyBurnTokens(_disable bool) (*types.Transaction, error) {
	return _Cstoken.Contract.DisableManuallyBurnTokens(&_Cstoken.TransactOpts, _disable)
}

// DisableTransfers is a paid mutator transaction binding the contract method 0x1608f18f.
//
// Solidity: function disableTransfers(bool _disable) returns()
func (_Cstoken *CstokenTransactor) DisableTransfers(opts *bind.TransactOpts, _disable bool) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "disableTransfers", _disable)
}

// DisableTransfers is a paid mutator transaction binding the contract method 0x1608f18f.
//
// Solidity: function disableTransfers(bool _disable) returns()
func (_Cstoken *CstokenSession) DisableTransfers(_disable bool) (*types.Transaction, error) {
	return _Cstoken.Contract.DisableTransfers(&_Cstoken.TransactOpts, _disable)
}

// DisableTransfers is a paid mutator transaction binding the contract method 0x1608f18f.
//
// Solidity: function disableTransfers(bool _disable) returns()
func (_Cstoken *CstokenTransactorSession) DisableTransfers(_disable bool) (*types.Transaction, error) {
	return _Cstoken.Contract.DisableTransfers(&_Cstoken.TransactOpts, _disable)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address _to, uint256 _amount) returns()
func (_Cstoken *CstokenTransactor) Issue(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "issue", _to, _amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address _to, uint256 _amount) returns()
func (_Cstoken *CstokenSession) Issue(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Issue(&_Cstoken.TransactOpts, _to, _amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address _to, uint256 _amount) returns()
func (_Cstoken *CstokenTransactorSession) Issue(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Issue(&_Cstoken.TransactOpts, _to, _amount)
}

// MintToken is a paid mutator transaction binding the contract method 0x23a36d2b.
//
// Solidity: function mintToken(address target, uint256 mintedAmount, uint256 agingTime) returns()
func (_Cstoken *CstokenTransactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int, agingTime *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "mintToken", target, mintedAmount, agingTime)
}

// MintToken is a paid mutator transaction binding the contract method 0x23a36d2b.
//
// Solidity: function mintToken(address target, uint256 mintedAmount, uint256 agingTime) returns()
func (_Cstoken *CstokenSession) MintToken(target common.Address, mintedAmount *big.Int, agingTime *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.MintToken(&_Cstoken.TransactOpts, target, mintedAmount, agingTime)
}

// MintToken is a paid mutator transaction binding the contract method 0x23a36d2b.
//
// Solidity: function mintToken(address target, uint256 mintedAmount, uint256 agingTime) returns()
func (_Cstoken *CstokenTransactorSession) MintToken(target common.Address, mintedAmount *big.Int, agingTime *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.MintToken(&_Cstoken.TransactOpts, target, mintedAmount, agingTime)
}

// ReceiveDividends is a paid mutator transaction binding the contract method 0x79fc4687.
//
// Solidity: function receiveDividends() returns()
func (_Cstoken *CstokenTransactor) ReceiveDividends(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "receiveDividends")
}

// ReceiveDividends is a paid mutator transaction binding the contract method 0x79fc4687.
//
// Solidity: function receiveDividends() returns()
func (_Cstoken *CstokenSession) ReceiveDividends() (*types.Transaction, error) {
	return _Cstoken.Contract.ReceiveDividends(&_Cstoken.TransactOpts)
}

// ReceiveDividends is a paid mutator transaction binding the contract method 0x79fc4687.
//
// Solidity: function receiveDividends() returns()
func (_Cstoken *CstokenTransactorSession) ReceiveDividends() (*types.Transaction, error) {
	return _Cstoken.Contract.ReceiveDividends(&_Cstoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Cstoken *CstokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Transfer(&_Cstoken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.Transfer(&_Cstoken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Cstoken *CstokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.TransferFrom(&_Cstoken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Cstoken *CstokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Cstoken.Contract.TransferFrom(&_Cstoken.TransactOpts, _from, _to, _value)
}

// CstokenAgingTransferIterator is returned from FilterAgingTransfer and is used to iterate over the raw logs and unpacked data for AgingTransfer events raised by the Cstoken contract.
type CstokenAgingTransferIterator struct {
	Event *CstokenAgingTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenAgingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenAgingTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenAgingTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenAgingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenAgingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenAgingTransfer represents a AgingTransfer event raised by the Cstoken contract.
type CstokenAgingTransfer struct {
	From      common.Address
	To        common.Address
	Value     *big.Int
	AgingTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgingTransfer is a free log retrieval operation binding the contract event 0x46a1749a7648b704d1ad3fe33741b13174a4b1641db362e808d00eab7250d106.
//
// Solidity: event AgingTransfer(address indexed from, address indexed to, uint256 value, uint256 agingTime)
func (_Cstoken *CstokenFilterer) FilterAgingTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CstokenAgingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "AgingTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CstokenAgingTransferIterator{contract: _Cstoken.contract, event: "AgingTransfer", logs: logs, sub: sub}, nil
}

// WatchAgingTransfer is a free log subscription operation binding the contract event 0x46a1749a7648b704d1ad3fe33741b13174a4b1641db362e808d00eab7250d106.
//
// Solidity: event AgingTransfer(address indexed from, address indexed to, uint256 value, uint256 agingTime)
func (_Cstoken *CstokenFilterer) WatchAgingTransfer(opts *bind.WatchOpts, sink chan<- *CstokenAgingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "AgingTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenAgingTransfer)
				if err := _Cstoken.contract.UnpackLog(event, "AgingTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgingTransfer is a log parse operation binding the contract event 0x46a1749a7648b704d1ad3fe33741b13174a4b1641db362e808d00eab7250d106.
//
// Solidity: event AgingTransfer(address indexed from, address indexed to, uint256 value, uint256 agingTime)
func (_Cstoken *CstokenFilterer) ParseAgingTransfer(log types.Log) (*CstokenAgingTransfer, error) {
	event := new(CstokenAgingTransfer)
	if err := _Cstoken.contract.UnpackLog(event, "AgingTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CstokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cstoken contract.
type CstokenApprovalIterator struct {
	Event *CstokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenApproval represents a Approval event raised by the Cstoken contract.
type CstokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Cstoken *CstokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CstokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CstokenApprovalIterator{contract: _Cstoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Cstoken *CstokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CstokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenApproval)
				if err := _Cstoken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Cstoken *CstokenFilterer) ParseApproval(log types.Log) (*CstokenApproval, error) {
	event := new(CstokenApproval)
	if err := _Cstoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CstokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Cstoken contract.
type CstokenBurnIterator struct {
	Event *CstokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenBurn represents a Burn event raised by the Cstoken contract.
type CstokenBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Cstoken *CstokenFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*CstokenBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &CstokenBurnIterator{contract: _Cstoken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Cstoken *CstokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *CstokenBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenBurn)
				if err := _Cstoken.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_Cstoken *CstokenFilterer) ParseBurn(log types.Log) (*CstokenBurn, error) {
	event := new(CstokenBurn)
	if err := _Cstoken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CstokenDestructionIterator is returned from FilterDestruction and is used to iterate over the raw logs and unpacked data for Destruction events raised by the Cstoken contract.
type CstokenDestructionIterator struct {
	Event *CstokenDestruction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenDestructionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenDestruction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenDestruction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenDestructionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenDestructionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenDestruction represents a Destruction event raised by the Cstoken contract.
type CstokenDestruction struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDestruction is a free log retrieval operation binding the contract event 0x9a1b418bc061a5d80270261562e6986a35d995f8051145f277be16103abd3453.
//
// Solidity: event Destruction(uint256 _amount)
func (_Cstoken *CstokenFilterer) FilterDestruction(opts *bind.FilterOpts) (*CstokenDestructionIterator, error) {

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "Destruction")
	if err != nil {
		return nil, err
	}
	return &CstokenDestructionIterator{contract: _Cstoken.contract, event: "Destruction", logs: logs, sub: sub}, nil
}

// WatchDestruction is a free log subscription operation binding the contract event 0x9a1b418bc061a5d80270261562e6986a35d995f8051145f277be16103abd3453.
//
// Solidity: event Destruction(uint256 _amount)
func (_Cstoken *CstokenFilterer) WatchDestruction(opts *bind.WatchOpts, sink chan<- *CstokenDestruction) (event.Subscription, error) {

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "Destruction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenDestruction)
				if err := _Cstoken.contract.UnpackLog(event, "Destruction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDestruction is a log parse operation binding the contract event 0x9a1b418bc061a5d80270261562e6986a35d995f8051145f277be16103abd3453.
//
// Solidity: event Destruction(uint256 _amount)
func (_Cstoken *CstokenFilterer) ParseDestruction(log types.Log) (*CstokenDestruction, error) {
	event := new(CstokenDestruction)
	if err := _Cstoken.contract.UnpackLog(event, "Destruction", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CstokenIssuanceIterator is returned from FilterIssuance and is used to iterate over the raw logs and unpacked data for Issuance events raised by the Cstoken contract.
type CstokenIssuanceIterator struct {
	Event *CstokenIssuance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenIssuanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenIssuance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenIssuance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenIssuanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenIssuanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenIssuance represents a Issuance event raised by the Cstoken contract.
type CstokenIssuance struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuance is a free log retrieval operation binding the contract event 0x9386c90217c323f58030f9dadcbc938f807a940f4ff41cd4cead9562f5da7dc3.
//
// Solidity: event Issuance(uint256 _amount)
func (_Cstoken *CstokenFilterer) FilterIssuance(opts *bind.FilterOpts) (*CstokenIssuanceIterator, error) {

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "Issuance")
	if err != nil {
		return nil, err
	}
	return &CstokenIssuanceIterator{contract: _Cstoken.contract, event: "Issuance", logs: logs, sub: sub}, nil
}

// WatchIssuance is a free log subscription operation binding the contract event 0x9386c90217c323f58030f9dadcbc938f807a940f4ff41cd4cead9562f5da7dc3.
//
// Solidity: event Issuance(uint256 _amount)
func (_Cstoken *CstokenFilterer) WatchIssuance(opts *bind.WatchOpts, sink chan<- *CstokenIssuance) (event.Subscription, error) {

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "Issuance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenIssuance)
				if err := _Cstoken.contract.UnpackLog(event, "Issuance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssuance is a log parse operation binding the contract event 0x9386c90217c323f58030f9dadcbc938f807a940f4ff41cd4cead9562f5da7dc3.
//
// Solidity: event Issuance(uint256 _amount)
func (_Cstoken *CstokenFilterer) ParseIssuance(log types.Log) (*CstokenIssuance, error) {
	event := new(CstokenIssuance)
	if err := _Cstoken.contract.UnpackLog(event, "Issuance", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CstokenNewSmartTokenIterator is returned from FilterNewSmartToken and is used to iterate over the raw logs and unpacked data for NewSmartToken events raised by the Cstoken contract.
type CstokenNewSmartTokenIterator struct {
	Event *CstokenNewSmartToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenNewSmartTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenNewSmartToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenNewSmartToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenNewSmartTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenNewSmartTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenNewSmartToken represents a NewSmartToken event raised by the Cstoken contract.
type CstokenNewSmartToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewSmartToken is a free log retrieval operation binding the contract event 0xf4cd1f8571e8d9c97ffcb81558807ab73f9803d54de5da6a0420593c82a4a9f0.
//
// Solidity: event NewSmartToken(address _token)
func (_Cstoken *CstokenFilterer) FilterNewSmartToken(opts *bind.FilterOpts) (*CstokenNewSmartTokenIterator, error) {

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "NewSmartToken")
	if err != nil {
		return nil, err
	}
	return &CstokenNewSmartTokenIterator{contract: _Cstoken.contract, event: "NewSmartToken", logs: logs, sub: sub}, nil
}

// WatchNewSmartToken is a free log subscription operation binding the contract event 0xf4cd1f8571e8d9c97ffcb81558807ab73f9803d54de5da6a0420593c82a4a9f0.
//
// Solidity: event NewSmartToken(address _token)
func (_Cstoken *CstokenFilterer) WatchNewSmartToken(opts *bind.WatchOpts, sink chan<- *CstokenNewSmartToken) (event.Subscription, error) {

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "NewSmartToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenNewSmartToken)
				if err := _Cstoken.contract.UnpackLog(event, "NewSmartToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSmartToken is a log parse operation binding the contract event 0xf4cd1f8571e8d9c97ffcb81558807ab73f9803d54de5da6a0420593c82a4a9f0.
//
// Solidity: event NewSmartToken(address _token)
func (_Cstoken *CstokenFilterer) ParseNewSmartToken(log types.Log) (*CstokenNewSmartToken, error) {
	event := new(CstokenNewSmartToken)
	if err := _Cstoken.contract.UnpackLog(event, "NewSmartToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CstokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cstoken contract.
type CstokenTransferIterator struct {
	Event *CstokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CstokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CstokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CstokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CstokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CstokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CstokenTransfer represents a Transfer event raised by the Cstoken contract.
type CstokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cstoken *CstokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CstokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cstoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CstokenTransferIterator{contract: _Cstoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cstoken *CstokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CstokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cstoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CstokenTransfer)
				if err := _Cstoken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cstoken *CstokenFilterer) ParseTransfer(log types.Log) (*CstokenTransfer, error) {
	event := new(CstokenTransfer)
	if err := _Cstoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

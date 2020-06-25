// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IpfsMarketABI is the input ABI used to generate the binding from.
const IpfsMarketABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"accessKey\",\"type\":\"bytes\"}],\"name\":\"answerPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getAccessKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"hasPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"accessKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IpfsMarket is an auto generated Go binding around an Ethereum contract.
type IpfsMarket struct {
	IpfsMarketCaller     // Read-only binding to the contract
	IpfsMarketTransactor // Write-only binding to the contract
	IpfsMarketFilterer   // Log filterer for contract events
}

// IpfsMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpfsMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpfsMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpfsMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpfsMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpfsMarketSession struct {
	Contract     *IpfsMarket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpfsMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpfsMarketCallerSession struct {
	Contract *IpfsMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IpfsMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpfsMarketTransactorSession struct {
	Contract     *IpfsMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IpfsMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpfsMarketRaw struct {
	Contract *IpfsMarket // Generic contract binding to access the raw methods on
}

// IpfsMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpfsMarketCallerRaw struct {
	Contract *IpfsMarketCaller // Generic read-only contract binding to access the raw methods on
}

// IpfsMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpfsMarketTransactorRaw struct {
	Contract *IpfsMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpfsMarket creates a new instance of IpfsMarket, bound to a specific deployed contract.
func NewIpfsMarket(address common.Address, backend bind.ContractBackend) (*IpfsMarket, error) {
	contract, err := bindIpfsMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IpfsMarket{IpfsMarketCaller: IpfsMarketCaller{contract: contract}, IpfsMarketTransactor: IpfsMarketTransactor{contract: contract}, IpfsMarketFilterer: IpfsMarketFilterer{contract: contract}}, nil
}

// NewIpfsMarketCaller creates a new read-only instance of IpfsMarket, bound to a specific deployed contract.
func NewIpfsMarketCaller(address common.Address, caller bind.ContractCaller) (*IpfsMarketCaller, error) {
	contract, err := bindIpfsMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpfsMarketCaller{contract: contract}, nil
}

// NewIpfsMarketTransactor creates a new write-only instance of IpfsMarket, bound to a specific deployed contract.
func NewIpfsMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*IpfsMarketTransactor, error) {
	contract, err := bindIpfsMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpfsMarketTransactor{contract: contract}, nil
}

// NewIpfsMarketFilterer creates a new log filterer instance of IpfsMarket, bound to a specific deployed contract.
func NewIpfsMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*IpfsMarketFilterer, error) {
	contract, err := bindIpfsMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpfsMarketFilterer{contract: contract}, nil
}

// bindIpfsMarket binds a generic wrapper to an already deployed contract.
func bindIpfsMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IpfsMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IpfsMarket *IpfsMarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IpfsMarket.Contract.IpfsMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IpfsMarket *IpfsMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IpfsMarket.Contract.IpfsMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IpfsMarket *IpfsMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IpfsMarket.Contract.IpfsMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IpfsMarket *IpfsMarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IpfsMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IpfsMarket *IpfsMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IpfsMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IpfsMarket *IpfsMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IpfsMarket.Contract.contract.Transact(opts, method, params...)
}

// GetAccessKey is a free data retrieval call binding the contract method 0x998ae61e.
//
// Solidity: function getAccessKey(string cid) constant returns(bytes)
func (_IpfsMarket *IpfsMarketCaller) GetAccessKey(opts *bind.CallOpts, cid string) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IpfsMarket.contract.Call(opts, out, "getAccessKey", cid)
	return *ret0, err
}

// GetAccessKey is a free data retrieval call binding the contract method 0x998ae61e.
//
// Solidity: function getAccessKey(string cid) constant returns(bytes)
func (_IpfsMarket *IpfsMarketSession) GetAccessKey(cid string) ([]byte, error) {
	return _IpfsMarket.Contract.GetAccessKey(&_IpfsMarket.CallOpts, cid)
}

// GetAccessKey is a free data retrieval call binding the contract method 0x998ae61e.
//
// Solidity: function getAccessKey(string cid) constant returns(bytes)
func (_IpfsMarket *IpfsMarketCallerSession) GetAccessKey(cid string) ([]byte, error) {
	return _IpfsMarket.Contract.GetAccessKey(&_IpfsMarket.CallOpts, cid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string cid) constant returns(uint256)
func (_IpfsMarket *IpfsMarketCaller) GetPrice(opts *bind.CallOpts, cid string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IpfsMarket.contract.Call(opts, out, "getPrice", cid)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string cid) constant returns(uint256)
func (_IpfsMarket *IpfsMarketSession) GetPrice(cid string) (*big.Int, error) {
	return _IpfsMarket.Contract.GetPrice(&_IpfsMarket.CallOpts, cid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string cid) constant returns(uint256)
func (_IpfsMarket *IpfsMarketCallerSession) GetPrice(cid string) (*big.Int, error) {
	return _IpfsMarket.Contract.GetPrice(&_IpfsMarket.CallOpts, cid)
}

// HasPurchased is a free data retrieval call binding the contract method 0x0512b8ab.
//
// Solidity: function hasPurchased(string cid) constant returns(bool)
func (_IpfsMarket *IpfsMarketCaller) HasPurchased(opts *bind.CallOpts, cid string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IpfsMarket.contract.Call(opts, out, "hasPurchased", cid)
	return *ret0, err
}

// HasPurchased is a free data retrieval call binding the contract method 0x0512b8ab.
//
// Solidity: function hasPurchased(string cid) constant returns(bool)
func (_IpfsMarket *IpfsMarketSession) HasPurchased(cid string) (bool, error) {
	return _IpfsMarket.Contract.HasPurchased(&_IpfsMarket.CallOpts, cid)
}

// HasPurchased is a free data retrieval call binding the contract method 0x0512b8ab.
//
// Solidity: function hasPurchased(string cid) constant returns(bool)
func (_IpfsMarket *IpfsMarketCallerSession) HasPurchased(cid string) (bool, error) {
	return _IpfsMarket.Contract.HasPurchased(&_IpfsMarket.CallOpts, cid)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) constant returns(uint256)
func (_IpfsMarket *IpfsMarketCaller) Payments(opts *bind.CallOpts, dest common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IpfsMarket.contract.Call(opts, out, "payments", dest)
	return *ret0, err
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) constant returns(uint256)
func (_IpfsMarket *IpfsMarketSession) Payments(dest common.Address) (*big.Int, error) {
	return _IpfsMarket.Contract.Payments(&_IpfsMarket.CallOpts, dest)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) constant returns(uint256)
func (_IpfsMarket *IpfsMarketCallerSession) Payments(dest common.Address) (*big.Int, error) {
	return _IpfsMarket.Contract.Payments(&_IpfsMarket.CallOpts, dest)
}

// AnswerPurchase is a paid mutator transaction binding the contract method 0xaf1cd463.
//
// Solidity: function answerPurchase(string cid, address purchaser, bytes accessKey) returns()
func (_IpfsMarket *IpfsMarketTransactor) AnswerPurchase(opts *bind.TransactOpts, cid string, purchaser common.Address, accessKey []byte) (*types.Transaction, error) {
	return _IpfsMarket.contract.Transact(opts, "answerPurchase", cid, purchaser, accessKey)
}

// AnswerPurchase is a paid mutator transaction binding the contract method 0xaf1cd463.
//
// Solidity: function answerPurchase(string cid, address purchaser, bytes accessKey) returns()
func (_IpfsMarket *IpfsMarketSession) AnswerPurchase(cid string, purchaser common.Address, accessKey []byte) (*types.Transaction, error) {
	return _IpfsMarket.Contract.AnswerPurchase(&_IpfsMarket.TransactOpts, cid, purchaser, accessKey)
}

// AnswerPurchase is a paid mutator transaction binding the contract method 0xaf1cd463.
//
// Solidity: function answerPurchase(string cid, address purchaser, bytes accessKey) returns()
func (_IpfsMarket *IpfsMarketTransactorSession) AnswerPurchase(cid string, purchaser common.Address, accessKey []byte) (*types.Transaction, error) {
	return _IpfsMarket.Contract.AnswerPurchase(&_IpfsMarket.TransactOpts, cid, purchaser, accessKey)
}

// Publish is a paid mutator transaction binding the contract method 0x449c72a9.
//
// Solidity: function publish(string cid, bytes accessKey, uint256 price) returns()
func (_IpfsMarket *IpfsMarketTransactor) Publish(opts *bind.TransactOpts, cid string, accessKey []byte, price *big.Int) (*types.Transaction, error) {
	return _IpfsMarket.contract.Transact(opts, "publish", cid, accessKey, price)
}

// Publish is a paid mutator transaction binding the contract method 0x449c72a9.
//
// Solidity: function publish(string cid, bytes accessKey, uint256 price) returns()
func (_IpfsMarket *IpfsMarketSession) Publish(cid string, accessKey []byte, price *big.Int) (*types.Transaction, error) {
	return _IpfsMarket.Contract.Publish(&_IpfsMarket.TransactOpts, cid, accessKey, price)
}

// Publish is a paid mutator transaction binding the contract method 0x449c72a9.
//
// Solidity: function publish(string cid, bytes accessKey, uint256 price) returns()
func (_IpfsMarket *IpfsMarketTransactorSession) Publish(cid string, accessKey []byte, price *big.Int) (*types.Transaction, error) {
	return _IpfsMarket.Contract.Publish(&_IpfsMarket.TransactOpts, cid, accessKey, price)
}

// Purchase is a paid mutator transaction binding the contract method 0xf5dd4e9c.
//
// Solidity: function purchase(string cid, bytes publicKey) returns()
func (_IpfsMarket *IpfsMarketTransactor) Purchase(opts *bind.TransactOpts, cid string, publicKey []byte) (*types.Transaction, error) {
	return _IpfsMarket.contract.Transact(opts, "purchase", cid, publicKey)
}

// Purchase is a paid mutator transaction binding the contract method 0xf5dd4e9c.
//
// Solidity: function purchase(string cid, bytes publicKey) returns()
func (_IpfsMarket *IpfsMarketSession) Purchase(cid string, publicKey []byte) (*types.Transaction, error) {
	return _IpfsMarket.Contract.Purchase(&_IpfsMarket.TransactOpts, cid, publicKey)
}

// Purchase is a paid mutator transaction binding the contract method 0xf5dd4e9c.
//
// Solidity: function purchase(string cid, bytes publicKey) returns()
func (_IpfsMarket *IpfsMarketTransactorSession) Purchase(cid string, publicKey []byte) (*types.Transaction, error) {
	return _IpfsMarket.Contract.Purchase(&_IpfsMarket.TransactOpts, cid, publicKey)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_IpfsMarket *IpfsMarketTransactor) WithdrawPayments(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _IpfsMarket.contract.Transact(opts, "withdrawPayments", payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_IpfsMarket *IpfsMarketSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _IpfsMarket.Contract.WithdrawPayments(&_IpfsMarket.TransactOpts, payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_IpfsMarket *IpfsMarketTransactorSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _IpfsMarket.Contract.WithdrawPayments(&_IpfsMarket.TransactOpts, payee)
}

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
const IpfsMarketABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"accessKey\",\"type\":\"bytes\"}],\"name\":\"PurchaseAnswered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"PurchaseCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"accessKey\",\"type\":\"bytes\"}],\"name\":\"answerPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getAccessKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"hasPurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"isAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"accessKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// IsAuthor is a free data retrieval call binding the contract method 0xff34ee29.
//
// Solidity: function isAuthor(string cid) constant returns(bool)
func (_IpfsMarket *IpfsMarketCaller) IsAuthor(opts *bind.CallOpts, cid string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IpfsMarket.contract.Call(opts, out, "isAuthor", cid)
	return *ret0, err
}

// IsAuthor is a free data retrieval call binding the contract method 0xff34ee29.
//
// Solidity: function isAuthor(string cid) constant returns(bool)
func (_IpfsMarket *IpfsMarketSession) IsAuthor(cid string) (bool, error) {
	return _IpfsMarket.Contract.IsAuthor(&_IpfsMarket.CallOpts, cid)
}

// IsAuthor is a free data retrieval call binding the contract method 0xff34ee29.
//
// Solidity: function isAuthor(string cid) constant returns(bool)
func (_IpfsMarket *IpfsMarketCallerSession) IsAuthor(cid string) (bool, error) {
	return _IpfsMarket.Contract.IsAuthor(&_IpfsMarket.CallOpts, cid)
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

// IpfsMarketPurchaseAnsweredIterator is returned from FilterPurchaseAnswered and is used to iterate over the raw logs and unpacked data for PurchaseAnswered events raised by the IpfsMarket contract.
type IpfsMarketPurchaseAnsweredIterator struct {
	Event *IpfsMarketPurchaseAnswered // Event containing the contract specifics and raw log

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
func (it *IpfsMarketPurchaseAnsweredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsMarketPurchaseAnswered)
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
		it.Event = new(IpfsMarketPurchaseAnswered)
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
func (it *IpfsMarketPurchaseAnsweredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsMarketPurchaseAnsweredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsMarketPurchaseAnswered represents a PurchaseAnswered event raised by the IpfsMarket contract.
type IpfsMarketPurchaseAnswered struct {
	Cid       string
	Purchaser common.Address
	AccessKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseAnswered is a free log retrieval operation binding the contract event 0x3cebc18d001ab9b316270cb0f8f6dc833a4a3b15bf81d2c51dfbff434079e9fb.
//
// Solidity: event PurchaseAnswered(string cid, address purchaser, bytes accessKey)
func (_IpfsMarket *IpfsMarketFilterer) FilterPurchaseAnswered(opts *bind.FilterOpts) (*IpfsMarketPurchaseAnsweredIterator, error) {

	logs, sub, err := _IpfsMarket.contract.FilterLogs(opts, "PurchaseAnswered")
	if err != nil {
		return nil, err
	}
	return &IpfsMarketPurchaseAnsweredIterator{contract: _IpfsMarket.contract, event: "PurchaseAnswered", logs: logs, sub: sub}, nil
}

// WatchPurchaseAnswered is a free log subscription operation binding the contract event 0x3cebc18d001ab9b316270cb0f8f6dc833a4a3b15bf81d2c51dfbff434079e9fb.
//
// Solidity: event PurchaseAnswered(string cid, address purchaser, bytes accessKey)
func (_IpfsMarket *IpfsMarketFilterer) WatchPurchaseAnswered(opts *bind.WatchOpts, sink chan<- *IpfsMarketPurchaseAnswered) (event.Subscription, error) {

	logs, sub, err := _IpfsMarket.contract.WatchLogs(opts, "PurchaseAnswered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsMarketPurchaseAnswered)
				if err := _IpfsMarket.contract.UnpackLog(event, "PurchaseAnswered", log); err != nil {
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

// ParsePurchaseAnswered is a log parse operation binding the contract event 0x3cebc18d001ab9b316270cb0f8f6dc833a4a3b15bf81d2c51dfbff434079e9fb.
//
// Solidity: event PurchaseAnswered(string cid, address purchaser, bytes accessKey)
func (_IpfsMarket *IpfsMarketFilterer) ParsePurchaseAnswered(log types.Log) (*IpfsMarketPurchaseAnswered, error) {
	event := new(IpfsMarketPurchaseAnswered)
	if err := _IpfsMarket.contract.UnpackLog(event, "PurchaseAnswered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IpfsMarketPurchaseCreatedIterator is returned from FilterPurchaseCreated and is used to iterate over the raw logs and unpacked data for PurchaseCreated events raised by the IpfsMarket contract.
type IpfsMarketPurchaseCreatedIterator struct {
	Event *IpfsMarketPurchaseCreated // Event containing the contract specifics and raw log

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
func (it *IpfsMarketPurchaseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IpfsMarketPurchaseCreated)
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
		it.Event = new(IpfsMarketPurchaseCreated)
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
func (it *IpfsMarketPurchaseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IpfsMarketPurchaseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IpfsMarketPurchaseCreated represents a PurchaseCreated event raised by the IpfsMarket contract.
type IpfsMarketPurchaseCreated struct {
	Cid       string
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCreated is a free log retrieval operation binding the contract event 0xe07f0ef600900faec69a755c496d0f8386fe613b182166d90e9de30d143ef5d4.
//
// Solidity: event PurchaseCreated(string cid, bytes publicKey)
func (_IpfsMarket *IpfsMarketFilterer) FilterPurchaseCreated(opts *bind.FilterOpts) (*IpfsMarketPurchaseCreatedIterator, error) {

	logs, sub, err := _IpfsMarket.contract.FilterLogs(opts, "PurchaseCreated")
	if err != nil {
		return nil, err
	}
	return &IpfsMarketPurchaseCreatedIterator{contract: _IpfsMarket.contract, event: "PurchaseCreated", logs: logs, sub: sub}, nil
}

// WatchPurchaseCreated is a free log subscription operation binding the contract event 0xe07f0ef600900faec69a755c496d0f8386fe613b182166d90e9de30d143ef5d4.
//
// Solidity: event PurchaseCreated(string cid, bytes publicKey)
func (_IpfsMarket *IpfsMarketFilterer) WatchPurchaseCreated(opts *bind.WatchOpts, sink chan<- *IpfsMarketPurchaseCreated) (event.Subscription, error) {

	logs, sub, err := _IpfsMarket.contract.WatchLogs(opts, "PurchaseCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IpfsMarketPurchaseCreated)
				if err := _IpfsMarket.contract.UnpackLog(event, "PurchaseCreated", log); err != nil {
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

// ParsePurchaseCreated is a log parse operation binding the contract event 0xe07f0ef600900faec69a755c496d0f8386fe613b182166d90e9de30d143ef5d4.
//
// Solidity: event PurchaseCreated(string cid, bytes publicKey)
func (_IpfsMarket *IpfsMarketFilterer) ParsePurchaseCreated(log types.Log) (*IpfsMarketPurchaseCreated, error) {
	event := new(IpfsMarketPurchaseCreated)
	if err := _IpfsMarket.contract.UnpackLog(event, "PurchaseCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

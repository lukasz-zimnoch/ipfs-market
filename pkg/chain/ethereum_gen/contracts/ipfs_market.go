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
const IpfsMarketABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"accessKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IpfsMarketFuncSigs maps the 4-byte function signature to its string representation.
var IpfsMarketFuncSigs = map[string]string{
	"449c72a9": "publish(string,bytes,uint256)",
}

// IpfsMarketBin is the compiled bytecode used for deploying new contracts.
var IpfsMarketBin = "0x608060405234801561001057600080fd5b506103da806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063449c72a914610030575b600080fd5b61015f6004803603606081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156100e857600080fd5b8201836020820111156100fa57600080fd5b8035906020019184600183028401116401000000008311171561011c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610161915050565b005b60006001600160a01b03166000846040518082805190602001908083835b6020831061019e5780518252601f19909201916020918201910161017f565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b031692909214915061021690505760405162461bcd60e51b815260040180806020018281038252602981526020018061037d6029913960400191505060405180910390fd5b6040518060600160405280336001600160a01b03168152602001838152602001828152506000846040518082805190602001908083835b6020831061026c5780518252601f19909201916020918201910161024d565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845181546001600160a01b0319166001600160a01b0390911617815584840151805191946102ce945060018601935001906102e1565b5060408201518160020155905050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061032257805160ff191683800117855561034f565b8280016001018555821561034f579182015b8281111561034f578251825591602001919060010190610334565b5061035b92915061035f565b5090565b61037991905b8082111561035b5760008155600101610365565b9056fe5075626c69636174696f6e207769746820676976656e2043494420616c726561647920657869737473a265627a7a723158204fc9ab343521a20476376cd13a61e8cb68dd2e85017005fdb4cdc7401f5bf00964736f6c63430005110032"

// DeployIpfsMarket deploys a new Ethereum contract, binding an instance of IpfsMarket to it.
func DeployIpfsMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IpfsMarket, error) {
	parsed, err := abi.JSON(strings.NewReader(IpfsMarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IpfsMarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IpfsMarket{IpfsMarketCaller: IpfsMarketCaller{contract: contract}, IpfsMarketTransactor: IpfsMarketTransactor{contract: contract}, IpfsMarketFilterer: IpfsMarketFilterer{contract: contract}}, nil
}

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

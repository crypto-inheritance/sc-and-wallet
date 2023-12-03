// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Inheritance

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// InheritanceMetaData contains all meta data concerning the Inheritance contract.
var InheritanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApprovalReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AssetsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inheritanceRequestTime\",\"type\":\"uint256\"}],\"name\":\"InheritanceRequestApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inheritanceRequestTime\",\"type\":\"uint256\"}],\"name\":\"InheritanceRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inheritanceRequestTime\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"InheritanceRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"SignatoryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"SignatoryApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"SignatoryRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"TransactionApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransactionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"name\":\"WeightsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signatory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"addSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signatory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"approveSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"approveTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approvedSignatoryCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelInheritanceRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkInheritanceRequestStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmInheritanceRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getApprovedTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastPendingTransactionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignatoryAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signatory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getWeightForAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signatory\",\"type\":\"address\"}],\"name\":\"getWithdrawLimitForAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governmentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inactivityThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inheritanceDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inheritanceDeadlineActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inheritanceRequestActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inheritanceRequestTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWalletAddressSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signatory\",\"type\":\"address\"}],\"name\":\"removeSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestInheritance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setApprovalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGovernmentAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setInactivityThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"setInheritanceDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setLastTransactionTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_walletAddress\",\"type\":\"address\"}],\"name\":\"setWalletAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signatory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"setWeightForSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signatories\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signatoryAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signatoryIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signatoryWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InheritanceABI is the input ABI used to generate the binding from.
// Deprecated: Use InheritanceMetaData.ABI instead.
var InheritanceABI = InheritanceMetaData.ABI

// Inheritance is an auto generated Go binding around an Ethereum contract.
type Inheritance struct {
	InheritanceCaller     // Read-only binding to the contract
	InheritanceTransactor // Write-only binding to the contract
	InheritanceFilterer   // Log filterer for contract events
}

// InheritanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type InheritanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InheritanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InheritanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InheritanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InheritanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InheritanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InheritanceSession struct {
	Contract     *Inheritance      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InheritanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InheritanceCallerSession struct {
	Contract *InheritanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InheritanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InheritanceTransactorSession struct {
	Contract     *InheritanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InheritanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type InheritanceRaw struct {
	Contract *Inheritance // Generic contract binding to access the raw methods on
}

// InheritanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InheritanceCallerRaw struct {
	Contract *InheritanceCaller // Generic read-only contract binding to access the raw methods on
}

// InheritanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InheritanceTransactorRaw struct {
	Contract *InheritanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInheritance creates a new instance of Inheritance, bound to a specific deployed contract.
func NewInheritance(address common.Address, backend bind.ContractBackend) (*Inheritance, error) {
	contract, err := bindInheritance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inheritance{InheritanceCaller: InheritanceCaller{contract: contract}, InheritanceTransactor: InheritanceTransactor{contract: contract}, InheritanceFilterer: InheritanceFilterer{contract: contract}}, nil
}

// NewInheritanceCaller creates a new read-only instance of Inheritance, bound to a specific deployed contract.
func NewInheritanceCaller(address common.Address, caller bind.ContractCaller) (*InheritanceCaller, error) {
	contract, err := bindInheritance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InheritanceCaller{contract: contract}, nil
}

// NewInheritanceTransactor creates a new write-only instance of Inheritance, bound to a specific deployed contract.
func NewInheritanceTransactor(address common.Address, transactor bind.ContractTransactor) (*InheritanceTransactor, error) {
	contract, err := bindInheritance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InheritanceTransactor{contract: contract}, nil
}

// NewInheritanceFilterer creates a new log filterer instance of Inheritance, bound to a specific deployed contract.
func NewInheritanceFilterer(address common.Address, filterer bind.ContractFilterer) (*InheritanceFilterer, error) {
	contract, err := bindInheritance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InheritanceFilterer{contract: contract}, nil
}

// bindInheritance binds a generic wrapper to an already deployed contract.
func bindInheritance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InheritanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inheritance *InheritanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inheritance.Contract.InheritanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inheritance *InheritanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.Contract.InheritanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inheritance *InheritanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inheritance.Contract.InheritanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inheritance *InheritanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inheritance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inheritance *InheritanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inheritance *InheritanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inheritance.Contract.contract.Transact(opts, method, params...)
}

// ApprovalThreshold is a free data retrieval call binding the contract method 0x7d0eef61.
//
// Solidity: function approvalThreshold() view returns(uint256)
func (_Inheritance *InheritanceCaller) ApprovalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "approvalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovalThreshold is a free data retrieval call binding the contract method 0x7d0eef61.
//
// Solidity: function approvalThreshold() view returns(uint256)
func (_Inheritance *InheritanceSession) ApprovalThreshold() (*big.Int, error) {
	return _Inheritance.Contract.ApprovalThreshold(&_Inheritance.CallOpts)
}

// ApprovalThreshold is a free data retrieval call binding the contract method 0x7d0eef61.
//
// Solidity: function approvalThreshold() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) ApprovalThreshold() (*big.Int, error) {
	return _Inheritance.Contract.ApprovalThreshold(&_Inheritance.CallOpts)
}

// ApprovedSignatoryCount is a free data retrieval call binding the contract method 0x14a5caea.
//
// Solidity: function approvedSignatoryCount() view returns(int256)
func (_Inheritance *InheritanceCaller) ApprovedSignatoryCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "approvedSignatoryCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedSignatoryCount is a free data retrieval call binding the contract method 0x14a5caea.
//
// Solidity: function approvedSignatoryCount() view returns(int256)
func (_Inheritance *InheritanceSession) ApprovedSignatoryCount() (*big.Int, error) {
	return _Inheritance.Contract.ApprovedSignatoryCount(&_Inheritance.CallOpts)
}

// ApprovedSignatoryCount is a free data retrieval call binding the contract method 0x14a5caea.
//
// Solidity: function approvedSignatoryCount() view returns(int256)
func (_Inheritance *InheritanceCallerSession) ApprovedSignatoryCount() (*big.Int, error) {
	return _Inheritance.Contract.ApprovedSignatoryCount(&_Inheritance.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Inheritance *InheritanceCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Inheritance *InheritanceSession) Balance() (*big.Int, error) {
	return _Inheritance.Contract.Balance(&_Inheritance.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) Balance() (*big.Int, error) {
	return _Inheritance.Contract.Balance(&_Inheritance.CallOpts)
}

// GetApprovedTransactionCount is a free data retrieval call binding the contract method 0xfbe96385.
//
// Solidity: function getApprovedTransactionCount() view returns(uint256)
func (_Inheritance *InheritanceCaller) GetApprovedTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "getApprovedTransactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetApprovedTransactionCount is a free data retrieval call binding the contract method 0xfbe96385.
//
// Solidity: function getApprovedTransactionCount() view returns(uint256)
func (_Inheritance *InheritanceSession) GetApprovedTransactionCount() (*big.Int, error) {
	return _Inheritance.Contract.GetApprovedTransactionCount(&_Inheritance.CallOpts)
}

// GetApprovedTransactionCount is a free data retrieval call binding the contract method 0xfbe96385.
//
// Solidity: function getApprovedTransactionCount() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) GetApprovedTransactionCount() (*big.Int, error) {
	return _Inheritance.Contract.GetApprovedTransactionCount(&_Inheritance.CallOpts)
}

// GetLastPendingTransactionId is a free data retrieval call binding the contract method 0x62d5d098.
//
// Solidity: function getLastPendingTransactionId() view returns(uint256)
func (_Inheritance *InheritanceCaller) GetLastPendingTransactionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "getLastPendingTransactionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPendingTransactionId is a free data retrieval call binding the contract method 0x62d5d098.
//
// Solidity: function getLastPendingTransactionId() view returns(uint256)
func (_Inheritance *InheritanceSession) GetLastPendingTransactionId() (*big.Int, error) {
	return _Inheritance.Contract.GetLastPendingTransactionId(&_Inheritance.CallOpts)
}

// GetLastPendingTransactionId is a free data retrieval call binding the contract method 0x62d5d098.
//
// Solidity: function getLastPendingTransactionId() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) GetLastPendingTransactionId() (*big.Int, error) {
	return _Inheritance.Contract.GetLastPendingTransactionId(&_Inheritance.CallOpts)
}

// GetSignatoryAddresses is a free data retrieval call binding the contract method 0x2d64258c.
//
// Solidity: function getSignatoryAddresses() view returns(address[])
func (_Inheritance *InheritanceCaller) GetSignatoryAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "getSignatoryAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSignatoryAddresses is a free data retrieval call binding the contract method 0x2d64258c.
//
// Solidity: function getSignatoryAddresses() view returns(address[])
func (_Inheritance *InheritanceSession) GetSignatoryAddresses() ([]common.Address, error) {
	return _Inheritance.Contract.GetSignatoryAddresses(&_Inheritance.CallOpts)
}

// GetSignatoryAddresses is a free data retrieval call binding the contract method 0x2d64258c.
//
// Solidity: function getSignatoryAddresses() view returns(address[])
func (_Inheritance *InheritanceCallerSession) GetSignatoryAddresses() ([]common.Address, error) {
	return _Inheritance.Contract.GetSignatoryAddresses(&_Inheritance.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Inheritance *InheritanceCaller) GetTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "getTransactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Inheritance *InheritanceSession) GetTransactionCount() (*big.Int, error) {
	return _Inheritance.Contract.GetTransactionCount(&_Inheritance.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) GetTransactionCount() (*big.Int, error) {
	return _Inheritance.Contract.GetTransactionCount(&_Inheritance.CallOpts)
}

// GetWeightForAmount is a free data retrieval call binding the contract method 0x1eb07371.
//
// Solidity: function getWeightForAmount(address _signatory, uint256 _amount) view returns(uint256)
func (_Inheritance *InheritanceCaller) GetWeightForAmount(opts *bind.CallOpts, _signatory common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "getWeightForAmount", _signatory, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWeightForAmount is a free data retrieval call binding the contract method 0x1eb07371.
//
// Solidity: function getWeightForAmount(address _signatory, uint256 _amount) view returns(uint256)
func (_Inheritance *InheritanceSession) GetWeightForAmount(_signatory common.Address, _amount *big.Int) (*big.Int, error) {
	return _Inheritance.Contract.GetWeightForAmount(&_Inheritance.CallOpts, _signatory, _amount)
}

// GetWeightForAmount is a free data retrieval call binding the contract method 0x1eb07371.
//
// Solidity: function getWeightForAmount(address _signatory, uint256 _amount) view returns(uint256)
func (_Inheritance *InheritanceCallerSession) GetWeightForAmount(_signatory common.Address, _amount *big.Int) (*big.Int, error) {
	return _Inheritance.Contract.GetWeightForAmount(&_Inheritance.CallOpts, _signatory, _amount)
}

// GetWithdrawLimitForAddress is a free data retrieval call binding the contract method 0x35f3be0c.
//
// Solidity: function getWithdrawLimitForAddress(address _signatory) view returns(uint256)
func (_Inheritance *InheritanceCaller) GetWithdrawLimitForAddress(opts *bind.CallOpts, _signatory common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "getWithdrawLimitForAddress", _signatory)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawLimitForAddress is a free data retrieval call binding the contract method 0x35f3be0c.
//
// Solidity: function getWithdrawLimitForAddress(address _signatory) view returns(uint256)
func (_Inheritance *InheritanceSession) GetWithdrawLimitForAddress(_signatory common.Address) (*big.Int, error) {
	return _Inheritance.Contract.GetWithdrawLimitForAddress(&_Inheritance.CallOpts, _signatory)
}

// GetWithdrawLimitForAddress is a free data retrieval call binding the contract method 0x35f3be0c.
//
// Solidity: function getWithdrawLimitForAddress(address _signatory) view returns(uint256)
func (_Inheritance *InheritanceCallerSession) GetWithdrawLimitForAddress(_signatory common.Address) (*big.Int, error) {
	return _Inheritance.Contract.GetWithdrawLimitForAddress(&_Inheritance.CallOpts, _signatory)
}

// GovernmentAddress is a free data retrieval call binding the contract method 0xd570ab34.
//
// Solidity: function governmentAddress() view returns(address)
func (_Inheritance *InheritanceCaller) GovernmentAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "governmentAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernmentAddress is a free data retrieval call binding the contract method 0xd570ab34.
//
// Solidity: function governmentAddress() view returns(address)
func (_Inheritance *InheritanceSession) GovernmentAddress() (common.Address, error) {
	return _Inheritance.Contract.GovernmentAddress(&_Inheritance.CallOpts)
}

// GovernmentAddress is a free data retrieval call binding the contract method 0xd570ab34.
//
// Solidity: function governmentAddress() view returns(address)
func (_Inheritance *InheritanceCallerSession) GovernmentAddress() (common.Address, error) {
	return _Inheritance.Contract.GovernmentAddress(&_Inheritance.CallOpts)
}

// InactivityThreshold is a free data retrieval call binding the contract method 0x9ec86276.
//
// Solidity: function inactivityThreshold() view returns(uint256)
func (_Inheritance *InheritanceCaller) InactivityThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "inactivityThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InactivityThreshold is a free data retrieval call binding the contract method 0x9ec86276.
//
// Solidity: function inactivityThreshold() view returns(uint256)
func (_Inheritance *InheritanceSession) InactivityThreshold() (*big.Int, error) {
	return _Inheritance.Contract.InactivityThreshold(&_Inheritance.CallOpts)
}

// InactivityThreshold is a free data retrieval call binding the contract method 0x9ec86276.
//
// Solidity: function inactivityThreshold() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) InactivityThreshold() (*big.Int, error) {
	return _Inheritance.Contract.InactivityThreshold(&_Inheritance.CallOpts)
}

// InheritanceDeadline is a free data retrieval call binding the contract method 0xf79b66af.
//
// Solidity: function inheritanceDeadline() view returns(uint256)
func (_Inheritance *InheritanceCaller) InheritanceDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "inheritanceDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InheritanceDeadline is a free data retrieval call binding the contract method 0xf79b66af.
//
// Solidity: function inheritanceDeadline() view returns(uint256)
func (_Inheritance *InheritanceSession) InheritanceDeadline() (*big.Int, error) {
	return _Inheritance.Contract.InheritanceDeadline(&_Inheritance.CallOpts)
}

// InheritanceDeadline is a free data retrieval call binding the contract method 0xf79b66af.
//
// Solidity: function inheritanceDeadline() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) InheritanceDeadline() (*big.Int, error) {
	return _Inheritance.Contract.InheritanceDeadline(&_Inheritance.CallOpts)
}

// InheritanceDeadlineActive is a free data retrieval call binding the contract method 0x2007c2b7.
//
// Solidity: function inheritanceDeadlineActive() view returns(bool)
func (_Inheritance *InheritanceCaller) InheritanceDeadlineActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "inheritanceDeadlineActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InheritanceDeadlineActive is a free data retrieval call binding the contract method 0x2007c2b7.
//
// Solidity: function inheritanceDeadlineActive() view returns(bool)
func (_Inheritance *InheritanceSession) InheritanceDeadlineActive() (bool, error) {
	return _Inheritance.Contract.InheritanceDeadlineActive(&_Inheritance.CallOpts)
}

// InheritanceDeadlineActive is a free data retrieval call binding the contract method 0x2007c2b7.
//
// Solidity: function inheritanceDeadlineActive() view returns(bool)
func (_Inheritance *InheritanceCallerSession) InheritanceDeadlineActive() (bool, error) {
	return _Inheritance.Contract.InheritanceDeadlineActive(&_Inheritance.CallOpts)
}

// InheritanceRequestActive is a free data retrieval call binding the contract method 0x2dac1c5b.
//
// Solidity: function inheritanceRequestActive() view returns(bool)
func (_Inheritance *InheritanceCaller) InheritanceRequestActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "inheritanceRequestActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InheritanceRequestActive is a free data retrieval call binding the contract method 0x2dac1c5b.
//
// Solidity: function inheritanceRequestActive() view returns(bool)
func (_Inheritance *InheritanceSession) InheritanceRequestActive() (bool, error) {
	return _Inheritance.Contract.InheritanceRequestActive(&_Inheritance.CallOpts)
}

// InheritanceRequestActive is a free data retrieval call binding the contract method 0x2dac1c5b.
//
// Solidity: function inheritanceRequestActive() view returns(bool)
func (_Inheritance *InheritanceCallerSession) InheritanceRequestActive() (bool, error) {
	return _Inheritance.Contract.InheritanceRequestActive(&_Inheritance.CallOpts)
}

// InheritanceRequestTime is a free data retrieval call binding the contract method 0x195e05f7.
//
// Solidity: function inheritanceRequestTime() view returns(uint256)
func (_Inheritance *InheritanceCaller) InheritanceRequestTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "inheritanceRequestTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InheritanceRequestTime is a free data retrieval call binding the contract method 0x195e05f7.
//
// Solidity: function inheritanceRequestTime() view returns(uint256)
func (_Inheritance *InheritanceSession) InheritanceRequestTime() (*big.Int, error) {
	return _Inheritance.Contract.InheritanceRequestTime(&_Inheritance.CallOpts)
}

// InheritanceRequestTime is a free data retrieval call binding the contract method 0x195e05f7.
//
// Solidity: function inheritanceRequestTime() view returns(uint256)
func (_Inheritance *InheritanceCallerSession) InheritanceRequestTime() (*big.Int, error) {
	return _Inheritance.Contract.InheritanceRequestTime(&_Inheritance.CallOpts)
}

// IsWalletAddressSet is a free data retrieval call binding the contract method 0xa7438db0.
//
// Solidity: function isWalletAddressSet() view returns(bool)
func (_Inheritance *InheritanceCaller) IsWalletAddressSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "isWalletAddressSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletAddressSet is a free data retrieval call binding the contract method 0xa7438db0.
//
// Solidity: function isWalletAddressSet() view returns(bool)
func (_Inheritance *InheritanceSession) IsWalletAddressSet() (bool, error) {
	return _Inheritance.Contract.IsWalletAddressSet(&_Inheritance.CallOpts)
}

// IsWalletAddressSet is a free data retrieval call binding the contract method 0xa7438db0.
//
// Solidity: function isWalletAddressSet() view returns(bool)
func (_Inheritance *InheritanceCallerSession) IsWalletAddressSet() (bool, error) {
	return _Inheritance.Contract.IsWalletAddressSet(&_Inheritance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Inheritance *InheritanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Inheritance *InheritanceSession) Owner() (common.Address, error) {
	return _Inheritance.Contract.Owner(&_Inheritance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Inheritance *InheritanceCallerSession) Owner() (common.Address, error) {
	return _Inheritance.Contract.Owner(&_Inheritance.CallOpts)
}

// Signatories is a free data retrieval call binding the contract method 0x7932f372.
//
// Solidity: function signatories(address ) view returns(bool)
func (_Inheritance *InheritanceCaller) Signatories(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "signatories", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signatories is a free data retrieval call binding the contract method 0x7932f372.
//
// Solidity: function signatories(address ) view returns(bool)
func (_Inheritance *InheritanceSession) Signatories(arg0 common.Address) (bool, error) {
	return _Inheritance.Contract.Signatories(&_Inheritance.CallOpts, arg0)
}

// Signatories is a free data retrieval call binding the contract method 0x7932f372.
//
// Solidity: function signatories(address ) view returns(bool)
func (_Inheritance *InheritanceCallerSession) Signatories(arg0 common.Address) (bool, error) {
	return _Inheritance.Contract.Signatories(&_Inheritance.CallOpts, arg0)
}

// SignatoryAddresses is a free data retrieval call binding the contract method 0x8232c223.
//
// Solidity: function signatoryAddresses(uint256 ) view returns(address)
func (_Inheritance *InheritanceCaller) SignatoryAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "signatoryAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignatoryAddresses is a free data retrieval call binding the contract method 0x8232c223.
//
// Solidity: function signatoryAddresses(uint256 ) view returns(address)
func (_Inheritance *InheritanceSession) SignatoryAddresses(arg0 *big.Int) (common.Address, error) {
	return _Inheritance.Contract.SignatoryAddresses(&_Inheritance.CallOpts, arg0)
}

// SignatoryAddresses is a free data retrieval call binding the contract method 0x8232c223.
//
// Solidity: function signatoryAddresses(uint256 ) view returns(address)
func (_Inheritance *InheritanceCallerSession) SignatoryAddresses(arg0 *big.Int) (common.Address, error) {
	return _Inheritance.Contract.SignatoryAddresses(&_Inheritance.CallOpts, arg0)
}

// SignatoryIndexes is a free data retrieval call binding the contract method 0xc5da53a0.
//
// Solidity: function signatoryIndexes(address ) view returns(uint256)
func (_Inheritance *InheritanceCaller) SignatoryIndexes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "signatoryIndexes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatoryIndexes is a free data retrieval call binding the contract method 0xc5da53a0.
//
// Solidity: function signatoryIndexes(address ) view returns(uint256)
func (_Inheritance *InheritanceSession) SignatoryIndexes(arg0 common.Address) (*big.Int, error) {
	return _Inheritance.Contract.SignatoryIndexes(&_Inheritance.CallOpts, arg0)
}

// SignatoryIndexes is a free data retrieval call binding the contract method 0xc5da53a0.
//
// Solidity: function signatoryIndexes(address ) view returns(uint256)
func (_Inheritance *InheritanceCallerSession) SignatoryIndexes(arg0 common.Address) (*big.Int, error) {
	return _Inheritance.Contract.SignatoryIndexes(&_Inheritance.CallOpts, arg0)
}

// SignatoryWeights is a free data retrieval call binding the contract method 0xdcd98fbc.
//
// Solidity: function signatoryWeights(address ) view returns(uint256)
func (_Inheritance *InheritanceCaller) SignatoryWeights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "signatoryWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatoryWeights is a free data retrieval call binding the contract method 0xdcd98fbc.
//
// Solidity: function signatoryWeights(address ) view returns(uint256)
func (_Inheritance *InheritanceSession) SignatoryWeights(arg0 common.Address) (*big.Int, error) {
	return _Inheritance.Contract.SignatoryWeights(&_Inheritance.CallOpts, arg0)
}

// SignatoryWeights is a free data retrieval call binding the contract method 0xdcd98fbc.
//
// Solidity: function signatoryWeights(address ) view returns(uint256)
func (_Inheritance *InheritanceCallerSession) SignatoryWeights(arg0 common.Address) (*big.Int, error) {
	return _Inheritance.Contract.SignatoryWeights(&_Inheritance.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address signatory, uint256 amount, bool approved, uint256 approvalCount)
func (_Inheritance *InheritanceCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Signatory     common.Address
	Amount        *big.Int
	Approved      bool
	ApprovalCount *big.Int
}, error) {
	var out []interface{}
	err := _Inheritance.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Signatory     common.Address
		Amount        *big.Int
		Approved      bool
		ApprovalCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signatory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Approved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.ApprovalCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address signatory, uint256 amount, bool approved, uint256 approvalCount)
func (_Inheritance *InheritanceSession) Transactions(arg0 *big.Int) (struct {
	Signatory     common.Address
	Amount        *big.Int
	Approved      bool
	ApprovalCount *big.Int
}, error) {
	return _Inheritance.Contract.Transactions(&_Inheritance.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address signatory, uint256 amount, bool approved, uint256 approvalCount)
func (_Inheritance *InheritanceCallerSession) Transactions(arg0 *big.Int) (struct {
	Signatory     common.Address
	Amount        *big.Int
	Approved      bool
	ApprovalCount *big.Int
}, error) {
	return _Inheritance.Contract.Transactions(&_Inheritance.CallOpts, arg0)
}

// AddBalance is a paid mutator transaction binding the contract method 0xd91921ed.
//
// Solidity: function addBalance(uint256 amount) returns()
func (_Inheritance *InheritanceTransactor) AddBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "addBalance", amount)
}

// AddBalance is a paid mutator transaction binding the contract method 0xd91921ed.
//
// Solidity: function addBalance(uint256 amount) returns()
func (_Inheritance *InheritanceSession) AddBalance(amount *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.AddBalance(&_Inheritance.TransactOpts, amount)
}

// AddBalance is a paid mutator transaction binding the contract method 0xd91921ed.
//
// Solidity: function addBalance(uint256 amount) returns()
func (_Inheritance *InheritanceTransactorSession) AddBalance(amount *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.AddBalance(&_Inheritance.TransactOpts, amount)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1f0da10f.
//
// Solidity: function addSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceTransactor) AddSignatory(opts *bind.TransactOpts, _signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "addSignatory", _signatory, _weight)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1f0da10f.
//
// Solidity: function addSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceSession) AddSignatory(_signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.AddSignatory(&_Inheritance.TransactOpts, _signatory, _weight)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1f0da10f.
//
// Solidity: function addSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceTransactorSession) AddSignatory(_signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.AddSignatory(&_Inheritance.TransactOpts, _signatory, _weight)
}

// ApproveSignatory is a paid mutator transaction binding the contract method 0xd3787c9a.
//
// Solidity: function approveSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceTransactor) ApproveSignatory(opts *bind.TransactOpts, _signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "approveSignatory", _signatory, _weight)
}

// ApproveSignatory is a paid mutator transaction binding the contract method 0xd3787c9a.
//
// Solidity: function approveSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceSession) ApproveSignatory(_signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.ApproveSignatory(&_Inheritance.TransactOpts, _signatory, _weight)
}

// ApproveSignatory is a paid mutator transaction binding the contract method 0xd3787c9a.
//
// Solidity: function approveSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceTransactorSession) ApproveSignatory(_signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.ApproveSignatory(&_Inheritance.TransactOpts, _signatory, _weight)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0x242232d1.
//
// Solidity: function approveTransaction(uint256 _transactionId) returns()
func (_Inheritance *InheritanceTransactor) ApproveTransaction(opts *bind.TransactOpts, _transactionId *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "approveTransaction", _transactionId)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0x242232d1.
//
// Solidity: function approveTransaction(uint256 _transactionId) returns()
func (_Inheritance *InheritanceSession) ApproveTransaction(_transactionId *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.ApproveTransaction(&_Inheritance.TransactOpts, _transactionId)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0x242232d1.
//
// Solidity: function approveTransaction(uint256 _transactionId) returns()
func (_Inheritance *InheritanceTransactorSession) ApproveTransaction(_transactionId *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.ApproveTransaction(&_Inheritance.TransactOpts, _transactionId)
}

// CancelInheritanceRequest is a paid mutator transaction binding the contract method 0xe76ec83f.
//
// Solidity: function cancelInheritanceRequest() returns()
func (_Inheritance *InheritanceTransactor) CancelInheritanceRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "cancelInheritanceRequest")
}

// CancelInheritanceRequest is a paid mutator transaction binding the contract method 0xe76ec83f.
//
// Solidity: function cancelInheritanceRequest() returns()
func (_Inheritance *InheritanceSession) CancelInheritanceRequest() (*types.Transaction, error) {
	return _Inheritance.Contract.CancelInheritanceRequest(&_Inheritance.TransactOpts)
}

// CancelInheritanceRequest is a paid mutator transaction binding the contract method 0xe76ec83f.
//
// Solidity: function cancelInheritanceRequest() returns()
func (_Inheritance *InheritanceTransactorSession) CancelInheritanceRequest() (*types.Transaction, error) {
	return _Inheritance.Contract.CancelInheritanceRequest(&_Inheritance.TransactOpts)
}

// CheckInheritanceRequestStatus is a paid mutator transaction binding the contract method 0x0680b651.
//
// Solidity: function checkInheritanceRequestStatus() returns(bool)
func (_Inheritance *InheritanceTransactor) CheckInheritanceRequestStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "checkInheritanceRequestStatus")
}

// CheckInheritanceRequestStatus is a paid mutator transaction binding the contract method 0x0680b651.
//
// Solidity: function checkInheritanceRequestStatus() returns(bool)
func (_Inheritance *InheritanceSession) CheckInheritanceRequestStatus() (*types.Transaction, error) {
	return _Inheritance.Contract.CheckInheritanceRequestStatus(&_Inheritance.TransactOpts)
}

// CheckInheritanceRequestStatus is a paid mutator transaction binding the contract method 0x0680b651.
//
// Solidity: function checkInheritanceRequestStatus() returns(bool)
func (_Inheritance *InheritanceTransactorSession) CheckInheritanceRequestStatus() (*types.Transaction, error) {
	return _Inheritance.Contract.CheckInheritanceRequestStatus(&_Inheritance.TransactOpts)
}

// ConfirmInheritanceRequest is a paid mutator transaction binding the contract method 0x28f40de9.
//
// Solidity: function confirmInheritanceRequest() returns()
func (_Inheritance *InheritanceTransactor) ConfirmInheritanceRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "confirmInheritanceRequest")
}

// ConfirmInheritanceRequest is a paid mutator transaction binding the contract method 0x28f40de9.
//
// Solidity: function confirmInheritanceRequest() returns()
func (_Inheritance *InheritanceSession) ConfirmInheritanceRequest() (*types.Transaction, error) {
	return _Inheritance.Contract.ConfirmInheritanceRequest(&_Inheritance.TransactOpts)
}

// ConfirmInheritanceRequest is a paid mutator transaction binding the contract method 0x28f40de9.
//
// Solidity: function confirmInheritanceRequest() returns()
func (_Inheritance *InheritanceTransactorSession) ConfirmInheritanceRequest() (*types.Transaction, error) {
	return _Inheritance.Contract.ConfirmInheritanceRequest(&_Inheritance.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Inheritance *InheritanceTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Inheritance *InheritanceSession) DepositFunds() (*types.Transaction, error) {
	return _Inheritance.Contract.DepositFunds(&_Inheritance.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Inheritance *InheritanceTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _Inheritance.Contract.DepositFunds(&_Inheritance.TransactOpts)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xd5ab92fd.
//
// Solidity: function removeSignatory(address _signatory) returns()
func (_Inheritance *InheritanceTransactor) RemoveSignatory(opts *bind.TransactOpts, _signatory common.Address) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "removeSignatory", _signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xd5ab92fd.
//
// Solidity: function removeSignatory(address _signatory) returns()
func (_Inheritance *InheritanceSession) RemoveSignatory(_signatory common.Address) (*types.Transaction, error) {
	return _Inheritance.Contract.RemoveSignatory(&_Inheritance.TransactOpts, _signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xd5ab92fd.
//
// Solidity: function removeSignatory(address _signatory) returns()
func (_Inheritance *InheritanceTransactorSession) RemoveSignatory(_signatory common.Address) (*types.Transaction, error) {
	return _Inheritance.Contract.RemoveSignatory(&_Inheritance.TransactOpts, _signatory)
}

// RequestInheritance is a paid mutator transaction binding the contract method 0x0b6946da.
//
// Solidity: function requestInheritance() returns()
func (_Inheritance *InheritanceTransactor) RequestInheritance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "requestInheritance")
}

// RequestInheritance is a paid mutator transaction binding the contract method 0x0b6946da.
//
// Solidity: function requestInheritance() returns()
func (_Inheritance *InheritanceSession) RequestInheritance() (*types.Transaction, error) {
	return _Inheritance.Contract.RequestInheritance(&_Inheritance.TransactOpts)
}

// RequestInheritance is a paid mutator transaction binding the contract method 0x0b6946da.
//
// Solidity: function requestInheritance() returns()
func (_Inheritance *InheritanceTransactorSession) RequestInheritance() (*types.Transaction, error) {
	return _Inheritance.Contract.RequestInheritance(&_Inheritance.TransactOpts)
}

// RequestTransaction is a paid mutator transaction binding the contract method 0x2232dc7d.
//
// Solidity: function requestTransaction(uint256 _amount) returns()
func (_Inheritance *InheritanceTransactor) RequestTransaction(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "requestTransaction", _amount)
}

// RequestTransaction is a paid mutator transaction binding the contract method 0x2232dc7d.
//
// Solidity: function requestTransaction(uint256 _amount) returns()
func (_Inheritance *InheritanceSession) RequestTransaction(_amount *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.RequestTransaction(&_Inheritance.TransactOpts, _amount)
}

// RequestTransaction is a paid mutator transaction binding the contract method 0x2232dc7d.
//
// Solidity: function requestTransaction(uint256 _amount) returns()
func (_Inheritance *InheritanceTransactorSession) RequestTransaction(_amount *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.RequestTransaction(&_Inheritance.TransactOpts, _amount)
}

// SetApprovalThreshold is a paid mutator transaction binding the contract method 0xa0016b8c.
//
// Solidity: function setApprovalThreshold(uint256 _threshold) returns()
func (_Inheritance *InheritanceTransactor) SetApprovalThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setApprovalThreshold", _threshold)
}

// SetApprovalThreshold is a paid mutator transaction binding the contract method 0xa0016b8c.
//
// Solidity: function setApprovalThreshold(uint256 _threshold) returns()
func (_Inheritance *InheritanceSession) SetApprovalThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetApprovalThreshold(&_Inheritance.TransactOpts, _threshold)
}

// SetApprovalThreshold is a paid mutator transaction binding the contract method 0xa0016b8c.
//
// Solidity: function setApprovalThreshold(uint256 _threshold) returns()
func (_Inheritance *InheritanceTransactorSession) SetApprovalThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetApprovalThreshold(&_Inheritance.TransactOpts, _threshold)
}

// SetGovernmentAddress is a paid mutator transaction binding the contract method 0x6a137219.
//
// Solidity: function setGovernmentAddress(address _address) returns()
func (_Inheritance *InheritanceTransactor) SetGovernmentAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setGovernmentAddress", _address)
}

// SetGovernmentAddress is a paid mutator transaction binding the contract method 0x6a137219.
//
// Solidity: function setGovernmentAddress(address _address) returns()
func (_Inheritance *InheritanceSession) SetGovernmentAddress(_address common.Address) (*types.Transaction, error) {
	return _Inheritance.Contract.SetGovernmentAddress(&_Inheritance.TransactOpts, _address)
}

// SetGovernmentAddress is a paid mutator transaction binding the contract method 0x6a137219.
//
// Solidity: function setGovernmentAddress(address _address) returns()
func (_Inheritance *InheritanceTransactorSession) SetGovernmentAddress(_address common.Address) (*types.Transaction, error) {
	return _Inheritance.Contract.SetGovernmentAddress(&_Inheritance.TransactOpts, _address)
}

// SetInactivityThreshold is a paid mutator transaction binding the contract method 0xd7eaef49.
//
// Solidity: function setInactivityThreshold(uint256 _threshold) returns()
func (_Inheritance *InheritanceTransactor) SetInactivityThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setInactivityThreshold", _threshold)
}

// SetInactivityThreshold is a paid mutator transaction binding the contract method 0xd7eaef49.
//
// Solidity: function setInactivityThreshold(uint256 _threshold) returns()
func (_Inheritance *InheritanceSession) SetInactivityThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetInactivityThreshold(&_Inheritance.TransactOpts, _threshold)
}

// SetInactivityThreshold is a paid mutator transaction binding the contract method 0xd7eaef49.
//
// Solidity: function setInactivityThreshold(uint256 _threshold) returns()
func (_Inheritance *InheritanceTransactorSession) SetInactivityThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetInactivityThreshold(&_Inheritance.TransactOpts, _threshold)
}

// SetInheritanceDeadline is a paid mutator transaction binding the contract method 0xf581aaed.
//
// Solidity: function setInheritanceDeadline(uint256 _deadline) returns()
func (_Inheritance *InheritanceTransactor) SetInheritanceDeadline(opts *bind.TransactOpts, _deadline *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setInheritanceDeadline", _deadline)
}

// SetInheritanceDeadline is a paid mutator transaction binding the contract method 0xf581aaed.
//
// Solidity: function setInheritanceDeadline(uint256 _deadline) returns()
func (_Inheritance *InheritanceSession) SetInheritanceDeadline(_deadline *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetInheritanceDeadline(&_Inheritance.TransactOpts, _deadline)
}

// SetInheritanceDeadline is a paid mutator transaction binding the contract method 0xf581aaed.
//
// Solidity: function setInheritanceDeadline(uint256 _deadline) returns()
func (_Inheritance *InheritanceTransactorSession) SetInheritanceDeadline(_deadline *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetInheritanceDeadline(&_Inheritance.TransactOpts, _deadline)
}

// SetLastTransactionTime is a paid mutator transaction binding the contract method 0x5980957b.
//
// Solidity: function setLastTransactionTime(uint256 _time) returns()
func (_Inheritance *InheritanceTransactor) SetLastTransactionTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setLastTransactionTime", _time)
}

// SetLastTransactionTime is a paid mutator transaction binding the contract method 0x5980957b.
//
// Solidity: function setLastTransactionTime(uint256 _time) returns()
func (_Inheritance *InheritanceSession) SetLastTransactionTime(_time *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetLastTransactionTime(&_Inheritance.TransactOpts, _time)
}

// SetLastTransactionTime is a paid mutator transaction binding the contract method 0x5980957b.
//
// Solidity: function setLastTransactionTime(uint256 _time) returns()
func (_Inheritance *InheritanceTransactorSession) SetLastTransactionTime(_time *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetLastTransactionTime(&_Inheritance.TransactOpts, _time)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0xac1a386a.
//
// Solidity: function setWalletAddress(address _walletAddress) returns()
func (_Inheritance *InheritanceTransactor) SetWalletAddress(opts *bind.TransactOpts, _walletAddress common.Address) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setWalletAddress", _walletAddress)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0xac1a386a.
//
// Solidity: function setWalletAddress(address _walletAddress) returns()
func (_Inheritance *InheritanceSession) SetWalletAddress(_walletAddress common.Address) (*types.Transaction, error) {
	return _Inheritance.Contract.SetWalletAddress(&_Inheritance.TransactOpts, _walletAddress)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0xac1a386a.
//
// Solidity: function setWalletAddress(address _walletAddress) returns()
func (_Inheritance *InheritanceTransactorSession) SetWalletAddress(_walletAddress common.Address) (*types.Transaction, error) {
	return _Inheritance.Contract.SetWalletAddress(&_Inheritance.TransactOpts, _walletAddress)
}

// SetWeightForSignatory is a paid mutator transaction binding the contract method 0x8c2c2d5e.
//
// Solidity: function setWeightForSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceTransactor) SetWeightForSignatory(opts *bind.TransactOpts, _signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.contract.Transact(opts, "setWeightForSignatory", _signatory, _weight)
}

// SetWeightForSignatory is a paid mutator transaction binding the contract method 0x8c2c2d5e.
//
// Solidity: function setWeightForSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceSession) SetWeightForSignatory(_signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetWeightForSignatory(&_Inheritance.TransactOpts, _signatory, _weight)
}

// SetWeightForSignatory is a paid mutator transaction binding the contract method 0x8c2c2d5e.
//
// Solidity: function setWeightForSignatory(address _signatory, uint256 _weight) returns()
func (_Inheritance *InheritanceTransactorSession) SetWeightForSignatory(_signatory common.Address, _weight *big.Int) (*types.Transaction, error) {
	return _Inheritance.Contract.SetWeightForSignatory(&_Inheritance.TransactOpts, _signatory, _weight)
}

// InheritanceApprovalReceivedIterator is returned from FilterApprovalReceived and is used to iterate over the raw logs and unpacked data for ApprovalReceived events raised by the Inheritance contract.
type InheritanceApprovalReceivedIterator struct {
	Event *InheritanceApprovalReceived // Event containing the contract specifics and raw log

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
func (it *InheritanceApprovalReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceApprovalReceived)
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
		it.Event = new(InheritanceApprovalReceived)
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
func (it *InheritanceApprovalReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceApprovalReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceApprovalReceived represents a ApprovalReceived event raised by the Inheritance contract.
type InheritanceApprovalReceived struct {
	TransactionId *big.Int
	Approver      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterApprovalReceived is a free log retrieval operation binding the contract event 0xbe556a88b31b59bf2916d67a1114497febeef512cc2e5fa4ad132f7865c56ac2.
//
// Solidity: event ApprovalReceived(uint256 indexed transactionId, address indexed approver)
func (_Inheritance *InheritanceFilterer) FilterApprovalReceived(opts *bind.FilterOpts, transactionId []*big.Int, approver []common.Address) (*InheritanceApprovalReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "ApprovalReceived", transactionIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceApprovalReceivedIterator{contract: _Inheritance.contract, event: "ApprovalReceived", logs: logs, sub: sub}, nil
}

// WatchApprovalReceived is a free log subscription operation binding the contract event 0xbe556a88b31b59bf2916d67a1114497febeef512cc2e5fa4ad132f7865c56ac2.
//
// Solidity: event ApprovalReceived(uint256 indexed transactionId, address indexed approver)
func (_Inheritance *InheritanceFilterer) WatchApprovalReceived(opts *bind.WatchOpts, sink chan<- *InheritanceApprovalReceived, transactionId []*big.Int, approver []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "ApprovalReceived", transactionIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceApprovalReceived)
				if err := _Inheritance.contract.UnpackLog(event, "ApprovalReceived", log); err != nil {
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

// ParseApprovalReceived is a log parse operation binding the contract event 0xbe556a88b31b59bf2916d67a1114497febeef512cc2e5fa4ad132f7865c56ac2.
//
// Solidity: event ApprovalReceived(uint256 indexed transactionId, address indexed approver)
func (_Inheritance *InheritanceFilterer) ParseApprovalReceived(log types.Log) (*InheritanceApprovalReceived, error) {
	event := new(InheritanceApprovalReceived)
	if err := _Inheritance.contract.UnpackLog(event, "ApprovalReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceAssetsDepositedIterator is returned from FilterAssetsDeposited and is used to iterate over the raw logs and unpacked data for AssetsDeposited events raised by the Inheritance contract.
type InheritanceAssetsDepositedIterator struct {
	Event *InheritanceAssetsDeposited // Event containing the contract specifics and raw log

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
func (it *InheritanceAssetsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceAssetsDeposited)
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
		it.Event = new(InheritanceAssetsDeposited)
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
func (it *InheritanceAssetsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceAssetsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceAssetsDeposited represents a AssetsDeposited event raised by the Inheritance contract.
type InheritanceAssetsDeposited struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetsDeposited is a free log retrieval operation binding the contract event 0x4394d1aa82c7930ef4a9916042dd27d578f39d015baa27b33969d08201d860b9.
//
// Solidity: event AssetsDeposited(address indexed recipient, uint256 indexed amount)
func (_Inheritance *InheritanceFilterer) FilterAssetsDeposited(opts *bind.FilterOpts, recipient []common.Address, amount []*big.Int) (*InheritanceAssetsDepositedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "AssetsDeposited", recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceAssetsDepositedIterator{contract: _Inheritance.contract, event: "AssetsDeposited", logs: logs, sub: sub}, nil
}

// WatchAssetsDeposited is a free log subscription operation binding the contract event 0x4394d1aa82c7930ef4a9916042dd27d578f39d015baa27b33969d08201d860b9.
//
// Solidity: event AssetsDeposited(address indexed recipient, uint256 indexed amount)
func (_Inheritance *InheritanceFilterer) WatchAssetsDeposited(opts *bind.WatchOpts, sink chan<- *InheritanceAssetsDeposited, recipient []common.Address, amount []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "AssetsDeposited", recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceAssetsDeposited)
				if err := _Inheritance.contract.UnpackLog(event, "AssetsDeposited", log); err != nil {
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

// ParseAssetsDeposited is a log parse operation binding the contract event 0x4394d1aa82c7930ef4a9916042dd27d578f39d015baa27b33969d08201d860b9.
//
// Solidity: event AssetsDeposited(address indexed recipient, uint256 indexed amount)
func (_Inheritance *InheritanceFilterer) ParseAssetsDeposited(log types.Log) (*InheritanceAssetsDeposited, error) {
	event := new(InheritanceAssetsDeposited)
	if err := _Inheritance.contract.UnpackLog(event, "AssetsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceInheritanceRequestApprovedIterator is returned from FilterInheritanceRequestApproved and is used to iterate over the raw logs and unpacked data for InheritanceRequestApproved events raised by the Inheritance contract.
type InheritanceInheritanceRequestApprovedIterator struct {
	Event *InheritanceInheritanceRequestApproved // Event containing the contract specifics and raw log

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
func (it *InheritanceInheritanceRequestApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceInheritanceRequestApproved)
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
		it.Event = new(InheritanceInheritanceRequestApproved)
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
func (it *InheritanceInheritanceRequestApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceInheritanceRequestApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceInheritanceRequestApproved represents a InheritanceRequestApproved event raised by the Inheritance contract.
type InheritanceInheritanceRequestApproved struct {
	InheritanceRequestTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterInheritanceRequestApproved is a free log retrieval operation binding the contract event 0x1d2c23efd7722cf7ac5a0cdb95d4d2b67b76bf7a583269ccce699b552bee1976.
//
// Solidity: event InheritanceRequestApproved(uint256 indexed inheritanceRequestTime)
func (_Inheritance *InheritanceFilterer) FilterInheritanceRequestApproved(opts *bind.FilterOpts, inheritanceRequestTime []*big.Int) (*InheritanceInheritanceRequestApprovedIterator, error) {

	var inheritanceRequestTimeRule []interface{}
	for _, inheritanceRequestTimeItem := range inheritanceRequestTime {
		inheritanceRequestTimeRule = append(inheritanceRequestTimeRule, inheritanceRequestTimeItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "InheritanceRequestApproved", inheritanceRequestTimeRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceInheritanceRequestApprovedIterator{contract: _Inheritance.contract, event: "InheritanceRequestApproved", logs: logs, sub: sub}, nil
}

// WatchInheritanceRequestApproved is a free log subscription operation binding the contract event 0x1d2c23efd7722cf7ac5a0cdb95d4d2b67b76bf7a583269ccce699b552bee1976.
//
// Solidity: event InheritanceRequestApproved(uint256 indexed inheritanceRequestTime)
func (_Inheritance *InheritanceFilterer) WatchInheritanceRequestApproved(opts *bind.WatchOpts, sink chan<- *InheritanceInheritanceRequestApproved, inheritanceRequestTime []*big.Int) (event.Subscription, error) {

	var inheritanceRequestTimeRule []interface{}
	for _, inheritanceRequestTimeItem := range inheritanceRequestTime {
		inheritanceRequestTimeRule = append(inheritanceRequestTimeRule, inheritanceRequestTimeItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "InheritanceRequestApproved", inheritanceRequestTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceInheritanceRequestApproved)
				if err := _Inheritance.contract.UnpackLog(event, "InheritanceRequestApproved", log); err != nil {
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

// ParseInheritanceRequestApproved is a log parse operation binding the contract event 0x1d2c23efd7722cf7ac5a0cdb95d4d2b67b76bf7a583269ccce699b552bee1976.
//
// Solidity: event InheritanceRequestApproved(uint256 indexed inheritanceRequestTime)
func (_Inheritance *InheritanceFilterer) ParseInheritanceRequestApproved(log types.Log) (*InheritanceInheritanceRequestApproved, error) {
	event := new(InheritanceInheritanceRequestApproved)
	if err := _Inheritance.contract.UnpackLog(event, "InheritanceRequestApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceInheritanceRequestCanceledIterator is returned from FilterInheritanceRequestCanceled and is used to iterate over the raw logs and unpacked data for InheritanceRequestCanceled events raised by the Inheritance contract.
type InheritanceInheritanceRequestCanceledIterator struct {
	Event *InheritanceInheritanceRequestCanceled // Event containing the contract specifics and raw log

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
func (it *InheritanceInheritanceRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceInheritanceRequestCanceled)
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
		it.Event = new(InheritanceInheritanceRequestCanceled)
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
func (it *InheritanceInheritanceRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceInheritanceRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceInheritanceRequestCanceled represents a InheritanceRequestCanceled event raised by the Inheritance contract.
type InheritanceInheritanceRequestCanceled struct {
	InheritanceRequestTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterInheritanceRequestCanceled is a free log retrieval operation binding the contract event 0xcac5ba8635c7845e362ae1bc94e17eb8d7dd6c7ec54b00e32360fcce2ec34365.
//
// Solidity: event InheritanceRequestCanceled(uint256 indexed inheritanceRequestTime)
func (_Inheritance *InheritanceFilterer) FilterInheritanceRequestCanceled(opts *bind.FilterOpts, inheritanceRequestTime []*big.Int) (*InheritanceInheritanceRequestCanceledIterator, error) {

	var inheritanceRequestTimeRule []interface{}
	for _, inheritanceRequestTimeItem := range inheritanceRequestTime {
		inheritanceRequestTimeRule = append(inheritanceRequestTimeRule, inheritanceRequestTimeItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "InheritanceRequestCanceled", inheritanceRequestTimeRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceInheritanceRequestCanceledIterator{contract: _Inheritance.contract, event: "InheritanceRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchInheritanceRequestCanceled is a free log subscription operation binding the contract event 0xcac5ba8635c7845e362ae1bc94e17eb8d7dd6c7ec54b00e32360fcce2ec34365.
//
// Solidity: event InheritanceRequestCanceled(uint256 indexed inheritanceRequestTime)
func (_Inheritance *InheritanceFilterer) WatchInheritanceRequestCanceled(opts *bind.WatchOpts, sink chan<- *InheritanceInheritanceRequestCanceled, inheritanceRequestTime []*big.Int) (event.Subscription, error) {

	var inheritanceRequestTimeRule []interface{}
	for _, inheritanceRequestTimeItem := range inheritanceRequestTime {
		inheritanceRequestTimeRule = append(inheritanceRequestTimeRule, inheritanceRequestTimeItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "InheritanceRequestCanceled", inheritanceRequestTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceInheritanceRequestCanceled)
				if err := _Inheritance.contract.UnpackLog(event, "InheritanceRequestCanceled", log); err != nil {
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

// ParseInheritanceRequestCanceled is a log parse operation binding the contract event 0xcac5ba8635c7845e362ae1bc94e17eb8d7dd6c7ec54b00e32360fcce2ec34365.
//
// Solidity: event InheritanceRequestCanceled(uint256 indexed inheritanceRequestTime)
func (_Inheritance *InheritanceFilterer) ParseInheritanceRequestCanceled(log types.Log) (*InheritanceInheritanceRequestCanceled, error) {
	event := new(InheritanceInheritanceRequestCanceled)
	if err := _Inheritance.contract.UnpackLog(event, "InheritanceRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceInheritanceRequestedIterator is returned from FilterInheritanceRequested and is used to iterate over the raw logs and unpacked data for InheritanceRequested events raised by the Inheritance contract.
type InheritanceInheritanceRequestedIterator struct {
	Event *InheritanceInheritanceRequested // Event containing the contract specifics and raw log

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
func (it *InheritanceInheritanceRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceInheritanceRequested)
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
		it.Event = new(InheritanceInheritanceRequested)
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
func (it *InheritanceInheritanceRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceInheritanceRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceInheritanceRequested represents a InheritanceRequested event raised by the Inheritance contract.
type InheritanceInheritanceRequested struct {
	InheritanceRequestTime *big.Int
	Signatory              common.Address
	Deadline               *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterInheritanceRequested is a free log retrieval operation binding the contract event 0x6fbf30cca1219d109dab754dd32c03442e0559fe73758f586dc6f490e2651519.
//
// Solidity: event InheritanceRequested(uint256 indexed inheritanceRequestTime, address indexed signatory, uint256 deadline)
func (_Inheritance *InheritanceFilterer) FilterInheritanceRequested(opts *bind.FilterOpts, inheritanceRequestTime []*big.Int, signatory []common.Address) (*InheritanceInheritanceRequestedIterator, error) {

	var inheritanceRequestTimeRule []interface{}
	for _, inheritanceRequestTimeItem := range inheritanceRequestTime {
		inheritanceRequestTimeRule = append(inheritanceRequestTimeRule, inheritanceRequestTimeItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "InheritanceRequested", inheritanceRequestTimeRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceInheritanceRequestedIterator{contract: _Inheritance.contract, event: "InheritanceRequested", logs: logs, sub: sub}, nil
}

// WatchInheritanceRequested is a free log subscription operation binding the contract event 0x6fbf30cca1219d109dab754dd32c03442e0559fe73758f586dc6f490e2651519.
//
// Solidity: event InheritanceRequested(uint256 indexed inheritanceRequestTime, address indexed signatory, uint256 deadline)
func (_Inheritance *InheritanceFilterer) WatchInheritanceRequested(opts *bind.WatchOpts, sink chan<- *InheritanceInheritanceRequested, inheritanceRequestTime []*big.Int, signatory []common.Address) (event.Subscription, error) {

	var inheritanceRequestTimeRule []interface{}
	for _, inheritanceRequestTimeItem := range inheritanceRequestTime {
		inheritanceRequestTimeRule = append(inheritanceRequestTimeRule, inheritanceRequestTimeItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "InheritanceRequested", inheritanceRequestTimeRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceInheritanceRequested)
				if err := _Inheritance.contract.UnpackLog(event, "InheritanceRequested", log); err != nil {
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

// ParseInheritanceRequested is a log parse operation binding the contract event 0x6fbf30cca1219d109dab754dd32c03442e0559fe73758f586dc6f490e2651519.
//
// Solidity: event InheritanceRequested(uint256 indexed inheritanceRequestTime, address indexed signatory, uint256 deadline)
func (_Inheritance *InheritanceFilterer) ParseInheritanceRequested(log types.Log) (*InheritanceInheritanceRequested, error) {
	event := new(InheritanceInheritanceRequested)
	if err := _Inheritance.contract.UnpackLog(event, "InheritanceRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceSignatoryAddedIterator is returned from FilterSignatoryAdded and is used to iterate over the raw logs and unpacked data for SignatoryAdded events raised by the Inheritance contract.
type InheritanceSignatoryAddedIterator struct {
	Event *InheritanceSignatoryAdded // Event containing the contract specifics and raw log

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
func (it *InheritanceSignatoryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceSignatoryAdded)
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
		it.Event = new(InheritanceSignatoryAdded)
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
func (it *InheritanceSignatoryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceSignatoryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceSignatoryAdded represents a SignatoryAdded event raised by the Inheritance contract.
type InheritanceSignatoryAdded struct {
	Signatory common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignatoryAdded is a free log retrieval operation binding the contract event 0x35e9f4e035b3b6d11ab65f031cdf634b53e98c81266f5dcb19ffa652533509f2.
//
// Solidity: event SignatoryAdded(address indexed signatory)
func (_Inheritance *InheritanceFilterer) FilterSignatoryAdded(opts *bind.FilterOpts, signatory []common.Address) (*InheritanceSignatoryAddedIterator, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "SignatoryAdded", signatoryRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceSignatoryAddedIterator{contract: _Inheritance.contract, event: "SignatoryAdded", logs: logs, sub: sub}, nil
}

// WatchSignatoryAdded is a free log subscription operation binding the contract event 0x35e9f4e035b3b6d11ab65f031cdf634b53e98c81266f5dcb19ffa652533509f2.
//
// Solidity: event SignatoryAdded(address indexed signatory)
func (_Inheritance *InheritanceFilterer) WatchSignatoryAdded(opts *bind.WatchOpts, sink chan<- *InheritanceSignatoryAdded, signatory []common.Address) (event.Subscription, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "SignatoryAdded", signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceSignatoryAdded)
				if err := _Inheritance.contract.UnpackLog(event, "SignatoryAdded", log); err != nil {
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

// ParseSignatoryAdded is a log parse operation binding the contract event 0x35e9f4e035b3b6d11ab65f031cdf634b53e98c81266f5dcb19ffa652533509f2.
//
// Solidity: event SignatoryAdded(address indexed signatory)
func (_Inheritance *InheritanceFilterer) ParseSignatoryAdded(log types.Log) (*InheritanceSignatoryAdded, error) {
	event := new(InheritanceSignatoryAdded)
	if err := _Inheritance.contract.UnpackLog(event, "SignatoryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceSignatoryApprovedIterator is returned from FilterSignatoryApproved and is used to iterate over the raw logs and unpacked data for SignatoryApproved events raised by the Inheritance contract.
type InheritanceSignatoryApprovedIterator struct {
	Event *InheritanceSignatoryApproved // Event containing the contract specifics and raw log

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
func (it *InheritanceSignatoryApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceSignatoryApproved)
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
		it.Event = new(InheritanceSignatoryApproved)
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
func (it *InheritanceSignatoryApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceSignatoryApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceSignatoryApproved represents a SignatoryApproved event raised by the Inheritance contract.
type InheritanceSignatoryApproved struct {
	Signatory common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignatoryApproved is a free log retrieval operation binding the contract event 0x0ae3a378066318e5e0e990e30bfa0c987ce6891cb3b2cc5b4a65e74e14ecc5b3.
//
// Solidity: event SignatoryApproved(address indexed signatory)
func (_Inheritance *InheritanceFilterer) FilterSignatoryApproved(opts *bind.FilterOpts, signatory []common.Address) (*InheritanceSignatoryApprovedIterator, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "SignatoryApproved", signatoryRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceSignatoryApprovedIterator{contract: _Inheritance.contract, event: "SignatoryApproved", logs: logs, sub: sub}, nil
}

// WatchSignatoryApproved is a free log subscription operation binding the contract event 0x0ae3a378066318e5e0e990e30bfa0c987ce6891cb3b2cc5b4a65e74e14ecc5b3.
//
// Solidity: event SignatoryApproved(address indexed signatory)
func (_Inheritance *InheritanceFilterer) WatchSignatoryApproved(opts *bind.WatchOpts, sink chan<- *InheritanceSignatoryApproved, signatory []common.Address) (event.Subscription, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "SignatoryApproved", signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceSignatoryApproved)
				if err := _Inheritance.contract.UnpackLog(event, "SignatoryApproved", log); err != nil {
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

// ParseSignatoryApproved is a log parse operation binding the contract event 0x0ae3a378066318e5e0e990e30bfa0c987ce6891cb3b2cc5b4a65e74e14ecc5b3.
//
// Solidity: event SignatoryApproved(address indexed signatory)
func (_Inheritance *InheritanceFilterer) ParseSignatoryApproved(log types.Log) (*InheritanceSignatoryApproved, error) {
	event := new(InheritanceSignatoryApproved)
	if err := _Inheritance.contract.UnpackLog(event, "SignatoryApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceSignatoryRemovedIterator is returned from FilterSignatoryRemoved and is used to iterate over the raw logs and unpacked data for SignatoryRemoved events raised by the Inheritance contract.
type InheritanceSignatoryRemovedIterator struct {
	Event *InheritanceSignatoryRemoved // Event containing the contract specifics and raw log

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
func (it *InheritanceSignatoryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceSignatoryRemoved)
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
		it.Event = new(InheritanceSignatoryRemoved)
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
func (it *InheritanceSignatoryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceSignatoryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceSignatoryRemoved represents a SignatoryRemoved event raised by the Inheritance contract.
type InheritanceSignatoryRemoved struct {
	Signatory common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignatoryRemoved is a free log retrieval operation binding the contract event 0x854bd90657c9840469c248f50889c6cb638355a3363d3c8bd8dd3048c5751eb3.
//
// Solidity: event SignatoryRemoved(address indexed signatory)
func (_Inheritance *InheritanceFilterer) FilterSignatoryRemoved(opts *bind.FilterOpts, signatory []common.Address) (*InheritanceSignatoryRemovedIterator, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "SignatoryRemoved", signatoryRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceSignatoryRemovedIterator{contract: _Inheritance.contract, event: "SignatoryRemoved", logs: logs, sub: sub}, nil
}

// WatchSignatoryRemoved is a free log subscription operation binding the contract event 0x854bd90657c9840469c248f50889c6cb638355a3363d3c8bd8dd3048c5751eb3.
//
// Solidity: event SignatoryRemoved(address indexed signatory)
func (_Inheritance *InheritanceFilterer) WatchSignatoryRemoved(opts *bind.WatchOpts, sink chan<- *InheritanceSignatoryRemoved, signatory []common.Address) (event.Subscription, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "SignatoryRemoved", signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceSignatoryRemoved)
				if err := _Inheritance.contract.UnpackLog(event, "SignatoryRemoved", log); err != nil {
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

// ParseSignatoryRemoved is a log parse operation binding the contract event 0x854bd90657c9840469c248f50889c6cb638355a3363d3c8bd8dd3048c5751eb3.
//
// Solidity: event SignatoryRemoved(address indexed signatory)
func (_Inheritance *InheritanceFilterer) ParseSignatoryRemoved(log types.Log) (*InheritanceSignatoryRemoved, error) {
	event := new(InheritanceSignatoryRemoved)
	if err := _Inheritance.contract.UnpackLog(event, "SignatoryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceTransactionApprovedIterator is returned from FilterTransactionApproved and is used to iterate over the raw logs and unpacked data for TransactionApproved events raised by the Inheritance contract.
type InheritanceTransactionApprovedIterator struct {
	Event *InheritanceTransactionApproved // Event containing the contract specifics and raw log

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
func (it *InheritanceTransactionApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceTransactionApproved)
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
		it.Event = new(InheritanceTransactionApproved)
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
func (it *InheritanceTransactionApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceTransactionApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceTransactionApproved represents a TransactionApproved event raised by the Inheritance contract.
type InheritanceTransactionApproved struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionApproved is a free log retrieval operation binding the contract event 0xbd3984aa34b5ed86a6d4f470c1c4b79e124b8801762705aa1887495eb991b130.
//
// Solidity: event TransactionApproved(uint256 indexed transactionId)
func (_Inheritance *InheritanceFilterer) FilterTransactionApproved(opts *bind.FilterOpts, transactionId []*big.Int) (*InheritanceTransactionApprovedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "TransactionApproved", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceTransactionApprovedIterator{contract: _Inheritance.contract, event: "TransactionApproved", logs: logs, sub: sub}, nil
}

// WatchTransactionApproved is a free log subscription operation binding the contract event 0xbd3984aa34b5ed86a6d4f470c1c4b79e124b8801762705aa1887495eb991b130.
//
// Solidity: event TransactionApproved(uint256 indexed transactionId)
func (_Inheritance *InheritanceFilterer) WatchTransactionApproved(opts *bind.WatchOpts, sink chan<- *InheritanceTransactionApproved, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "TransactionApproved", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceTransactionApproved)
				if err := _Inheritance.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
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

// ParseTransactionApproved is a log parse operation binding the contract event 0xbd3984aa34b5ed86a6d4f470c1c4b79e124b8801762705aa1887495eb991b130.
//
// Solidity: event TransactionApproved(uint256 indexed transactionId)
func (_Inheritance *InheritanceFilterer) ParseTransactionApproved(log types.Log) (*InheritanceTransactionApproved, error) {
	event := new(InheritanceTransactionApproved)
	if err := _Inheritance.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceTransactionRequestedIterator is returned from FilterTransactionRequested and is used to iterate over the raw logs and unpacked data for TransactionRequested events raised by the Inheritance contract.
type InheritanceTransactionRequestedIterator struct {
	Event *InheritanceTransactionRequested // Event containing the contract specifics and raw log

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
func (it *InheritanceTransactionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceTransactionRequested)
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
		it.Event = new(InheritanceTransactionRequested)
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
func (it *InheritanceTransactionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceTransactionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceTransactionRequested represents a TransactionRequested event raised by the Inheritance contract.
type InheritanceTransactionRequested struct {
	TransactionId *big.Int
	Signatory     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionRequested is a free log retrieval operation binding the contract event 0x46a217ae59dad352469f44a6b35285964a9f93b456f0e3cd13475797025047e6.
//
// Solidity: event TransactionRequested(uint256 indexed transactionId, address indexed signatory, uint256 amount)
func (_Inheritance *InheritanceFilterer) FilterTransactionRequested(opts *bind.FilterOpts, transactionId []*big.Int, signatory []common.Address) (*InheritanceTransactionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "TransactionRequested", transactionIdRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceTransactionRequestedIterator{contract: _Inheritance.contract, event: "TransactionRequested", logs: logs, sub: sub}, nil
}

// WatchTransactionRequested is a free log subscription operation binding the contract event 0x46a217ae59dad352469f44a6b35285964a9f93b456f0e3cd13475797025047e6.
//
// Solidity: event TransactionRequested(uint256 indexed transactionId, address indexed signatory, uint256 amount)
func (_Inheritance *InheritanceFilterer) WatchTransactionRequested(opts *bind.WatchOpts, sink chan<- *InheritanceTransactionRequested, transactionId []*big.Int, signatory []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "TransactionRequested", transactionIdRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceTransactionRequested)
				if err := _Inheritance.contract.UnpackLog(event, "TransactionRequested", log); err != nil {
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

// ParseTransactionRequested is a log parse operation binding the contract event 0x46a217ae59dad352469f44a6b35285964a9f93b456f0e3cd13475797025047e6.
//
// Solidity: event TransactionRequested(uint256 indexed transactionId, address indexed signatory, uint256 amount)
func (_Inheritance *InheritanceFilterer) ParseTransactionRequested(log types.Log) (*InheritanceTransactionRequested, error) {
	event := new(InheritanceTransactionRequested)
	if err := _Inheritance.contract.UnpackLog(event, "TransactionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InheritanceWeightsUpdatedIterator is returned from FilterWeightsUpdated and is used to iterate over the raw logs and unpacked data for WeightsUpdated events raised by the Inheritance contract.
type InheritanceWeightsUpdatedIterator struct {
	Event *InheritanceWeightsUpdated // Event containing the contract specifics and raw log

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
func (it *InheritanceWeightsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InheritanceWeightsUpdated)
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
		it.Event = new(InheritanceWeightsUpdated)
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
func (it *InheritanceWeightsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InheritanceWeightsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InheritanceWeightsUpdated represents a WeightsUpdated event raised by the Inheritance contract.
type InheritanceWeightsUpdated struct {
	Signatory common.Address
	MaxShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWeightsUpdated is a free log retrieval operation binding the contract event 0x86e2be4c9d5228453468bb65d798139c5683f19ba27c007106863d0d2737b3b1.
//
// Solidity: event WeightsUpdated(address indexed signatory, uint256 maxShares)
func (_Inheritance *InheritanceFilterer) FilterWeightsUpdated(opts *bind.FilterOpts, signatory []common.Address) (*InheritanceWeightsUpdatedIterator, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.FilterLogs(opts, "WeightsUpdated", signatoryRule)
	if err != nil {
		return nil, err
	}
	return &InheritanceWeightsUpdatedIterator{contract: _Inheritance.contract, event: "WeightsUpdated", logs: logs, sub: sub}, nil
}

// WatchWeightsUpdated is a free log subscription operation binding the contract event 0x86e2be4c9d5228453468bb65d798139c5683f19ba27c007106863d0d2737b3b1.
//
// Solidity: event WeightsUpdated(address indexed signatory, uint256 maxShares)
func (_Inheritance *InheritanceFilterer) WatchWeightsUpdated(opts *bind.WatchOpts, sink chan<- *InheritanceWeightsUpdated, signatory []common.Address) (event.Subscription, error) {

	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Inheritance.contract.WatchLogs(opts, "WeightsUpdated", signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InheritanceWeightsUpdated)
				if err := _Inheritance.contract.UnpackLog(event, "WeightsUpdated", log); err != nil {
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

// ParseWeightsUpdated is a log parse operation binding the contract event 0x86e2be4c9d5228453468bb65d798139c5683f19ba27c007106863d0d2737b3b1.
//
// Solidity: event WeightsUpdated(address indexed signatory, uint256 maxShares)
func (_Inheritance *InheritanceFilterer) ParseWeightsUpdated(log types.Log) (*InheritanceWeightsUpdated, error) {
	event := new(InheritanceWeightsUpdated)
	if err := _Inheritance.contract.UnpackLog(event, "WeightsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

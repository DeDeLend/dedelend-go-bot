// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DDLoptionInfo is an auto generated low-level Go binding around an user-defined struct.
type DDLoptionInfo struct {
	StrategyAddress common.Address
	Amount          *big.Int
	Strike          *big.Int
	Expiration      *big.Int
	IsLong          bool
}

// DDLMetaData contains all meta data concerning the DDL contract.
var DDLMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[4]\",\"name\":\"_arrLongHegicStrategy\",\"type\":\"address[4]\"},{\"internalType\":\"address[4]\",\"name\":\"_arrShortHegicStrategy\",\"type\":\"address[4]\"},{\"internalType\":\"contractIERC721\",\"name\":\"_collateralToken\",\"type\":\"address\"},{\"internalType\":\"contractIHegicOperationalTreasury\",\"name\":\"_operationalPool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_USDC\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minBorrowLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_COLLATERAL_DECIMALS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_PriorLiqPriceCoef\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userReturn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolReturn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqFee\",\"type\":\"uint256\"}],\"name\":\"ExerciseByPriorLiqPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolProfit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqFee\",\"type\":\"uint256\"}],\"name\":\"ForcedExercise\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolProfit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqFee\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionID\",\"type\":\"uint256\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLLATERAL_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INTEREST_RATE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LTV_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PriorLiqPriceCoef\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"borrowedByOption\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newBorrowTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"calculateUpcomingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"upcomingFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"collateralInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"strategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strike\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLong\",\"type\":\"bool\"}],\"internalType\":\"structDDL.optionInfo\",\"name\":\"strategy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralToken\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"currentLiqPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exerciseByPriorLiqPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"forcedExercise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"liqPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"loanState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"loanStateByPriorLiqPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"lockOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"maxBorrowLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBorrowLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operationalPool\",\"outputs\":[{\"internalType\":\"contractIHegicOperationalTreasury\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIPoolDDL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"priorLiqPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"profitByOption\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setInterestRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setInterestRateDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setLTVDecimals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinBorrowLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"strategyType\",\"outputs\":[{\"internalType\":\"enumDDL.HegicStrategyType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DDLABI is the input ABI used to generate the binding from.
// Deprecated: Use DDLMetaData.ABI instead.
var DDLABI = DDLMetaData.ABI

// DDL is an auto generated Go binding around an Ethereum contract.
type DDL struct {
	DDLCaller     // Read-only binding to the contract
	DDLTransactor // Write-only binding to the contract
	DDLFilterer   // Log filterer for contract events
}

// DDLCaller is an auto generated read-only Go binding around an Ethereum contract.
type DDLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DDLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DDLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DDLSession struct {
	Contract     *DDL              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DDLCallerSession struct {
	Contract *DDLCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DDLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DDLTransactorSession struct {
	Contract     *DDLTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDLRaw is an auto generated low-level Go binding around an Ethereum contract.
type DDLRaw struct {
	Contract *DDL // Generic contract binding to access the raw methods on
}

// DDLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DDLCallerRaw struct {
	Contract *DDLCaller // Generic read-only contract binding to access the raw methods on
}

// DDLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DDLTransactorRaw struct {
	Contract *DDLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDDL creates a new instance of DDL, bound to a specific deployed contract.
func NewDDL(address common.Address, backend bind.ContractBackend) (*DDL, error) {
	contract, err := bindDDL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DDL{DDLCaller: DDLCaller{contract: contract}, DDLTransactor: DDLTransactor{contract: contract}, DDLFilterer: DDLFilterer{contract: contract}}, nil
}

// NewDDLCaller creates a new read-only instance of DDL, bound to a specific deployed contract.
func NewDDLCaller(address common.Address, caller bind.ContractCaller) (*DDLCaller, error) {
	contract, err := bindDDL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DDLCaller{contract: contract}, nil
}

// NewDDLTransactor creates a new write-only instance of DDL, bound to a specific deployed contract.
func NewDDLTransactor(address common.Address, transactor bind.ContractTransactor) (*DDLTransactor, error) {
	contract, err := bindDDL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DDLTransactor{contract: contract}, nil
}

// NewDDLFilterer creates a new log filterer instance of DDL, bound to a specific deployed contract.
func NewDDLFilterer(address common.Address, filterer bind.ContractFilterer) (*DDLFilterer, error) {
	contract, err := bindDDL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DDLFilterer{contract: contract}, nil
}

// bindDDL binds a generic wrapper to an already deployed contract.
func bindDDL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DDLABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDL *DDLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DDL.Contract.DDLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDL *DDLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDL.Contract.DDLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDL *DDLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDL.Contract.DDLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDL *DDLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DDL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDL *DDLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDL *DDLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDL.Contract.contract.Transact(opts, method, params...)
}

// COLLATERALDECIMALS is a free data retrieval call binding the contract method 0x4ddde78d.
//
// Solidity: function COLLATERAL_DECIMALS() view returns(uint256)
func (_DDL *DDLCaller) COLLATERALDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "COLLATERAL_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COLLATERALDECIMALS is a free data retrieval call binding the contract method 0x4ddde78d.
//
// Solidity: function COLLATERAL_DECIMALS() view returns(uint256)
func (_DDL *DDLSession) COLLATERALDECIMALS() (*big.Int, error) {
	return _DDL.Contract.COLLATERALDECIMALS(&_DDL.CallOpts)
}

// COLLATERALDECIMALS is a free data retrieval call binding the contract method 0x4ddde78d.
//
// Solidity: function COLLATERAL_DECIMALS() view returns(uint256)
func (_DDL *DDLCallerSession) COLLATERALDECIMALS() (*big.Int, error) {
	return _DDL.Contract.COLLATERALDECIMALS(&_DDL.CallOpts)
}

// INTERESTRATEDECIMALS is a free data retrieval call binding the contract method 0x66234ffa.
//
// Solidity: function INTEREST_RATE_DECIMALS() view returns(uint256)
func (_DDL *DDLCaller) INTERESTRATEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "INTEREST_RATE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INTERESTRATEDECIMALS is a free data retrieval call binding the contract method 0x66234ffa.
//
// Solidity: function INTEREST_RATE_DECIMALS() view returns(uint256)
func (_DDL *DDLSession) INTERESTRATEDECIMALS() (*big.Int, error) {
	return _DDL.Contract.INTERESTRATEDECIMALS(&_DDL.CallOpts)
}

// INTERESTRATEDECIMALS is a free data retrieval call binding the contract method 0x66234ffa.
//
// Solidity: function INTEREST_RATE_DECIMALS() view returns(uint256)
func (_DDL *DDLCallerSession) INTERESTRATEDECIMALS() (*big.Int, error) {
	return _DDL.Contract.INTERESTRATEDECIMALS(&_DDL.CallOpts)
}

// LTV is a free data retrieval call binding the contract method 0x62965d8e.
//
// Solidity: function LTV() view returns(uint256)
func (_DDL *DDLCaller) LTV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "LTV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LTV is a free data retrieval call binding the contract method 0x62965d8e.
//
// Solidity: function LTV() view returns(uint256)
func (_DDL *DDLSession) LTV() (*big.Int, error) {
	return _DDL.Contract.LTV(&_DDL.CallOpts)
}

// LTV is a free data retrieval call binding the contract method 0x62965d8e.
//
// Solidity: function LTV() view returns(uint256)
func (_DDL *DDLCallerSession) LTV() (*big.Int, error) {
	return _DDL.Contract.LTV(&_DDL.CallOpts)
}

// LTVDECIMALS is a free data retrieval call binding the contract method 0x9d38d064.
//
// Solidity: function LTV_DECIMALS() view returns(uint256)
func (_DDL *DDLCaller) LTVDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "LTV_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LTVDECIMALS is a free data retrieval call binding the contract method 0x9d38d064.
//
// Solidity: function LTV_DECIMALS() view returns(uint256)
func (_DDL *DDLSession) LTVDECIMALS() (*big.Int, error) {
	return _DDL.Contract.LTVDECIMALS(&_DDL.CallOpts)
}

// LTVDECIMALS is a free data retrieval call binding the contract method 0x9d38d064.
//
// Solidity: function LTV_DECIMALS() view returns(uint256)
func (_DDL *DDLCallerSession) LTVDECIMALS() (*big.Int, error) {
	return _DDL.Contract.LTVDECIMALS(&_DDL.CallOpts)
}

// PriorLiqPriceCoef is a free data retrieval call binding the contract method 0x425378b4.
//
// Solidity: function PriorLiqPriceCoef() view returns(uint256)
func (_DDL *DDLCaller) PriorLiqPriceCoef(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "PriorLiqPriceCoef")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriorLiqPriceCoef is a free data retrieval call binding the contract method 0x425378b4.
//
// Solidity: function PriorLiqPriceCoef() view returns(uint256)
func (_DDL *DDLSession) PriorLiqPriceCoef() (*big.Int, error) {
	return _DDL.Contract.PriorLiqPriceCoef(&_DDL.CallOpts)
}

// PriorLiqPriceCoef is a free data retrieval call binding the contract method 0x425378b4.
//
// Solidity: function PriorLiqPriceCoef() view returns(uint256)
func (_DDL *DDLCallerSession) PriorLiqPriceCoef() (*big.Int, error) {
	return _DDL.Contract.PriorLiqPriceCoef(&_DDL.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_DDL *DDLCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_DDL *DDLSession) USDC() (common.Address, error) {
	return _DDL.Contract.USDC(&_DDL.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_DDL *DDLCallerSession) USDC() (common.Address, error) {
	return _DDL.Contract.USDC(&_DDL.CallOpts)
}

// BorrowedByOption is a free data retrieval call binding the contract method 0x6c016c00.
//
// Solidity: function borrowedByOption(uint256 ) view returns(uint256 borrowed, uint256 newBorrowTimestamp)
func (_DDL *DDLCaller) BorrowedByOption(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Borrowed           *big.Int
	NewBorrowTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "borrowedByOption", arg0)

	outstruct := new(struct {
		Borrowed           *big.Int
		NewBorrowTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Borrowed = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NewBorrowTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BorrowedByOption is a free data retrieval call binding the contract method 0x6c016c00.
//
// Solidity: function borrowedByOption(uint256 ) view returns(uint256 borrowed, uint256 newBorrowTimestamp)
func (_DDL *DDLSession) BorrowedByOption(arg0 *big.Int) (struct {
	Borrowed           *big.Int
	NewBorrowTimestamp *big.Int
}, error) {
	return _DDL.Contract.BorrowedByOption(&_DDL.CallOpts, arg0)
}

// BorrowedByOption is a free data retrieval call binding the contract method 0x6c016c00.
//
// Solidity: function borrowedByOption(uint256 ) view returns(uint256 borrowed, uint256 newBorrowTimestamp)
func (_DDL *DDLCallerSession) BorrowedByOption(arg0 *big.Int) (struct {
	Borrowed           *big.Int
	NewBorrowTimestamp *big.Int
}, error) {
	return _DDL.Contract.BorrowedByOption(&_DDL.CallOpts, arg0)
}

// CalculateUpcomingFee is a free data retrieval call binding the contract method 0x484f1371.
//
// Solidity: function calculateUpcomingFee(uint256 id) view returns(uint256 upcomingFee)
func (_DDL *DDLCaller) CalculateUpcomingFee(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "calculateUpcomingFee", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateUpcomingFee is a free data retrieval call binding the contract method 0x484f1371.
//
// Solidity: function calculateUpcomingFee(uint256 id) view returns(uint256 upcomingFee)
func (_DDL *DDLSession) CalculateUpcomingFee(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.CalculateUpcomingFee(&_DDL.CallOpts, id)
}

// CalculateUpcomingFee is a free data retrieval call binding the contract method 0x484f1371.
//
// Solidity: function calculateUpcomingFee(uint256 id) view returns(uint256 upcomingFee)
func (_DDL *DDLCallerSession) CalculateUpcomingFee(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.CalculateUpcomingFee(&_DDL.CallOpts, id)
}

// CollateralInfo is a free data retrieval call binding the contract method 0x24a6665e.
//
// Solidity: function collateralInfo(uint256 ) view returns(address owner, (address,uint256,uint256,uint256,bool) strategy)
func (_DDL *DDLCaller) CollateralInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner    common.Address
	Strategy DDLoptionInfo
}, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "collateralInfo", arg0)

	outstruct := new(struct {
		Owner    common.Address
		Strategy DDLoptionInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Strategy = *abi.ConvertType(out[1], new(DDLoptionInfo)).(*DDLoptionInfo)

	return *outstruct, err

}

// CollateralInfo is a free data retrieval call binding the contract method 0x24a6665e.
//
// Solidity: function collateralInfo(uint256 ) view returns(address owner, (address,uint256,uint256,uint256,bool) strategy)
func (_DDL *DDLSession) CollateralInfo(arg0 *big.Int) (struct {
	Owner    common.Address
	Strategy DDLoptionInfo
}, error) {
	return _DDL.Contract.CollateralInfo(&_DDL.CallOpts, arg0)
}

// CollateralInfo is a free data retrieval call binding the contract method 0x24a6665e.
//
// Solidity: function collateralInfo(uint256 ) view returns(address owner, (address,uint256,uint256,uint256,bool) strategy)
func (_DDL *DDLCallerSession) CollateralInfo(arg0 *big.Int) (struct {
	Owner    common.Address
	Strategy DDLoptionInfo
}, error) {
	return _DDL.Contract.CollateralInfo(&_DDL.CallOpts, arg0)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_DDL *DDLCaller) CollateralToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "collateralToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_DDL *DDLSession) CollateralToken() (common.Address, error) {
	return _DDL.Contract.CollateralToken(&_DDL.CallOpts)
}

// CollateralToken is a free data retrieval call binding the contract method 0xb2016bd4.
//
// Solidity: function collateralToken() view returns(address)
func (_DDL *DDLCallerSession) CollateralToken() (common.Address, error) {
	return _DDL.Contract.CollateralToken(&_DDL.CallOpts)
}

// CurrentLiqPrice is a free data retrieval call binding the contract method 0x2cb2f334.
//
// Solidity: function currentLiqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCaller) CurrentLiqPrice(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "currentLiqPrice", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentLiqPrice is a free data retrieval call binding the contract method 0x2cb2f334.
//
// Solidity: function currentLiqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLSession) CurrentLiqPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.CurrentLiqPrice(&_DDL.CallOpts, id)
}

// CurrentLiqPrice is a free data retrieval call binding the contract method 0x2cb2f334.
//
// Solidity: function currentLiqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCallerSession) CurrentLiqPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.CurrentLiqPrice(&_DDL.CallOpts, id)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x7a3c4c17.
//
// Solidity: function currentPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCaller) CurrentPrice(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "currentPrice", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x7a3c4c17.
//
// Solidity: function currentPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLSession) CurrentPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.CurrentPrice(&_DDL.CallOpts, id)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x7a3c4c17.
//
// Solidity: function currentPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCallerSession) CurrentPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.CurrentPrice(&_DDL.CallOpts, id)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_DDL *DDLCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_DDL *DDLSession) InterestRate() (*big.Int, error) {
	return _DDL.Contract.InterestRate(&_DDL.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_DDL *DDLCallerSession) InterestRate() (*big.Int, error) {
	return _DDL.Contract.InterestRate(&_DDL.CallOpts)
}

// LiqPrice is a free data retrieval call binding the contract method 0xd88c6910.
//
// Solidity: function liqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCaller) LiqPrice(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "liqPrice", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiqPrice is a free data retrieval call binding the contract method 0xd88c6910.
//
// Solidity: function liqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLSession) LiqPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.LiqPrice(&_DDL.CallOpts, id)
}

// LiqPrice is a free data retrieval call binding the contract method 0xd88c6910.
//
// Solidity: function liqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCallerSession) LiqPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.LiqPrice(&_DDL.CallOpts, id)
}

// LoanState is a free data retrieval call binding the contract method 0x49b274ed.
//
// Solidity: function loanState(uint256 id) view returns(bool)
func (_DDL *DDLCaller) LoanState(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "loanState", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LoanState is a free data retrieval call binding the contract method 0x49b274ed.
//
// Solidity: function loanState(uint256 id) view returns(bool)
func (_DDL *DDLSession) LoanState(id *big.Int) (bool, error) {
	return _DDL.Contract.LoanState(&_DDL.CallOpts, id)
}

// LoanState is a free data retrieval call binding the contract method 0x49b274ed.
//
// Solidity: function loanState(uint256 id) view returns(bool)
func (_DDL *DDLCallerSession) LoanState(id *big.Int) (bool, error) {
	return _DDL.Contract.LoanState(&_DDL.CallOpts, id)
}

// LoanStateByPriorLiqPrice is a free data retrieval call binding the contract method 0x365c7504.
//
// Solidity: function loanStateByPriorLiqPrice(uint256 id) view returns(bool)
func (_DDL *DDLCaller) LoanStateByPriorLiqPrice(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "loanStateByPriorLiqPrice", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LoanStateByPriorLiqPrice is a free data retrieval call binding the contract method 0x365c7504.
//
// Solidity: function loanStateByPriorLiqPrice(uint256 id) view returns(bool)
func (_DDL *DDLSession) LoanStateByPriorLiqPrice(id *big.Int) (bool, error) {
	return _DDL.Contract.LoanStateByPriorLiqPrice(&_DDL.CallOpts, id)
}

// LoanStateByPriorLiqPrice is a free data retrieval call binding the contract method 0x365c7504.
//
// Solidity: function loanStateByPriorLiqPrice(uint256 id) view returns(bool)
func (_DDL *DDLCallerSession) LoanStateByPriorLiqPrice(id *big.Int) (bool, error) {
	return _DDL.Contract.LoanStateByPriorLiqPrice(&_DDL.CallOpts, id)
}

// MaxBorrowLimit is a free data retrieval call binding the contract method 0x9e3aa5b1.
//
// Solidity: function maxBorrowLimit(uint256 id) view returns(uint256)
func (_DDL *DDLCaller) MaxBorrowLimit(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "maxBorrowLimit", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBorrowLimit is a free data retrieval call binding the contract method 0x9e3aa5b1.
//
// Solidity: function maxBorrowLimit(uint256 id) view returns(uint256)
func (_DDL *DDLSession) MaxBorrowLimit(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.MaxBorrowLimit(&_DDL.CallOpts, id)
}

// MaxBorrowLimit is a free data retrieval call binding the contract method 0x9e3aa5b1.
//
// Solidity: function maxBorrowLimit(uint256 id) view returns(uint256)
func (_DDL *DDLCallerSession) MaxBorrowLimit(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.MaxBorrowLimit(&_DDL.CallOpts, id)
}

// MinBorrowLimit is a free data retrieval call binding the contract method 0xe260d747.
//
// Solidity: function minBorrowLimit() view returns(uint256)
func (_DDL *DDLCaller) MinBorrowLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "minBorrowLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBorrowLimit is a free data retrieval call binding the contract method 0xe260d747.
//
// Solidity: function minBorrowLimit() view returns(uint256)
func (_DDL *DDLSession) MinBorrowLimit() (*big.Int, error) {
	return _DDL.Contract.MinBorrowLimit(&_DDL.CallOpts)
}

// MinBorrowLimit is a free data retrieval call binding the contract method 0xe260d747.
//
// Solidity: function minBorrowLimit() view returns(uint256)
func (_DDL *DDLCallerSession) MinBorrowLimit() (*big.Int, error) {
	return _DDL.Contract.MinBorrowLimit(&_DDL.CallOpts)
}

// OperationalPool is a free data retrieval call binding the contract method 0x91319872.
//
// Solidity: function operationalPool() view returns(address)
func (_DDL *DDLCaller) OperationalPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "operationalPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperationalPool is a free data retrieval call binding the contract method 0x91319872.
//
// Solidity: function operationalPool() view returns(address)
func (_DDL *DDLSession) OperationalPool() (common.Address, error) {
	return _DDL.Contract.OperationalPool(&_DDL.CallOpts)
}

// OperationalPool is a free data retrieval call binding the contract method 0x91319872.
//
// Solidity: function operationalPool() view returns(address)
func (_DDL *DDLCallerSession) OperationalPool() (common.Address, error) {
	return _DDL.Contract.OperationalPool(&_DDL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDL *DDLCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDL *DDLSession) Owner() (common.Address, error) {
	return _DDL.Contract.Owner(&_DDL.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDL *DDLCallerSession) Owner() (common.Address, error) {
	return _DDL.Contract.Owner(&_DDL.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_DDL *DDLCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_DDL *DDLSession) Pool() (common.Address, error) {
	return _DDL.Contract.Pool(&_DDL.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_DDL *DDLCallerSession) Pool() (common.Address, error) {
	return _DDL.Contract.Pool(&_DDL.CallOpts)
}

// PriorLiqPrice is a free data retrieval call binding the contract method 0xf05735db.
//
// Solidity: function priorLiqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCaller) PriorLiqPrice(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "priorLiqPrice", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriorLiqPrice is a free data retrieval call binding the contract method 0xf05735db.
//
// Solidity: function priorLiqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLSession) PriorLiqPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.PriorLiqPrice(&_DDL.CallOpts, id)
}

// PriorLiqPrice is a free data retrieval call binding the contract method 0xf05735db.
//
// Solidity: function priorLiqPrice(uint256 id) view returns(uint256 price)
func (_DDL *DDLCallerSession) PriorLiqPrice(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.PriorLiqPrice(&_DDL.CallOpts, id)
}

// ProfitByOption is a free data retrieval call binding the contract method 0x0833a6c6.
//
// Solidity: function profitByOption(uint256 id) view returns(uint256 profit)
func (_DDL *DDLCaller) ProfitByOption(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "profitByOption", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitByOption is a free data retrieval call binding the contract method 0x0833a6c6.
//
// Solidity: function profitByOption(uint256 id) view returns(uint256 profit)
func (_DDL *DDLSession) ProfitByOption(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.ProfitByOption(&_DDL.CallOpts, id)
}

// ProfitByOption is a free data retrieval call binding the contract method 0x0833a6c6.
//
// Solidity: function profitByOption(uint256 id) view returns(uint256 profit)
func (_DDL *DDLCallerSession) ProfitByOption(id *big.Int) (*big.Int, error) {
	return _DDL.Contract.ProfitByOption(&_DDL.CallOpts, id)
}

// StrategyType is a free data retrieval call binding the contract method 0x357bcef0.
//
// Solidity: function strategyType(address ) view returns(uint8)
func (_DDL *DDLCaller) StrategyType(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _DDL.contract.Call(opts, &out, "strategyType", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StrategyType is a free data retrieval call binding the contract method 0x357bcef0.
//
// Solidity: function strategyType(address ) view returns(uint8)
func (_DDL *DDLSession) StrategyType(arg0 common.Address) (uint8, error) {
	return _DDL.Contract.StrategyType(&_DDL.CallOpts, arg0)
}

// StrategyType is a free data retrieval call binding the contract method 0x357bcef0.
//
// Solidity: function strategyType(address ) view returns(uint8)
func (_DDL *DDLCallerSession) StrategyType(arg0 common.Address) (uint8, error) {
	return _DDL.Contract.StrategyType(&_DDL.CallOpts, arg0)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 id, uint256 amount) returns()
func (_DDL *DDLTransactor) Borrow(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "borrow", id, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 id, uint256 amount) returns()
func (_DDL *DDLSession) Borrow(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Borrow(&_DDL.TransactOpts, id, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 id, uint256 amount) returns()
func (_DDL *DDLTransactorSession) Borrow(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Borrow(&_DDL.TransactOpts, id, amount)
}

// ExerciseByPriorLiqPrice is a paid mutator transaction binding the contract method 0x333daaa0.
//
// Solidity: function exerciseByPriorLiqPrice(uint256 id) returns()
func (_DDL *DDLTransactor) ExerciseByPriorLiqPrice(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "exerciseByPriorLiqPrice", id)
}

// ExerciseByPriorLiqPrice is a paid mutator transaction binding the contract method 0x333daaa0.
//
// Solidity: function exerciseByPriorLiqPrice(uint256 id) returns()
func (_DDL *DDLSession) ExerciseByPriorLiqPrice(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.ExerciseByPriorLiqPrice(&_DDL.TransactOpts, id)
}

// ExerciseByPriorLiqPrice is a paid mutator transaction binding the contract method 0x333daaa0.
//
// Solidity: function exerciseByPriorLiqPrice(uint256 id) returns()
func (_DDL *DDLTransactorSession) ExerciseByPriorLiqPrice(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.ExerciseByPriorLiqPrice(&_DDL.TransactOpts, id)
}

// ForcedExercise is a paid mutator transaction binding the contract method 0x01d73641.
//
// Solidity: function forcedExercise(uint256 id) returns()
func (_DDL *DDLTransactor) ForcedExercise(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "forcedExercise", id)
}

// ForcedExercise is a paid mutator transaction binding the contract method 0x01d73641.
//
// Solidity: function forcedExercise(uint256 id) returns()
func (_DDL *DDLSession) ForcedExercise(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.ForcedExercise(&_DDL.TransactOpts, id)
}

// ForcedExercise is a paid mutator transaction binding the contract method 0x01d73641.
//
// Solidity: function forcedExercise(uint256 id) returns()
func (_DDL *DDLTransactorSession) ForcedExercise(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.ForcedExercise(&_DDL.TransactOpts, id)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 id) returns()
func (_DDL *DDLTransactor) Liquidate(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "liquidate", id)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 id) returns()
func (_DDL *DDLSession) Liquidate(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Liquidate(&_DDL.TransactOpts, id)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 id) returns()
func (_DDL *DDLTransactorSession) Liquidate(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Liquidate(&_DDL.TransactOpts, id)
}

// LockOption is a paid mutator transaction binding the contract method 0x8ac2edc1.
//
// Solidity: function lockOption(uint256 id) returns()
func (_DDL *DDLTransactor) LockOption(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "lockOption", id)
}

// LockOption is a paid mutator transaction binding the contract method 0x8ac2edc1.
//
// Solidity: function lockOption(uint256 id) returns()
func (_DDL *DDLSession) LockOption(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.LockOption(&_DDL.TransactOpts, id)
}

// LockOption is a paid mutator transaction binding the contract method 0x8ac2edc1.
//
// Solidity: function lockOption(uint256 id) returns()
func (_DDL *DDLTransactorSession) LockOption(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.LockOption(&_DDL.TransactOpts, id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDL *DDLTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDL *DDLSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDL.Contract.RenounceOwnership(&_DDL.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDL *DDLTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDL.Contract.RenounceOwnership(&_DDL.TransactOpts)
}

// Repay is a paid mutator transaction binding the contract method 0xd8aed145.
//
// Solidity: function repay(uint256 id, uint256 amount) returns()
func (_DDL *DDLTransactor) Repay(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "repay", id, amount)
}

// Repay is a paid mutator transaction binding the contract method 0xd8aed145.
//
// Solidity: function repay(uint256 id, uint256 amount) returns()
func (_DDL *DDLSession) Repay(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Repay(&_DDL.TransactOpts, id, amount)
}

// Repay is a paid mutator transaction binding the contract method 0xd8aed145.
//
// Solidity: function repay(uint256 id, uint256 amount) returns()
func (_DDL *DDLTransactorSession) Repay(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Repay(&_DDL.TransactOpts, id, amount)
}

// SetInterestRate is a paid mutator transaction binding the contract method 0x5f84f302.
//
// Solidity: function setInterestRate(uint256 value) returns()
func (_DDL *DDLTransactor) SetInterestRate(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "setInterestRate", value)
}

// SetInterestRate is a paid mutator transaction binding the contract method 0x5f84f302.
//
// Solidity: function setInterestRate(uint256 value) returns()
func (_DDL *DDLSession) SetInterestRate(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetInterestRate(&_DDL.TransactOpts, value)
}

// SetInterestRate is a paid mutator transaction binding the contract method 0x5f84f302.
//
// Solidity: function setInterestRate(uint256 value) returns()
func (_DDL *DDLTransactorSession) SetInterestRate(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetInterestRate(&_DDL.TransactOpts, value)
}

// SetInterestRateDecimals is a paid mutator transaction binding the contract method 0x587a03c2.
//
// Solidity: function setInterestRateDecimals(uint256 value) returns()
func (_DDL *DDLTransactor) SetInterestRateDecimals(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "setInterestRateDecimals", value)
}

// SetInterestRateDecimals is a paid mutator transaction binding the contract method 0x587a03c2.
//
// Solidity: function setInterestRateDecimals(uint256 value) returns()
func (_DDL *DDLSession) SetInterestRateDecimals(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetInterestRateDecimals(&_DDL.TransactOpts, value)
}

// SetInterestRateDecimals is a paid mutator transaction binding the contract method 0x587a03c2.
//
// Solidity: function setInterestRateDecimals(uint256 value) returns()
func (_DDL *DDLTransactorSession) SetInterestRateDecimals(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetInterestRateDecimals(&_DDL.TransactOpts, value)
}

// SetLTV is a paid mutator transaction binding the contract method 0xb13de266.
//
// Solidity: function setLTV(uint256 value) returns()
func (_DDL *DDLTransactor) SetLTV(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "setLTV", value)
}

// SetLTV is a paid mutator transaction binding the contract method 0xb13de266.
//
// Solidity: function setLTV(uint256 value) returns()
func (_DDL *DDLSession) SetLTV(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetLTV(&_DDL.TransactOpts, value)
}

// SetLTV is a paid mutator transaction binding the contract method 0xb13de266.
//
// Solidity: function setLTV(uint256 value) returns()
func (_DDL *DDLTransactorSession) SetLTV(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetLTV(&_DDL.TransactOpts, value)
}

// SetLTVDecimals is a paid mutator transaction binding the contract method 0xcb46ed97.
//
// Solidity: function setLTVDecimals(uint256 value) returns()
func (_DDL *DDLTransactor) SetLTVDecimals(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "setLTVDecimals", value)
}

// SetLTVDecimals is a paid mutator transaction binding the contract method 0xcb46ed97.
//
// Solidity: function setLTVDecimals(uint256 value) returns()
func (_DDL *DDLSession) SetLTVDecimals(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetLTVDecimals(&_DDL.TransactOpts, value)
}

// SetLTVDecimals is a paid mutator transaction binding the contract method 0xcb46ed97.
//
// Solidity: function setLTVDecimals(uint256 value) returns()
func (_DDL *DDLTransactorSession) SetLTVDecimals(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetLTVDecimals(&_DDL.TransactOpts, value)
}

// SetMinBorrowLimit is a paid mutator transaction binding the contract method 0xc3707709.
//
// Solidity: function setMinBorrowLimit(uint256 value) returns()
func (_DDL *DDLTransactor) SetMinBorrowLimit(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "setMinBorrowLimit", value)
}

// SetMinBorrowLimit is a paid mutator transaction binding the contract method 0xc3707709.
//
// Solidity: function setMinBorrowLimit(uint256 value) returns()
func (_DDL *DDLSession) SetMinBorrowLimit(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetMinBorrowLimit(&_DDL.TransactOpts, value)
}

// SetMinBorrowLimit is a paid mutator transaction binding the contract method 0xc3707709.
//
// Solidity: function setMinBorrowLimit(uint256 value) returns()
func (_DDL *DDLTransactorSession) SetMinBorrowLimit(value *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.SetMinBorrowLimit(&_DDL.TransactOpts, value)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address value) returns()
func (_DDL *DDLTransactor) SetPool(opts *bind.TransactOpts, value common.Address) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "setPool", value)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address value) returns()
func (_DDL *DDLSession) SetPool(value common.Address) (*types.Transaction, error) {
	return _DDL.Contract.SetPool(&_DDL.TransactOpts, value)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address value) returns()
func (_DDL *DDLTransactorSession) SetPool(value common.Address) (*types.Transaction, error) {
	return _DDL.Contract.SetPool(&_DDL.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDL *DDLTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDL *DDLSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDL.Contract.TransferOwnership(&_DDL.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDL *DDLTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDL.Contract.TransferOwnership(&_DDL.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 id) returns()
func (_DDL *DDLTransactor) Unlock(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DDL.contract.Transact(opts, "unlock", id)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 id) returns()
func (_DDL *DDLSession) Unlock(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Unlock(&_DDL.TransactOpts, id)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 id) returns()
func (_DDL *DDLTransactorSession) Unlock(id *big.Int) (*types.Transaction, error) {
	return _DDL.Contract.Unlock(&_DDL.TransactOpts, id)
}

// DDLBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the DDL contract.
type DDLBorrowIterator struct {
	Event *DDLBorrow // Event containing the contract specifics and raw log

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
func (it *DDLBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLBorrow)
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
		it.Event = new(DDLBorrow)
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
func (it *DDLBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLBorrow represents a Borrow event raised by the DDL contract.
type DDLBorrow struct {
	User      common.Address
	OptionID  *big.Int
	Amount    *big.Int
	Strategy  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x1fdbee5d1bd2901710e8945918b1a052ce67662661a2ccbf68420e4e430f8dab.
//
// Solidity: event Borrow(address indexed user, uint256 indexed optionID, uint256 amount, address strategy, uint256 timestamp)
func (_DDL *DDLFilterer) FilterBorrow(opts *bind.FilterOpts, user []common.Address, optionID []*big.Int) (*DDLBorrowIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "Borrow", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return &DDLBorrowIterator{contract: _DDL.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x1fdbee5d1bd2901710e8945918b1a052ce67662661a2ccbf68420e4e430f8dab.
//
// Solidity: event Borrow(address indexed user, uint256 indexed optionID, uint256 amount, address strategy, uint256 timestamp)
func (_DDL *DDLFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *DDLBorrow, user []common.Address, optionID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "Borrow", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLBorrow)
				if err := _DDL.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x1fdbee5d1bd2901710e8945918b1a052ce67662661a2ccbf68420e4e430f8dab.
//
// Solidity: event Borrow(address indexed user, uint256 indexed optionID, uint256 amount, address strategy, uint256 timestamp)
func (_DDL *DDLFilterer) ParseBorrow(log types.Log) (*DDLBorrow, error) {
	event := new(DDLBorrow)
	if err := _DDL.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDLExerciseByPriorLiqPriceIterator is returned from FilterExerciseByPriorLiqPrice and is used to iterate over the raw logs and unpacked data for ExerciseByPriorLiqPrice events raised by the DDL contract.
type DDLExerciseByPriorLiqPriceIterator struct {
	Event *DDLExerciseByPriorLiqPrice // Event containing the contract specifics and raw log

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
func (it *DDLExerciseByPriorLiqPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLExerciseByPriorLiqPrice)
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
		it.Event = new(DDLExerciseByPriorLiqPrice)
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
func (it *DDLExerciseByPriorLiqPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLExerciseByPriorLiqPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLExerciseByPriorLiqPrice represents a ExerciseByPriorLiqPrice event raised by the DDL contract.
type DDLExerciseByPriorLiqPrice struct {
	User       common.Address
	OptionID   *big.Int
	UserReturn *big.Int
	PoolReturn *big.Int
	LiqFee     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExerciseByPriorLiqPrice is a free log retrieval operation binding the contract event 0xf28455920a5bcb3aeff7a6bbdd4fed497e6f159330fe161026128b5c381bdd90.
//
// Solidity: event ExerciseByPriorLiqPrice(address indexed user, uint256 indexed optionID, uint256 userReturn, uint256 poolReturn, uint256 liqFee)
func (_DDL *DDLFilterer) FilterExerciseByPriorLiqPrice(opts *bind.FilterOpts, user []common.Address, optionID []*big.Int) (*DDLExerciseByPriorLiqPriceIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "ExerciseByPriorLiqPrice", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return &DDLExerciseByPriorLiqPriceIterator{contract: _DDL.contract, event: "ExerciseByPriorLiqPrice", logs: logs, sub: sub}, nil
}

// WatchExerciseByPriorLiqPrice is a free log subscription operation binding the contract event 0xf28455920a5bcb3aeff7a6bbdd4fed497e6f159330fe161026128b5c381bdd90.
//
// Solidity: event ExerciseByPriorLiqPrice(address indexed user, uint256 indexed optionID, uint256 userReturn, uint256 poolReturn, uint256 liqFee)
func (_DDL *DDLFilterer) WatchExerciseByPriorLiqPrice(opts *bind.WatchOpts, sink chan<- *DDLExerciseByPriorLiqPrice, user []common.Address, optionID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "ExerciseByPriorLiqPrice", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLExerciseByPriorLiqPrice)
				if err := _DDL.contract.UnpackLog(event, "ExerciseByPriorLiqPrice", log); err != nil {
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

// ParseExerciseByPriorLiqPrice is a log parse operation binding the contract event 0xf28455920a5bcb3aeff7a6bbdd4fed497e6f159330fe161026128b5c381bdd90.
//
// Solidity: event ExerciseByPriorLiqPrice(address indexed user, uint256 indexed optionID, uint256 userReturn, uint256 poolReturn, uint256 liqFee)
func (_DDL *DDLFilterer) ParseExerciseByPriorLiqPrice(log types.Log) (*DDLExerciseByPriorLiqPrice, error) {
	event := new(DDLExerciseByPriorLiqPrice)
	if err := _DDL.contract.UnpackLog(event, "ExerciseByPriorLiqPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDLForcedExerciseIterator is returned from FilterForcedExercise and is used to iterate over the raw logs and unpacked data for ForcedExercise events raised by the DDL contract.
type DDLForcedExerciseIterator struct {
	Event *DDLForcedExercise // Event containing the contract specifics and raw log

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
func (it *DDLForcedExerciseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLForcedExercise)
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
		it.Event = new(DDLForcedExercise)
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
func (it *DDLForcedExerciseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLForcedExerciseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLForcedExercise represents a ForcedExercise event raised by the DDL contract.
type DDLForcedExercise struct {
	User       common.Address
	OptionID   *big.Int
	Amount     *big.Int
	PoolProfit *big.Int
	LiqFee     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterForcedExercise is a free log retrieval operation binding the contract event 0x7510f2e218f218a66429c54a712775c000c398117e38b240be7ef52310a51e82.
//
// Solidity: event ForcedExercise(address indexed user, uint256 indexed optionID, uint256 amount, uint256 poolProfit, uint256 liqFee)
func (_DDL *DDLFilterer) FilterForcedExercise(opts *bind.FilterOpts, user []common.Address, optionID []*big.Int) (*DDLForcedExerciseIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "ForcedExercise", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return &DDLForcedExerciseIterator{contract: _DDL.contract, event: "ForcedExercise", logs: logs, sub: sub}, nil
}

// WatchForcedExercise is a free log subscription operation binding the contract event 0x7510f2e218f218a66429c54a712775c000c398117e38b240be7ef52310a51e82.
//
// Solidity: event ForcedExercise(address indexed user, uint256 indexed optionID, uint256 amount, uint256 poolProfit, uint256 liqFee)
func (_DDL *DDLFilterer) WatchForcedExercise(opts *bind.WatchOpts, sink chan<- *DDLForcedExercise, user []common.Address, optionID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "ForcedExercise", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLForcedExercise)
				if err := _DDL.contract.UnpackLog(event, "ForcedExercise", log); err != nil {
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

// ParseForcedExercise is a log parse operation binding the contract event 0x7510f2e218f218a66429c54a712775c000c398117e38b240be7ef52310a51e82.
//
// Solidity: event ForcedExercise(address indexed user, uint256 indexed optionID, uint256 amount, uint256 poolProfit, uint256 liqFee)
func (_DDL *DDLFilterer) ParseForcedExercise(log types.Log) (*DDLForcedExercise, error) {
	event := new(DDLForcedExercise)
	if err := _DDL.contract.UnpackLog(event, "ForcedExercise", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDLLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the DDL contract.
type DDLLiquidateIterator struct {
	Event *DDLLiquidate // Event containing the contract specifics and raw log

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
func (it *DDLLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLLiquidate)
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
		it.Event = new(DDLLiquidate)
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
func (it *DDLLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLLiquidate represents a Liquidate event raised by the DDL contract.
type DDLLiquidate struct {
	User       common.Address
	OptionID   *big.Int
	Amount     *big.Int
	PoolProfit *big.Int
	LiqFee     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x4e91025e01b2329df1ec1067d27aafec4d1c41f682ccd794dee04321e0b1e8dc.
//
// Solidity: event Liquidate(address indexed user, uint256 indexed optionID, uint256 amount, uint256 poolProfit, uint256 liqFee)
func (_DDL *DDLFilterer) FilterLiquidate(opts *bind.FilterOpts, user []common.Address, optionID []*big.Int) (*DDLLiquidateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "Liquidate", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return &DDLLiquidateIterator{contract: _DDL.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x4e91025e01b2329df1ec1067d27aafec4d1c41f682ccd794dee04321e0b1e8dc.
//
// Solidity: event Liquidate(address indexed user, uint256 indexed optionID, uint256 amount, uint256 poolProfit, uint256 liqFee)
func (_DDL *DDLFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *DDLLiquidate, user []common.Address, optionID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "Liquidate", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLLiquidate)
				if err := _DDL.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x4e91025e01b2329df1ec1067d27aafec4d1c41f682ccd794dee04321e0b1e8dc.
//
// Solidity: event Liquidate(address indexed user, uint256 indexed optionID, uint256 amount, uint256 poolProfit, uint256 liqFee)
func (_DDL *DDLFilterer) ParseLiquidate(log types.Log) (*DDLLiquidate, error) {
	event := new(DDLLiquidate)
	if err := _DDL.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDLOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DDL contract.
type DDLOwnershipTransferredIterator struct {
	Event *DDLOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DDLOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLOwnershipTransferred)
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
		it.Event = new(DDLOwnershipTransferred)
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
func (it *DDLOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLOwnershipTransferred represents a OwnershipTransferred event raised by the DDL contract.
type DDLOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDL *DDLFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DDLOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DDLOwnershipTransferredIterator{contract: _DDL.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDL *DDLFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DDLOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLOwnershipTransferred)
				if err := _DDL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDL *DDLFilterer) ParseOwnershipTransferred(log types.Log) (*DDLOwnershipTransferred, error) {
	event := new(DDLOwnershipTransferred)
	if err := _DDL.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDLRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the DDL contract.
type DDLRepayIterator struct {
	Event *DDLRepay // Event containing the contract specifics and raw log

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
func (it *DDLRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLRepay)
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
		it.Event = new(DDLRepay)
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
func (it *DDLRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLRepay represents a Repay event raised by the DDL contract.
type DDLRepay struct {
	User     common.Address
	OptionID *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 indexed optionID, uint256 amount)
func (_DDL *DDLFilterer) FilterRepay(opts *bind.FilterOpts, user []common.Address, optionID []*big.Int) (*DDLRepayIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "Repay", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return &DDLRepayIterator{contract: _DDL.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 indexed optionID, uint256 amount)
func (_DDL *DDLFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *DDLRepay, user []common.Address, optionID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "Repay", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLRepay)
				if err := _DDL.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x77c6871227e5d2dec8dadd5354f78453203e22e669cd0ec4c19d9a8c5edb31d0.
//
// Solidity: event Repay(address indexed user, uint256 indexed optionID, uint256 amount)
func (_DDL *DDLFilterer) ParseRepay(log types.Log) (*DDLRepay, error) {
	event := new(DDLRepay)
	if err := _DDL.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDLUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the DDL contract.
type DDLUnlockIterator struct {
	Event *DDLUnlock // Event containing the contract specifics and raw log

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
func (it *DDLUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDLUnlock)
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
		it.Event = new(DDLUnlock)
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
func (it *DDLUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDLUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDLUnlock represents a Unlock event raised by the DDL contract.
type DDLUnlock struct {
	User     common.Address
	OptionID *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x6381d9813cabeb57471b5a7e05078e64845ccdb563146a6911d536f24ce960f1.
//
// Solidity: event Unlock(address indexed user, uint256 indexed optionID)
func (_DDL *DDLFilterer) FilterUnlock(opts *bind.FilterOpts, user []common.Address, optionID []*big.Int) (*DDLUnlockIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.FilterLogs(opts, "Unlock", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return &DDLUnlockIterator{contract: _DDL.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x6381d9813cabeb57471b5a7e05078e64845ccdb563146a6911d536f24ce960f1.
//
// Solidity: event Unlock(address indexed user, uint256 indexed optionID)
func (_DDL *DDLFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *DDLUnlock, user []common.Address, optionID []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIDRule []interface{}
	for _, optionIDItem := range optionID {
		optionIDRule = append(optionIDRule, optionIDItem)
	}

	logs, sub, err := _DDL.contract.WatchLogs(opts, "Unlock", userRule, optionIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDLUnlock)
				if err := _DDL.contract.UnpackLog(event, "Unlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0x6381d9813cabeb57471b5a7e05078e64845ccdb563146a6911d536f24ce960f1.
//
// Solidity: event Unlock(address indexed user, uint256 indexed optionID)
func (_DDL *DDLFilterer) ParseUnlock(log types.Log) (*DDLUnlock, error) {
	event := new(DDLUnlock)
	if err := _DDL.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

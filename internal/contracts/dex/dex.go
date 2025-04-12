// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package odx_dex

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
	_ = abi.ConvertType
)

// OdxDexMetaData contains all meta data concerning the OdxDex contract.
var OdxDexMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OdxDexABI is the input ABI used to generate the binding from.
// Deprecated: Use OdxDexMetaData.ABI instead.
var OdxDexABI = OdxDexMetaData.ABI

// OdxDex is an auto generated Go binding around an Ethereum contract.
type OdxDex struct {
	OdxDexCaller     // Read-only binding to the contract
	OdxDexTransactor // Write-only binding to the contract
	OdxDexFilterer   // Log filterer for contract events
}

// OdxDexCaller is an auto generated read-only Go binding around an Ethereum contract.
type OdxDexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdxDexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OdxDexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdxDexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OdxDexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdxDexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OdxDexSession struct {
	Contract     *OdxDex           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OdxDexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OdxDexCallerSession struct {
	Contract *OdxDexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OdxDexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OdxDexTransactorSession struct {
	Contract     *OdxDexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OdxDexRaw is an auto generated low-level Go binding around an Ethereum contract.
type OdxDexRaw struct {
	Contract *OdxDex // Generic contract binding to access the raw methods on
}

// OdxDexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OdxDexCallerRaw struct {
	Contract *OdxDexCaller // Generic read-only contract binding to access the raw methods on
}

// OdxDexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OdxDexTransactorRaw struct {
	Contract *OdxDexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOdxDex creates a new instance of OdxDex, bound to a specific deployed contract.
func NewOdxDex(address common.Address, backend bind.ContractBackend) (*OdxDex, error) {
	contract, err := bindOdxDex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OdxDex{OdxDexCaller: OdxDexCaller{contract: contract}, OdxDexTransactor: OdxDexTransactor{contract: contract}, OdxDexFilterer: OdxDexFilterer{contract: contract}}, nil
}

// NewOdxDexCaller creates a new read-only instance of OdxDex, bound to a specific deployed contract.
func NewOdxDexCaller(address common.Address, caller bind.ContractCaller) (*OdxDexCaller, error) {
	contract, err := bindOdxDex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OdxDexCaller{contract: contract}, nil
}

// NewOdxDexTransactor creates a new write-only instance of OdxDex, bound to a specific deployed contract.
func NewOdxDexTransactor(address common.Address, transactor bind.ContractTransactor) (*OdxDexTransactor, error) {
	contract, err := bindOdxDex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OdxDexTransactor{contract: contract}, nil
}

// NewOdxDexFilterer creates a new log filterer instance of OdxDex, bound to a specific deployed contract.
func NewOdxDexFilterer(address common.Address, filterer bind.ContractFilterer) (*OdxDexFilterer, error) {
	contract, err := bindOdxDex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OdxDexFilterer{contract: contract}, nil
}

// bindOdxDex binds a generic wrapper to an already deployed contract.
func bindOdxDex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OdxDexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OdxDex *OdxDexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OdxDex.Contract.OdxDexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OdxDex *OdxDexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdxDex.Contract.OdxDexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OdxDex *OdxDexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OdxDex.Contract.OdxDexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OdxDex *OdxDexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OdxDex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OdxDex *OdxDexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdxDex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OdxDex *OdxDexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OdxDex.Contract.contract.Transact(opts, method, params...)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256)
func (_OdxDex *OdxDexCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OdxDex.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256)
func (_OdxDex *OdxDexSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _OdxDex.Contract.GetAmountOut(&_OdxDex.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256)
func (_OdxDex *OdxDexCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _OdxDex.Contract.GetAmountOut(&_OdxDex.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_OdxDex *OdxDexCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _OdxDex.contract.Call(opts, &out, "getReserves", tokenA, tokenB)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_OdxDex *OdxDexSession) GetReserves(tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _OdxDex.Contract.GetReserves(&_OdxDex.CallOpts, tokenA, tokenB)
}

// GetReserves is a free data retrieval call binding the contract method 0xd52bb6f4.
//
// Solidity: function getReserves(address tokenA, address tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_OdxDex *OdxDexCallerSession) GetReserves(tokenA common.Address, tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _OdxDex.Contract.GetReserves(&_OdxDex.CallOpts, tokenA, tokenB)
}

// Reserves is a free data retrieval call binding the contract method 0xea4afc54.
//
// Solidity: function reserves(address , address ) view returns(uint256)
func (_OdxDex *OdxDexCaller) Reserves(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OdxDex.contract.Call(opts, &out, "reserves", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserves is a free data retrieval call binding the contract method 0xea4afc54.
//
// Solidity: function reserves(address , address ) view returns(uint256)
func (_OdxDex *OdxDexSession) Reserves(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _OdxDex.Contract.Reserves(&_OdxDex.CallOpts, arg0, arg1)
}

// Reserves is a free data retrieval call binding the contract method 0xea4afc54.
//
// Solidity: function reserves(address , address ) view returns(uint256)
func (_OdxDex *OdxDexCallerSession) Reserves(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _OdxDex.Contract.Reserves(&_OdxDex.CallOpts, arg0, arg1)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcf6c62ea.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountA, uint256 amountB) returns()
func (_OdxDex *OdxDexTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _OdxDex.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountA, amountB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcf6c62ea.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountA, uint256 amountB) returns()
func (_OdxDex *OdxDexSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _OdxDex.Contract.AddLiquidity(&_OdxDex.TransactOpts, tokenA, tokenB, amountA, amountB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcf6c62ea.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountA, uint256 amountB) returns()
func (_OdxDex *OdxDexTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _OdxDex.Contract.AddLiquidity(&_OdxDex.TransactOpts, tokenA, tokenB, amountA, amountB)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_OdxDex *OdxDexTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _OdxDex.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_OdxDex *OdxDexSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _OdxDex.Contract.SwapExactTokensForTokens(&_OdxDex.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_OdxDex *OdxDexTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _OdxDex.Contract.SwapExactTokensForTokens(&_OdxDex.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// OdxDexSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the OdxDex contract.
type OdxDexSwapIterator struct {
	Event *OdxDexSwap // Event containing the contract specifics and raw log

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
func (it *OdxDexSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OdxDexSwap)
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
		it.Event = new(OdxDexSwap)
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
func (it *OdxDexSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OdxDexSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OdxDexSwap represents a Swap event raised by the OdxDex contract.
type OdxDexSwap struct {
	Sender    common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	TokenIn   common.Address
	TokenOut  common.Address
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x1621cb2414b25cfb014ed2e1e8051310c0f691ac8d2ed92928e804595df0553b.
//
// Solidity: event Swap(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed tokenIn, address indexed tokenOut, address to)
func (_OdxDex *OdxDexFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*OdxDexSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _OdxDex.contract.FilterLogs(opts, "Swap", senderRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &OdxDexSwapIterator{contract: _OdxDex.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x1621cb2414b25cfb014ed2e1e8051310c0f691ac8d2ed92928e804595df0553b.
//
// Solidity: event Swap(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed tokenIn, address indexed tokenOut, address to)
func (_OdxDex *OdxDexFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *OdxDexSwap, sender []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _OdxDex.contract.WatchLogs(opts, "Swap", senderRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OdxDexSwap)
				if err := _OdxDex.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x1621cb2414b25cfb014ed2e1e8051310c0f691ac8d2ed92928e804595df0553b.
//
// Solidity: event Swap(address indexed sender, uint256 amountIn, uint256 amountOut, address indexed tokenIn, address indexed tokenOut, address to)
func (_OdxDex *OdxDexFilterer) ParseSwap(log types.Log) (*OdxDexSwap, error) {
	event := new(OdxDexSwap)
	if err := _OdxDex.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

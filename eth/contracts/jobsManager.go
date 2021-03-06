// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// JobsManagerABI is the input ABI used to generate the binding from.
const JobsManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_streamId\",\"type\":\"string\"},{\"name\":\"_transcodingOptions\",\"type\":\"string\"},{\"name\":\"_maxPricePerSegment\",\"type\":\"uint256\"}],\"name\":\"job\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"name\":\"jobId\",\"type\":\"uint256\"},{\"name\":\"streamId\",\"type\":\"string\"},{\"name\":\"transcodingOptions\",\"type\":\"string\"},{\"name\":\"maxPricePerSegment\",\"type\":\"uint256\"},{\"name\":\"broadcasterAddress\",\"type\":\"address\"},{\"name\":\"transcoderAddress\",\"type\":\"address\"},{\"name\":\"endBlock\",\"type\":\"uint256\"},{\"name\":\"escrow\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"},{\"name\":\"_segmentNumber\",\"type\":\"uint256\"},{\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"receiveVerification\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finderFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobStreamId\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"missedVerificationSlashAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobEscrow\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_segmentRange\",\"type\":\"uint256[2]\"},{\"name\":\"_claimRoot\",\"type\":\"bytes32\"}],\"name\":\"claimWork\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetContractId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimEndSegment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"},{\"name\":\"_segmentNumber\",\"type\":\"uint256\"},{\"name\":\"_dataStorageHash\",\"type\":\"string\"},{\"name\":\"_dataHashes\",\"type\":\"bytes32[2]\"},{\"name\":\"_broadcasterSig\",\"type\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobTranscodingOptions\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimTranscoderTotalStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"verificationRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"distributeFees\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimIds\",\"type\":\"uint256[]\"}],\"name\":\"batchDistributeFees\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimEndSlashingBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobTranscoderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobTotalClaims\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numJobs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobMaxPricePerSegment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimEndVerificationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"jobStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"verificationPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"failedVerificationSlashAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobEndBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"verificationFailureThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"},{\"name\":\"_segmentNumber\",\"type\":\"uint256\"}],\"name\":\"missedVerificationSlash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"getJobBroadcasterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\"}],\"name\":\"getClaimStartSegment\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"broadcasterDeposits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"jobEndingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_jobId\",\"type\":\"uint256\"}],\"name\":\"endJob\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_verificationRate\",\"type\":\"uint64\"},{\"name\":\"_jobEndingPeriod\",\"type\":\"uint256\"},{\"name\":\"_verificationPeriod\",\"type\":\"uint256\"},{\"name\":\"_slashingPeriod\",\"type\":\"uint256\"},{\"name\":\"_failedVerificationSlashAmount\",\"type\":\"uint64\"},{\"name\":\"_missedVerificationSlashAmount\",\"type\":\"uint64\"},{\"name\":\"_finderFee\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_controller\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"broadcaster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"streamId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"transcodingOptions\",\"type\":\"string\"}],\"name\":\"NewJob\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"claimId\",\"type\":\"uint256\"}],\"name\":\"NewClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"claimId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"segmentNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ReceivedVerification\",\"type\":\"event\"}]"

// JobsManager is an auto generated Go binding around an Ethereum contract.
type JobsManager struct {
	JobsManagerCaller     // Read-only binding to the contract
	JobsManagerTransactor // Write-only binding to the contract
}

// JobsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobsManagerSession struct {
	Contract     *JobsManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobsManagerCallerSession struct {
	Contract *JobsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// JobsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobsManagerTransactorSession struct {
	Contract     *JobsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// JobsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobsManagerRaw struct {
	Contract *JobsManager // Generic contract binding to access the raw methods on
}

// JobsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobsManagerCallerRaw struct {
	Contract *JobsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// JobsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobsManagerTransactorRaw struct {
	Contract *JobsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobsManager creates a new instance of JobsManager, bound to a specific deployed contract.
func NewJobsManager(address common.Address, backend bind.ContractBackend) (*JobsManager, error) {
	contract, err := bindJobsManager(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JobsManager{JobsManagerCaller: JobsManagerCaller{contract: contract}, JobsManagerTransactor: JobsManagerTransactor{contract: contract}}, nil
}

// NewJobsManagerCaller creates a new read-only instance of JobsManager, bound to a specific deployed contract.
func NewJobsManagerCaller(address common.Address, caller bind.ContractCaller) (*JobsManagerCaller, error) {
	contract, err := bindJobsManager(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &JobsManagerCaller{contract: contract}, nil
}

// NewJobsManagerTransactor creates a new write-only instance of JobsManager, bound to a specific deployed contract.
func NewJobsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*JobsManagerTransactor, error) {
	contract, err := bindJobsManager(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &JobsManagerTransactor{contract: contract}, nil
}

// bindJobsManager binds a generic wrapper to an already deployed contract.
func bindJobsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobsManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobsManager *JobsManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JobsManager.Contract.JobsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobsManager *JobsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobsManager.Contract.JobsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobsManager *JobsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobsManager.Contract.JobsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JobsManager *JobsManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _JobsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JobsManager *JobsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JobsManager *JobsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JobsManager.Contract.contract.Transact(opts, method, params...)
}

// BroadcasterDeposits is a free data retrieval call binding the contract method 0xf0b8720c.
//
// Solidity: function broadcasterDeposits( address) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) BroadcasterDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "broadcasterDeposits", arg0)
	return *ret0, err
}

// BroadcasterDeposits is a free data retrieval call binding the contract method 0xf0b8720c.
//
// Solidity: function broadcasterDeposits( address) constant returns(uint256)
func (_JobsManager *JobsManagerSession) BroadcasterDeposits(arg0 common.Address) (*big.Int, error) {
	return _JobsManager.Contract.BroadcasterDeposits(&_JobsManager.CallOpts, arg0)
}

// BroadcasterDeposits is a free data retrieval call binding the contract method 0xf0b8720c.
//
// Solidity: function broadcasterDeposits( address) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) BroadcasterDeposits(arg0 common.Address) (*big.Int, error) {
	return _JobsManager.Contract.BroadcasterDeposits(&_JobsManager.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_JobsManager *JobsManagerCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "controller")
	return *ret0, err
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_JobsManager *JobsManagerSession) Controller() (common.Address, error) {
	return _JobsManager.Contract.Controller(&_JobsManager.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() constant returns(address)
func (_JobsManager *JobsManagerCallerSession) Controller() (common.Address, error) {
	return _JobsManager.Contract.Controller(&_JobsManager.CallOpts)
}

// FailedVerificationSlashAmount is a free data retrieval call binding the contract method 0xbe5c2423.
//
// Solidity: function failedVerificationSlashAmount() constant returns(uint64)
func (_JobsManager *JobsManagerCaller) FailedVerificationSlashAmount(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "failedVerificationSlashAmount")
	return *ret0, err
}

// FailedVerificationSlashAmount is a free data retrieval call binding the contract method 0xbe5c2423.
//
// Solidity: function failedVerificationSlashAmount() constant returns(uint64)
func (_JobsManager *JobsManagerSession) FailedVerificationSlashAmount() (uint64, error) {
	return _JobsManager.Contract.FailedVerificationSlashAmount(&_JobsManager.CallOpts)
}

// FailedVerificationSlashAmount is a free data retrieval call binding the contract method 0xbe5c2423.
//
// Solidity: function failedVerificationSlashAmount() constant returns(uint64)
func (_JobsManager *JobsManagerCallerSession) FailedVerificationSlashAmount() (uint64, error) {
	return _JobsManager.Contract.FailedVerificationSlashAmount(&_JobsManager.CallOpts)
}

// FinderFee is a free data retrieval call binding the contract method 0x1e6b0e44.
//
// Solidity: function finderFee() constant returns(uint64)
func (_JobsManager *JobsManagerCaller) FinderFee(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "finderFee")
	return *ret0, err
}

// FinderFee is a free data retrieval call binding the contract method 0x1e6b0e44.
//
// Solidity: function finderFee() constant returns(uint64)
func (_JobsManager *JobsManagerSession) FinderFee() (uint64, error) {
	return _JobsManager.Contract.FinderFee(&_JobsManager.CallOpts)
}

// FinderFee is a free data retrieval call binding the contract method 0x1e6b0e44.
//
// Solidity: function finderFee() constant returns(uint64)
func (_JobsManager *JobsManagerCallerSession) FinderFee() (uint64, error) {
	return _JobsManager.Contract.FinderFee(&_JobsManager.CallOpts)
}

// GetClaimBlock is a free data retrieval call binding the contract method 0x5831fd65.
//
// Solidity: function getClaimBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetClaimBlock(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimBlock", _jobId, _claimId)
	return *ret0, err
}

// GetClaimBlock is a free data retrieval call binding the contract method 0x5831fd65.
//
// Solidity: function getClaimBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetClaimBlock(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimBlock(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimBlock is a free data retrieval call binding the contract method 0x5831fd65.
//
// Solidity: function getClaimBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetClaimBlock(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimBlock(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimEndSegment is a free data retrieval call binding the contract method 0x52bde16c.
//
// Solidity: function getClaimEndSegment(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetClaimEndSegment(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimEndSegment", _jobId, _claimId)
	return *ret0, err
}

// GetClaimEndSegment is a free data retrieval call binding the contract method 0x52bde16c.
//
// Solidity: function getClaimEndSegment(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetClaimEndSegment(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimEndSegment(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimEndSegment is a free data retrieval call binding the contract method 0x52bde16c.
//
// Solidity: function getClaimEndSegment(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetClaimEndSegment(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimEndSegment(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimEndSlashingBlock is a free data retrieval call binding the contract method 0x8d9805b7.
//
// Solidity: function getClaimEndSlashingBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetClaimEndSlashingBlock(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimEndSlashingBlock", _jobId, _claimId)
	return *ret0, err
}

// GetClaimEndSlashingBlock is a free data retrieval call binding the contract method 0x8d9805b7.
//
// Solidity: function getClaimEndSlashingBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetClaimEndSlashingBlock(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimEndSlashingBlock(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimEndSlashingBlock is a free data retrieval call binding the contract method 0x8d9805b7.
//
// Solidity: function getClaimEndSlashingBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetClaimEndSlashingBlock(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimEndSlashingBlock(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimEndVerificationBlock is a free data retrieval call binding the contract method 0xa74e1b17.
//
// Solidity: function getClaimEndVerificationBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetClaimEndVerificationBlock(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimEndVerificationBlock", _jobId, _claimId)
	return *ret0, err
}

// GetClaimEndVerificationBlock is a free data retrieval call binding the contract method 0xa74e1b17.
//
// Solidity: function getClaimEndVerificationBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetClaimEndVerificationBlock(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimEndVerificationBlock(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimEndVerificationBlock is a free data retrieval call binding the contract method 0xa74e1b17.
//
// Solidity: function getClaimEndVerificationBlock(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetClaimEndVerificationBlock(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimEndVerificationBlock(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimRoot is a free data retrieval call binding the contract method 0x94fd2229.
//
// Solidity: function getClaimRoot(_jobId uint256, _claimId uint256) constant returns(bytes32)
func (_JobsManager *JobsManagerCaller) GetClaimRoot(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimRoot", _jobId, _claimId)
	return *ret0, err
}

// GetClaimRoot is a free data retrieval call binding the contract method 0x94fd2229.
//
// Solidity: function getClaimRoot(_jobId uint256, _claimId uint256) constant returns(bytes32)
func (_JobsManager *JobsManagerSession) GetClaimRoot(_jobId *big.Int, _claimId *big.Int) ([32]byte, error) {
	return _JobsManager.Contract.GetClaimRoot(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimRoot is a free data retrieval call binding the contract method 0x94fd2229.
//
// Solidity: function getClaimRoot(_jobId uint256, _claimId uint256) constant returns(bytes32)
func (_JobsManager *JobsManagerCallerSession) GetClaimRoot(_jobId *big.Int, _claimId *big.Int) ([32]byte, error) {
	return _JobsManager.Contract.GetClaimRoot(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimStartSegment is a free data retrieval call binding the contract method 0xe6d24add.
//
// Solidity: function getClaimStartSegment(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetClaimStartSegment(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimStartSegment", _jobId, _claimId)
	return *ret0, err
}

// GetClaimStartSegment is a free data retrieval call binding the contract method 0xe6d24add.
//
// Solidity: function getClaimStartSegment(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetClaimStartSegment(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimStartSegment(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimStartSegment is a free data retrieval call binding the contract method 0xe6d24add.
//
// Solidity: function getClaimStartSegment(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetClaimStartSegment(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimStartSegment(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimStatus is a free data retrieval call binding the contract method 0x72d9b690.
//
// Solidity: function getClaimStatus(_jobId uint256, _claimId uint256) constant returns(uint8)
func (_JobsManager *JobsManagerCaller) GetClaimStatus(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimStatus", _jobId, _claimId)
	return *ret0, err
}

// GetClaimStatus is a free data retrieval call binding the contract method 0x72d9b690.
//
// Solidity: function getClaimStatus(_jobId uint256, _claimId uint256) constant returns(uint8)
func (_JobsManager *JobsManagerSession) GetClaimStatus(_jobId *big.Int, _claimId *big.Int) (uint8, error) {
	return _JobsManager.Contract.GetClaimStatus(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimStatus is a free data retrieval call binding the contract method 0x72d9b690.
//
// Solidity: function getClaimStatus(_jobId uint256, _claimId uint256) constant returns(uint8)
func (_JobsManager *JobsManagerCallerSession) GetClaimStatus(_jobId *big.Int, _claimId *big.Int) (uint8, error) {
	return _JobsManager.Contract.GetClaimStatus(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimTranscoderTotalStake is a free data retrieval call binding the contract method 0x709cd50c.
//
// Solidity: function getClaimTranscoderTotalStake(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetClaimTranscoderTotalStake(opts *bind.CallOpts, _jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getClaimTranscoderTotalStake", _jobId, _claimId)
	return *ret0, err
}

// GetClaimTranscoderTotalStake is a free data retrieval call binding the contract method 0x709cd50c.
//
// Solidity: function getClaimTranscoderTotalStake(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetClaimTranscoderTotalStake(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimTranscoderTotalStake(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetClaimTranscoderTotalStake is a free data retrieval call binding the contract method 0x709cd50c.
//
// Solidity: function getClaimTranscoderTotalStake(_jobId uint256, _claimId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetClaimTranscoderTotalStake(_jobId *big.Int, _claimId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetClaimTranscoderTotalStake(&_JobsManager.CallOpts, _jobId, _claimId)
}

// GetJobBroadcasterAddress is a free data retrieval call binding the contract method 0xd41637be.
//
// Solidity: function getJobBroadcasterAddress(_jobId uint256) constant returns(address)
func (_JobsManager *JobsManagerCaller) GetJobBroadcasterAddress(opts *bind.CallOpts, _jobId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobBroadcasterAddress", _jobId)
	return *ret0, err
}

// GetJobBroadcasterAddress is a free data retrieval call binding the contract method 0xd41637be.
//
// Solidity: function getJobBroadcasterAddress(_jobId uint256) constant returns(address)
func (_JobsManager *JobsManagerSession) GetJobBroadcasterAddress(_jobId *big.Int) (common.Address, error) {
	return _JobsManager.Contract.GetJobBroadcasterAddress(&_JobsManager.CallOpts, _jobId)
}

// GetJobBroadcasterAddress is a free data retrieval call binding the contract method 0xd41637be.
//
// Solidity: function getJobBroadcasterAddress(_jobId uint256) constant returns(address)
func (_JobsManager *JobsManagerCallerSession) GetJobBroadcasterAddress(_jobId *big.Int) (common.Address, error) {
	return _JobsManager.Contract.GetJobBroadcasterAddress(&_JobsManager.CallOpts, _jobId)
}

// GetJobEndBlock is a free data retrieval call binding the contract method 0xbf5d2ccb.
//
// Solidity: function getJobEndBlock(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetJobEndBlock(opts *bind.CallOpts, _jobId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobEndBlock", _jobId)
	return *ret0, err
}

// GetJobEndBlock is a free data retrieval call binding the contract method 0xbf5d2ccb.
//
// Solidity: function getJobEndBlock(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetJobEndBlock(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobEndBlock(&_JobsManager.CallOpts, _jobId)
}

// GetJobEndBlock is a free data retrieval call binding the contract method 0xbf5d2ccb.
//
// Solidity: function getJobEndBlock(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetJobEndBlock(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobEndBlock(&_JobsManager.CallOpts, _jobId)
}

// GetJobEscrow is a free data retrieval call binding the contract method 0x3561c340.
//
// Solidity: function getJobEscrow(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetJobEscrow(opts *bind.CallOpts, _jobId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobEscrow", _jobId)
	return *ret0, err
}

// GetJobEscrow is a free data retrieval call binding the contract method 0x3561c340.
//
// Solidity: function getJobEscrow(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetJobEscrow(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobEscrow(&_JobsManager.CallOpts, _jobId)
}

// GetJobEscrow is a free data retrieval call binding the contract method 0x3561c340.
//
// Solidity: function getJobEscrow(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetJobEscrow(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobEscrow(&_JobsManager.CallOpts, _jobId)
}

// GetJobMaxPricePerSegment is a free data retrieval call binding the contract method 0x9eba91d5.
//
// Solidity: function getJobMaxPricePerSegment(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetJobMaxPricePerSegment(opts *bind.CallOpts, _jobId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobMaxPricePerSegment", _jobId)
	return *ret0, err
}

// GetJobMaxPricePerSegment is a free data retrieval call binding the contract method 0x9eba91d5.
//
// Solidity: function getJobMaxPricePerSegment(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetJobMaxPricePerSegment(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobMaxPricePerSegment(&_JobsManager.CallOpts, _jobId)
}

// GetJobMaxPricePerSegment is a free data retrieval call binding the contract method 0x9eba91d5.
//
// Solidity: function getJobMaxPricePerSegment(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetJobMaxPricePerSegment(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobMaxPricePerSegment(&_JobsManager.CallOpts, _jobId)
}

// GetJobStreamId is a free data retrieval call binding the contract method 0x2929140a.
//
// Solidity: function getJobStreamId(_jobId uint256) constant returns(string)
func (_JobsManager *JobsManagerCaller) GetJobStreamId(opts *bind.CallOpts, _jobId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobStreamId", _jobId)
	return *ret0, err
}

// GetJobStreamId is a free data retrieval call binding the contract method 0x2929140a.
//
// Solidity: function getJobStreamId(_jobId uint256) constant returns(string)
func (_JobsManager *JobsManagerSession) GetJobStreamId(_jobId *big.Int) (string, error) {
	return _JobsManager.Contract.GetJobStreamId(&_JobsManager.CallOpts, _jobId)
}

// GetJobStreamId is a free data retrieval call binding the contract method 0x2929140a.
//
// Solidity: function getJobStreamId(_jobId uint256) constant returns(string)
func (_JobsManager *JobsManagerCallerSession) GetJobStreamId(_jobId *big.Int) (string, error) {
	return _JobsManager.Contract.GetJobStreamId(&_JobsManager.CallOpts, _jobId)
}

// GetJobTotalClaims is a free data retrieval call binding the contract method 0x8fafceac.
//
// Solidity: function getJobTotalClaims(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCaller) GetJobTotalClaims(opts *bind.CallOpts, _jobId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobTotalClaims", _jobId)
	return *ret0, err
}

// GetJobTotalClaims is a free data retrieval call binding the contract method 0x8fafceac.
//
// Solidity: function getJobTotalClaims(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerSession) GetJobTotalClaims(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobTotalClaims(&_JobsManager.CallOpts, _jobId)
}

// GetJobTotalClaims is a free data retrieval call binding the contract method 0x8fafceac.
//
// Solidity: function getJobTotalClaims(_jobId uint256) constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) GetJobTotalClaims(_jobId *big.Int) (*big.Int, error) {
	return _JobsManager.Contract.GetJobTotalClaims(&_JobsManager.CallOpts, _jobId)
}

// GetJobTranscoderAddress is a free data retrieval call binding the contract method 0x8d980fe4.
//
// Solidity: function getJobTranscoderAddress(_jobId uint256) constant returns(address)
func (_JobsManager *JobsManagerCaller) GetJobTranscoderAddress(opts *bind.CallOpts, _jobId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobTranscoderAddress", _jobId)
	return *ret0, err
}

// GetJobTranscoderAddress is a free data retrieval call binding the contract method 0x8d980fe4.
//
// Solidity: function getJobTranscoderAddress(_jobId uint256) constant returns(address)
func (_JobsManager *JobsManagerSession) GetJobTranscoderAddress(_jobId *big.Int) (common.Address, error) {
	return _JobsManager.Contract.GetJobTranscoderAddress(&_JobsManager.CallOpts, _jobId)
}

// GetJobTranscoderAddress is a free data retrieval call binding the contract method 0x8d980fe4.
//
// Solidity: function getJobTranscoderAddress(_jobId uint256) constant returns(address)
func (_JobsManager *JobsManagerCallerSession) GetJobTranscoderAddress(_jobId *big.Int) (common.Address, error) {
	return _JobsManager.Contract.GetJobTranscoderAddress(&_JobsManager.CallOpts, _jobId)
}

// GetJobTranscodingOptions is a free data retrieval call binding the contract method 0x668dfce7.
//
// Solidity: function getJobTranscodingOptions(_jobId uint256) constant returns(string)
func (_JobsManager *JobsManagerCaller) GetJobTranscodingOptions(opts *bind.CallOpts, _jobId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "getJobTranscodingOptions", _jobId)
	return *ret0, err
}

// GetJobTranscodingOptions is a free data retrieval call binding the contract method 0x668dfce7.
//
// Solidity: function getJobTranscodingOptions(_jobId uint256) constant returns(string)
func (_JobsManager *JobsManagerSession) GetJobTranscodingOptions(_jobId *big.Int) (string, error) {
	return _JobsManager.Contract.GetJobTranscodingOptions(&_JobsManager.CallOpts, _jobId)
}

// GetJobTranscodingOptions is a free data retrieval call binding the contract method 0x668dfce7.
//
// Solidity: function getJobTranscodingOptions(_jobId uint256) constant returns(string)
func (_JobsManager *JobsManagerCallerSession) GetJobTranscodingOptions(_jobId *big.Int) (string, error) {
	return _JobsManager.Contract.GetJobTranscodingOptions(&_JobsManager.CallOpts, _jobId)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_JobsManager *JobsManagerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_JobsManager *JobsManagerSession) IsInitialized() (bool, error) {
	return _JobsManager.Contract.IsInitialized(&_JobsManager.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_JobsManager *JobsManagerCallerSession) IsInitialized() (bool, error) {
	return _JobsManager.Contract.IsInitialized(&_JobsManager.CallOpts)
}

// JobEndingPeriod is a free data retrieval call binding the contract method 0xf1e126e6.
//
// Solidity: function jobEndingPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerCaller) JobEndingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "jobEndingPeriod")
	return *ret0, err
}

// JobEndingPeriod is a free data retrieval call binding the contract method 0xf1e126e6.
//
// Solidity: function jobEndingPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerSession) JobEndingPeriod() (*big.Int, error) {
	return _JobsManager.Contract.JobEndingPeriod(&_JobsManager.CallOpts)
}

// JobEndingPeriod is a free data retrieval call binding the contract method 0xf1e126e6.
//
// Solidity: function jobEndingPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) JobEndingPeriod() (*big.Int, error) {
	return _JobsManager.Contract.JobEndingPeriod(&_JobsManager.CallOpts)
}

// JobStatus is a free data retrieval call binding the contract method 0xa8e5e219.
//
// Solidity: function jobStatus(_jobId uint256) constant returns(uint8)
func (_JobsManager *JobsManagerCaller) JobStatus(opts *bind.CallOpts, _jobId *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "jobStatus", _jobId)
	return *ret0, err
}

// JobStatus is a free data retrieval call binding the contract method 0xa8e5e219.
//
// Solidity: function jobStatus(_jobId uint256) constant returns(uint8)
func (_JobsManager *JobsManagerSession) JobStatus(_jobId *big.Int) (uint8, error) {
	return _JobsManager.Contract.JobStatus(&_JobsManager.CallOpts, _jobId)
}

// JobStatus is a free data retrieval call binding the contract method 0xa8e5e219.
//
// Solidity: function jobStatus(_jobId uint256) constant returns(uint8)
func (_JobsManager *JobsManagerCallerSession) JobStatus(_jobId *big.Int) (uint8, error) {
	return _JobsManager.Contract.JobStatus(&_JobsManager.CallOpts, _jobId)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs( uint256) constant returns(jobId uint256, streamId string, transcodingOptions string, maxPricePerSegment uint256, broadcasterAddress address, transcoderAddress address, endBlock uint256, escrow uint256)
func (_JobsManager *JobsManagerCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	JobId              *big.Int
	StreamId           string
	TranscodingOptions string
	MaxPricePerSegment *big.Int
	BroadcasterAddress common.Address
	TranscoderAddress  common.Address
	EndBlock           *big.Int
	Escrow             *big.Int
}, error) {
	ret := new(struct {
		JobId              *big.Int
		StreamId           string
		TranscodingOptions string
		MaxPricePerSegment *big.Int
		BroadcasterAddress common.Address
		TranscoderAddress  common.Address
		EndBlock           *big.Int
		Escrow             *big.Int
	})
	out := ret
	err := _JobsManager.contract.Call(opts, out, "jobs", arg0)
	return *ret, err
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs( uint256) constant returns(jobId uint256, streamId string, transcodingOptions string, maxPricePerSegment uint256, broadcasterAddress address, transcoderAddress address, endBlock uint256, escrow uint256)
func (_JobsManager *JobsManagerSession) Jobs(arg0 *big.Int) (struct {
	JobId              *big.Int
	StreamId           string
	TranscodingOptions string
	MaxPricePerSegment *big.Int
	BroadcasterAddress common.Address
	TranscoderAddress  common.Address
	EndBlock           *big.Int
	Escrow             *big.Int
}, error) {
	return _JobsManager.Contract.Jobs(&_JobsManager.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs( uint256) constant returns(jobId uint256, streamId string, transcodingOptions string, maxPricePerSegment uint256, broadcasterAddress address, transcoderAddress address, endBlock uint256, escrow uint256)
func (_JobsManager *JobsManagerCallerSession) Jobs(arg0 *big.Int) (struct {
	JobId              *big.Int
	StreamId           string
	TranscodingOptions string
	MaxPricePerSegment *big.Int
	BroadcasterAddress common.Address
	TranscoderAddress  common.Address
	EndBlock           *big.Int
	Escrow             *big.Int
}, error) {
	return _JobsManager.Contract.Jobs(&_JobsManager.CallOpts, arg0)
}

// MissedVerificationSlashAmount is a free data retrieval call binding the contract method 0x32b5b2d1.
//
// Solidity: function missedVerificationSlashAmount() constant returns(uint64)
func (_JobsManager *JobsManagerCaller) MissedVerificationSlashAmount(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "missedVerificationSlashAmount")
	return *ret0, err
}

// MissedVerificationSlashAmount is a free data retrieval call binding the contract method 0x32b5b2d1.
//
// Solidity: function missedVerificationSlashAmount() constant returns(uint64)
func (_JobsManager *JobsManagerSession) MissedVerificationSlashAmount() (uint64, error) {
	return _JobsManager.Contract.MissedVerificationSlashAmount(&_JobsManager.CallOpts)
}

// MissedVerificationSlashAmount is a free data retrieval call binding the contract method 0x32b5b2d1.
//
// Solidity: function missedVerificationSlashAmount() constant returns(uint64)
func (_JobsManager *JobsManagerCallerSession) MissedVerificationSlashAmount() (uint64, error) {
	return _JobsManager.Contract.MissedVerificationSlashAmount(&_JobsManager.CallOpts)
}

// NumJobs is a free data retrieval call binding the contract method 0x9212051c.
//
// Solidity: function numJobs() constant returns(uint256)
func (_JobsManager *JobsManagerCaller) NumJobs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "numJobs")
	return *ret0, err
}

// NumJobs is a free data retrieval call binding the contract method 0x9212051c.
//
// Solidity: function numJobs() constant returns(uint256)
func (_JobsManager *JobsManagerSession) NumJobs() (*big.Int, error) {
	return _JobsManager.Contract.NumJobs(&_JobsManager.CallOpts)
}

// NumJobs is a free data retrieval call binding the contract method 0x9212051c.
//
// Solidity: function numJobs() constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) NumJobs() (*big.Int, error) {
	return _JobsManager.Contract.NumJobs(&_JobsManager.CallOpts)
}

// SlashingPeriod is a free data retrieval call binding the contract method 0xed24f661.
//
// Solidity: function slashingPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerCaller) SlashingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "slashingPeriod")
	return *ret0, err
}

// SlashingPeriod is a free data retrieval call binding the contract method 0xed24f661.
//
// Solidity: function slashingPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerSession) SlashingPeriod() (*big.Int, error) {
	return _JobsManager.Contract.SlashingPeriod(&_JobsManager.CallOpts)
}

// SlashingPeriod is a free data retrieval call binding the contract method 0xed24f661.
//
// Solidity: function slashingPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) SlashingPeriod() (*big.Int, error) {
	return _JobsManager.Contract.SlashingPeriod(&_JobsManager.CallOpts)
}

// TargetContractId is a free data retrieval call binding the contract method 0x51720b41.
//
// Solidity: function targetContractId() constant returns(bytes32)
func (_JobsManager *JobsManagerCaller) TargetContractId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "targetContractId")
	return *ret0, err
}

// TargetContractId is a free data retrieval call binding the contract method 0x51720b41.
//
// Solidity: function targetContractId() constant returns(bytes32)
func (_JobsManager *JobsManagerSession) TargetContractId() ([32]byte, error) {
	return _JobsManager.Contract.TargetContractId(&_JobsManager.CallOpts)
}

// TargetContractId is a free data retrieval call binding the contract method 0x51720b41.
//
// Solidity: function targetContractId() constant returns(bytes32)
func (_JobsManager *JobsManagerCallerSession) TargetContractId() ([32]byte, error) {
	return _JobsManager.Contract.TargetContractId(&_JobsManager.CallOpts)
}

// VerificationFailureThreshold is a free data retrieval call binding the contract method 0xc4d4895a.
//
// Solidity: function verificationFailureThreshold() constant returns(uint64)
func (_JobsManager *JobsManagerCaller) VerificationFailureThreshold(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "verificationFailureThreshold")
	return *ret0, err
}

// VerificationFailureThreshold is a free data retrieval call binding the contract method 0xc4d4895a.
//
// Solidity: function verificationFailureThreshold() constant returns(uint64)
func (_JobsManager *JobsManagerSession) VerificationFailureThreshold() (uint64, error) {
	return _JobsManager.Contract.VerificationFailureThreshold(&_JobsManager.CallOpts)
}

// VerificationFailureThreshold is a free data retrieval call binding the contract method 0xc4d4895a.
//
// Solidity: function verificationFailureThreshold() constant returns(uint64)
func (_JobsManager *JobsManagerCallerSession) VerificationFailureThreshold() (uint64, error) {
	return _JobsManager.Contract.VerificationFailureThreshold(&_JobsManager.CallOpts)
}

// VerificationPeriod is a free data retrieval call binding the contract method 0xb1bb7e0f.
//
// Solidity: function verificationPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerCaller) VerificationPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "verificationPeriod")
	return *ret0, err
}

// VerificationPeriod is a free data retrieval call binding the contract method 0xb1bb7e0f.
//
// Solidity: function verificationPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerSession) VerificationPeriod() (*big.Int, error) {
	return _JobsManager.Contract.VerificationPeriod(&_JobsManager.CallOpts)
}

// VerificationPeriod is a free data retrieval call binding the contract method 0xb1bb7e0f.
//
// Solidity: function verificationPeriod() constant returns(uint256)
func (_JobsManager *JobsManagerCallerSession) VerificationPeriod() (*big.Int, error) {
	return _JobsManager.Contract.VerificationPeriod(&_JobsManager.CallOpts)
}

// VerificationRate is a free data retrieval call binding the contract method 0x7af8b87d.
//
// Solidity: function verificationRate() constant returns(uint64)
func (_JobsManager *JobsManagerCaller) VerificationRate(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _JobsManager.contract.Call(opts, out, "verificationRate")
	return *ret0, err
}

// VerificationRate is a free data retrieval call binding the contract method 0x7af8b87d.
//
// Solidity: function verificationRate() constant returns(uint64)
func (_JobsManager *JobsManagerSession) VerificationRate() (uint64, error) {
	return _JobsManager.Contract.VerificationRate(&_JobsManager.CallOpts)
}

// VerificationRate is a free data retrieval call binding the contract method 0x7af8b87d.
//
// Solidity: function verificationRate() constant returns(uint64)
func (_JobsManager *JobsManagerCallerSession) VerificationRate() (uint64, error) {
	return _JobsManager.Contract.VerificationRate(&_JobsManager.CallOpts)
}

// BatchDistributeFees is a paid mutator transaction binding the contract method 0x8978fc79.
//
// Solidity: function batchDistributeFees(_jobId uint256, _claimIds uint256[]) returns(bool)
func (_JobsManager *JobsManagerTransactor) BatchDistributeFees(opts *bind.TransactOpts, _jobId *big.Int, _claimIds []*big.Int) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "batchDistributeFees", _jobId, _claimIds)
}

// BatchDistributeFees is a paid mutator transaction binding the contract method 0x8978fc79.
//
// Solidity: function batchDistributeFees(_jobId uint256, _claimIds uint256[]) returns(bool)
func (_JobsManager *JobsManagerSession) BatchDistributeFees(_jobId *big.Int, _claimIds []*big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.BatchDistributeFees(&_JobsManager.TransactOpts, _jobId, _claimIds)
}

// BatchDistributeFees is a paid mutator transaction binding the contract method 0x8978fc79.
//
// Solidity: function batchDistributeFees(_jobId uint256, _claimIds uint256[]) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) BatchDistributeFees(_jobId *big.Int, _claimIds []*big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.BatchDistributeFees(&_JobsManager.TransactOpts, _jobId, _claimIds)
}

// ClaimWork is a paid mutator transaction binding the contract method 0x3ffe5eb7.
//
// Solidity: function claimWork(_jobId uint256, _segmentRange uint256[2], _claimRoot bytes32) returns(bool)
func (_JobsManager *JobsManagerTransactor) ClaimWork(opts *bind.TransactOpts, _jobId *big.Int, _segmentRange [2]*big.Int, _claimRoot [32]byte) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "claimWork", _jobId, _segmentRange, _claimRoot)
}

// ClaimWork is a paid mutator transaction binding the contract method 0x3ffe5eb7.
//
// Solidity: function claimWork(_jobId uint256, _segmentRange uint256[2], _claimRoot bytes32) returns(bool)
func (_JobsManager *JobsManagerSession) ClaimWork(_jobId *big.Int, _segmentRange [2]*big.Int, _claimRoot [32]byte) (*types.Transaction, error) {
	return _JobsManager.Contract.ClaimWork(&_JobsManager.TransactOpts, _jobId, _segmentRange, _claimRoot)
}

// ClaimWork is a paid mutator transaction binding the contract method 0x3ffe5eb7.
//
// Solidity: function claimWork(_jobId uint256, _segmentRange uint256[2], _claimRoot bytes32) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) ClaimWork(_jobId *big.Int, _segmentRange [2]*big.Int, _claimRoot [32]byte) (*types.Transaction, error) {
	return _JobsManager.Contract.ClaimWork(&_JobsManager.TransactOpts, _jobId, _segmentRange, _claimRoot)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_JobsManager *JobsManagerTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_JobsManager *JobsManagerSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.Deposit(&_JobsManager.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.Deposit(&_JobsManager.TransactOpts, _amount)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x7e69671a.
//
// Solidity: function distributeFees(_jobId uint256, _claimId uint256) returns(bool)
func (_JobsManager *JobsManagerTransactor) DistributeFees(opts *bind.TransactOpts, _jobId *big.Int, _claimId *big.Int) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "distributeFees", _jobId, _claimId)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x7e69671a.
//
// Solidity: function distributeFees(_jobId uint256, _claimId uint256) returns(bool)
func (_JobsManager *JobsManagerSession) DistributeFees(_jobId *big.Int, _claimId *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.DistributeFees(&_JobsManager.TransactOpts, _jobId, _claimId)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x7e69671a.
//
// Solidity: function distributeFees(_jobId uint256, _claimId uint256) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) DistributeFees(_jobId *big.Int, _claimId *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.DistributeFees(&_JobsManager.TransactOpts, _jobId, _claimId)
}

// EndJob is a paid mutator transaction binding the contract method 0xf8e888b4.
//
// Solidity: function endJob(_jobId uint256) returns(bool)
func (_JobsManager *JobsManagerTransactor) EndJob(opts *bind.TransactOpts, _jobId *big.Int) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "endJob", _jobId)
}

// EndJob is a paid mutator transaction binding the contract method 0xf8e888b4.
//
// Solidity: function endJob(_jobId uint256) returns(bool)
func (_JobsManager *JobsManagerSession) EndJob(_jobId *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.EndJob(&_JobsManager.TransactOpts, _jobId)
}

// EndJob is a paid mutator transaction binding the contract method 0xf8e888b4.
//
// Solidity: function endJob(_jobId uint256) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) EndJob(_jobId *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.EndJob(&_JobsManager.TransactOpts, _jobId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe1f9154.
//
// Solidity: function initialize(_verificationRate uint64, _jobEndingPeriod uint256, _verificationPeriod uint256, _slashingPeriod uint256, _failedVerificationSlashAmount uint64, _missedVerificationSlashAmount uint64, _finderFee uint64) returns(bool)
func (_JobsManager *JobsManagerTransactor) Initialize(opts *bind.TransactOpts, _verificationRate uint64, _jobEndingPeriod *big.Int, _verificationPeriod *big.Int, _slashingPeriod *big.Int, _failedVerificationSlashAmount uint64, _missedVerificationSlashAmount uint64, _finderFee uint64) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "initialize", _verificationRate, _jobEndingPeriod, _verificationPeriod, _slashingPeriod, _failedVerificationSlashAmount, _missedVerificationSlashAmount, _finderFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe1f9154.
//
// Solidity: function initialize(_verificationRate uint64, _jobEndingPeriod uint256, _verificationPeriod uint256, _slashingPeriod uint256, _failedVerificationSlashAmount uint64, _missedVerificationSlashAmount uint64, _finderFee uint64) returns(bool)
func (_JobsManager *JobsManagerSession) Initialize(_verificationRate uint64, _jobEndingPeriod *big.Int, _verificationPeriod *big.Int, _slashingPeriod *big.Int, _failedVerificationSlashAmount uint64, _missedVerificationSlashAmount uint64, _finderFee uint64) (*types.Transaction, error) {
	return _JobsManager.Contract.Initialize(&_JobsManager.TransactOpts, _verificationRate, _jobEndingPeriod, _verificationPeriod, _slashingPeriod, _failedVerificationSlashAmount, _missedVerificationSlashAmount, _finderFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe1f9154.
//
// Solidity: function initialize(_verificationRate uint64, _jobEndingPeriod uint256, _verificationPeriod uint256, _slashingPeriod uint256, _failedVerificationSlashAmount uint64, _missedVerificationSlashAmount uint64, _finderFee uint64) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) Initialize(_verificationRate uint64, _jobEndingPeriod *big.Int, _verificationPeriod *big.Int, _slashingPeriod *big.Int, _failedVerificationSlashAmount uint64, _missedVerificationSlashAmount uint64, _finderFee uint64) (*types.Transaction, error) {
	return _JobsManager.Contract.Initialize(&_JobsManager.TransactOpts, _verificationRate, _jobEndingPeriod, _verificationPeriod, _slashingPeriod, _failedVerificationSlashAmount, _missedVerificationSlashAmount, _finderFee)
}

// Job is a paid mutator transaction binding the contract method 0x009fe64b.
//
// Solidity: function job(_streamId string, _transcodingOptions string, _maxPricePerSegment uint256) returns(bool)
func (_JobsManager *JobsManagerTransactor) Job(opts *bind.TransactOpts, _streamId string, _transcodingOptions string, _maxPricePerSegment *big.Int) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "job", _streamId, _transcodingOptions, _maxPricePerSegment)
}

// Job is a paid mutator transaction binding the contract method 0x009fe64b.
//
// Solidity: function job(_streamId string, _transcodingOptions string, _maxPricePerSegment uint256) returns(bool)
func (_JobsManager *JobsManagerSession) Job(_streamId string, _transcodingOptions string, _maxPricePerSegment *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.Job(&_JobsManager.TransactOpts, _streamId, _transcodingOptions, _maxPricePerSegment)
}

// Job is a paid mutator transaction binding the contract method 0x009fe64b.
//
// Solidity: function job(_streamId string, _transcodingOptions string, _maxPricePerSegment uint256) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) Job(_streamId string, _transcodingOptions string, _maxPricePerSegment *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.Job(&_JobsManager.TransactOpts, _streamId, _transcodingOptions, _maxPricePerSegment)
}

// MissedVerificationSlash is a paid mutator transaction binding the contract method 0xc8e8f487.
//
// Solidity: function missedVerificationSlash(_jobId uint256, _claimId uint256, _segmentNumber uint256) returns(bool)
func (_JobsManager *JobsManagerTransactor) MissedVerificationSlash(opts *bind.TransactOpts, _jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "missedVerificationSlash", _jobId, _claimId, _segmentNumber)
}

// MissedVerificationSlash is a paid mutator transaction binding the contract method 0xc8e8f487.
//
// Solidity: function missedVerificationSlash(_jobId uint256, _claimId uint256, _segmentNumber uint256) returns(bool)
func (_JobsManager *JobsManagerSession) MissedVerificationSlash(_jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.MissedVerificationSlash(&_JobsManager.TransactOpts, _jobId, _claimId, _segmentNumber)
}

// MissedVerificationSlash is a paid mutator transaction binding the contract method 0xc8e8f487.
//
// Solidity: function missedVerificationSlash(_jobId uint256, _claimId uint256, _segmentNumber uint256) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) MissedVerificationSlash(_jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int) (*types.Transaction, error) {
	return _JobsManager.Contract.MissedVerificationSlash(&_JobsManager.TransactOpts, _jobId, _claimId, _segmentNumber)
}

// ReceiveVerification is a paid mutator transaction binding the contract method 0x1e0976f3.
//
// Solidity: function receiveVerification(_jobId uint256, _claimId uint256, _segmentNumber uint256, _result bool) returns(bool)
func (_JobsManager *JobsManagerTransactor) ReceiveVerification(opts *bind.TransactOpts, _jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int, _result bool) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "receiveVerification", _jobId, _claimId, _segmentNumber, _result)
}

// ReceiveVerification is a paid mutator transaction binding the contract method 0x1e0976f3.
//
// Solidity: function receiveVerification(_jobId uint256, _claimId uint256, _segmentNumber uint256, _result bool) returns(bool)
func (_JobsManager *JobsManagerSession) ReceiveVerification(_jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int, _result bool) (*types.Transaction, error) {
	return _JobsManager.Contract.ReceiveVerification(&_JobsManager.TransactOpts, _jobId, _claimId, _segmentNumber, _result)
}

// ReceiveVerification is a paid mutator transaction binding the contract method 0x1e0976f3.
//
// Solidity: function receiveVerification(_jobId uint256, _claimId uint256, _segmentNumber uint256, _result bool) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) ReceiveVerification(_jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int, _result bool) (*types.Transaction, error) {
	return _JobsManager.Contract.ReceiveVerification(&_JobsManager.TransactOpts, _jobId, _claimId, _segmentNumber, _result)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_JobsManager *JobsManagerTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_JobsManager *JobsManagerSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _JobsManager.Contract.SetController(&_JobsManager.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(_controller address) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _JobsManager.Contract.SetController(&_JobsManager.TransactOpts, _controller)
}

// Verify is a paid mutator transaction binding the contract method 0x5a40ec7e.
//
// Solidity: function verify(_jobId uint256, _claimId uint256, _segmentNumber uint256, _dataStorageHash string, _dataHashes bytes32[2], _broadcasterSig bytes, _proof bytes) returns(bool)
func (_JobsManager *JobsManagerTransactor) Verify(opts *bind.TransactOpts, _jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int, _dataStorageHash string, _dataHashes [2][32]byte, _broadcasterSig []byte, _proof []byte) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "verify", _jobId, _claimId, _segmentNumber, _dataStorageHash, _dataHashes, _broadcasterSig, _proof)
}

// Verify is a paid mutator transaction binding the contract method 0x5a40ec7e.
//
// Solidity: function verify(_jobId uint256, _claimId uint256, _segmentNumber uint256, _dataStorageHash string, _dataHashes bytes32[2], _broadcasterSig bytes, _proof bytes) returns(bool)
func (_JobsManager *JobsManagerSession) Verify(_jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int, _dataStorageHash string, _dataHashes [2][32]byte, _broadcasterSig []byte, _proof []byte) (*types.Transaction, error) {
	return _JobsManager.Contract.Verify(&_JobsManager.TransactOpts, _jobId, _claimId, _segmentNumber, _dataStorageHash, _dataHashes, _broadcasterSig, _proof)
}

// Verify is a paid mutator transaction binding the contract method 0x5a40ec7e.
//
// Solidity: function verify(_jobId uint256, _claimId uint256, _segmentNumber uint256, _dataStorageHash string, _dataHashes bytes32[2], _broadcasterSig bytes, _proof bytes) returns(bool)
func (_JobsManager *JobsManagerTransactorSession) Verify(_jobId *big.Int, _claimId *big.Int, _segmentNumber *big.Int, _dataStorageHash string, _dataHashes [2][32]byte, _broadcasterSig []byte, _proof []byte) (*types.Transaction, error) {
	return _JobsManager.Contract.Verify(&_JobsManager.TransactOpts, _jobId, _claimId, _segmentNumber, _dataStorageHash, _dataHashes, _broadcasterSig, _proof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_JobsManager *JobsManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JobsManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_JobsManager *JobsManagerSession) Withdraw() (*types.Transaction, error) {
	return _JobsManager.Contract.Withdraw(&_JobsManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_JobsManager *JobsManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _JobsManager.Contract.Withdraw(&_JobsManager.TransactOpts)
}

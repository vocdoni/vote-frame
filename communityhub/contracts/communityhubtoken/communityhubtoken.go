// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CommunityHubToken

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

// ICommunityHubCensus is an auto generated low-level Go binding around an user-defined struct.
type ICommunityHubCensus struct {
	CensusType uint8
	Tokens     []ICommunityHubToken
	Channel    string
}

// ICommunityHubCommunity is an auto generated low-level Go binding around an user-defined struct.
type ICommunityHubCommunity struct {
	Metadata                 ICommunityHubCommunityMetadata
	Census                   ICommunityHubCensus
	Guardians                []*big.Int
	CreateElectionPermission uint8
	Disabled                 bool
	Funds                    *big.Int
}

// ICommunityHubCommunityMetadata is an auto generated low-level Go binding around an user-defined struct.
type ICommunityHubCommunityMetadata struct {
	Name          string
	ImageURI      string
	GroupChatURL  string
	Channels      []string
	Notifications bool
}

// ICommunityHubToken is an auto generated low-level Go binding around an user-defined struct.
type ICommunityHubToken struct {
	Blockchain      string
	ContractAddress common.Address
}

// IResultResult is an auto generated low-level Go binding around an user-defined struct.
type IResultResult struct {
	Question         string
	Options          []string
	Date             string
	Tally            [][]*big.Int
	Turnout          *big.Int
	TotalVotingPower *big.Int
	Participants     []*big.Int
	CensusRoot       [32]byte
	CensusURI        string
}

// CommunityHubTokenMetaData contains all meta data concerning the CommunityHubToken contract.
var CommunityHubTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"AmountMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"guardian\",\"type\":\"uint256\"}],\"name\":\"GuardianNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCreateElectionPermission\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"}],\"name\":\"AdminCommunityManaged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumICommunityHub.CensusType\",\"name\":\"censusType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"blockchain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structICommunityHub.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structICommunityHub.Census\",\"name\":\"census\",\"type\":\"tuple\"}],\"name\":\"CensusSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CommunityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"}],\"name\":\"CommunityDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"}],\"name\":\"CommunityDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"}],\"name\":\"CommunityEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"CreateCommunityPriceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumICommunityHub.CreateElectionPermission\",\"name\":\"createElectionPermission\",\"type\":\"uint8\"}],\"name\":\"CreateElectionPermissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"guardian\",\"type\":\"uint256\"}],\"name\":\"GuardianAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"guardian\",\"type\":\"uint256\"}],\"name\":\"GuardianRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"groupChatURL\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"notifications\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structICommunityHub.CommunityMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"name\":\"MetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"notifiableElections\",\"type\":\"bool\"}],\"name\":\"NotifiableElectionsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PricePerElectionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"communityId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"electionId\",\"type\":\"bytes32\"}],\"name\":\"ResultsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_guardian\",\"type\":\"uint256\"}],\"name\":\"addGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"groupChatURL\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"notifications\",\"type\":\"bool\"}],\"internalType\":\"structICommunityHub.CommunityMetadata\",\"name\":\"_metadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICommunityHub.CensusType\",\"name\":\"censusType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"blockchain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structICommunityHub.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"internalType\":\"structICommunityHub.Census\",\"name\":\"_census\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"_guardians\",\"type\":\"uint256[]\"},{\"internalType\":\"enumICommunityHub.CreateElectionPermission\",\"name\":\"_createElectionPermission\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_disabled\",\"type\":\"bool\"}],\"name\":\"adminManageCommunity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"adminSetCommunityPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"adminSetPricePerElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"groupChatURL\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"notifications\",\"type\":\"bool\"}],\"internalType\":\"structICommunityHub.CommunityMetadata\",\"name\":\"_metadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICommunityHub.CensusType\",\"name\":\"censusType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"blockchain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structICommunityHub.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"internalType\":\"structICommunityHub.Census\",\"name\":\"_census\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"_guardians\",\"type\":\"uint256[]\"},{\"internalType\":\"enumICommunityHub.CreateElectionPermission\",\"name\":\"_createElectionPermission\",\"type\":\"uint8\"}],\"name\":\"createCommunity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"}],\"name\":\"getCommunity\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"groupChatURL\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"notifications\",\"type\":\"bool\"}],\"internalType\":\"structICommunityHub.CommunityMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICommunityHub.CensusType\",\"name\":\"censusType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"blockchain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structICommunityHub.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"internalType\":\"structICommunityHub.Census\",\"name\":\"census\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"guardians\",\"type\":\"uint256[]\"},{\"internalType\":\"enumICommunityHub.CreateElectionPermission\",\"name\":\"createElectionPermission\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"disabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"internalType\":\"structICommunityHub.Community\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreateCommunityPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCommunityId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerElection\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_electionId\",\"type\":\"bytes32\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"question\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"options\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256[][]\",\"name\":\"tally\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"turnout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"participants\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIResult.Result\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_guardian\",\"type\":\"uint256\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumICommunityHub.CensusType\",\"name\":\"censusType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"blockchain\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"internalType\":\"structICommunityHub.Token[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"channel\",\"type\":\"string\"}],\"internalType\":\"structICommunityHub.Census\",\"name\":\"_census\",\"type\":\"tuple\"}],\"name\":\"setCensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"internalType\":\"enumICommunityHub.CreateElectionPermission\",\"name\":\"_createElectionPermission\",\"type\":\"uint8\"}],\"name\":\"setCreateElectionPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"groupChatURL\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"channels\",\"type\":\"string[]\"},{\"internalType\":\"bool\",\"name\":\"notifications\",\"type\":\"bool\"}],\"internalType\":\"structICommunityHub.CommunityMetadata\",\"name\":\"_metadata\",\"type\":\"tuple\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_notifiableElections\",\"type\":\"bool\"}],\"name\":\"setNotifiableElections\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_communityId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_electionId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"question\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"options\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256[][]\",\"name\":\"tally\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"turnout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotingPower\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"participants\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"censusRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"censusURI\",\"type\":\"string\"}],\"internalType\":\"structIResult.Result\",\"name\":\"_result\",\"type\":\"tuple\"}],\"name\":\"setResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CommunityHubTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CommunityHubTokenMetaData.ABI instead.
var CommunityHubTokenABI = CommunityHubTokenMetaData.ABI

// CommunityHubToken is an auto generated Go binding around an Ethereum contract.
type CommunityHubToken struct {
	CommunityHubTokenCaller     // Read-only binding to the contract
	CommunityHubTokenTransactor // Write-only binding to the contract
	CommunityHubTokenFilterer   // Log filterer for contract events
}

// CommunityHubTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommunityHubTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommunityHubTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommunityHubTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommunityHubTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommunityHubTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommunityHubTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommunityHubTokenSession struct {
	Contract     *CommunityHubToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CommunityHubTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommunityHubTokenCallerSession struct {
	Contract *CommunityHubTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CommunityHubTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommunityHubTokenTransactorSession struct {
	Contract     *CommunityHubTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CommunityHubTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommunityHubTokenRaw struct {
	Contract *CommunityHubToken // Generic contract binding to access the raw methods on
}

// CommunityHubTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommunityHubTokenCallerRaw struct {
	Contract *CommunityHubTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CommunityHubTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommunityHubTokenTransactorRaw struct {
	Contract *CommunityHubTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommunityHubToken creates a new instance of CommunityHubToken, bound to a specific deployed contract.
func NewCommunityHubToken(address common.Address, backend bind.ContractBackend) (*CommunityHubToken, error) {
	contract, err := bindCommunityHubToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommunityHubToken{CommunityHubTokenCaller: CommunityHubTokenCaller{contract: contract}, CommunityHubTokenTransactor: CommunityHubTokenTransactor{contract: contract}, CommunityHubTokenFilterer: CommunityHubTokenFilterer{contract: contract}}, nil
}

// NewCommunityHubTokenCaller creates a new read-only instance of CommunityHubToken, bound to a specific deployed contract.
func NewCommunityHubTokenCaller(address common.Address, caller bind.ContractCaller) (*CommunityHubTokenCaller, error) {
	contract, err := bindCommunityHubToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCaller{contract: contract}, nil
}

// NewCommunityHubTokenTransactor creates a new write-only instance of CommunityHubToken, bound to a specific deployed contract.
func NewCommunityHubTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CommunityHubTokenTransactor, error) {
	contract, err := bindCommunityHubToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenTransactor{contract: contract}, nil
}

// NewCommunityHubTokenFilterer creates a new log filterer instance of CommunityHubToken, bound to a specific deployed contract.
func NewCommunityHubTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CommunityHubTokenFilterer, error) {
	contract, err := bindCommunityHubToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenFilterer{contract: contract}, nil
}

// bindCommunityHubToken binds a generic wrapper to an already deployed contract.
func bindCommunityHubToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommunityHubTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommunityHubToken *CommunityHubTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommunityHubToken.Contract.CommunityHubTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommunityHubToken *CommunityHubTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.CommunityHubTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommunityHubToken *CommunityHubTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.CommunityHubTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommunityHubToken *CommunityHubTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommunityHubToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommunityHubToken *CommunityHubTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommunityHubToken *CommunityHubTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.contract.Transact(opts, method, params...)
}

// GetCommunity is a free data retrieval call binding the contract method 0x370603f8.
//
// Solidity: function getCommunity(uint256 _communityId) view returns(((string,string,string,string[],bool),(uint8,(string,address)[],string),uint256[],uint8,bool,uint256))
func (_CommunityHubToken *CommunityHubTokenCaller) GetCommunity(opts *bind.CallOpts, _communityId *big.Int) (ICommunityHubCommunity, error) {
	var out []interface{}
	err := _CommunityHubToken.contract.Call(opts, &out, "getCommunity", _communityId)

	if err != nil {
		return *new(ICommunityHubCommunity), err
	}

	out0 := *abi.ConvertType(out[0], new(ICommunityHubCommunity)).(*ICommunityHubCommunity)

	return out0, err

}

// GetCommunity is a free data retrieval call binding the contract method 0x370603f8.
//
// Solidity: function getCommunity(uint256 _communityId) view returns(((string,string,string,string[],bool),(uint8,(string,address)[],string),uint256[],uint8,bool,uint256))
func (_CommunityHubToken *CommunityHubTokenSession) GetCommunity(_communityId *big.Int) (ICommunityHubCommunity, error) {
	return _CommunityHubToken.Contract.GetCommunity(&_CommunityHubToken.CallOpts, _communityId)
}

// GetCommunity is a free data retrieval call binding the contract method 0x370603f8.
//
// Solidity: function getCommunity(uint256 _communityId) view returns(((string,string,string,string[],bool),(uint8,(string,address)[],string),uint256[],uint8,bool,uint256))
func (_CommunityHubToken *CommunityHubTokenCallerSession) GetCommunity(_communityId *big.Int) (ICommunityHubCommunity, error) {
	return _CommunityHubToken.Contract.GetCommunity(&_CommunityHubToken.CallOpts, _communityId)
}

// GetCreateCommunityPrice is a free data retrieval call binding the contract method 0xc5bf44ab.
//
// Solidity: function getCreateCommunityPrice() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenCaller) GetCreateCommunityPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommunityHubToken.contract.Call(opts, &out, "getCreateCommunityPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreateCommunityPrice is a free data retrieval call binding the contract method 0xc5bf44ab.
//
// Solidity: function getCreateCommunityPrice() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenSession) GetCreateCommunityPrice() (*big.Int, error) {
	return _CommunityHubToken.Contract.GetCreateCommunityPrice(&_CommunityHubToken.CallOpts)
}

// GetCreateCommunityPrice is a free data retrieval call binding the contract method 0xc5bf44ab.
//
// Solidity: function getCreateCommunityPrice() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenCallerSession) GetCreateCommunityPrice() (*big.Int, error) {
	return _CommunityHubToken.Contract.GetCreateCommunityPrice(&_CommunityHubToken.CallOpts)
}

// GetNextCommunityId is a free data retrieval call binding the contract method 0x9532c8c3.
//
// Solidity: function getNextCommunityId() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenCaller) GetNextCommunityId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommunityHubToken.contract.Call(opts, &out, "getNextCommunityId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCommunityId is a free data retrieval call binding the contract method 0x9532c8c3.
//
// Solidity: function getNextCommunityId() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenSession) GetNextCommunityId() (*big.Int, error) {
	return _CommunityHubToken.Contract.GetNextCommunityId(&_CommunityHubToken.CallOpts)
}

// GetNextCommunityId is a free data retrieval call binding the contract method 0x9532c8c3.
//
// Solidity: function getNextCommunityId() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenCallerSession) GetNextCommunityId() (*big.Int, error) {
	return _CommunityHubToken.Contract.GetNextCommunityId(&_CommunityHubToken.CallOpts)
}

// GetPricePerElection is a free data retrieval call binding the contract method 0xd62b5c59.
//
// Solidity: function getPricePerElection() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenCaller) GetPricePerElection(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CommunityHubToken.contract.Call(opts, &out, "getPricePerElection")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerElection is a free data retrieval call binding the contract method 0xd62b5c59.
//
// Solidity: function getPricePerElection() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenSession) GetPricePerElection() (*big.Int, error) {
	return _CommunityHubToken.Contract.GetPricePerElection(&_CommunityHubToken.CallOpts)
}

// GetPricePerElection is a free data retrieval call binding the contract method 0xd62b5c59.
//
// Solidity: function getPricePerElection() view returns(uint256)
func (_CommunityHubToken *CommunityHubTokenCallerSession) GetPricePerElection() (*big.Int, error) {
	return _CommunityHubToken.Contract.GetPricePerElection(&_CommunityHubToken.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0x13e86265.
//
// Solidity: function getResult(uint256 _communityId, bytes32 _electionId) view returns((string,string[],string,uint256[][],uint256,uint256,uint256[],bytes32,string) result)
func (_CommunityHubToken *CommunityHubTokenCaller) GetResult(opts *bind.CallOpts, _communityId *big.Int, _electionId [32]byte) (IResultResult, error) {
	var out []interface{}
	err := _CommunityHubToken.contract.Call(opts, &out, "getResult", _communityId, _electionId)

	if err != nil {
		return *new(IResultResult), err
	}

	out0 := *abi.ConvertType(out[0], new(IResultResult)).(*IResultResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x13e86265.
//
// Solidity: function getResult(uint256 _communityId, bytes32 _electionId) view returns((string,string[],string,uint256[][],uint256,uint256,uint256[],bytes32,string) result)
func (_CommunityHubToken *CommunityHubTokenSession) GetResult(_communityId *big.Int, _electionId [32]byte) (IResultResult, error) {
	return _CommunityHubToken.Contract.GetResult(&_CommunityHubToken.CallOpts, _communityId, _electionId)
}

// GetResult is a free data retrieval call binding the contract method 0x13e86265.
//
// Solidity: function getResult(uint256 _communityId, bytes32 _electionId) view returns((string,string[],string,uint256[][],uint256,uint256,uint256[],bytes32,string) result)
func (_CommunityHubToken *CommunityHubTokenCallerSession) GetResult(_communityId *big.Int, _electionId [32]byte) (IResultResult, error) {
	return _CommunityHubToken.Contract.GetResult(&_CommunityHubToken.CallOpts, _communityId, _electionId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CommunityHubToken *CommunityHubTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CommunityHubToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CommunityHubToken *CommunityHubTokenSession) Owner() (common.Address, error) {
	return _CommunityHubToken.Contract.Owner(&_CommunityHubToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CommunityHubToken *CommunityHubTokenCallerSession) Owner() (common.Address, error) {
	return _CommunityHubToken.Contract.Owner(&_CommunityHubToken.CallOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xc8b785f3.
//
// Solidity: function addGuardian(uint256 _communityId, uint256 _guardian) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) AddGuardian(opts *bind.TransactOpts, _communityId *big.Int, _guardian *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "addGuardian", _communityId, _guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xc8b785f3.
//
// Solidity: function addGuardian(uint256 _communityId, uint256 _guardian) returns()
func (_CommunityHubToken *CommunityHubTokenSession) AddGuardian(_communityId *big.Int, _guardian *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AddGuardian(&_CommunityHubToken.TransactOpts, _communityId, _guardian)
}

// AddGuardian is a paid mutator transaction binding the contract method 0xc8b785f3.
//
// Solidity: function addGuardian(uint256 _communityId, uint256 _guardian) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) AddGuardian(_communityId *big.Int, _guardian *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AddGuardian(&_CommunityHubToken.TransactOpts, _communityId, _guardian)
}

// AdminManageCommunity is a paid mutator transaction binding the contract method 0x040cdd0a.
//
// Solidity: function adminManageCommunity(uint256 _communityId, (string,string,string,string[],bool) _metadata, (uint8,(string,address)[],string) _census, uint256[] _guardians, uint8 _createElectionPermission, bool _disabled) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) AdminManageCommunity(opts *bind.TransactOpts, _communityId *big.Int, _metadata ICommunityHubCommunityMetadata, _census ICommunityHubCensus, _guardians []*big.Int, _createElectionPermission uint8, _disabled bool) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "adminManageCommunity", _communityId, _metadata, _census, _guardians, _createElectionPermission, _disabled)
}

// AdminManageCommunity is a paid mutator transaction binding the contract method 0x040cdd0a.
//
// Solidity: function adminManageCommunity(uint256 _communityId, (string,string,string,string[],bool) _metadata, (uint8,(string,address)[],string) _census, uint256[] _guardians, uint8 _createElectionPermission, bool _disabled) returns()
func (_CommunityHubToken *CommunityHubTokenSession) AdminManageCommunity(_communityId *big.Int, _metadata ICommunityHubCommunityMetadata, _census ICommunityHubCensus, _guardians []*big.Int, _createElectionPermission uint8, _disabled bool) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AdminManageCommunity(&_CommunityHubToken.TransactOpts, _communityId, _metadata, _census, _guardians, _createElectionPermission, _disabled)
}

// AdminManageCommunity is a paid mutator transaction binding the contract method 0x040cdd0a.
//
// Solidity: function adminManageCommunity(uint256 _communityId, (string,string,string,string[],bool) _metadata, (uint8,(string,address)[],string) _census, uint256[] _guardians, uint8 _createElectionPermission, bool _disabled) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) AdminManageCommunity(_communityId *big.Int, _metadata ICommunityHubCommunityMetadata, _census ICommunityHubCensus, _guardians []*big.Int, _createElectionPermission uint8, _disabled bool) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AdminManageCommunity(&_CommunityHubToken.TransactOpts, _communityId, _metadata, _census, _guardians, _createElectionPermission, _disabled)
}

// AdminSetCommunityPrice is a paid mutator transaction binding the contract method 0x873ceb2e.
//
// Solidity: function adminSetCommunityPrice(uint256 _price) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) AdminSetCommunityPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "adminSetCommunityPrice", _price)
}

// AdminSetCommunityPrice is a paid mutator transaction binding the contract method 0x873ceb2e.
//
// Solidity: function adminSetCommunityPrice(uint256 _price) returns()
func (_CommunityHubToken *CommunityHubTokenSession) AdminSetCommunityPrice(_price *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AdminSetCommunityPrice(&_CommunityHubToken.TransactOpts, _price)
}

// AdminSetCommunityPrice is a paid mutator transaction binding the contract method 0x873ceb2e.
//
// Solidity: function adminSetCommunityPrice(uint256 _price) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) AdminSetCommunityPrice(_price *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AdminSetCommunityPrice(&_CommunityHubToken.TransactOpts, _price)
}

// AdminSetPricePerElection is a paid mutator transaction binding the contract method 0x3e512ebc.
//
// Solidity: function adminSetPricePerElection(uint256 _price) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) AdminSetPricePerElection(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "adminSetPricePerElection", _price)
}

// AdminSetPricePerElection is a paid mutator transaction binding the contract method 0x3e512ebc.
//
// Solidity: function adminSetPricePerElection(uint256 _price) returns()
func (_CommunityHubToken *CommunityHubTokenSession) AdminSetPricePerElection(_price *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AdminSetPricePerElection(&_CommunityHubToken.TransactOpts, _price)
}

// AdminSetPricePerElection is a paid mutator transaction binding the contract method 0x3e512ebc.
//
// Solidity: function adminSetPricePerElection(uint256 _price) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) AdminSetPricePerElection(_price *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.AdminSetPricePerElection(&_CommunityHubToken.TransactOpts, _price)
}

// CreateCommunity is a paid mutator transaction binding the contract method 0x6061a460.
//
// Solidity: function createCommunity((string,string,string,string[],bool) _metadata, (uint8,(string,address)[],string) _census, uint256[] _guardians, uint8 _createElectionPermission) payable returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) CreateCommunity(opts *bind.TransactOpts, _metadata ICommunityHubCommunityMetadata, _census ICommunityHubCensus, _guardians []*big.Int, _createElectionPermission uint8) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "createCommunity", _metadata, _census, _guardians, _createElectionPermission)
}

// CreateCommunity is a paid mutator transaction binding the contract method 0x6061a460.
//
// Solidity: function createCommunity((string,string,string,string[],bool) _metadata, (uint8,(string,address)[],string) _census, uint256[] _guardians, uint8 _createElectionPermission) payable returns()
func (_CommunityHubToken *CommunityHubTokenSession) CreateCommunity(_metadata ICommunityHubCommunityMetadata, _census ICommunityHubCensus, _guardians []*big.Int, _createElectionPermission uint8) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.CreateCommunity(&_CommunityHubToken.TransactOpts, _metadata, _census, _guardians, _createElectionPermission)
}

// CreateCommunity is a paid mutator transaction binding the contract method 0x6061a460.
//
// Solidity: function createCommunity((string,string,string,string[],bool) _metadata, (uint8,(string,address)[],string) _census, uint256[] _guardians, uint8 _createElectionPermission) payable returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) CreateCommunity(_metadata ICommunityHubCommunityMetadata, _census ICommunityHubCensus, _guardians []*big.Int, _createElectionPermission uint8) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.CreateCommunity(&_CommunityHubToken.TransactOpts, _metadata, _census, _guardians, _createElectionPermission)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _communityId) payable returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) Deposit(opts *bind.TransactOpts, _communityId *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "deposit", _communityId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _communityId) payable returns()
func (_CommunityHubToken *CommunityHubTokenSession) Deposit(_communityId *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.Deposit(&_CommunityHubToken.TransactOpts, _communityId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _communityId) payable returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) Deposit(_communityId *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.Deposit(&_CommunityHubToken.TransactOpts, _communityId)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0xaba675b0.
//
// Solidity: function removeGuardian(uint256 _communityId, uint256 _guardian) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) RemoveGuardian(opts *bind.TransactOpts, _communityId *big.Int, _guardian *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "removeGuardian", _communityId, _guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0xaba675b0.
//
// Solidity: function removeGuardian(uint256 _communityId, uint256 _guardian) returns()
func (_CommunityHubToken *CommunityHubTokenSession) RemoveGuardian(_communityId *big.Int, _guardian *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.RemoveGuardian(&_CommunityHubToken.TransactOpts, _communityId, _guardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0xaba675b0.
//
// Solidity: function removeGuardian(uint256 _communityId, uint256 _guardian) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) RemoveGuardian(_communityId *big.Int, _guardian *big.Int) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.RemoveGuardian(&_CommunityHubToken.TransactOpts, _communityId, _guardian)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CommunityHubToken *CommunityHubTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _CommunityHubToken.Contract.RenounceOwnership(&_CommunityHubToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CommunityHubToken.Contract.RenounceOwnership(&_CommunityHubToken.TransactOpts)
}

// SetCensus is a paid mutator transaction binding the contract method 0x56c4e4ba.
//
// Solidity: function setCensus(uint256 _communityId, (uint8,(string,address)[],string) _census) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) SetCensus(opts *bind.TransactOpts, _communityId *big.Int, _census ICommunityHubCensus) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "setCensus", _communityId, _census)
}

// SetCensus is a paid mutator transaction binding the contract method 0x56c4e4ba.
//
// Solidity: function setCensus(uint256 _communityId, (uint8,(string,address)[],string) _census) returns()
func (_CommunityHubToken *CommunityHubTokenSession) SetCensus(_communityId *big.Int, _census ICommunityHubCensus) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetCensus(&_CommunityHubToken.TransactOpts, _communityId, _census)
}

// SetCensus is a paid mutator transaction binding the contract method 0x56c4e4ba.
//
// Solidity: function setCensus(uint256 _communityId, (uint8,(string,address)[],string) _census) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) SetCensus(_communityId *big.Int, _census ICommunityHubCensus) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetCensus(&_CommunityHubToken.TransactOpts, _communityId, _census)
}

// SetCreateElectionPermission is a paid mutator transaction binding the contract method 0x7824c05c.
//
// Solidity: function setCreateElectionPermission(uint256 _communityId, uint8 _createElectionPermission) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) SetCreateElectionPermission(opts *bind.TransactOpts, _communityId *big.Int, _createElectionPermission uint8) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "setCreateElectionPermission", _communityId, _createElectionPermission)
}

// SetCreateElectionPermission is a paid mutator transaction binding the contract method 0x7824c05c.
//
// Solidity: function setCreateElectionPermission(uint256 _communityId, uint8 _createElectionPermission) returns()
func (_CommunityHubToken *CommunityHubTokenSession) SetCreateElectionPermission(_communityId *big.Int, _createElectionPermission uint8) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetCreateElectionPermission(&_CommunityHubToken.TransactOpts, _communityId, _createElectionPermission)
}

// SetCreateElectionPermission is a paid mutator transaction binding the contract method 0x7824c05c.
//
// Solidity: function setCreateElectionPermission(uint256 _communityId, uint8 _createElectionPermission) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) SetCreateElectionPermission(_communityId *big.Int, _createElectionPermission uint8) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetCreateElectionPermission(&_CommunityHubToken.TransactOpts, _communityId, _createElectionPermission)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x82e6d638.
//
// Solidity: function setMetadata(uint256 _communityId, (string,string,string,string[],bool) _metadata) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) SetMetadata(opts *bind.TransactOpts, _communityId *big.Int, _metadata ICommunityHubCommunityMetadata) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "setMetadata", _communityId, _metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x82e6d638.
//
// Solidity: function setMetadata(uint256 _communityId, (string,string,string,string[],bool) _metadata) returns()
func (_CommunityHubToken *CommunityHubTokenSession) SetMetadata(_communityId *big.Int, _metadata ICommunityHubCommunityMetadata) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetMetadata(&_CommunityHubToken.TransactOpts, _communityId, _metadata)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x82e6d638.
//
// Solidity: function setMetadata(uint256 _communityId, (string,string,string,string[],bool) _metadata) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) SetMetadata(_communityId *big.Int, _metadata ICommunityHubCommunityMetadata) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetMetadata(&_CommunityHubToken.TransactOpts, _communityId, _metadata)
}

// SetNotifiableElections is a paid mutator transaction binding the contract method 0xe6160fa3.
//
// Solidity: function setNotifiableElections(uint256 _communityId, bool _notifiableElections) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) SetNotifiableElections(opts *bind.TransactOpts, _communityId *big.Int, _notifiableElections bool) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "setNotifiableElections", _communityId, _notifiableElections)
}

// SetNotifiableElections is a paid mutator transaction binding the contract method 0xe6160fa3.
//
// Solidity: function setNotifiableElections(uint256 _communityId, bool _notifiableElections) returns()
func (_CommunityHubToken *CommunityHubTokenSession) SetNotifiableElections(_communityId *big.Int, _notifiableElections bool) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetNotifiableElections(&_CommunityHubToken.TransactOpts, _communityId, _notifiableElections)
}

// SetNotifiableElections is a paid mutator transaction binding the contract method 0xe6160fa3.
//
// Solidity: function setNotifiableElections(uint256 _communityId, bool _notifiableElections) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) SetNotifiableElections(_communityId *big.Int, _notifiableElections bool) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetNotifiableElections(&_CommunityHubToken.TransactOpts, _communityId, _notifiableElections)
}

// SetResult is a paid mutator transaction binding the contract method 0x3c973f75.
//
// Solidity: function setResult(uint256 _communityId, bytes32 _electionId, (string,string[],string,uint256[][],uint256,uint256,uint256[],bytes32,string) _result) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) SetResult(opts *bind.TransactOpts, _communityId *big.Int, _electionId [32]byte, _result IResultResult) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "setResult", _communityId, _electionId, _result)
}

// SetResult is a paid mutator transaction binding the contract method 0x3c973f75.
//
// Solidity: function setResult(uint256 _communityId, bytes32 _electionId, (string,string[],string,uint256[][],uint256,uint256,uint256[],bytes32,string) _result) returns()
func (_CommunityHubToken *CommunityHubTokenSession) SetResult(_communityId *big.Int, _electionId [32]byte, _result IResultResult) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetResult(&_CommunityHubToken.TransactOpts, _communityId, _electionId, _result)
}

// SetResult is a paid mutator transaction binding the contract method 0x3c973f75.
//
// Solidity: function setResult(uint256 _communityId, bytes32 _electionId, (string,string[],string,uint256[][],uint256,uint256,uint256[],bytes32,string) _result) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) SetResult(_communityId *big.Int, _electionId [32]byte, _result IResultResult) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.SetResult(&_CommunityHubToken.TransactOpts, _communityId, _electionId, _result)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CommunityHubToken *CommunityHubTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.TransferOwnership(&_CommunityHubToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CommunityHubToken.Contract.TransferOwnership(&_CommunityHubToken.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityHubToken.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CommunityHubToken *CommunityHubTokenSession) Withdraw() (*types.Transaction, error) {
	return _CommunityHubToken.Contract.Withdraw(&_CommunityHubToken.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CommunityHubToken.Contract.Withdraw(&_CommunityHubToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CommunityHubToken *CommunityHubTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommunityHubToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CommunityHubToken *CommunityHubTokenSession) Receive() (*types.Transaction, error) {
	return _CommunityHubToken.Contract.Receive(&_CommunityHubToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CommunityHubToken *CommunityHubTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _CommunityHubToken.Contract.Receive(&_CommunityHubToken.TransactOpts)
}

// CommunityHubTokenAdminCommunityManagedIterator is returned from FilterAdminCommunityManaged and is used to iterate over the raw logs and unpacked data for AdminCommunityManaged events raised by the CommunityHubToken contract.
type CommunityHubTokenAdminCommunityManagedIterator struct {
	Event *CommunityHubTokenAdminCommunityManaged // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenAdminCommunityManagedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenAdminCommunityManaged)
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
		it.Event = new(CommunityHubTokenAdminCommunityManaged)
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
func (it *CommunityHubTokenAdminCommunityManagedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenAdminCommunityManagedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenAdminCommunityManaged represents a AdminCommunityManaged event raised by the CommunityHubToken contract.
type CommunityHubTokenAdminCommunityManaged struct {
	CommunityId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAdminCommunityManaged is a free log retrieval operation binding the contract event 0x887cc86755221db77fb9a404d4834a6ca76a28df7c71c7650318819ad0f46a3c.
//
// Solidity: event AdminCommunityManaged(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterAdminCommunityManaged(opts *bind.FilterOpts) (*CommunityHubTokenAdminCommunityManagedIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "AdminCommunityManaged")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenAdminCommunityManagedIterator{contract: _CommunityHubToken.contract, event: "AdminCommunityManaged", logs: logs, sub: sub}, nil
}

// WatchAdminCommunityManaged is a free log subscription operation binding the contract event 0x887cc86755221db77fb9a404d4834a6ca76a28df7c71c7650318819ad0f46a3c.
//
// Solidity: event AdminCommunityManaged(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchAdminCommunityManaged(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenAdminCommunityManaged) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "AdminCommunityManaged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenAdminCommunityManaged)
				if err := _CommunityHubToken.contract.UnpackLog(event, "AdminCommunityManaged", log); err != nil {
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

// ParseAdminCommunityManaged is a log parse operation binding the contract event 0x887cc86755221db77fb9a404d4834a6ca76a28df7c71c7650318819ad0f46a3c.
//
// Solidity: event AdminCommunityManaged(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseAdminCommunityManaged(log types.Log) (*CommunityHubTokenAdminCommunityManaged, error) {
	event := new(CommunityHubTokenAdminCommunityManaged)
	if err := _CommunityHubToken.contract.UnpackLog(event, "AdminCommunityManaged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCensusSetIterator is returned from FilterCensusSet and is used to iterate over the raw logs and unpacked data for CensusSet events raised by the CommunityHubToken contract.
type CommunityHubTokenCensusSetIterator struct {
	Event *CommunityHubTokenCensusSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCensusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCensusSet)
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
		it.Event = new(CommunityHubTokenCensusSet)
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
func (it *CommunityHubTokenCensusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCensusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCensusSet represents a CensusSet event raised by the CommunityHubToken contract.
type CommunityHubTokenCensusSet struct {
	CommunityId *big.Int
	Census      ICommunityHubCensus
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCensusSet is a free log retrieval operation binding the contract event 0xf768bb0dd5bd987cdbc2706c837263b29e7326b0651e4fe20c16372d24cc02ac.
//
// Solidity: event CensusSet(uint256 communityId, (uint8,(string,address)[],string) census)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCensusSet(opts *bind.FilterOpts) (*CommunityHubTokenCensusSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CensusSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCensusSetIterator{contract: _CommunityHubToken.contract, event: "CensusSet", logs: logs, sub: sub}, nil
}

// WatchCensusSet is a free log subscription operation binding the contract event 0xf768bb0dd5bd987cdbc2706c837263b29e7326b0651e4fe20c16372d24cc02ac.
//
// Solidity: event CensusSet(uint256 communityId, (uint8,(string,address)[],string) census)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCensusSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCensusSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CensusSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCensusSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CensusSet", log); err != nil {
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

// ParseCensusSet is a log parse operation binding the contract event 0xf768bb0dd5bd987cdbc2706c837263b29e7326b0651e4fe20c16372d24cc02ac.
//
// Solidity: event CensusSet(uint256 communityId, (uint8,(string,address)[],string) census)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCensusSet(log types.Log) (*CommunityHubTokenCensusSet, error) {
	event := new(CommunityHubTokenCensusSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CensusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCommunityCreatedIterator is returned from FilterCommunityCreated and is used to iterate over the raw logs and unpacked data for CommunityCreated events raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityCreatedIterator struct {
	Event *CommunityHubTokenCommunityCreated // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCommunityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCommunityCreated)
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
		it.Event = new(CommunityHubTokenCommunityCreated)
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
func (it *CommunityHubTokenCommunityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCommunityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCommunityCreated represents a CommunityCreated event raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityCreated struct {
	CommunityId *big.Int
	Creator     common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommunityCreated is a free log retrieval operation binding the contract event 0x42c24a3323433819443a75d0e2651a9c84f696fba638d730042960045ef27241.
//
// Solidity: event CommunityCreated(uint256 communityId, address creator)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCommunityCreated(opts *bind.FilterOpts) (*CommunityHubTokenCommunityCreatedIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CommunityCreated")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCommunityCreatedIterator{contract: _CommunityHubToken.contract, event: "CommunityCreated", logs: logs, sub: sub}, nil
}

// WatchCommunityCreated is a free log subscription operation binding the contract event 0x42c24a3323433819443a75d0e2651a9c84f696fba638d730042960045ef27241.
//
// Solidity: event CommunityCreated(uint256 communityId, address creator)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCommunityCreated(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCommunityCreated) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CommunityCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCommunityCreated)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityCreated", log); err != nil {
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

// ParseCommunityCreated is a log parse operation binding the contract event 0x42c24a3323433819443a75d0e2651a9c84f696fba638d730042960045ef27241.
//
// Solidity: event CommunityCreated(uint256 communityId, address creator)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCommunityCreated(log types.Log) (*CommunityHubTokenCommunityCreated, error) {
	event := new(CommunityHubTokenCommunityCreated)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCommunityDepositIterator is returned from FilterCommunityDeposit and is used to iterate over the raw logs and unpacked data for CommunityDeposit events raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityDepositIterator struct {
	Event *CommunityHubTokenCommunityDeposit // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCommunityDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCommunityDeposit)
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
		it.Event = new(CommunityHubTokenCommunityDeposit)
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
func (it *CommunityHubTokenCommunityDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCommunityDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCommunityDeposit represents a CommunityDeposit event raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityDeposit struct {
	Sender      common.Address
	Amount      *big.Int
	CommunityId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommunityDeposit is a free log retrieval operation binding the contract event 0xd6e9bd49f32ce5202a849c6ab581631b3e952257363b06975f32872e9e1656c9.
//
// Solidity: event CommunityDeposit(address sender, uint256 amount, uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCommunityDeposit(opts *bind.FilterOpts) (*CommunityHubTokenCommunityDepositIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CommunityDeposit")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCommunityDepositIterator{contract: _CommunityHubToken.contract, event: "CommunityDeposit", logs: logs, sub: sub}, nil
}

// WatchCommunityDeposit is a free log subscription operation binding the contract event 0xd6e9bd49f32ce5202a849c6ab581631b3e952257363b06975f32872e9e1656c9.
//
// Solidity: event CommunityDeposit(address sender, uint256 amount, uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCommunityDeposit(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCommunityDeposit) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CommunityDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCommunityDeposit)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityDeposit", log); err != nil {
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

// ParseCommunityDeposit is a log parse operation binding the contract event 0xd6e9bd49f32ce5202a849c6ab581631b3e952257363b06975f32872e9e1656c9.
//
// Solidity: event CommunityDeposit(address sender, uint256 amount, uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCommunityDeposit(log types.Log) (*CommunityHubTokenCommunityDeposit, error) {
	event := new(CommunityHubTokenCommunityDeposit)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCommunityDisabledIterator is returned from FilterCommunityDisabled and is used to iterate over the raw logs and unpacked data for CommunityDisabled events raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityDisabledIterator struct {
	Event *CommunityHubTokenCommunityDisabled // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCommunityDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCommunityDisabled)
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
		it.Event = new(CommunityHubTokenCommunityDisabled)
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
func (it *CommunityHubTokenCommunityDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCommunityDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCommunityDisabled represents a CommunityDisabled event raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityDisabled struct {
	CommunityId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommunityDisabled is a free log retrieval operation binding the contract event 0xd5d88098de5ba30573f08926c4e6445cab93ce09d98c1147a64045608a50797f.
//
// Solidity: event CommunityDisabled(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCommunityDisabled(opts *bind.FilterOpts) (*CommunityHubTokenCommunityDisabledIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CommunityDisabled")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCommunityDisabledIterator{contract: _CommunityHubToken.contract, event: "CommunityDisabled", logs: logs, sub: sub}, nil
}

// WatchCommunityDisabled is a free log subscription operation binding the contract event 0xd5d88098de5ba30573f08926c4e6445cab93ce09d98c1147a64045608a50797f.
//
// Solidity: event CommunityDisabled(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCommunityDisabled(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCommunityDisabled) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CommunityDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCommunityDisabled)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityDisabled", log); err != nil {
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

// ParseCommunityDisabled is a log parse operation binding the contract event 0xd5d88098de5ba30573f08926c4e6445cab93ce09d98c1147a64045608a50797f.
//
// Solidity: event CommunityDisabled(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCommunityDisabled(log types.Log) (*CommunityHubTokenCommunityDisabled, error) {
	event := new(CommunityHubTokenCommunityDisabled)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCommunityEnabledIterator is returned from FilterCommunityEnabled and is used to iterate over the raw logs and unpacked data for CommunityEnabled events raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityEnabledIterator struct {
	Event *CommunityHubTokenCommunityEnabled // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCommunityEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCommunityEnabled)
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
		it.Event = new(CommunityHubTokenCommunityEnabled)
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
func (it *CommunityHubTokenCommunityEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCommunityEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCommunityEnabled represents a CommunityEnabled event raised by the CommunityHubToken contract.
type CommunityHubTokenCommunityEnabled struct {
	CommunityId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommunityEnabled is a free log retrieval operation binding the contract event 0x8da007fc61f3601008aa139c728ea9bf16672668b03c07f6c402ab74c034fc81.
//
// Solidity: event CommunityEnabled(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCommunityEnabled(opts *bind.FilterOpts) (*CommunityHubTokenCommunityEnabledIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CommunityEnabled")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCommunityEnabledIterator{contract: _CommunityHubToken.contract, event: "CommunityEnabled", logs: logs, sub: sub}, nil
}

// WatchCommunityEnabled is a free log subscription operation binding the contract event 0x8da007fc61f3601008aa139c728ea9bf16672668b03c07f6c402ab74c034fc81.
//
// Solidity: event CommunityEnabled(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCommunityEnabled(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCommunityEnabled) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CommunityEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCommunityEnabled)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityEnabled", log); err != nil {
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

// ParseCommunityEnabled is a log parse operation binding the contract event 0x8da007fc61f3601008aa139c728ea9bf16672668b03c07f6c402ab74c034fc81.
//
// Solidity: event CommunityEnabled(uint256 communityId)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCommunityEnabled(log types.Log) (*CommunityHubTokenCommunityEnabled, error) {
	event := new(CommunityHubTokenCommunityEnabled)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CommunityEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCreateCommunityPriceSetIterator is returned from FilterCreateCommunityPriceSet and is used to iterate over the raw logs and unpacked data for CreateCommunityPriceSet events raised by the CommunityHubToken contract.
type CommunityHubTokenCreateCommunityPriceSetIterator struct {
	Event *CommunityHubTokenCreateCommunityPriceSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCreateCommunityPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCreateCommunityPriceSet)
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
		it.Event = new(CommunityHubTokenCreateCommunityPriceSet)
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
func (it *CommunityHubTokenCreateCommunityPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCreateCommunityPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCreateCommunityPriceSet represents a CreateCommunityPriceSet event raised by the CommunityHubToken contract.
type CommunityHubTokenCreateCommunityPriceSet struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateCommunityPriceSet is a free log retrieval operation binding the contract event 0xac9c054628ba106d4664c2c08386354e161eab546a0a47faf040bfc1062845e7.
//
// Solidity: event CreateCommunityPriceSet(uint256 price)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCreateCommunityPriceSet(opts *bind.FilterOpts) (*CommunityHubTokenCreateCommunityPriceSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CreateCommunityPriceSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCreateCommunityPriceSetIterator{contract: _CommunityHubToken.contract, event: "CreateCommunityPriceSet", logs: logs, sub: sub}, nil
}

// WatchCreateCommunityPriceSet is a free log subscription operation binding the contract event 0xac9c054628ba106d4664c2c08386354e161eab546a0a47faf040bfc1062845e7.
//
// Solidity: event CreateCommunityPriceSet(uint256 price)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCreateCommunityPriceSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCreateCommunityPriceSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CreateCommunityPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCreateCommunityPriceSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CreateCommunityPriceSet", log); err != nil {
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

// ParseCreateCommunityPriceSet is a log parse operation binding the contract event 0xac9c054628ba106d4664c2c08386354e161eab546a0a47faf040bfc1062845e7.
//
// Solidity: event CreateCommunityPriceSet(uint256 price)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCreateCommunityPriceSet(log types.Log) (*CommunityHubTokenCreateCommunityPriceSet, error) {
	event := new(CommunityHubTokenCreateCommunityPriceSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CreateCommunityPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenCreateElectionPermissionSetIterator is returned from FilterCreateElectionPermissionSet and is used to iterate over the raw logs and unpacked data for CreateElectionPermissionSet events raised by the CommunityHubToken contract.
type CommunityHubTokenCreateElectionPermissionSetIterator struct {
	Event *CommunityHubTokenCreateElectionPermissionSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenCreateElectionPermissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenCreateElectionPermissionSet)
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
		it.Event = new(CommunityHubTokenCreateElectionPermissionSet)
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
func (it *CommunityHubTokenCreateElectionPermissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenCreateElectionPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenCreateElectionPermissionSet represents a CreateElectionPermissionSet event raised by the CommunityHubToken contract.
type CommunityHubTokenCreateElectionPermissionSet struct {
	CommunityId              *big.Int
	CreateElectionPermission uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterCreateElectionPermissionSet is a free log retrieval operation binding the contract event 0xeaeee81ca4b132c1f2699cf2e7f71c26adb6ffe780097ac7d6e9ddf978398b62.
//
// Solidity: event CreateElectionPermissionSet(uint256 communityId, uint8 createElectionPermission)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterCreateElectionPermissionSet(opts *bind.FilterOpts) (*CommunityHubTokenCreateElectionPermissionSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "CreateElectionPermissionSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenCreateElectionPermissionSetIterator{contract: _CommunityHubToken.contract, event: "CreateElectionPermissionSet", logs: logs, sub: sub}, nil
}

// WatchCreateElectionPermissionSet is a free log subscription operation binding the contract event 0xeaeee81ca4b132c1f2699cf2e7f71c26adb6ffe780097ac7d6e9ddf978398b62.
//
// Solidity: event CreateElectionPermissionSet(uint256 communityId, uint8 createElectionPermission)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchCreateElectionPermissionSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenCreateElectionPermissionSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "CreateElectionPermissionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenCreateElectionPermissionSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "CreateElectionPermissionSet", log); err != nil {
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

// ParseCreateElectionPermissionSet is a log parse operation binding the contract event 0xeaeee81ca4b132c1f2699cf2e7f71c26adb6ffe780097ac7d6e9ddf978398b62.
//
// Solidity: event CreateElectionPermissionSet(uint256 communityId, uint8 createElectionPermission)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseCreateElectionPermissionSet(log types.Log) (*CommunityHubTokenCreateElectionPermissionSet, error) {
	event := new(CommunityHubTokenCreateElectionPermissionSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "CreateElectionPermissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CommunityHubToken contract.
type CommunityHubTokenDepositIterator struct {
	Event *CommunityHubTokenDeposit // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenDeposit)
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
		it.Event = new(CommunityHubTokenDeposit)
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
func (it *CommunityHubTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenDeposit represents a Deposit event raised by the CommunityHubToken contract.
type CommunityHubTokenDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 amount)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterDeposit(opts *bind.FilterOpts) (*CommunityHubTokenDepositIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenDepositIterator{contract: _CommunityHubToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 amount)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenDeposit) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenDeposit)
				if err := _CommunityHubToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 amount)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseDeposit(log types.Log) (*CommunityHubTokenDeposit, error) {
	event := new(CommunityHubTokenDeposit)
	if err := _CommunityHubToken.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenGuardianAddedIterator is returned from FilterGuardianAdded and is used to iterate over the raw logs and unpacked data for GuardianAdded events raised by the CommunityHubToken contract.
type CommunityHubTokenGuardianAddedIterator struct {
	Event *CommunityHubTokenGuardianAdded // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenGuardianAdded)
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
		it.Event = new(CommunityHubTokenGuardianAdded)
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
func (it *CommunityHubTokenGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenGuardianAdded represents a GuardianAdded event raised by the CommunityHubToken contract.
type CommunityHubTokenGuardianAdded struct {
	CommunityId *big.Int
	Guardian    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdded is a free log retrieval operation binding the contract event 0xfcdfd5aeb97f499ca134ed43f010f2a4f5b0ab73d317ac27246a066a050a73af.
//
// Solidity: event GuardianAdded(uint256 communityId, uint256 guardian)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterGuardianAdded(opts *bind.FilterOpts) (*CommunityHubTokenGuardianAddedIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "GuardianAdded")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenGuardianAddedIterator{contract: _CommunityHubToken.contract, event: "GuardianAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianAdded is a free log subscription operation binding the contract event 0xfcdfd5aeb97f499ca134ed43f010f2a4f5b0ab73d317ac27246a066a050a73af.
//
// Solidity: event GuardianAdded(uint256 communityId, uint256 guardian)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchGuardianAdded(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenGuardianAdded) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "GuardianAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenGuardianAdded)
				if err := _CommunityHubToken.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
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

// ParseGuardianAdded is a log parse operation binding the contract event 0xfcdfd5aeb97f499ca134ed43f010f2a4f5b0ab73d317ac27246a066a050a73af.
//
// Solidity: event GuardianAdded(uint256 communityId, uint256 guardian)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseGuardianAdded(log types.Log) (*CommunityHubTokenGuardianAdded, error) {
	event := new(CommunityHubTokenGuardianAdded)
	if err := _CommunityHubToken.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenGuardianRemovedIterator is returned from FilterGuardianRemoved and is used to iterate over the raw logs and unpacked data for GuardianRemoved events raised by the CommunityHubToken contract.
type CommunityHubTokenGuardianRemovedIterator struct {
	Event *CommunityHubTokenGuardianRemoved // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenGuardianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenGuardianRemoved)
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
		it.Event = new(CommunityHubTokenGuardianRemoved)
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
func (it *CommunityHubTokenGuardianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenGuardianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenGuardianRemoved represents a GuardianRemoved event raised by the CommunityHubToken contract.
type CommunityHubTokenGuardianRemoved struct {
	CommunityId *big.Int
	Guardian    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGuardianRemoved is a free log retrieval operation binding the contract event 0x2264fa092e5819982f4edca5b7e9c6318a16dc3e03673510429c078656dea45c.
//
// Solidity: event GuardianRemoved(uint256 communityId, uint256 guardian)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterGuardianRemoved(opts *bind.FilterOpts) (*CommunityHubTokenGuardianRemovedIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "GuardianRemoved")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenGuardianRemovedIterator{contract: _CommunityHubToken.contract, event: "GuardianRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardianRemoved is a free log subscription operation binding the contract event 0x2264fa092e5819982f4edca5b7e9c6318a16dc3e03673510429c078656dea45c.
//
// Solidity: event GuardianRemoved(uint256 communityId, uint256 guardian)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchGuardianRemoved(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenGuardianRemoved) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "GuardianRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenGuardianRemoved)
				if err := _CommunityHubToken.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
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

// ParseGuardianRemoved is a log parse operation binding the contract event 0x2264fa092e5819982f4edca5b7e9c6318a16dc3e03673510429c078656dea45c.
//
// Solidity: event GuardianRemoved(uint256 communityId, uint256 guardian)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseGuardianRemoved(log types.Log) (*CommunityHubTokenGuardianRemoved, error) {
	event := new(CommunityHubTokenGuardianRemoved)
	if err := _CommunityHubToken.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenMetadataSetIterator is returned from FilterMetadataSet and is used to iterate over the raw logs and unpacked data for MetadataSet events raised by the CommunityHubToken contract.
type CommunityHubTokenMetadataSetIterator struct {
	Event *CommunityHubTokenMetadataSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenMetadataSet)
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
		it.Event = new(CommunityHubTokenMetadataSet)
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
func (it *CommunityHubTokenMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenMetadataSet represents a MetadataSet event raised by the CommunityHubToken contract.
type CommunityHubTokenMetadataSet struct {
	CommunityId *big.Int
	Metadata    ICommunityHubCommunityMetadata
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMetadataSet is a free log retrieval operation binding the contract event 0x440c55e385a61cd472c0cd76ee015cab9bd0760aef9634633e6597fdea583fd5.
//
// Solidity: event MetadataSet(uint256 communityId, (string,string,string,string[],bool) metadata)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterMetadataSet(opts *bind.FilterOpts) (*CommunityHubTokenMetadataSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "MetadataSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenMetadataSetIterator{contract: _CommunityHubToken.contract, event: "MetadataSet", logs: logs, sub: sub}, nil
}

// WatchMetadataSet is a free log subscription operation binding the contract event 0x440c55e385a61cd472c0cd76ee015cab9bd0760aef9634633e6597fdea583fd5.
//
// Solidity: event MetadataSet(uint256 communityId, (string,string,string,string[],bool) metadata)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchMetadataSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenMetadataSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "MetadataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenMetadataSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "MetadataSet", log); err != nil {
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

// ParseMetadataSet is a log parse operation binding the contract event 0x440c55e385a61cd472c0cd76ee015cab9bd0760aef9634633e6597fdea583fd5.
//
// Solidity: event MetadataSet(uint256 communityId, (string,string,string,string[],bool) metadata)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseMetadataSet(log types.Log) (*CommunityHubTokenMetadataSet, error) {
	event := new(CommunityHubTokenMetadataSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "MetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenNotifiableElectionsSetIterator is returned from FilterNotifiableElectionsSet and is used to iterate over the raw logs and unpacked data for NotifiableElectionsSet events raised by the CommunityHubToken contract.
type CommunityHubTokenNotifiableElectionsSetIterator struct {
	Event *CommunityHubTokenNotifiableElectionsSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenNotifiableElectionsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenNotifiableElectionsSet)
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
		it.Event = new(CommunityHubTokenNotifiableElectionsSet)
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
func (it *CommunityHubTokenNotifiableElectionsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenNotifiableElectionsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenNotifiableElectionsSet represents a NotifiableElectionsSet event raised by the CommunityHubToken contract.
type CommunityHubTokenNotifiableElectionsSet struct {
	CommunityId         *big.Int
	NotifiableElections bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNotifiableElectionsSet is a free log retrieval operation binding the contract event 0xed2f2c7a2316ddc2f46a4581250611f807810052bb38bef2fbd1a81d087e064b.
//
// Solidity: event NotifiableElectionsSet(uint256 communityId, bool notifiableElections)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterNotifiableElectionsSet(opts *bind.FilterOpts) (*CommunityHubTokenNotifiableElectionsSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "NotifiableElectionsSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenNotifiableElectionsSetIterator{contract: _CommunityHubToken.contract, event: "NotifiableElectionsSet", logs: logs, sub: sub}, nil
}

// WatchNotifiableElectionsSet is a free log subscription operation binding the contract event 0xed2f2c7a2316ddc2f46a4581250611f807810052bb38bef2fbd1a81d087e064b.
//
// Solidity: event NotifiableElectionsSet(uint256 communityId, bool notifiableElections)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchNotifiableElectionsSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenNotifiableElectionsSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "NotifiableElectionsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenNotifiableElectionsSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "NotifiableElectionsSet", log); err != nil {
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

// ParseNotifiableElectionsSet is a log parse operation binding the contract event 0xed2f2c7a2316ddc2f46a4581250611f807810052bb38bef2fbd1a81d087e064b.
//
// Solidity: event NotifiableElectionsSet(uint256 communityId, bool notifiableElections)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseNotifiableElectionsSet(log types.Log) (*CommunityHubTokenNotifiableElectionsSet, error) {
	event := new(CommunityHubTokenNotifiableElectionsSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "NotifiableElectionsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CommunityHubToken contract.
type CommunityHubTokenOwnershipTransferredIterator struct {
	Event *CommunityHubTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenOwnershipTransferred)
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
		it.Event = new(CommunityHubTokenOwnershipTransferred)
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
func (it *CommunityHubTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CommunityHubToken contract.
type CommunityHubTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CommunityHubTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenOwnershipTransferredIterator{contract: _CommunityHubToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenOwnershipTransferred)
				if err := _CommunityHubToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseOwnershipTransferred(log types.Log) (*CommunityHubTokenOwnershipTransferred, error) {
	event := new(CommunityHubTokenOwnershipTransferred)
	if err := _CommunityHubToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenPricePerElectionSetIterator is returned from FilterPricePerElectionSet and is used to iterate over the raw logs and unpacked data for PricePerElectionSet events raised by the CommunityHubToken contract.
type CommunityHubTokenPricePerElectionSetIterator struct {
	Event *CommunityHubTokenPricePerElectionSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenPricePerElectionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenPricePerElectionSet)
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
		it.Event = new(CommunityHubTokenPricePerElectionSet)
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
func (it *CommunityHubTokenPricePerElectionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenPricePerElectionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenPricePerElectionSet represents a PricePerElectionSet event raised by the CommunityHubToken contract.
type CommunityHubTokenPricePerElectionSet struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPricePerElectionSet is a free log retrieval operation binding the contract event 0x0840c2296c4c97ecdc9b8dc544cd67f45a8d629b9d6109410530bb58b7f17b92.
//
// Solidity: event PricePerElectionSet(uint256 price)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterPricePerElectionSet(opts *bind.FilterOpts) (*CommunityHubTokenPricePerElectionSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "PricePerElectionSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenPricePerElectionSetIterator{contract: _CommunityHubToken.contract, event: "PricePerElectionSet", logs: logs, sub: sub}, nil
}

// WatchPricePerElectionSet is a free log subscription operation binding the contract event 0x0840c2296c4c97ecdc9b8dc544cd67f45a8d629b9d6109410530bb58b7f17b92.
//
// Solidity: event PricePerElectionSet(uint256 price)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchPricePerElectionSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenPricePerElectionSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "PricePerElectionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenPricePerElectionSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "PricePerElectionSet", log); err != nil {
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

// ParsePricePerElectionSet is a log parse operation binding the contract event 0x0840c2296c4c97ecdc9b8dc544cd67f45a8d629b9d6109410530bb58b7f17b92.
//
// Solidity: event PricePerElectionSet(uint256 price)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParsePricePerElectionSet(log types.Log) (*CommunityHubTokenPricePerElectionSet, error) {
	event := new(CommunityHubTokenPricePerElectionSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "PricePerElectionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenResultsSetIterator is returned from FilterResultsSet and is used to iterate over the raw logs and unpacked data for ResultsSet events raised by the CommunityHubToken contract.
type CommunityHubTokenResultsSetIterator struct {
	Event *CommunityHubTokenResultsSet // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenResultsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenResultsSet)
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
		it.Event = new(CommunityHubTokenResultsSet)
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
func (it *CommunityHubTokenResultsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenResultsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenResultsSet represents a ResultsSet event raised by the CommunityHubToken contract.
type CommunityHubTokenResultsSet struct {
	CommunityId *big.Int
	ElectionId  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterResultsSet is a free log retrieval operation binding the contract event 0x77deb32519991dda7821b0b9367f9124aa3fd934f4845bcfd5dc6fa3f1922663.
//
// Solidity: event ResultsSet(uint256 communityId, bytes32 electionId)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterResultsSet(opts *bind.FilterOpts) (*CommunityHubTokenResultsSetIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "ResultsSet")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenResultsSetIterator{contract: _CommunityHubToken.contract, event: "ResultsSet", logs: logs, sub: sub}, nil
}

// WatchResultsSet is a free log subscription operation binding the contract event 0x77deb32519991dda7821b0b9367f9124aa3fd934f4845bcfd5dc6fa3f1922663.
//
// Solidity: event ResultsSet(uint256 communityId, bytes32 electionId)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchResultsSet(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenResultsSet) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "ResultsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenResultsSet)
				if err := _CommunityHubToken.contract.UnpackLog(event, "ResultsSet", log); err != nil {
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

// ParseResultsSet is a log parse operation binding the contract event 0x77deb32519991dda7821b0b9367f9124aa3fd934f4845bcfd5dc6fa3f1922663.
//
// Solidity: event ResultsSet(uint256 communityId, bytes32 electionId)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseResultsSet(log types.Log) (*CommunityHubTokenResultsSet, error) {
	event := new(CommunityHubTokenResultsSet)
	if err := _CommunityHubToken.contract.UnpackLog(event, "ResultsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommunityHubTokenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the CommunityHubToken contract.
type CommunityHubTokenWithdrawalIterator struct {
	Event *CommunityHubTokenWithdrawal // Event containing the contract specifics and raw log

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
func (it *CommunityHubTokenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommunityHubTokenWithdrawal)
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
		it.Event = new(CommunityHubTokenWithdrawal)
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
func (it *CommunityHubTokenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommunityHubTokenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommunityHubTokenWithdrawal represents a Withdrawal event raised by the CommunityHubToken contract.
type CommunityHubTokenWithdrawal struct {
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xd964a27d45f595739c13d8b1160b57491050cacf3a2e5602207277d6228f64ee.
//
// Solidity: event Withdrawal(uint256 amount, address to)
func (_CommunityHubToken *CommunityHubTokenFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*CommunityHubTokenWithdrawalIterator, error) {

	logs, sub, err := _CommunityHubToken.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &CommunityHubTokenWithdrawalIterator{contract: _CommunityHubToken.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xd964a27d45f595739c13d8b1160b57491050cacf3a2e5602207277d6228f64ee.
//
// Solidity: event Withdrawal(uint256 amount, address to)
func (_CommunityHubToken *CommunityHubTokenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *CommunityHubTokenWithdrawal) (event.Subscription, error) {

	logs, sub, err := _CommunityHubToken.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommunityHubTokenWithdrawal)
				if err := _CommunityHubToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xd964a27d45f595739c13d8b1160b57491050cacf3a2e5602207277d6228f64ee.
//
// Solidity: event Withdrawal(uint256 amount, address to)
func (_CommunityHubToken *CommunityHubTokenFilterer) ParseWithdrawal(log types.Log) (*CommunityHubTokenWithdrawal, error) {
	event := new(CommunityHubTokenWithdrawal)
	if err := _CommunityHubToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

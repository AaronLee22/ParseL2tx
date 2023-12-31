// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cannonical

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

// LibOVMCodecQueueElement is an auto generated low-level Go binding around an user-defined struct.
type LibOVMCodecQueueElement struct {
	TransactionHash [32]byte
	Timestamp       *big.Int
	BlockNumber     *big.Int
}

// CannonicalMetaData contains all meta data concerning the Cannonical contract.
var CannonicalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_libAddressManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxTransactionGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_enqueueGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enqueueGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"enqueueL2GasPrepaid\",\"type\":\"uint256\"}],\"name\":\"L2GasParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingQueueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numQueueElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"name\":\"QueueBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingQueueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numQueueElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"name\":\"SequencerBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_prevTotalElements\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"TransactionBatchAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1TxOrigin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"TransactionEnqueued\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ROLLUP_TX_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_ROLLUP_TX_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appendSequencerBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"contractIChainStorageContainer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"enqueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enqueueGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enqueueL2GasPrepaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastTimestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextQueueIndex\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumPendingQueueElements\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getQueueElement\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"}],\"internalType\":\"structLib_OVMCodec.QueueElement\",\"name\":\"_element\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalBatches\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalElements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalElements\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2GasDiscountDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"libAddressManager\",\"outputs\":[{\"internalType\":\"contractLib_AddressManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTransactionGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2GasDiscountDivisor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_enqueueGasCost\",\"type\":\"uint256\"}],\"name\":\"setGasParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200259e3803806200259e833981810160405281019062000036919062000147565b83805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600481905550816002819055508060018190555080826200009a9190620001e3565b600381905550505050506200022d565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620000d982620000ae565b9050919050565b620000eb81620000cd565b8114620000f6575f80fd5b50565b5f815190506200010981620000e0565b92915050565b5f819050919050565b62000123816200010f565b81146200012e575f80fd5b50565b5f81519050620001418162000118565b92915050565b5f805f8060808587031215620001625762000161620000aa565b5b5f6200017187828801620000f9565b9450506020620001848782880162000131565b9350506040620001978782880162000131565b9250506060620001aa8782880162000131565b91505092959194509250565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f620001ef826200010f565b9150620001fc836200010f565b92508282026200020c816200010f565b91508282048414831517620002265762000225620001b6565b5b5092915050565b612363806200023b5f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c8063876ed5cb116100ab578063d0f893441161006f578063d0f8934414610312578063e561dddc1461031c578063e654b1fb1461033a578063edcc4a4514610358578063f722b41a146103745761012a565b8063876ed5cb1461027c5780638d38c6c11461029a578063b8f77005146102b8578063ccf987c8146102d6578063cfdf677e146102f45761012a565b80635ae6256d116100f25780635ae6256d146101e85780636fee07e01461020657806378f4b2f2146102225780637a167a8a146102405780637aa63a861461025e5761012a565b80630b3dfa971461012e578063299ca4781461014c5780632a7f18be1461016a578063378997701461019a578063461a4478146101b8575b5f80fd5b610136610392565b60405161014391906112b2565b60405180910390f35b610154610398565b6040516101619190611345565b60405180910390f35b610184600480360381019061017f9190611399565b6103bb565b604051610191919061143b565b60405180910390f35b6101a261044f565b6040516101af9190611463565b60405180910390f35b6101d260048036038101906101cd91906115b8565b610465565b6040516101df919061161f565b60405180910390f35b6101f0610505565b6040516101fd9190611463565b60405180910390f35b610220600480360381019061021b9190611700565b61051b565b005b61022a610831565b60405161023791906112b2565b60405180910390f35b610248610838565b6040516102559190611463565b60405180910390f35b610266610851565b60405161027391906112b2565b60405180910390f35b61028461086e565b60405161029191906112b2565b60405180910390f35b6102a2610874565b6040516102af91906112b2565b60405180910390f35b6102c061087a565b6040516102cd9190611463565b60405180910390f35b6102de610886565b6040516102eb91906112b2565b60405180910390f35b6102fc61088c565b604051610309919061178c565b60405180910390f35b61031a6108b3565b005b610324610c87565b60405161033191906112b2565b60405180910390f35b610342610d01565b60405161034f91906112b2565b60405180910390f35b610372600480360381019061036d91906117a5565b610d07565b005b61037c610e65565b6040516103899190611463565b60405180910390f35b60035481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103c3611247565b600682815481106103d7576103d66117e3565b5b905f5260205f2090600202016040518060600160405290815f8201548152602001600182015f9054906101000a900464ffffffffff1664ffffffffff1664ffffffffff1681526020016001820160059054906101000a900464ffffffffff1664ffffffffff1664ffffffffff16815250509050919050565b5f80610459610e8e565b50925050508091505090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bf40fac1836040518263ffffffff1660e01b81526004016104bf919061188a565b602060405180830381865afa1580156104da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104fe91906118be565b9050919050565b5f8061050f610e8e565b93505050508091505090565b61c35081511115610561576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055890611959565b60405180910390fd5b6004548211156105a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059d906119e7565b60405180910390fd5b620186a08210156105ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e390611a75565b60405180910390fd5b600354821115610686575f600254600354846106089190611ac0565b6106129190611b20565b90505f5a905081811161065a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065190611bc0565b60405180910390fd5b5f5b825a836106699190611ac0565b101561068257808061067a90611bde565b91505061065c565b5050505b5f3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036106c2573390506106ce565b6106cb33610f79565b90505b5f818585856040516020016106e69493929190611c77565b604051602081830303815290604052805190602001209050600660405180606001604052808381526020014264ffffffffff1681526020014364ffffffffff16815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015f6101000a81548164ffffffffff021916908364ffffffffff16021790555060408201518160010160056101000a81548164ffffffffff021916908364ffffffffff16021790555050505f60016006805490506107bd9190611ac0565b9050808673ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb588884260405161082193929190611cc1565b60405180910390a4505050505050565b620186a081565b5f60055f9054906101000a900464ffffffffff16905090565b5f8061085b610e8e565b50505090508064ffffffffff1691505090565b61c35081565b60045481565b5f600680549050905090565b60025481565b5f6108ae60405180606001604052806021815260200161230d60219139610465565b905090565b5f805f60043560d81c925060093560e81c9150600c3560e81c90506108d6610851565b8364ffffffffff161461091e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091590611d6d565b60405180910390fd5b61095c6040518060400160405280600d81526020017f4f564d5f53657175656e63657200000000000000000000000000000000000000815250610465565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c090611dfb565b60405180910390fd5b5f8162ffffff1660106109dc9190611e19565b600f6109e89190611e5a565b90508064ffffffffff165f3690501015610a37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2e90611efd565b60405180910390fd5b5f8060055f9054906101000a900464ffffffffff169050610a56611276565b5f5b8562ffffff168163ffffffff161015610ab1575f610a7b8263ffffffff16610f98565b9050809250825f015185610a8f9190611f2a565b9450826020015184610aa19190611f61565b9350508080600101915050610a58565b506006805490508264ffffffffff161115610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af89061202f565b60405180910390fd5b5f838762ffffff16610b13919061204d565b63ffffffff1690505f805f846020015103610b3b578360400151915083606001519050610be6565b5f6006600187610b4b9190612084565b64ffffffffff1681548110610b6357610b626117e3565b5b905f5260205f2090600202016040518060600160405290815f8201548152602001600182015f9054906101000a900464ffffffffff1664ffffffffff1664ffffffffff1681526020016001820160059054906101000a900464ffffffffff1664ffffffffff1664ffffffffff168152505090508060200151925080604001519150505b610c0c600143610bf69190611ac0565b408a62ffffff168564ffffffffff168585611012565b7f602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f8998386610c399190612084565b84610c42610851565b604051610c51939291906120ec565b60405180910390a18460055f6101000a81548164ffffffffff021916908364ffffffffff16021790555050505050505050505050565b5f610c9061088c565b73ffffffffffffffffffffffffffffffffffffffff16631f7b6d326040518163ffffffff1660e01b8152600401602060405180830381865afa158015610cd8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cfc9190612135565b905090565b60015481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d6f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d9391906118be565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e00576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df7906121aa565b60405180910390fd5b80600181905550816002819055508082610e1a9190611e19565b6003819055507fc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e600254600154600354604051610e59939291906121c8565b60405180910390a15050565b5f60055f9054906101000a900464ffffffffff16600680549050610e899190612084565b905090565b5f805f805f610e9b61088c565b73ffffffffffffffffffffffffffffffffffffffff1663ccf8f9696040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f079190612252565b90505f805f808460281c945064ffffffffff8516935069ffffffffff0000000000851660281c92506effffffffff00000000000000000000851660501c915073ffffffffff000000000000000000000000000000851660781c9050838383839850985098509850505050505090919293565b5f73111100000000000000000000000000000000111182019050919050565b610fa0611276565b5f601083610fae9190611e19565b600f610fba9190611e5a565b90505f805f80843560e81c9350600385013560e81c9250600685013560d81c9150600b85013560d81c905060405180608001604052808581526020018481526020018381526020018281525095505050505050919050565b5f61101b61088c565b90505f80611027610e8e565b5050915091505f6040518060a001604052808573ffffffffffffffffffffffffffffffffffffffff16631f7b6d326040518163ffffffff1660e01b8152600401602060405180830381865afa158015611082573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110a69190612135565b81526020018a81526020018981526020018464ffffffffff16815260200160405180602001604052805f8152508152509050805f01517f127186556e7be68c7e31263195225b4de02820707889540969f62c05cf73525e8260200151836040015184606001518560800151604051611121949392919061228c565b60405180910390a25f611133826111d3565b90505f61115c8360400151866111499190611f61565b8a866111559190611f61565b8a8a611218565b90508573ffffffffffffffffffffffffffffffffffffffff16632015276c83836040518363ffffffff1660e01b81526004016111999291906122e5565b5f604051808303815f87803b1580156111b0575f80fd5b505af11580156111c2573d5f803e3d5ffd5b505050505050505050505050505050565b5f81602001518260400151836060015184608001516040516020016111fb949392919061228c565b604051602081830303815290604052805190602001209050919050565b5f808590508460281b811790508360501b811790508260781b811790508060281b905080915050949350505050565b60405180606001604052805f80191681526020015f64ffffffffff1681526020015f64ffffffffff1681525090565b60405180608001604052805f81526020015f81526020015f81526020015f81525090565b5f819050919050565b6112ac8161129a565b82525050565b5f6020820190506112c55f8301846112a3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61130d611308611303846112cb565b6112ea565b6112cb565b9050919050565b5f61131e826112f3565b9050919050565b5f61132f82611314565b9050919050565b61133f81611325565b82525050565b5f6020820190506113585f830184611336565b92915050565b5f604051905090565b5f80fd5b5f80fd5b6113788161129a565b8114611382575f80fd5b50565b5f813590506113938161136f565b92915050565b5f602082840312156113ae576113ad611367565b5b5f6113bb84828501611385565b91505092915050565b5f819050919050565b6113d6816113c4565b82525050565b5f64ffffffffff82169050919050565b6113f5816113dc565b82525050565b606082015f82015161140f5f8501826113cd565b50602082015161142260208501826113ec565b50604082015161143560408501826113ec565b50505050565b5f60608201905061144e5f8301846113fb565b92915050565b61145d816113dc565b82525050565b5f6020820190506114765f830184611454565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6114ca82611484565b810181811067ffffffffffffffff821117156114e9576114e8611494565b5b80604052505050565b5f6114fb61135e565b905061150782826114c1565b919050565b5f67ffffffffffffffff82111561152657611525611494565b5b61152f82611484565b9050602081019050919050565b828183375f83830152505050565b5f61155c6115578461150c565b6114f2565b90508281526020810184848401111561157857611577611480565b5b61158384828561153c565b509392505050565b5f82601f83011261159f5761159e61147c565b5b81356115af84826020860161154a565b91505092915050565b5f602082840312156115cd576115cc611367565b5b5f82013567ffffffffffffffff8111156115ea576115e961136b565b5b6115f68482850161158b565b91505092915050565b5f611609826112cb565b9050919050565b611619816115ff565b82525050565b5f6020820190506116325f830184611610565b92915050565b611641816115ff565b811461164b575f80fd5b50565b5f8135905061165c81611638565b92915050565b5f67ffffffffffffffff82111561167c5761167b611494565b5b61168582611484565b9050602081019050919050565b5f6116a461169f84611662565b6114f2565b9050828152602081018484840111156116c0576116bf611480565b5b6116cb84828561153c565b509392505050565b5f82601f8301126116e7576116e661147c565b5b81356116f7848260208601611692565b91505092915050565b5f805f6060848603121561171757611716611367565b5b5f6117248682870161164e565b935050602061173586828701611385565b925050604084013567ffffffffffffffff8111156117565761175561136b565b5b611762868287016116d3565b9150509250925092565b5f61177682611314565b9050919050565b6117868161176c565b82525050565b5f60208201905061179f5f83018461177d565b92915050565b5f80604083850312156117bb576117ba611367565b5b5f6117c885828601611385565b92505060206117d985828601611385565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561184757808201518184015260208101905061182c565b5f8484015250505050565b5f61185c82611810565b611866818561181a565b935061187681856020860161182a565b61187f81611484565b840191505092915050565b5f6020820190508181035f8301526118a28184611852565b905092915050565b5f815190506118b881611638565b92915050565b5f602082840312156118d3576118d2611367565b5b5f6118e0848285016118aa565b91505092915050565b7f5472616e73616374696f6e20646174612073697a652065786365656473206d615f8201527f78696d756d20666f7220726f6c6c7570207472616e73616374696f6e2e000000602082015250565b5f611943603d8361181a565b915061194e826118e9565b604082019050919050565b5f6020820190508181035f83015261197081611937565b9050919050565b7f5472616e73616374696f6e20676173206c696d69742065786365656473206d615f8201527f78696d756d20666f7220726f6c6c7570207472616e73616374696f6e2e000000602082015250565b5f6119d1603d8361181a565b91506119dc82611977565b604082019050919050565b5f6020820190508181035f8301526119fe816119c5565b9050919050565b7f5472616e73616374696f6e20676173206c696d697420746f6f206c6f7720746f5f8201527f20656e71756575652e0000000000000000000000000000000000000000000000602082015250565b5f611a5f60298361181a565b9150611a6a82611a05565b604082019050919050565b5f6020820190508181035f830152611a8c81611a53565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611aca8261129a565b9150611ad58361129a565b9250828203905081811115611aed57611aec611a93565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611b2a8261129a565b9150611b358361129a565b925082611b4557611b44611af3565b5b828204905092915050565b7f496e73756666696369656e742067617320666f72204c322072617465206c696d5f8201527f6974696e67206275726e2e000000000000000000000000000000000000000000602082015250565b5f611baa602b8361181a565b9150611bb582611b50565b604082019050919050565b5f6020820190508181035f830152611bd781611b9e565b9050919050565b5f611be88261129a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c1a57611c19611a93565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f611c4982611c25565b611c538185611c2f565b9350611c6381856020860161182a565b611c6c81611484565b840191505092915050565b5f608082019050611c8a5f830187611610565b611c976020830186611610565b611ca460408301856112a3565b8181036060830152611cb68184611c3f565b905095945050505050565b5f606082019050611cd45f8301866112a3565b8181036020830152611ce68185611c3f565b9050611cf560408301846112a3565b949350505050565b7f41637475616c20626174636820737461727420696e64657820646f6573206e6f5f8201527f74206d6174636820657870656374656420737461727420696e6465782e000000602082015250565b5f611d57603d8361181a565b9150611d6282611cfd565b604082019050919050565b5f6020820190508181035f830152611d8481611d4b565b9050919050565b7f46756e6374696f6e2063616e206f6e6c792062652063616c6c656420627920745f8201527f68652053657175656e6365722e00000000000000000000000000000000000000602082015250565b5f611de5602d8361181a565b9150611df082611d8b565b604082019050919050565b5f6020820190508181035f830152611e1281611dd9565b9050919050565b5f611e238261129a565b9150611e2e8361129a565b9250828202611e3c8161129a565b91508282048414831517611e5357611e52611a93565b5b5092915050565b5f611e648261129a565b9150611e6f8361129a565b9250828201905080821115611e8757611e86611a93565b5b92915050565b7f4e6f7420656e6f756768204261746368436f6e74657874732070726f766964655f8201527f642e000000000000000000000000000000000000000000000000000000000000602082015250565b5f611ee760228361181a565b9150611ef282611e8d565b604082019050919050565b5f6020820190508181035f830152611f1481611edb565b9050919050565b5f63ffffffff82169050919050565b5f611f3482611f1b565b9150611f3f83611f1b565b9250828201905063ffffffff811115611f5b57611f5a611a93565b5b92915050565b5f611f6b826113dc565b9150611f76836113dc565b9250828201905064ffffffffff811115611f9357611f92611a93565b5b92915050565b7f417474656d7074656420746f20617070656e64206d6f726520656c656d656e745f8201527f73207468616e2061726520617661696c61626c6520696e20746865207175657560208201527f652e000000000000000000000000000000000000000000000000000000000000604082015250565b5f61201960428361181a565b915061202482611f99565b606082019050919050565b5f6020820190508181035f8301526120468161200d565b9050919050565b5f61205782611f1b565b915061206283611f1b565b9250828203905063ffffffff81111561207e5761207d611a93565b5b92915050565b5f61208e826113dc565b9150612099836113dc565b9250828203905064ffffffffff8111156120b6576120b5611a93565b5b92915050565b5f6120d66120d16120cc846113dc565b6112ea565b61129a565b9050919050565b6120e6816120bc565b82525050565b5f6060820190506120ff5f8301866120dd565b61210c60208301856120dd565b61211960408301846112a3565b949350505050565b5f8151905061212f8161136f565b92915050565b5f6020828403121561214a57612149611367565b5b5f61215784828501612121565b91505092915050565b7f4f6e6c792063616c6c61626c6520627920746865204275726e2041646d696e2e5f82015250565b5f61219460208361181a565b915061219f82612160565b602082019050919050565b5f6020820190508181035f8301526121c181612188565b9050919050565b5f6060820190506121db5f8301866112a3565b6121e860208301856112a3565b6121f560408301846112a3565b949350505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000082169050919050565b612231816121fd565b811461223b575f80fd5b50565b5f8151905061224c81612228565b92915050565b5f6020828403121561226757612266611367565b5b5f6122748482850161223e565b91505092915050565b612286816113c4565b82525050565b5f60808201905061229f5f83018761227d565b6122ac60208301866112a3565b6122b960408301856112a3565b81810360608301526122cb8184611c3f565b905095945050505050565b6122df816121fd565b82525050565b5f6040820190506122f85f83018561227d565b61230560208301846122d6565b939250505056fe436861696e53746f72616765436f6e7461696e65722d4354432d62617463686573a26469706673582212205af07a87c10c0c07ca59ada348dcdd44389cbd222b456f03280ff9df008c7de264736f6c63430008170033",
}

// CannonicalABI is the input ABI used to generate the binding from.
// Deprecated: Use CannonicalMetaData.ABI instead.
var CannonicalABI = CannonicalMetaData.ABI

// CannonicalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CannonicalMetaData.Bin instead.
var CannonicalBin = CannonicalMetaData.Bin

// DeployCannonical deploys a new Ethereum contract, binding an instance of Cannonical to it.
func DeployCannonical(auth *bind.TransactOpts, backend bind.ContractBackend, _libAddressManager common.Address, _maxTransactionGasLimit *big.Int, _l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (common.Address, *types.Transaction, *Cannonical, error) {
	parsed, err := CannonicalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CannonicalBin), backend, _libAddressManager, _maxTransactionGasLimit, _l2GasDiscountDivisor, _enqueueGasCost)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cannonical{CannonicalCaller: CannonicalCaller{contract: contract}, CannonicalTransactor: CannonicalTransactor{contract: contract}, CannonicalFilterer: CannonicalFilterer{contract: contract}}, nil
}

// Cannonical is an auto generated Go binding around an Ethereum contract.
type Cannonical struct {
	CannonicalCaller     // Read-only binding to the contract
	CannonicalTransactor // Write-only binding to the contract
	CannonicalFilterer   // Log filterer for contract events
}

// CannonicalCaller is an auto generated read-only Go binding around an Ethereum contract.
type CannonicalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CannonicalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CannonicalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CannonicalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CannonicalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CannonicalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CannonicalSession struct {
	Contract     *Cannonical       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CannonicalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CannonicalCallerSession struct {
	Contract *CannonicalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CannonicalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CannonicalTransactorSession struct {
	Contract     *CannonicalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CannonicalRaw is an auto generated low-level Go binding around an Ethereum contract.
type CannonicalRaw struct {
	Contract *Cannonical // Generic contract binding to access the raw methods on
}

// CannonicalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CannonicalCallerRaw struct {
	Contract *CannonicalCaller // Generic read-only contract binding to access the raw methods on
}

// CannonicalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CannonicalTransactorRaw struct {
	Contract *CannonicalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCannonical creates a new instance of Cannonical, bound to a specific deployed contract.
func NewCannonical(address common.Address, backend bind.ContractBackend) (*Cannonical, error) {
	contract, err := bindCannonical(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cannonical{CannonicalCaller: CannonicalCaller{contract: contract}, CannonicalTransactor: CannonicalTransactor{contract: contract}, CannonicalFilterer: CannonicalFilterer{contract: contract}}, nil
}

// NewCannonicalCaller creates a new read-only instance of Cannonical, bound to a specific deployed contract.
func NewCannonicalCaller(address common.Address, caller bind.ContractCaller) (*CannonicalCaller, error) {
	contract, err := bindCannonical(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CannonicalCaller{contract: contract}, nil
}

// NewCannonicalTransactor creates a new write-only instance of Cannonical, bound to a specific deployed contract.
func NewCannonicalTransactor(address common.Address, transactor bind.ContractTransactor) (*CannonicalTransactor, error) {
	contract, err := bindCannonical(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CannonicalTransactor{contract: contract}, nil
}

// NewCannonicalFilterer creates a new log filterer instance of Cannonical, bound to a specific deployed contract.
func NewCannonicalFilterer(address common.Address, filterer bind.ContractFilterer) (*CannonicalFilterer, error) {
	contract, err := bindCannonical(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CannonicalFilterer{contract: contract}, nil
}

// bindCannonical binds a generic wrapper to an already deployed contract.
func bindCannonical(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CannonicalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cannonical *CannonicalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cannonical.Contract.CannonicalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cannonical *CannonicalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cannonical.Contract.CannonicalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cannonical *CannonicalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cannonical.Contract.CannonicalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cannonical *CannonicalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cannonical.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cannonical *CannonicalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cannonical.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cannonical *CannonicalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cannonical.Contract.contract.Transact(opts, method, params...)
}

// MAXROLLUPTXSIZE is a free data retrieval call binding the contract method 0x876ed5cb.
//
// Solidity: function MAX_ROLLUP_TX_SIZE() view returns(uint256)
func (_Cannonical *CannonicalCaller) MAXROLLUPTXSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "MAX_ROLLUP_TX_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXROLLUPTXSIZE is a free data retrieval call binding the contract method 0x876ed5cb.
//
// Solidity: function MAX_ROLLUP_TX_SIZE() view returns(uint256)
func (_Cannonical *CannonicalSession) MAXROLLUPTXSIZE() (*big.Int, error) {
	return _Cannonical.Contract.MAXROLLUPTXSIZE(&_Cannonical.CallOpts)
}

// MAXROLLUPTXSIZE is a free data retrieval call binding the contract method 0x876ed5cb.
//
// Solidity: function MAX_ROLLUP_TX_SIZE() view returns(uint256)
func (_Cannonical *CannonicalCallerSession) MAXROLLUPTXSIZE() (*big.Int, error) {
	return _Cannonical.Contract.MAXROLLUPTXSIZE(&_Cannonical.CallOpts)
}

// MINROLLUPTXGAS is a free data retrieval call binding the contract method 0x78f4b2f2.
//
// Solidity: function MIN_ROLLUP_TX_GAS() view returns(uint256)
func (_Cannonical *CannonicalCaller) MINROLLUPTXGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "MIN_ROLLUP_TX_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINROLLUPTXGAS is a free data retrieval call binding the contract method 0x78f4b2f2.
//
// Solidity: function MIN_ROLLUP_TX_GAS() view returns(uint256)
func (_Cannonical *CannonicalSession) MINROLLUPTXGAS() (*big.Int, error) {
	return _Cannonical.Contract.MINROLLUPTXGAS(&_Cannonical.CallOpts)
}

// MINROLLUPTXGAS is a free data retrieval call binding the contract method 0x78f4b2f2.
//
// Solidity: function MIN_ROLLUP_TX_GAS() view returns(uint256)
func (_Cannonical *CannonicalCallerSession) MINROLLUPTXGAS() (*big.Int, error) {
	return _Cannonical.Contract.MINROLLUPTXGAS(&_Cannonical.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_Cannonical *CannonicalCaller) Batches(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "batches")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_Cannonical *CannonicalSession) Batches() (common.Address, error) {
	return _Cannonical.Contract.Batches(&_Cannonical.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xcfdf677e.
//
// Solidity: function batches() view returns(address)
func (_Cannonical *CannonicalCallerSession) Batches() (common.Address, error) {
	return _Cannonical.Contract.Batches(&_Cannonical.CallOpts)
}

// EnqueueGasCost is a free data retrieval call binding the contract method 0xe654b1fb.
//
// Solidity: function enqueueGasCost() view returns(uint256)
func (_Cannonical *CannonicalCaller) EnqueueGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "enqueueGasCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnqueueGasCost is a free data retrieval call binding the contract method 0xe654b1fb.
//
// Solidity: function enqueueGasCost() view returns(uint256)
func (_Cannonical *CannonicalSession) EnqueueGasCost() (*big.Int, error) {
	return _Cannonical.Contract.EnqueueGasCost(&_Cannonical.CallOpts)
}

// EnqueueGasCost is a free data retrieval call binding the contract method 0xe654b1fb.
//
// Solidity: function enqueueGasCost() view returns(uint256)
func (_Cannonical *CannonicalCallerSession) EnqueueGasCost() (*big.Int, error) {
	return _Cannonical.Contract.EnqueueGasCost(&_Cannonical.CallOpts)
}

// EnqueueL2GasPrepaid is a free data retrieval call binding the contract method 0x0b3dfa97.
//
// Solidity: function enqueueL2GasPrepaid() view returns(uint256)
func (_Cannonical *CannonicalCaller) EnqueueL2GasPrepaid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "enqueueL2GasPrepaid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnqueueL2GasPrepaid is a free data retrieval call binding the contract method 0x0b3dfa97.
//
// Solidity: function enqueueL2GasPrepaid() view returns(uint256)
func (_Cannonical *CannonicalSession) EnqueueL2GasPrepaid() (*big.Int, error) {
	return _Cannonical.Contract.EnqueueL2GasPrepaid(&_Cannonical.CallOpts)
}

// EnqueueL2GasPrepaid is a free data retrieval call binding the contract method 0x0b3dfa97.
//
// Solidity: function enqueueL2GasPrepaid() view returns(uint256)
func (_Cannonical *CannonicalCallerSession) EnqueueL2GasPrepaid() (*big.Int, error) {
	return _Cannonical.Contract.EnqueueL2GasPrepaid(&_Cannonical.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_Cannonical *CannonicalCaller) GetLastBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getLastBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_Cannonical *CannonicalSession) GetLastBlockNumber() (*big.Int, error) {
	return _Cannonical.Contract.GetLastBlockNumber(&_Cannonical.CallOpts)
}

// GetLastBlockNumber is a free data retrieval call binding the contract method 0x5ae6256d.
//
// Solidity: function getLastBlockNumber() view returns(uint40)
func (_Cannonical *CannonicalCallerSession) GetLastBlockNumber() (*big.Int, error) {
	return _Cannonical.Contract.GetLastBlockNumber(&_Cannonical.CallOpts)
}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_Cannonical *CannonicalCaller) GetLastTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getLastTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_Cannonical *CannonicalSession) GetLastTimestamp() (*big.Int, error) {
	return _Cannonical.Contract.GetLastTimestamp(&_Cannonical.CallOpts)
}

// GetLastTimestamp is a free data retrieval call binding the contract method 0x37899770.
//
// Solidity: function getLastTimestamp() view returns(uint40)
func (_Cannonical *CannonicalCallerSession) GetLastTimestamp() (*big.Int, error) {
	return _Cannonical.Contract.GetLastTimestamp(&_Cannonical.CallOpts)
}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_Cannonical *CannonicalCaller) GetNextQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getNextQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_Cannonical *CannonicalSession) GetNextQueueIndex() (*big.Int, error) {
	return _Cannonical.Contract.GetNextQueueIndex(&_Cannonical.CallOpts)
}

// GetNextQueueIndex is a free data retrieval call binding the contract method 0x7a167a8a.
//
// Solidity: function getNextQueueIndex() view returns(uint40)
func (_Cannonical *CannonicalCallerSession) GetNextQueueIndex() (*big.Int, error) {
	return _Cannonical.Contract.GetNextQueueIndex(&_Cannonical.CallOpts)
}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_Cannonical *CannonicalCaller) GetNumPendingQueueElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getNumPendingQueueElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_Cannonical *CannonicalSession) GetNumPendingQueueElements() (*big.Int, error) {
	return _Cannonical.Contract.GetNumPendingQueueElements(&_Cannonical.CallOpts)
}

// GetNumPendingQueueElements is a free data retrieval call binding the contract method 0xf722b41a.
//
// Solidity: function getNumPendingQueueElements() view returns(uint40)
func (_Cannonical *CannonicalCallerSession) GetNumPendingQueueElements() (*big.Int, error) {
	return _Cannonical.Contract.GetNumPendingQueueElements(&_Cannonical.CallOpts)
}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_Cannonical *CannonicalCaller) GetQueueElement(opts *bind.CallOpts, _index *big.Int) (LibOVMCodecQueueElement, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getQueueElement", _index)

	if err != nil {
		return *new(LibOVMCodecQueueElement), err
	}

	out0 := *abi.ConvertType(out[0], new(LibOVMCodecQueueElement)).(*LibOVMCodecQueueElement)

	return out0, err

}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_Cannonical *CannonicalSession) GetQueueElement(_index *big.Int) (LibOVMCodecQueueElement, error) {
	return _Cannonical.Contract.GetQueueElement(&_Cannonical.CallOpts, _index)
}

// GetQueueElement is a free data retrieval call binding the contract method 0x2a7f18be.
//
// Solidity: function getQueueElement(uint256 _index) view returns((bytes32,uint40,uint40) _element)
func (_Cannonical *CannonicalCallerSession) GetQueueElement(_index *big.Int) (LibOVMCodecQueueElement, error) {
	return _Cannonical.Contract.GetQueueElement(&_Cannonical.CallOpts, _index)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_Cannonical *CannonicalCaller) GetQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_Cannonical *CannonicalSession) GetQueueLength() (*big.Int, error) {
	return _Cannonical.Contract.GetQueueLength(&_Cannonical.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint40)
func (_Cannonical *CannonicalCallerSession) GetQueueLength() (*big.Int, error) {
	return _Cannonical.Contract.GetQueueLength(&_Cannonical.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_Cannonical *CannonicalCaller) GetTotalBatches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getTotalBatches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_Cannonical *CannonicalSession) GetTotalBatches() (*big.Int, error) {
	return _Cannonical.Contract.GetTotalBatches(&_Cannonical.CallOpts)
}

// GetTotalBatches is a free data retrieval call binding the contract method 0xe561dddc.
//
// Solidity: function getTotalBatches() view returns(uint256 _totalBatches)
func (_Cannonical *CannonicalCallerSession) GetTotalBatches() (*big.Int, error) {
	return _Cannonical.Contract.GetTotalBatches(&_Cannonical.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_Cannonical *CannonicalCaller) GetTotalElements(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "getTotalElements")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_Cannonical *CannonicalSession) GetTotalElements() (*big.Int, error) {
	return _Cannonical.Contract.GetTotalElements(&_Cannonical.CallOpts)
}

// GetTotalElements is a free data retrieval call binding the contract method 0x7aa63a86.
//
// Solidity: function getTotalElements() view returns(uint256 _totalElements)
func (_Cannonical *CannonicalCallerSession) GetTotalElements() (*big.Int, error) {
	return _Cannonical.Contract.GetTotalElements(&_Cannonical.CallOpts)
}

// L2GasDiscountDivisor is a free data retrieval call binding the contract method 0xccf987c8.
//
// Solidity: function l2GasDiscountDivisor() view returns(uint256)
func (_Cannonical *CannonicalCaller) L2GasDiscountDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "l2GasDiscountDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2GasDiscountDivisor is a free data retrieval call binding the contract method 0xccf987c8.
//
// Solidity: function l2GasDiscountDivisor() view returns(uint256)
func (_Cannonical *CannonicalSession) L2GasDiscountDivisor() (*big.Int, error) {
	return _Cannonical.Contract.L2GasDiscountDivisor(&_Cannonical.CallOpts)
}

// L2GasDiscountDivisor is a free data retrieval call binding the contract method 0xccf987c8.
//
// Solidity: function l2GasDiscountDivisor() view returns(uint256)
func (_Cannonical *CannonicalCallerSession) L2GasDiscountDivisor() (*big.Int, error) {
	return _Cannonical.Contract.L2GasDiscountDivisor(&_Cannonical.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_Cannonical *CannonicalCaller) LibAddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "libAddressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_Cannonical *CannonicalSession) LibAddressManager() (common.Address, error) {
	return _Cannonical.Contract.LibAddressManager(&_Cannonical.CallOpts)
}

// LibAddressManager is a free data retrieval call binding the contract method 0x299ca478.
//
// Solidity: function libAddressManager() view returns(address)
func (_Cannonical *CannonicalCallerSession) LibAddressManager() (common.Address, error) {
	return _Cannonical.Contract.LibAddressManager(&_Cannonical.CallOpts)
}

// MaxTransactionGasLimit is a free data retrieval call binding the contract method 0x8d38c6c1.
//
// Solidity: function maxTransactionGasLimit() view returns(uint256)
func (_Cannonical *CannonicalCaller) MaxTransactionGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "maxTransactionGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTransactionGasLimit is a free data retrieval call binding the contract method 0x8d38c6c1.
//
// Solidity: function maxTransactionGasLimit() view returns(uint256)
func (_Cannonical *CannonicalSession) MaxTransactionGasLimit() (*big.Int, error) {
	return _Cannonical.Contract.MaxTransactionGasLimit(&_Cannonical.CallOpts)
}

// MaxTransactionGasLimit is a free data retrieval call binding the contract method 0x8d38c6c1.
//
// Solidity: function maxTransactionGasLimit() view returns(uint256)
func (_Cannonical *CannonicalCallerSession) MaxTransactionGasLimit() (*big.Int, error) {
	return _Cannonical.Contract.MaxTransactionGasLimit(&_Cannonical.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_Cannonical *CannonicalCaller) Resolve(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _Cannonical.contract.Call(opts, &out, "resolve", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_Cannonical *CannonicalSession) Resolve(_name string) (common.Address, error) {
	return _Cannonical.Contract.Resolve(&_Cannonical.CallOpts, _name)
}

// Resolve is a free data retrieval call binding the contract method 0x461a4478.
//
// Solidity: function resolve(string _name) view returns(address)
func (_Cannonical *CannonicalCallerSession) Resolve(_name string) (common.Address, error) {
	return _Cannonical.Contract.Resolve(&_Cannonical.CallOpts, _name)
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_Cannonical *CannonicalTransactor) AppendSequencerBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cannonical.contract.Transact(opts, "appendSequencerBatch")
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_Cannonical *CannonicalSession) AppendSequencerBatch() (*types.Transaction, error) {
	return _Cannonical.Contract.AppendSequencerBatch(&_Cannonical.TransactOpts)
}

// AppendSequencerBatch is a paid mutator transaction binding the contract method 0xd0f89344.
//
// Solidity: function appendSequencerBatch() returns()
func (_Cannonical *CannonicalTransactorSession) AppendSequencerBatch() (*types.Transaction, error) {
	return _Cannonical.Contract.AppendSequencerBatch(&_Cannonical.TransactOpts)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_Cannonical *CannonicalTransactor) Enqueue(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cannonical.contract.Transact(opts, "enqueue", _target, _gasLimit, _data)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_Cannonical *CannonicalSession) Enqueue(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cannonical.Contract.Enqueue(&_Cannonical.TransactOpts, _target, _gasLimit, _data)
}

// Enqueue is a paid mutator transaction binding the contract method 0x6fee07e0.
//
// Solidity: function enqueue(address _target, uint256 _gasLimit, bytes _data) returns()
func (_Cannonical *CannonicalTransactorSession) Enqueue(_target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _Cannonical.Contract.Enqueue(&_Cannonical.TransactOpts, _target, _gasLimit, _data)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_Cannonical *CannonicalTransactor) SetGasParams(opts *bind.TransactOpts, _l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _Cannonical.contract.Transact(opts, "setGasParams", _l2GasDiscountDivisor, _enqueueGasCost)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_Cannonical *CannonicalSession) SetGasParams(_l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _Cannonical.Contract.SetGasParams(&_Cannonical.TransactOpts, _l2GasDiscountDivisor, _enqueueGasCost)
}

// SetGasParams is a paid mutator transaction binding the contract method 0xedcc4a45.
//
// Solidity: function setGasParams(uint256 _l2GasDiscountDivisor, uint256 _enqueueGasCost) returns()
func (_Cannonical *CannonicalTransactorSession) SetGasParams(_l2GasDiscountDivisor *big.Int, _enqueueGasCost *big.Int) (*types.Transaction, error) {
	return _Cannonical.Contract.SetGasParams(&_Cannonical.TransactOpts, _l2GasDiscountDivisor, _enqueueGasCost)
}

// CannonicalL2GasParamsUpdatedIterator is returned from FilterL2GasParamsUpdated and is used to iterate over the raw logs and unpacked data for L2GasParamsUpdated events raised by the Cannonical contract.
type CannonicalL2GasParamsUpdatedIterator struct {
	Event *CannonicalL2GasParamsUpdated // Event containing the contract specifics and raw log

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
func (it *CannonicalL2GasParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CannonicalL2GasParamsUpdated)
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
		it.Event = new(CannonicalL2GasParamsUpdated)
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
func (it *CannonicalL2GasParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CannonicalL2GasParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CannonicalL2GasParamsUpdated represents a L2GasParamsUpdated event raised by the Cannonical contract.
type CannonicalL2GasParamsUpdated struct {
	L2GasDiscountDivisor *big.Int
	EnqueueGasCost       *big.Int
	EnqueueL2GasPrepaid  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterL2GasParamsUpdated is a free log retrieval operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_Cannonical *CannonicalFilterer) FilterL2GasParamsUpdated(opts *bind.FilterOpts) (*CannonicalL2GasParamsUpdatedIterator, error) {

	logs, sub, err := _Cannonical.contract.FilterLogs(opts, "L2GasParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &CannonicalL2GasParamsUpdatedIterator{contract: _Cannonical.contract, event: "L2GasParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchL2GasParamsUpdated is a free log subscription operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_Cannonical *CannonicalFilterer) WatchL2GasParamsUpdated(opts *bind.WatchOpts, sink chan<- *CannonicalL2GasParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _Cannonical.contract.WatchLogs(opts, "L2GasParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CannonicalL2GasParamsUpdated)
				if err := _Cannonical.contract.UnpackLog(event, "L2GasParamsUpdated", log); err != nil {
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

// ParseL2GasParamsUpdated is a log parse operation binding the contract event 0xc6ed75e96b8b18b71edc1a6e82a9d677f8268c774a262c624eeb2cf0a8b3e07e.
//
// Solidity: event L2GasParamsUpdated(uint256 l2GasDiscountDivisor, uint256 enqueueGasCost, uint256 enqueueL2GasPrepaid)
func (_Cannonical *CannonicalFilterer) ParseL2GasParamsUpdated(log types.Log) (*CannonicalL2GasParamsUpdated, error) {
	event := new(CannonicalL2GasParamsUpdated)
	if err := _Cannonical.contract.UnpackLog(event, "L2GasParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CannonicalQueueBatchAppendedIterator is returned from FilterQueueBatchAppended and is used to iterate over the raw logs and unpacked data for QueueBatchAppended events raised by the Cannonical contract.
type CannonicalQueueBatchAppendedIterator struct {
	Event *CannonicalQueueBatchAppended // Event containing the contract specifics and raw log

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
func (it *CannonicalQueueBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CannonicalQueueBatchAppended)
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
		it.Event = new(CannonicalQueueBatchAppended)
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
func (it *CannonicalQueueBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CannonicalQueueBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CannonicalQueueBatchAppended represents a QueueBatchAppended event raised by the Cannonical contract.
type CannonicalQueueBatchAppended struct {
	StartingQueueIndex *big.Int
	NumQueueElements   *big.Int
	TotalElements      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQueueBatchAppended is a free log retrieval operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_Cannonical *CannonicalFilterer) FilterQueueBatchAppended(opts *bind.FilterOpts) (*CannonicalQueueBatchAppendedIterator, error) {

	logs, sub, err := _Cannonical.contract.FilterLogs(opts, "QueueBatchAppended")
	if err != nil {
		return nil, err
	}
	return &CannonicalQueueBatchAppendedIterator{contract: _Cannonical.contract, event: "QueueBatchAppended", logs: logs, sub: sub}, nil
}

// WatchQueueBatchAppended is a free log subscription operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_Cannonical *CannonicalFilterer) WatchQueueBatchAppended(opts *bind.WatchOpts, sink chan<- *CannonicalQueueBatchAppended) (event.Subscription, error) {

	logs, sub, err := _Cannonical.contract.WatchLogs(opts, "QueueBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CannonicalQueueBatchAppended)
				if err := _Cannonical.contract.UnpackLog(event, "QueueBatchAppended", log); err != nil {
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

// ParseQueueBatchAppended is a log parse operation binding the contract event 0x64d7f508348c70dea42d5302a393987e4abc20e45954ab3f9d320207751956f0.
//
// Solidity: event QueueBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_Cannonical *CannonicalFilterer) ParseQueueBatchAppended(log types.Log) (*CannonicalQueueBatchAppended, error) {
	event := new(CannonicalQueueBatchAppended)
	if err := _Cannonical.contract.UnpackLog(event, "QueueBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CannonicalSequencerBatchAppendedIterator is returned from FilterSequencerBatchAppended and is used to iterate over the raw logs and unpacked data for SequencerBatchAppended events raised by the Cannonical contract.
type CannonicalSequencerBatchAppendedIterator struct {
	Event *CannonicalSequencerBatchAppended // Event containing the contract specifics and raw log

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
func (it *CannonicalSequencerBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CannonicalSequencerBatchAppended)
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
		it.Event = new(CannonicalSequencerBatchAppended)
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
func (it *CannonicalSequencerBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CannonicalSequencerBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CannonicalSequencerBatchAppended represents a SequencerBatchAppended event raised by the Cannonical contract.
type CannonicalSequencerBatchAppended struct {
	StartingQueueIndex *big.Int
	NumQueueElements   *big.Int
	TotalElements      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchAppended is a free log retrieval operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_Cannonical *CannonicalFilterer) FilterSequencerBatchAppended(opts *bind.FilterOpts) (*CannonicalSequencerBatchAppendedIterator, error) {

	logs, sub, err := _Cannonical.contract.FilterLogs(opts, "SequencerBatchAppended")
	if err != nil {
		return nil, err
	}
	return &CannonicalSequencerBatchAppendedIterator{contract: _Cannonical.contract, event: "SequencerBatchAppended", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchAppended is a free log subscription operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_Cannonical *CannonicalFilterer) WatchSequencerBatchAppended(opts *bind.WatchOpts, sink chan<- *CannonicalSequencerBatchAppended) (event.Subscription, error) {

	logs, sub, err := _Cannonical.contract.WatchLogs(opts, "SequencerBatchAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CannonicalSequencerBatchAppended)
				if err := _Cannonical.contract.UnpackLog(event, "SequencerBatchAppended", log); err != nil {
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

// ParseSequencerBatchAppended is a log parse operation binding the contract event 0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899.
//
// Solidity: event SequencerBatchAppended(uint256 _startingQueueIndex, uint256 _numQueueElements, uint256 _totalElements)
func (_Cannonical *CannonicalFilterer) ParseSequencerBatchAppended(log types.Log) (*CannonicalSequencerBatchAppended, error) {
	event := new(CannonicalSequencerBatchAppended)
	if err := _Cannonical.contract.UnpackLog(event, "SequencerBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CannonicalTransactionBatchAppendedIterator is returned from FilterTransactionBatchAppended and is used to iterate over the raw logs and unpacked data for TransactionBatchAppended events raised by the Cannonical contract.
type CannonicalTransactionBatchAppendedIterator struct {
	Event *CannonicalTransactionBatchAppended // Event containing the contract specifics and raw log

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
func (it *CannonicalTransactionBatchAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CannonicalTransactionBatchAppended)
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
		it.Event = new(CannonicalTransactionBatchAppended)
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
func (it *CannonicalTransactionBatchAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CannonicalTransactionBatchAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CannonicalTransactionBatchAppended represents a TransactionBatchAppended event raised by the Cannonical contract.
type CannonicalTransactionBatchAppended struct {
	BatchIndex        *big.Int
	BatchRoot         [32]byte
	BatchSize         *big.Int
	PrevTotalElements *big.Int
	ExtraData         []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchAppended is a free log retrieval operation binding the contract event 0x127186556e7be68c7e31263195225b4de02820707889540969f62c05cf73525e.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _extraData)
func (_Cannonical *CannonicalFilterer) FilterTransactionBatchAppended(opts *bind.FilterOpts, _batchIndex []*big.Int) (*CannonicalTransactionBatchAppendedIterator, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Cannonical.contract.FilterLogs(opts, "TransactionBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &CannonicalTransactionBatchAppendedIterator{contract: _Cannonical.contract, event: "TransactionBatchAppended", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchAppended is a free log subscription operation binding the contract event 0x127186556e7be68c7e31263195225b4de02820707889540969f62c05cf73525e.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _extraData)
func (_Cannonical *CannonicalFilterer) WatchTransactionBatchAppended(opts *bind.WatchOpts, sink chan<- *CannonicalTransactionBatchAppended, _batchIndex []*big.Int) (event.Subscription, error) {

	var _batchIndexRule []interface{}
	for _, _batchIndexItem := range _batchIndex {
		_batchIndexRule = append(_batchIndexRule, _batchIndexItem)
	}

	logs, sub, err := _Cannonical.contract.WatchLogs(opts, "TransactionBatchAppended", _batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CannonicalTransactionBatchAppended)
				if err := _Cannonical.contract.UnpackLog(event, "TransactionBatchAppended", log); err != nil {
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

// ParseTransactionBatchAppended is a log parse operation binding the contract event 0x127186556e7be68c7e31263195225b4de02820707889540969f62c05cf73525e.
//
// Solidity: event TransactionBatchAppended(uint256 indexed _batchIndex, bytes32 _batchRoot, uint256 _batchSize, uint256 _prevTotalElements, bytes _extraData)
func (_Cannonical *CannonicalFilterer) ParseTransactionBatchAppended(log types.Log) (*CannonicalTransactionBatchAppended, error) {
	event := new(CannonicalTransactionBatchAppended)
	if err := _Cannonical.contract.UnpackLog(event, "TransactionBatchAppended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CannonicalTransactionEnqueuedIterator is returned from FilterTransactionEnqueued and is used to iterate over the raw logs and unpacked data for TransactionEnqueued events raised by the Cannonical contract.
type CannonicalTransactionEnqueuedIterator struct {
	Event *CannonicalTransactionEnqueued // Event containing the contract specifics and raw log

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
func (it *CannonicalTransactionEnqueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CannonicalTransactionEnqueued)
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
		it.Event = new(CannonicalTransactionEnqueued)
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
func (it *CannonicalTransactionEnqueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CannonicalTransactionEnqueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CannonicalTransactionEnqueued represents a TransactionEnqueued event raised by the Cannonical contract.
type CannonicalTransactionEnqueued struct {
	L1TxOrigin common.Address
	Target     common.Address
	GasLimit   *big.Int
	Data       []byte
	QueueIndex *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionEnqueued is a free log retrieval operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_Cannonical *CannonicalFilterer) FilterTransactionEnqueued(opts *bind.FilterOpts, _l1TxOrigin []common.Address, _target []common.Address, _queueIndex []*big.Int) (*CannonicalTransactionEnqueuedIterator, error) {

	var _l1TxOriginRule []interface{}
	for _, _l1TxOriginItem := range _l1TxOrigin {
		_l1TxOriginRule = append(_l1TxOriginRule, _l1TxOriginItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	var _queueIndexRule []interface{}
	for _, _queueIndexItem := range _queueIndex {
		_queueIndexRule = append(_queueIndexRule, _queueIndexItem)
	}

	logs, sub, err := _Cannonical.contract.FilterLogs(opts, "TransactionEnqueued", _l1TxOriginRule, _targetRule, _queueIndexRule)
	if err != nil {
		return nil, err
	}
	return &CannonicalTransactionEnqueuedIterator{contract: _Cannonical.contract, event: "TransactionEnqueued", logs: logs, sub: sub}, nil
}

// WatchTransactionEnqueued is a free log subscription operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_Cannonical *CannonicalFilterer) WatchTransactionEnqueued(opts *bind.WatchOpts, sink chan<- *CannonicalTransactionEnqueued, _l1TxOrigin []common.Address, _target []common.Address, _queueIndex []*big.Int) (event.Subscription, error) {

	var _l1TxOriginRule []interface{}
	for _, _l1TxOriginItem := range _l1TxOrigin {
		_l1TxOriginRule = append(_l1TxOriginRule, _l1TxOriginItem)
	}
	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}

	var _queueIndexRule []interface{}
	for _, _queueIndexItem := range _queueIndex {
		_queueIndexRule = append(_queueIndexRule, _queueIndexItem)
	}

	logs, sub, err := _Cannonical.contract.WatchLogs(opts, "TransactionEnqueued", _l1TxOriginRule, _targetRule, _queueIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CannonicalTransactionEnqueued)
				if err := _Cannonical.contract.UnpackLog(event, "TransactionEnqueued", log); err != nil {
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

// ParseTransactionEnqueued is a log parse operation binding the contract event 0x4b388aecf9fa6cc92253704e5975a6129a4f735bdbd99567df4ed0094ee4ceb5.
//
// Solidity: event TransactionEnqueued(address indexed _l1TxOrigin, address indexed _target, uint256 _gasLimit, bytes _data, uint256 indexed _queueIndex, uint256 _timestamp)
func (_Cannonical *CannonicalFilterer) ParseTransactionEnqueued(log types.Log) (*CannonicalTransactionEnqueued, error) {
	event := new(CannonicalTransactionEnqueued)
	if err := _Cannonical.contract.UnpackLog(event, "TransactionEnqueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

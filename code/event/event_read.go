package main

import (
	cannonical "Parse_eventlog/code/code/contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Event struct {
	StartingQueueIndex uint
	NumQueueElements   uint
	TotalElements      uint
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a39c40e31fb64ef8bdf7b0f6feccb277")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x4A1941f18874Df01e5CAA1CD3DA4b1803CBD32C2")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(18609761),
		ToBlock:   big.NewInt(18654431),
		Addresses: []common.Address{
			contractAddress,
		},
		Topics: [][]common.Hash{
			{common.HexToHash("0x602f1aeac0ca2e7a13e281a9ef0ad7838542712ce16780fa2ecffd351f05f899")},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(cannonical.CannonicalMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("BlockHash", vLog.BlockHash.Hex())
		fmt.Println("BlockNumber", vLog.BlockNumber)
		fmt.Println("TxHASH", vLog.TxHash.Hex())
		fmt.Printf("TxData %+v", vLog.Data)
		common.Bytes2Hex(vLog.Data)
		fmt.Printf("TxData %+v", common.Bytes2Hex(vLog.Data))

		event, err := contractAbi.Unpack("SequencerBatchAppended", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Event Data : ", event, "\n")

		//fmt.Println("Key", string(event.StartingQueueIndex)) // foo
		//fmt.Println("value", string(event.NumQueueElements)) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("SequencerBatchAppended(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4

}

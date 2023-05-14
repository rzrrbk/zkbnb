package desertexit

import (
	"context"
	"encoding/json"
	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
	"github.com/bnb-chain/zkbnb/common/abicoder"
	monitor2 "github.com/bnb-chain/zkbnb/common/monitor"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeromicro/go-zero/core/logx"
	"math/big"
	"testing"
)

//const networkRpc = "https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd"
const networkRpc = "https://data-seed-prebsc-1-s3.binance.org:8545"

const commitHash = "0x68c377009c46578ae1377d0da601813615a4e21b6d1cf52f0d39fd2d1c9b9433"

const verifyHash = "0x8b90ebe41dd69ea95135b7f7b5984326a2c419f436635e7ff5aaacf0f17cc8c7"

const revertHash = "0x8b90ebe41dd69ea95135b7f7b5984326a2c419f436635e7ff5aaacf0f17cc8c7"

func TestDesertExit_getCommitBlocksCallData(t *testing.T) {
	client, err := rpc.NewClient(networkRpc)
	if err != nil {
		logx.Severef("failed to create rpc client, %v", err)
		return
	}
	newABIDecoder := abicoder.NewABIDecoder(monitor2.ZkBNBContractAbi)
	transaction, _, err := client.TransactionByHash(context.Background(), common.HexToHash(commitHash))
	if err != nil {
		logx.Severe(err)
		return
	}

	receipt, err := client.GetTransactionReceipt(commitHash)
	if err != nil {
		logx.Errorf("query transaction receipt %s failed, err: %v", commitHash, err)
	} else {
		json, _ := receipt.MarshalJSON()
		logx.Infof(string(json))
	}

	storageStoredBlockInfo := StorageStoredBlockInfo{}
	newBlocksData := make([]ZkBNBCommitBlockInfo, 0)
	callData := CommitBlocksCallData{LastCommittedBlockData: &storageStoredBlockInfo, NewBlocksData: newBlocksData}
	if err := newABIDecoder.UnpackIntoInterface(&callData, "commitBlocks", transaction.Data()[4:]); err != nil {
		logx.Severe(err)
		return
	}
	jsonBytes, err := json.Marshal(callData)
	logx.Infof("callData=%s", string(jsonBytes))

	storageStoredBlockInfoDTO := StorageStoredBlockInfoDTO{
		BlockNumber:                  callData.LastCommittedBlockData.BlockNumber,
		PriorityOperations:           callData.LastCommittedBlockData.PriorityOperations,
		PendingOnchainOperationsHash: common.Bytes2Hex(callData.LastCommittedBlockData.PendingOnchainOperationsHash[:]),
		Timestamp:                    callData.LastCommittedBlockData.Timestamp,
		StateRoot:                    common.Bytes2Hex(callData.LastCommittedBlockData.StateRoot[:]),
		Commitment:                   common.Bytes2Hex(callData.LastCommittedBlockData.Commitment[:]),
		BlockSize:                    callData.LastCommittedBlockData.BlockSize,
	}
	jsonBytes, _ = json.Marshal(storageStoredBlockInfoDTO)
	logx.Infof("StorageStoredBlockInfo:%s", jsonBytes)

}

func TestDesertExit_getVerifyAndExecuteBlocksCallData(t *testing.T) {
	client, err := rpc.NewClient(networkRpc)
	if err != nil {
		logx.Severef("failed to create rpc client, %v", err)
		return
	}
	newABIDecoder := abicoder.NewABIDecoder(monitor2.ZkBNBContractAbi)
	transaction, _, err := client.TransactionByHash(context.Background(), common.HexToHash(verifyHash))
	if err != nil {
		logx.Severe(err)
		return
	}

	receipt, err := client.GetTransactionReceipt(verifyHash)
	if err != nil {
		logx.Errorf("query transaction receipt %s failed, err: %v", verifyHash, err)
	} else {
		json, _ := receipt.MarshalJSON()
		logx.Infof(string(json))
	}

	newBlocksData := make([]ZkBNBVerifyAndExecuteBlockInfo, 0)
	proofs := make([]*big.Int, 0)
	callData := VerifyAndExecuteBlocksCallData{Proofs: proofs, VerifyAndExecuteBlocksInfo: newBlocksData}
	if err := newABIDecoder.UnpackIntoInterface(&callData, "verifyAndExecuteBlocks", transaction.Data()[4:]); err != nil {
		logx.Severe(err)
		return
	}
	jsonBytes, err := json.Marshal(callData)
	logx.Infof("callData=%s", string(jsonBytes))

	for _, verifyAndExecuteBlockInfo := range callData.VerifyAndExecuteBlocksInfo {
		storageStoredBlockInfo := StorageStoredBlockInfoDTO{
			BlockNumber:                  verifyAndExecuteBlockInfo.BlockHeader.BlockNumber,
			PriorityOperations:           verifyAndExecuteBlockInfo.BlockHeader.PriorityOperations,
			PendingOnchainOperationsHash: common.Bytes2Hex(verifyAndExecuteBlockInfo.BlockHeader.PendingOnchainOperationsHash[:]),
			Timestamp:                    verifyAndExecuteBlockInfo.BlockHeader.Timestamp,
			StateRoot:                    common.Bytes2Hex(verifyAndExecuteBlockInfo.BlockHeader.StateRoot[:]),
			Commitment:                   common.Bytes2Hex(verifyAndExecuteBlockInfo.BlockHeader.Commitment[:]),
			BlockSize:                    verifyAndExecuteBlockInfo.BlockHeader.BlockSize,
		}
		jsonBytes, _ := json.Marshal(storageStoredBlockInfo)
		logx.Infof("verifyAndExecuteBlockInfo.BlockHeader:%s", jsonBytes)

		jsonBytes, err = json.Marshal(verifyAndExecuteBlockInfo.PendingOnchainOpsPubData)
		logx.Infof("PendingOnchainOpsPubData:%s", jsonBytes)
	}

}

func TestDesertExit_getRevertBlocksCallData(t *testing.T) {
	client, err := rpc.NewClient(networkRpc)
	if err != nil {
		logx.Severef("failed to create rpc client, %v", err)
		return
	}
	newABIDecoder := abicoder.NewABIDecoder(monitor2.ZkBNBContractAbi)
	transaction, _, err := client.TransactionByHash(context.Background(), common.HexToHash(revertHash))
	if err != nil {
		logx.Severe(err)
		return
	}

	receipt, err := client.GetTransactionReceipt(revertHash)
	if err != nil {
		logx.Errorf("query transaction receipt %s failed, err: %v", revertHash, err)
	} else {
		json, _ := receipt.MarshalJSON()
		logx.Infof(string(json))
	}

	blocksToRevertData := make([]StorageStoredBlockInfo, 0)
	callData := RevertBlocksCallData{BlocksToRevert: blocksToRevertData}
	if err := newABIDecoder.UnpackIntoInterface(&callData, "revertBlocks", transaction.Data()[4:]); err != nil {
		logx.Severe(err)
		return
	}
	jsonBytes, err := json.Marshal(callData)
	logx.Infof("callData=%s", string(jsonBytes))
}

type RevertBlocksCallData struct {
	BlocksToRevert []StorageStoredBlockInfo `abi:"_blocksToRevert"`
}

type StorageStoredBlockInfoDTO struct {
	BlockSize                    uint16
	BlockNumber                  uint32
	PriorityOperations           uint64
	PendingOnchainOperationsHash string
	Timestamp                    *big.Int
	StateRoot                    string
	Commitment                   string
}

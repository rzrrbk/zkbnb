package signature

import (
	"fmt"
	"github.com/bnb-chain/zkbnb/service/apiserver/internal/logic/utils"
	"github.com/bnb-chain/zkbnb/types"
	"github.com/zeromicro/go-zero/core/logx"
)

const (

	// SignatureTemplateWithdrawal Withdrawal ${amount} to: ${to.toLowerCase()}\nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateWithdrawal = "Withdrawal %s to: %s\nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateTransfer /* Transfer ${amount} ${tokenAddress} to: ${to.toLowerCase()}\nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce} */
	SignatureTemplateTransfer = "Transfer %s %d to: %d\nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateCreateCollection CreateCollection ${accountIndex} ${collectionName} \nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateCreateCollection = "CreateCollection %d %s \nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateMintNft MintNFT ${contentHash} for: ${recipient.toLowerCase()}\nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateMintNft = "MintNFT %s for: %d\nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateTransferNft TransferNFT ${NftIndex} ${fromAccountIndex} to ${toAccountIndex} \nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateTransferNft = "TransferNFT %d %d to %d \nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateWithdrawalNft Withdrawal ${tokenIndex} to: ${to.toLowerCase()}\nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateWithdrawalNft = "Withdrawal %d to: %s\nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateCancelOffer CancelOffer ${offerId} by: ${accountIndex} \nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateCancelOffer = "CancelOffer %d by: %d \nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"
	// SignatureTemplateAtomicMatch AtomicMatch ${amount} ${offerId} ${nftIndex} ${accountIndex} \nGasFee: ${fee} \nGasAccountIndex: ${feeTokenAddress}\nNonce: ${nonce}
	SignatureTemplateAtomicMatch = "AtomicMatch %s %d %d %d \nGasFee: %s \nGasAccountIndex: %d\nNonce: %d"

	// SignatureTemplateAccount AccountIndex:{AccountIndex}\nNftIndex:{NftIndex}\nNonce:{Nonce}
	SignatureTemplateAccount = "AccountIndex:%d\nNftIndex:%d\nNonce:%d"
)

var SignatureFunctionMap = make(map[uint32]func(txInfo string) (string, error), 0)

func GenerateSignatureBody(txType uint32, txInfo string) (string, error) {
	if len(SignatureFunctionMap) == 0 {
		ConstructSignatureFunction()
	}

	SignatureFunc := SignatureFunctionMap[txType]
	if SignatureFunc == nil {
		logx.Errorf("Can not find Signature Function for TxType:%d", txType)
		return "", types.AppErrNoSignFunctionForTxType
	}

	signatureBody, err := SignatureFunc(txInfo)
	if err != nil {
		return "", err
	}
	return signatureBody, nil
}

func ConstructSignatureFunction() {
	SignatureFunctionMap[types.TxTypeWithdraw] = SignatureForWithdrawal
	SignatureFunctionMap[types.TxTypeTransfer] = SignatureForTransfer
	SignatureFunctionMap[types.TxTypeCreateCollection] = SignatureForCreateCollection
	SignatureFunctionMap[types.TxTypeMintNft] = SignatureForMintNft
	SignatureFunctionMap[types.TxTypeTransferNft] = SignatureForTransferNft
	SignatureFunctionMap[types.TxTypeWithdrawNft] = SignatureForWithdrawalNft
	SignatureFunctionMap[types.TxTypeCancelOffer] = SignatureForCancelOffer
	SignatureFunctionMap[types.TxTypeAtomicMatch] = SignatureForAtomicMatch
	SignatureFunctionMap[types.TxTypeEmpty] = SignatureForAccount

}

func SignatureForWithdrawal(txInfo string) (string, error) {
	transaction, err := types.ParseWithdrawTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse withdrawal tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateWithdrawal, utils.FormatWeiToEtherStr(transaction.AssetAmount), transaction.ToAddress,
		utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForTransfer(txInfo string) (string, error) {
	transaction, err := types.ParseTransferTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse transfer tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateTransfer, utils.FormatWeiToEtherStr(transaction.AssetAmount), transaction.FromAccountIndex,
		transaction.ToAccountIndex, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForCreateCollection(txInfo string) (string, error) {
	transaction, err := types.ParseCreateCollectionTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse create collection tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateCreateCollection, transaction.AccountIndex,
		transaction.Name, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForMintNft(txInfo string) (string, error) {
	transaction, err := types.ParseMintNftTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse mint nft tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateMintNft, transaction.ToAccountNameHash,
		transaction.ToAccountIndex, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForTransferNft(txInfo string) (string, error) {
	transaction, err := types.ParseTransferNftTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse cancel offer tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateTransferNft, transaction.NftIndex, transaction.FromAccountIndex,
		transaction.ToAccountIndex, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForWithdrawalNft(txInfo string) (string, error) {
	transaction, err := types.ParseWithdrawNftTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse withdrawal nft tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateWithdrawalNft, transaction.NftIndex,
		transaction.ToAddress, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForCancelOffer(txInfo string) (string, error) {
	transaction, err := types.ParseCancelOfferTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse cancel offer tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	signatureBody := fmt.Sprintf(SignatureTemplateCancelOffer, transaction.OfferId,
		transaction.AccountIndex, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForAtomicMatch(txInfo string) (string, error) {
	transaction, err := types.ParseAtomicMatchTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse atomic match tx failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}

	offer := transaction.BuyOffer
	if offer == nil {
		offer = transaction.SellOffer
	}
	if offer == nil {
		return "", types.AppErrBothOfferNotExist
	}

	signatureBody := fmt.Sprintf(SignatureTemplateAtomicMatch, utils.FormatWeiToEtherStr(offer.AssetAmount), offer.OfferId, offer.NftIndex,
		transaction.AccountIndex, utils.FormatWeiToEtherStr(transaction.GasFeeAssetAmount), transaction.GasAccountIndex, transaction.Nonce)
	return signatureBody, nil
}

func SignatureForAccount(txInfo string) (string, error) {
	transaction, err := types.ParseUpdateNftTxInfo(txInfo)
	if err != nil {
		logx.Errorf("parse atomic match nft info failed: %s", err.Error())
		return "", types.AppErrInvalidTxInfo
	}
	signatureBody := fmt.Sprintf(SignatureTemplateAccount, transaction.AccountIndex, transaction.NftIndex, transaction.Nonce)
	return signatureBody, nil
}

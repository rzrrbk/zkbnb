// check option don't need to return error, because error is determined by business logic
// sometimes, the true case represents error. sometimes the false case indicates error, ever if that two case is generated by same checkFunc
package utils

import (
	"math"
	"strings"
)

const (
	maxAccountNameLength          = 30
	maxAccountNameLengthOmitSpace = 20

	minAssetId = 0
	maxAssetId = math.MaxUint32

	minAccountIndex = 0
	maxAccountIndex = math.MaxUint32

	minPublicKeyLength = 20
	maxPublicKeyLength = 50

	minPairIndex = 0
	maxPairIndex = math.MaxUint32

	minLPAmount = 0
	maxLPAmount = math.MaxUint32

	minTxtype = 0
	maxTxtype = 8

	minLimit = 0
	maxLimit = 50

	minOffset = 0
	maxOffset = math.MaxUint32
)

func CheckAccountName(accountName string) bool {
	return len(accountName) > maxAccountNameLength
}

func CheckFormatAccountName(accountName string) bool {
	return len(accountName) > maxAccountNameLengthOmitSpace
}

func CheckAccountPK(accountPK string) bool {
	return len(accountPK) > maxPublicKeyLength
}

func CheckAssetId(assetId uint32) bool {
	return assetId > maxAssetId
}

func CheckAccountIndex(accountIndex uint32) bool {
	return accountIndex > maxAccountIndex
}

func CheckPairIndex(pairIndex uint32) bool {
	return pairIndex > maxAccountIndex
}

func CheckAmount(amount string) bool { //true:report errors   false:continue
	return false
}

func CheckTxType(txType uint32) bool {
	return txType > maxTxtype
}

func CheckTypeLimit(limit uint32) bool {
	return limit > maxLimit
}

func CheckTypeOffset(offset uint32) bool {
	return offset > maxOffset
}

func CheckLPAmount(lPAmount uint32) bool {
	return lPAmount > maxLPAmount
}

func CheckOfferset(offerset, total uint32) bool {
	return offerset < total && offerset > 0
}

// Format AccountName and
func FormatSting(name string) string {
	name = strings.ToLower(name)
	name = strings.Replace(name, "\n", "", -1)
	return name
}

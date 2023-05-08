package recovery

import (
	"github.com/bnb-chain/zkbnb/common/log"
	"github.com/bnb-chain/zkbnb/tools/query"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"

	bsmt "github.com/bnb-chain/zkbnb-smt"
	"github.com/bnb-chain/zkbnb/tools/query/svc"
	"github.com/bnb-chain/zkbnb/tree"
)

func RecoveryTreeDB(
	configFile string,
	serviceName string,
	batchSize int,
) {
	configInfo := query.BuildConfig(configFile, serviceName)
	ctx := svc.NewServiceContext(configInfo)
	logx.MustSetup(configInfo.LogConf)
	logx.DisableStat()
	proc.AddShutdownListener(func() {
		logx.Close()
	})

	latestVerifiedBlockNr, err := ctx.BlockModel.GetLatestVerifiedHeight()
	if err != nil {
		logx.Errorf("get latest verified fromHeight failed: %v", err)
		return
	}

	ctxLog := log.NewCtxWithKV(log.BlockHeightContext, latestVerifiedBlockNr)

	logx.WithContext(ctxLog).Infof("the latest verified height is %d,recovery to blockHeight=%d", latestVerifiedBlockNr)

	// dbinitializer tree database
	treeCtx, err := tree.NewContext(serviceName, configInfo.TreeDB.Driver, true, false, configInfo.TreeDB.RoutinePoolSize, &configInfo.TreeDB.LevelDBOption, &configInfo.TreeDB.RedisDBOption)
	if err != nil {
		logx.WithContext(ctxLog).Errorf("Init tree database failed: %s", err)
		return
	}

	treeCtx.SetOptions(bsmt.InitializeVersion(0))
	treeCtx.SetBatchReloadSize(batchSize)
	err = tree.SetupTreeDB(treeCtx)
	if err != nil {
		logx.WithContext(ctxLog).Errorf("Init tree database failed: %s", err)
		return
	}

	// dbinitializer accountTree and accountStateTrees
	_, _, err = tree.InitAccountTree(
		ctx.AccountModel,
		ctx.AccountHistoryModel,
		make([]int64, 0),
		latestVerifiedBlockNr,
		treeCtx,
		configInfo.TreeDB.AssetTreeCacheSize,
		true,
	)
	if err != nil {
		logx.WithContext(ctxLog).Error("InitMerkleTree error:", err)
		return
	}
	logx.WithContext(ctxLog).Infof("recovery account smt successfully")

	// dbinitializer nftTree
	_, err = tree.InitNftTree(
		ctx.NftModel,
		ctx.NftHistoryModel,
		latestVerifiedBlockNr,
		treeCtx, true)
	if err != nil {
		logx.WithContext(ctxLog).Errorf("InitNftTree error: %s", err.Error())
		return
	}
	logx.WithContext(ctxLog).Infof("recovery nft smt successfully")

}

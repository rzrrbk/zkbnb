package rollbackwitnesssmt

import (
	"fmt"
	"github.com/bnb-chain/zkbnb/service/witness/config"
	"github.com/bnb-chain/zkbnb/service/witness/witness"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/proc"
)

func RollbackWitnessSmt(
	configFile string,
	height int64,
) error {
	c := config.Config{}
	if err := config.InitSystemConfiguration(&c, configFile); err != nil {
		logx.Severef("failed to initiate system configuration, %v", err)
		panic("failed to initiate system configuration, err:" + err.Error())
	}
	logx.MustSetup(c.LogConf)
	logx.DisableStat()
	proc.AddShutdownListener(func() {
		logx.Close()
	})

	w, err := witness.NewWitness(c)
	if err != nil {
		return fmt.Errorf("failed to create witness instance, %v", err)
	}
	err = w.Rollback(height)
	if err != nil {
		return fmt.Errorf("failed to rollback smt, %v", err)
	}
	logx.Infof("rollback smt success,the new smt version is %d now", height-1)
	return nil
}
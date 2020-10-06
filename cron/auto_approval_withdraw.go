package cron

import (
	"time"
	"xtek/exchange/commission/internal/domain/withdraw"

	"github.com/richard-xtek/go-grpc-micro-kit/log"
)

// NewAutoApprovalWithdrawCron ...
func NewAutoApprovalWithdrawCron(logger log.Factory, withdrawDomain *withdraw.Withdraw) *AutoApprovalWithdrawCron {
	return &AutoApprovalWithdrawCron{
		logger:         logger,
		withdrawDomain: withdrawDomain,
	}
}

// AutoApprovalWithdrawCron ...
type AutoApprovalWithdrawCron struct {
	withdrawDomain *withdraw.Withdraw
	logger         log.Factory
	stopped        chan bool

	isRunning bool
	isStopped bool
}

// Start ...
func (c *AutoApprovalWithdrawCron) Start() error {
	c.isStopped = false
	go func() {
		for {
			if c.isStopped {
				return
			}

			c.isRunning = true
			c.logger.Bg().Info("[AutoApprovalWithdrawCron] Run ...")
			c.withdrawDomain.AutoApprovalWithdrawCron(c.logger)
			c.logger.Bg().Info("[AutoApprovalWithdrawCron] Finished")
			c.isRunning = false

			time.Sleep(30 * time.Second)
		}
	}()
	return nil
}

// Stop ...
func (c *AutoApprovalWithdrawCron) Stop() error {

	c.logger.Bg().Info("[AutoApprovalWithdrawCron] Stopping ...")
	c.stop()
	c.logger.Bg().Info("[AutoApprovalWithdrawCron] Stopped")
	return nil
}

func (c *AutoApprovalWithdrawCron) stop() {
	if c.isRunning {
		c.logger.Bg().Warn("[AutoApprovalWithdrawCron] Running, please wait")
		time.Sleep(time.Second * 1)
		c.stop()
	}
	c.isStopped = true
}

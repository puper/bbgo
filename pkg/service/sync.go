package service

import (
	"context"
	"errors"
	"time"

	"github.com/puper/bbgo/pkg/cache"
	"github.com/puper/bbgo/pkg/util"

	log "github.com/sirupsen/logrus"

	"github.com/puper/bbgo/pkg/types"
)

var ErrNotImplemented = errors.New("not implemented")
var ErrExchangeRewardServiceNotImplemented = errors.New("exchange does not implement ExchangeRewardService interface")

type SyncService struct {
	TradeService    *TradeService
	OrderService    *OrderService
	RewardService   *RewardService
	WithdrawService *WithdrawService
	DepositService  *DepositService
	MarginService   *MarginService
}

// SyncSessionSymbols syncs the trades from the given exchange session
func (s *SyncService) SyncSessionSymbols(
	ctx context.Context, exchange types.Exchange, startTime time.Time, symbols ...string,
) error {
	markets, err := cache.LoadExchangeMarketsWithCache(ctx, exchange)
	if err != nil {
		return err
	}

	for _, symbol := range symbols {
		// skip symbols do not exist in the market info
		if _, ok := markets[symbol]; !ok {
			continue
		}

		log.Infof("syncing %s %s trades from %s...", exchange.Name(), symbol, startTime)
		if err := s.TradeService.Sync(ctx, exchange, symbol, startTime); err != nil {
			return err
		}

		log.Infof("syncing %s %s orders from %s...", exchange.Name(), symbol, startTime)
		if err := s.OrderService.Sync(ctx, exchange, symbol, startTime); err != nil {
			return err
		}
	}

	return nil
}

func (s *SyncService) SyncMarginHistory(
	ctx context.Context, exchange types.Exchange, startTime time.Time, assets ...string,
) error {
	if _, implemented := exchange.(types.MarginHistoryService); !implemented {
		log.Debugf("exchange %T does not support types.MarginHistoryService", exchange)
		return nil
	}

	if marginExchange, implemented := exchange.(types.MarginExchange); !implemented {
		log.Debugf("exchange %T does not implement types.MarginExchange", exchange)
		return nil
	} else {
		marginSettings := marginExchange.GetMarginSettings()
		if !marginSettings.IsMargin {
			log.Debugf("exchange %T is not using margin", exchange)
			return nil
		}
	}

	log.Infof("syncing %s margin history: %v...", exchange.Name(), assets)
	for _, asset := range assets {
		if err := s.MarginService.Sync(ctx, exchange, asset, startTime); err != nil {
			return err
		}
	}

	return nil
}

func (s *SyncService) SyncRewardHistory(ctx context.Context, exchange types.Exchange, startTime time.Time) error {
	if _, implemented := exchange.(types.ExchangeRewardService); !implemented {
		return nil
	}

	log.Infof("syncing %s reward records...", exchange.Name())
	if util.IsPaperTrade() {
		log.Info("reward is not supported in paper trading")
		return nil
	}
	if err := s.RewardService.Sync(ctx, exchange, startTime); err != nil {
		return err
	}

	return nil
}

func (s *SyncService) SyncDepositHistory(ctx context.Context, exchange types.Exchange, startTime time.Time) error {
	log.Infof("syncing %s deposit records...", exchange.Name())
	if util.IsPaperTrade() {
		log.Info("deposit is not supported in paper trading")
		return nil
	}

	if err := s.DepositService.Sync(ctx, exchange, startTime); err != nil {
		if err != ErrNotImplemented {
			log.Warnf("%s deposit service is not supported", exchange.Name())
			return err
		}
	}

	return nil
}

func (s *SyncService) SyncWithdrawHistory(ctx context.Context, exchange types.Exchange, startTime time.Time) error {
	log.Infof("syncing %s withdraw records...", exchange.Name())
	if util.IsPaperTrade() {
		log.Info("withdraw is not supported in paper trading")
		return nil
	}

	if err := s.WithdrawService.Sync(ctx, exchange, startTime); err != nil {
		if err != ErrNotImplemented {
			log.Warnf("%s withdraw service is not supported", exchange.Name())
			return err
		}
	}

	return nil
}

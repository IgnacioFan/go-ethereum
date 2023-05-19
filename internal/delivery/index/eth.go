package index

import (
	"context"
	"fmt"
	"go-ethereum/internal/service"
	"go-ethereum/internal/service/eth"
	"go-ethereum/internal/util"
	"go-ethereum/pkg/logger"
	"go-ethereum/pkg/postgres"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

type EthIndex struct {
	Service service.Eth
	Logger  *logrus.Logger
}

func NewEthIndexer() (*EthIndex, error) {
	logger := logger.NewLogger()
	db, err := postgres.NewPostgres()
	eth, err := eth.NewService(db)

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &EthIndex{Service: eth, Logger: logger}, nil
}

func (eth *EthIndex) Run(start, window, end int64) {
	ctx := context.Background()
	chainId, err := eth.Service.NetworkIDRPC(ctx)
	if err != nil {
		eth.Logger.Error(err)
		return
	}
	for {
		if start <= end {
			next := util.Min(start+window-1, end).(int64)
			existed, err := eth.Service.BlocksExist(ctx, start, next)
			if err != nil {
				eth.Logger.Error(err)
				return
			}
			if existed {
				start += window
				continue
			}
			for number := start; number <= next; number++ {
				go eth.ScanBlock(ctx, number, chainId)
				wg.Add(1)
			}
			wg.Wait()

			eth.Logger.Info(fmt.Sprintf("Scanned blocks from %v to %v", start, next))
			start += window
		}
		start = util.Min(start, end).(int64)
		eth.Logger.Info(fmt.Sprintf("Stop at block %v", start))
		time.Sleep(3 * time.Second)
	}
}

func (eth *EthIndex) ScanBlock(ctx context.Context, number, chainId int64) {
	defer wg.Done()
	if err := eth.Service.SaveBlockRPC(ctx, number, chainId); err != nil {
		eth.Logger.Error(err)
	}
	return
}

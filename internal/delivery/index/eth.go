package index

import (
	"context"
	"go-ethereum/internal/client/eth_client"
	"go-ethereum/internal/service"
	"go-ethereum/internal/service/eth"
	"go-ethereum/pkg/logger"
	"go-ethereum/pkg/postgres"
	"time"

	"github.com/sirupsen/logrus"
)

type EthIndex struct {
	Service service.Eth
	Logger  *logrus.Logger
}

func NewEthIndexer() (*EthIndex, error) {
	logger := logger.NewLogger()
	client, err := eth_client.NewClient()
	db, err := postgres.NewPostgres()
	eth, err := eth.NewService(db, client)
	if err != nil {
		logger.Error("Failed at NewEthIndexer", err)
		return nil, err
	}

	return &EthIndex{Service: eth, Logger: logger}, nil
}

func (i *EthIndex) Run(start, end int) {
	blockNumCh := make(chan int64, 10000)
	go i.Process(blockNumCh)
	for {
		if start == end {
			time.Sleep(time.Second)
			continue
		}
		blockNumCh <- int64(start)
		start++
	}
}

func (i *EthIndex) Process(blockNumCh <-chan int64) {
	for blockId := range blockNumCh {
		ctx := context.Background()
		err := i.Service.SaveBlock(ctx, blockId)
		if err != nil {
			i.Logger.Error("Failed at Process", err)
			continue
		}
	}
}

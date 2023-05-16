package index

import (
	"context"
	"fmt"
	"go-ethereum/internal/service/eth"
	"time"
)

type EthIndex struct {
	Service eth.EthService
}

func NewEthIndexer() *EthIndex {
	eth, err := eth.NewService()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &EthIndex{Service: eth}
}

func (i *EthIndex) Run() {
	blockNumCh := make(chan int64, 10000)
	startBlock := 1
	endBlock := 5
	go i.Process(blockNumCh)
	for {
		if endBlock == startBlock {
			time.Sleep(time.Second)
			continue
		}
		blockNumCh <- int64(startBlock)
		startBlock++
	}
}

func (i *EthIndex) Process(blockNumCh <-chan int64) {
	for n := range blockNumCh {
		ctx := context.Background()
		err := i.Service.SaveBlock(ctx, n)
		if err != nil {
			fmt.Println("Process", err)
			continue
		}
	}
}

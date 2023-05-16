package eth_client

import (
	"context"
	"fmt"
	"go-ethereum/internal/entity"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://ethereum-mainnet-rpc.allthatnode.com"

type Impl struct {
	Client *ethclient.Client
}

func NewClient() (Client, error) {
	c, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("ethclient.Dial failed: %v\n", err)
	}
	return &Impl{Client: c}, nil
}

func (i *Impl) BlockByNumber(ctx context.Context, blockId int64) (*entity.Block, error) {
	bigNum := big.NewInt(blockId)
	res, err := i.Client.BlockByNumber(ctx, bigNum)
	if err != nil {
		return nil, fmt.Errorf("ethclient.BlockByNumber failed: %v\n", err)
	}

	return &entity.Block{
		Number:     res.NumberU64(),
		Hash:       res.Hash().String(),
		Timestamp:  res.Time(),
		ParentHash: res.ParentHash().String(),
	}, nil
}

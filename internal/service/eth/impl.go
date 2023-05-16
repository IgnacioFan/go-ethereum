package eth

import (
	"context"
	"fmt"
	"go-ethereum/internal/repository/eth_client"
	"go-ethereum/pkg/postgres"
)

type Impl struct {
	Client eth_client.Client
	DB     *postgres.Postgres
}

func NewService() (EthService, error) {
	client, err := eth_client.NewClient()
	if err != nil {
		return nil, err
	}
	db, err := postgres.NewPostgres()
	if err != nil {
		return nil, err
	}
	return &Impl{
		Client: client,
		DB:     db,
	}, nil
}

func (i *Impl) GetBlocks() {

}

func (i *Impl) GetBlock() {
	// by id
}

func (i *Impl) GetTransaction() {
	// by hash
}

func (i *Impl) SaveBlock(ctx context.Context, blockId int64) error {
	block, err := i.Client.BlockByNumber(ctx, blockId)
	if err != nil {
		// add system log
		return err
	}
	fmt.Println("SaveBlock", block)
	res := i.DB.DB.Create(block)
	fmt.Println("SaveBlock", res)
	if res.Error != nil {
		// add system log
		return res.Error
	}
	return nil
}

func (i *Impl) SaveTransaction() {

}

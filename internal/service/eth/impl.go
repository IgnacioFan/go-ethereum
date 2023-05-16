package eth

import (
	"context"
	"fmt"
	"go-ethereum/internal/client"
	"go-ethereum/internal/client/eth_client"
	"go-ethereum/internal/repository"
	"go-ethereum/internal/repository/eth_repo"
	"go-ethereum/internal/service"
	"go-ethereum/pkg/postgres"
)

type Impl struct {
	Client client.Eth
	Repo   repository.Eth
}

func NewService(db *postgres.Postgres) (service.Eth, error) {
	client, err := eth_client.NewClient()
	if err != nil {
		return nil, err
	}
	return &Impl{
		Client: client,
		Repo:   eth_repo.NewRepo(db),
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
	return i.Repo.SaveBlock(block)
}

func (i *Impl) SaveTransaction() {

}

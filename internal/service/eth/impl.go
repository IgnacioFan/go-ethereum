package eth

import (
	"context"
	"go-ethereum/internal/entity"
	"go-ethereum/internal/repository"
	"go-ethereum/internal/repository/eth_repo"
	"go-ethereum/internal/service"
	"go-ethereum/pkg/postgres"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "test"
)

type Impl struct {
	Client *ethclient.Client
	Repo   repository.Eth
}

func NewService(db *postgres.Postgres) (service.Eth, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Impl{
		Client: client,
		Repo:   eth_repo.NewRepo(db),
	}, nil
}

func (i *Impl) BlocksExist(ctx context.Context, startNumber, endNumber int64) (bool, error) {
	return i.Repo.BlocksExist(startNumber, endNumber)
}

func (i *Impl) BlockExist(ctx context.Context, number int64) (bool, error) {
	return i.Repo.BlockExist(number)
}

func (i *Impl) NetworkIDRPC(ctx context.Context) (int64, error) {
	chainId, err := i.Client.NetworkID(ctx)
	return chainId.Int64(), err
}

func (i *Impl) SaveBlockRPC(ctx context.Context, number, chainId int64) error {
	exist, err := i.BlockExist(ctx, number)
	if exist {
		return nil
	} else if err != nil {
		return err
	}

	block, err := i.Client.BlockByNumber(ctx, big.NewInt(number))
	if err != nil {
		return err
	}

	transactions := make([]entity.Transaction, 0)
	signer := types.LatestSignerForChainID(big.NewInt(chainId))

	for _, tx := range block.Transactions() {
		sender, err := types.Sender(signer, tx)
		if err != nil {
			return err
		}
		var val int64
		if tx.Value() != nil {
			val = tx.Value().Int64()
		}
		var to string
		if tx.To() != nil {
			to = tx.To().String()
		}

		transactions = append(transactions, entity.Transaction{
			Hash:  tx.Hash().String(),
			From:  sender.String(),
			To:    to,
			Nonce: tx.Nonce(),
			Data:  common.Bytes2Hex(tx.Data()),
			Value: val,
		})
	}

	return i.Repo.SaveBlock(&entity.Block{
		Number:       block.NumberU64(),
		Hash:         block.Hash().String(),
		Timestamp:    block.Time(),
		ParentHash:   block.ParentHash().String(),
		Transactions: transactions,
	})
}

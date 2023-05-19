package http

import (
	"fmt"
	"go-ethereum/internal/service"
	"go-ethereum/internal/service/eth"
	"go-ethereum/pkg/logger"
	"go-ethereum/pkg/postgres"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Http struct {
	*gin.Engine
	Service service.Eth
	Logger  *logrus.Logger
}

func NewHttpServer() (*Http, error) {
	logger := logger.NewLogger()
	db, err := postgres.NewPostgres()
	eth, err := eth.NewService(db)
	if err != nil {
		logger.Error("Failed at NewEthIndexer", err)
		return nil, err
	}

	server := &Http{
		Engine:  gin.Default(),
		Service: eth,
		Logger:  logger,
	}
	server.SetRoute()
	return server, nil
}

func (h *Http) Start(port int) error {
	return h.Run(fmt.Sprintf(":%d", port))
}

func (h *Http) SetRoute() {
	blocks := h.Group("/blocks")
	blocks.GET("/", h.getBlocks)
	blocks.GET("/:id", h.getBlock)

	transaction := h.Group("/transaction")
	transaction.GET("/:txHash", h.getTransaction)
}

package http

import (
	"fmt"
	"go-ethereum/internal/service"
	"go-ethereum/internal/service/eth"
	"go-ethereum/internal/util"
	"go-ethereum/pkg/logger"
	"go-ethereum/pkg/postgres"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	MIN_LIMIT int = 10
	MAX_LIMIT int = 100
)

type Server struct {
	*gin.Engine
	Service service.Eth
	Logger  *logrus.Logger
}

func NewHttpServer() (*Server, error) {
	logger := logger.NewLogger()
	db, err := postgres.NewPostgres()
	eth, err := eth.NewService(db)
	if err != nil {
		logger.Error("Failed at NewEthIndexer", err)
		return nil, err
	}

	server := &Server{
		Engine:  gin.Default(),
		Service: eth,
		Logger:  logger,
	}
	server.SetRoute()
	return server, nil
}

func (s *Server) Start(port int) error {
	return s.Run(fmt.Sprintf(":%d", port))
}

func (s *Server) SetRoute() {
	v1 := s.Group("api/v1")
	{
		v1.GET("blocks", func(ctx *gin.Context) {
			var limit int
			if limitStr := ctx.Query("limit"); len(limitStr) == 0 {
				limit = MIN_LIMIT
			} else {
				num, err := strconv.Atoi(limitStr)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, s.errMsg(err))
					return
				}
				limit = util.Min(num, MAX_LIMIT).(int)
			}

			blocks, err := s.Service.GetBlocks(ctx, limit)
			if err != nil {
				s.Logger.Debug("blocks, error: ", err)
				ctx.JSON(http.StatusInternalServerError, s.errMsg(err))
				return
			}

			ctx.JSON(http.StatusOK, blocks)
		})
		v1.GET("blocks/:id", func(ctx *gin.Context) {
			blockId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
			block, err := s.Service.GetBlock(ctx, blockId)
			if err != nil {
				s.Logger.Debug("blocks/:id, error: ", err)
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			ctx.JSON(http.StatusOK, block)
		})
		v1.GET("transaction/:txHash", func(ctx *gin.Context) {
			txHash := ctx.Param("txHash")
			tx, err := s.Service.GetTransaction(ctx, txHash)
			if err != nil {
				s.Logger.Debug("transaction/:txHash, error: ", err)
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, tx)
		})
	}
}

func (s *Server) errMsg(err error) interface{} {
	return gin.H{
		"error": err.Error(),
	}
}

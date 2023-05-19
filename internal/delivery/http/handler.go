package http

import (
	"go-ethereum/internal/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MIN_LIMIT int = 10
	MAX_LIMIT int = 100
)

func (h *Http) getBlocks(ctx *gin.Context) {
	var limit int
	if limitStr := ctx.Query("limit"); len(limitStr) == 0 {
		limit = MIN_LIMIT
	} else {
		num, err := strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errMsg(err))
			return
		}
		limit = util.Min(num, MAX_LIMIT).(int)
	}

	blocks, err := h.Service.GetBlocks(ctx, limit)
	if err != nil {
		h.Logger.Debug("blocks, error: ", err)
		ctx.JSON(http.StatusInternalServerError, errMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, blocks)
}

func (h *Http) getBlock(ctx *gin.Context) {
	blockId, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	block, err := h.Service.GetBlock(ctx, blockId)
	if err != nil {
		h.Logger.Debug("blocks/:id, error: ", err)
		ctx.JSON(http.StatusInternalServerError, errMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, block)
}

func (h *Http) getTransaction(ctx *gin.Context) {
	txHash := ctx.Param("txHash")
	tx, err := h.Service.GetTransaction(ctx, txHash)
	if err != nil {
		h.Logger.Debug("transaction/:txHash, error: ", err)
		ctx.JSON(http.StatusInternalServerError, errMsg(err))
		return
	}
	ctx.JSON(http.StatusOK, tx)
}

func errMsg(err error) interface{} {
	return gin.H{
		"error": err.Error(),
	}
}

package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewHttpServer() *Server {
	server := &Server{
		Engine: gin.Default(),
	}
	server.SetRoute()
	return server
}

func (s *Server) Start(port int) error {
	return s.Run(fmt.Sprintf(":%d", port))
}

func (s *Server) SetRoute() {
	v1 := s.Group("api/v1")
	{
		v1.GET("blocks", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "returns latest n blocks")
		})
		v1.GET("blocks/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "returns block with id")
		})
		v1.GET("transaction/:txHash", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "returns TX data with event logs")
		})
	}
}

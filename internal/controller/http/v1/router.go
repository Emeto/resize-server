package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"resize-server/internal/usecase"
	"resize-server/pkg/logger"
)

func NewRouter(handler *gin.Engine, l logger.Interface, s *usecase.ScalingUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	h := handler.Group("/v1")
	{
		newScalingRoutes(h, s, l)
	}
}

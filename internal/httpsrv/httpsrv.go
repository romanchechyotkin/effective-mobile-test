package httpsrv

import (
	"net/http"

	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/api"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	log    *zap.Logger
	base   *http.Server
	router *gin.Engine
	client *api.Client
	db     *db.QBuilder
}

type HTTPServer interface {
	RegisterRoutes()
}

func NewServer(cfg *HTTPConfig, log *zap.Logger, apiClient *api.Client, builder *db.QBuilder) (*Server, error) {
	instance := Server{
		log:    log,
		router: gin.Default(),
		client: apiClient,
		db:     builder,
	}

	instance.base = &http.Server{
		Addr:    cfg.Bind + ":" + cfg.Port,
		Handler: instance.router,
	}

	instance.log.Debug("http server configuration", zap.Any("cfg", cfg))

	return &instance, nil
}

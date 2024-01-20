package httpsrv

import (
	"net/http"

	"github.com/romanchechyotkin/effective-mobile-test-task/internal/storage"
	"github.com/romanchechyotkin/effective-mobile-test-task/internal/storage/repo"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UsersCollection interface {
	Users() *repo.Users
}

type Server struct {
	log     *zap.Logger
	base    *http.Server
	router  *gin.Engine
	client  *api.Client
	storage UsersCollection
}

func NewServer(cfg *HTTPConfig, log *zap.Logger, apiClient *api.Client, collection *storage.Collection) (*Server, error) {
	instance := Server{
		log:     log,
		router:  gin.Default(),
		client:  apiClient,
		storage: collection,
	}

	instance.base = &http.Server{
		Addr:    cfg.Bind + ":" + cfg.Port,
		Handler: instance.router,
	}

	instance.log.Debug("http server configuration", zap.Any("cfg", cfg))

	return &instance, nil
}

package httpsrv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	log    *zap.Logger
	base   http.Server
	Router *gin.Engine
}

type Endpoints interface {
	Target() http.HandlerFunc
}

func NewServer(cfg *HTTPConfig, log *zap.Logger) (*Server, error) {
	instance := Server{
		log:    log,
		Router: gin.Default(),
	}

	instance.base = http.Server{
		Addr:        cfg.Bind + ":" + cfg.Port,
		IdleTimeout: cfg.IdleTimeout,
		Handler:     instance.Router,
	}

	return &instance, nil
}

package v1

import (
	"github.com/SayKonstantin/metrika-service/internal/config"
	"github.com/SayKonstantin/metrika-service/pb"
	"github.com/nikoksr/notify"
	"github.com/rs/zerolog"
)

type Server struct {
	cfg      config.ServerConfig
	logger   *zerolog.Logger
	notifier notify.Notifier
	pb.UnimplementedMetrikaServiceServer
}

func NewServer(cfg config.ServerConfig, logger *zerolog.Logger, srv pb.UnimplementedMetrikaServiceServer, notify notify.Notifier) *Server {
	apiLogger := logger.With().Str("api", "grpc").Logger()

	return &Server{
		cfg:                               cfg,
		logger:                            &apiLogger,
		notifier:                          notify,
		UnimplementedMetrikaServiceServer: srv,
	}
}

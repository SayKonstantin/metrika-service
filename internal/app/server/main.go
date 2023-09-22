package server

import (
	"context"
	"fmt"
	"github.com/SayKonstantin/metrika-service/internal/config"
	v1 "github.com/SayKonstantin/metrika-service/internal/controllers/api/grpc/v1"
	"github.com/SayKonstantin/metrika-service/pb"
	"github.com/nikoksr/notify"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	cfg                  *config.ServerConfig
	grpcServer           *grpc.Server
	productServiceServer pb.MetrikaServiceServer
	logger               *zerolog.Logger
	Notify               notify.Notifier
}

func NewApp(cfg *config.ServerConfig, logger *zerolog.Logger, notify notify.Notifier) *App {
	srv := v1.NewServer(*cfg, logger, pb.UnimplementedMetrikaServiceServer{}, notify)

	return &App{
		cfg:                  cfg,
		grpcServer:           nil,
		productServiceServer: srv,
		logger:               logger,
		Notify:               notify,
	}
}

func (a App) startGRPC(server pb.MetrikaServiceServer) error {
	a.logger.Info().Msg(fmt.Sprintf("GRPC запущен на %s:%d", a.cfg.GRPC.Ip, a.cfg.GRPC.Port))
	if a.cfg.IsEnabled {
		err := a.Notify.Send(context.Background(), "Metrika Service", fmt.Sprintf("gRPC запущен на %v:%v", a.cfg.GRPC.Ip, a.cfg.GRPC.Port))
		if err != nil {
			a.logger.Fatal().Err(err).Msg("ошибка отправки уведомления")
		}
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.Ip, a.cfg.GRPC.Port))
	if err != nil {
		a.logger.Fatal().Err(err).Msg("failed to create listener")
	}
	a.grpcServer = grpc.NewServer()
	pb.RegisterMetrikaServiceServer(a.grpcServer, server)
	if err := a.grpcServer.Serve(lis); err != nil {
		a.logger.Fatal().Err(err).Msg("failed to serve")
	}
	return nil
}

func (a App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startGRPC(a.productServiceServer)
	})
	return grp.Wait()
}

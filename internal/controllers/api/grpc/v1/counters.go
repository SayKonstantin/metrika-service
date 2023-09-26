package v1

import (
	"context"
	mp "github.com/SayKonstantin/metrika-service/internal/domain/policy"
	"github.com/SayKonstantin/metrika-service/internal/domain/service"
	"github.com/SayKonstantin/metrika-service/internal/repository/metrika"
	"github.com/SayKonstantin/metrika-service/pb"
	sdk "github.com/mg-realcom/metrika-sdk"
	"net/http"
	"time"
)

func (s Server) GetCounters(ctx context.Context, req *pb.GetCountersRequest) (*pb.GetCountersResponse, error) {
	s.logger.Info().Msg("GetCounters")
	metrikaClient := sdk.Client{
		Tr:    &http.Client{Timeout: 20 * time.Minute},
		Token: req.YandexToken,
	}
	metrikaRepo := metrika.NewCounterRepository(&metrikaClient, s.logger)
	srv := service.NewCounterService(metrikaRepo, s.logger)
	metrikaPolicy := mp.NewCounterPolicy(*srv)
	counters, err := metrikaPolicy.GetCounters(ctx)
	if err != nil {
		return &pb.GetCountersResponse{Counters: ""}, err
	}
	return &pb.GetCountersResponse{Counters: counters}, nil

}

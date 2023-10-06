package v1

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	mp "github.com/SayKonstantin/metrika-service/internal/domain/policy"
	"github.com/SayKonstantin/metrika-service/internal/domain/service"
	"github.com/SayKonstantin/metrika-service/internal/repository/bq"
	"github.com/SayKonstantin/metrika-service/internal/repository/cs"
	"github.com/SayKonstantin/metrika-service/internal/repository/metrika"
	"github.com/SayKonstantin/metrika-service/pb"
	sdk "github.com/mg-realcom/metrika-sdk"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

func (s Server) PushHitsToBQ(ctx context.Context, req *pb.PushLogRequest) (*pb.PushLogResponse, error) {
	methodLogger := s.logger.With().Str("method", "PushHitsToBQ").Logger()
	methodLogger.Info().Msg(msgMethodPrepared)
	defer methodLogger.Info().Msg(msgMethodFinished)

	dates := metrika.DateRange{DateTo: req.MetrikaConfig.Period.DateTo, DateFrom: req.MetrikaConfig.Period.DateFrom}
	metrikaClient := sdk.Client{
		Tr:        &http.Client{Timeout: 20 * time.Minute},
		Token:     req.MetrikaConfig.YandexToken,
		CounterId: req.MetrikaConfig.CounterId,
	}
	metrikaRepo := metrika.NewHitRepository(&metrikaClient, s.cfg.AttachmentsDir, metrika.HitsFields, metrika.HitsSource, &methodLogger)
	cloudClient := option.WithCredentialsFile(s.cfg.KeysDir + "/" + req.BqConfig.ServiceKey)

	bqClient, err := bigquery.NewClient(ctx, req.BqConfig.ProjectId, cloudClient)
	if err != nil {
		s.notifier.Send(ctx, "Metrika Service", fmt.Sprintf("PushHitsToBQ: Creation BQ client error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	defer func(bqClient *bigquery.Client) {
		err := bqClient.Close()
		if err != nil {
			methodLogger.Error().Err(err).Msg(msgErrMethod)
		}
	}(bqClient)
	bqRepo := bq.NewHitRepository(bqClient, req.BqConfig.DatasetId, req.BqConfig.TableId, &methodLogger)
	csClient, err := storage.NewClient(ctx, cloudClient)
	if err != nil {
		s.notifier.Send(ctx, "Metrika Service", fmt.Sprintf("PushHits: Error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	defer func(csClient *storage.Client) {
		err := csClient.Close()
		if err != nil {
			methodLogger.Error().Err(err).Msg(msgErrMethod)
		}
	}(csClient)
	csRepo := cs.NewHitRepository(req.CsConfig.BucketName, req.BqConfig.ProjectId, csClient, &methodLogger)
	srv := service.NewHitService(metrikaRepo, csRepo, bqRepo, dates, &methodLogger)
	metrikaPolicy := mp.NewHitPolicy(*srv)
	methodLogger.Info().Msg(msgMethodStarted)
	err = metrikaPolicy.PushHitsToBQ(ctx)
	if err != nil {
		methodLogger.Error().Err(err).Msg(msgErrMethod)
		s.notifier.Send(ctx, "Metrika Service", fmt.Sprintf("PushHitsToBQ: execution error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	return &pb.PushLogResponse{Success: true}, nil

}

package v1

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	mp "github.com/SayKonstantin/metrika-service/internal/domain/policy"
	service "github.com/SayKonstantin/metrika-service/internal/domain/service"
	"github.com/SayKonstantin/metrika-service/internal/repository/bq"
	"github.com/SayKonstantin/metrika-service/internal/repository/cs"
	"github.com/SayKonstantin/metrika-service/internal/repository/metrika"
	"github.com/SayKonstantin/metrika-service/pb"
	sdk "github.com/mg-realcom/metrika-sdk"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

func (s Server) PushVisitsToBQ(ctx context.Context, req *pb.PushLogRequest) (*pb.PushLogResponse, error) {
	methodLogger := s.logger.With().Str("method", "PushHitsToBQ").Logger()
	methodLogger.Info().Msg(msgMethodPrepared)
	defer methodLogger.Info().Msg(msgMethodFinished)

	dates := metrika.DateRange{DateTo: req.MetrikaConfig.Period.DateTo, DateFrom: req.MetrikaConfig.Period.DateFrom}
	metrikaClient := sdk.Client{
		Tr:        &http.Client{Timeout: 20 * time.Minute},
		Token:     req.MetrikaConfig.YandexToken,
		CounterId: req.MetrikaConfig.CounterId,
	}
	metrikaRepo := metrika.NewVisitRepository(&metrikaClient, s.cfg.AttachmentsDir, metrika.VisitsFields, metrika.VisitsSource, &methodLogger)
	cloudClient := option.WithCredentialsFile(s.cfg.KeysDir + "/" + req.BqConfig.ServiceKey)

	bqClient, err := bigquery.NewClient(ctx, req.BqConfig.ProjectId, cloudClient)
	if err != nil {
		s.notifier.Send(ctx, "Metrika Service", fmt.Sprintf("PushVisits: Error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	defer func(bqClient *bigquery.Client) {
		err := bqClient.Close()
		if err != nil {
			methodLogger.Error().Err(err).Msg(msgErrMethod)
		}
	}(bqClient)
	bqRepo := bq.NewVisitRepository(bqClient, req.BqConfig.DatasetId, req.BqConfig.TableId, &methodLogger)
	csClient, err := storage.NewClient(ctx, cloudClient)
	if err != nil {
		s.notifier.Send(ctx, "Metrika Service", fmt.Sprintf("PushVisitsToBQ: execution error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	defer func(csClient *storage.Client) {
		err := csClient.Close()
		if err != nil {
			methodLogger.Error().Err(err).Msg(msgErrMethod)
		}
	}(csClient)

	csRepo := cs.NewVisitRepository(req.CsConfig.BucketName, req.BqConfig.ProjectId, csClient, &methodLogger)
	srv := service.NewVisitService(metrikaRepo, csRepo, bqRepo, dates, &methodLogger)
	metrikaPolicy := mp.NewVisitPolicy(*srv)
	methodLogger.Info().Msg(msgMethodStarted)
	err = metrikaPolicy.PushVisits(ctx)
	if err != nil {
		s.notifier.Send(ctx, "Metrika Service", fmt.Sprintf("PushVisits: Error: %s", err))
		return &pb.PushLogResponse{Success: false}, err
	}
	return &pb.PushLogResponse{Success: true}, nil
}
